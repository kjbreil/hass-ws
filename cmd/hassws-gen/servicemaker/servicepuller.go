package servicemaker

import (
	"fmt"
	"github.com/Jeffail/gabs/v2"
	"github.com/iancoleman/strcase"
	"sort"
	"strings"
)

// ServiceList represents a list of Home Assistant service domains and their metadata, as parsed from JSON.
type ServiceList struct {
	services     []Domain
	serviceNames []string
	json         *gabs.Container
	enums        map[string]map[string][]string
}

// Domain represents a Home Assistant service domain, containing its name and a set of services.
type Domain struct {
	name        string
	camelName   string
	services    map[string]*Service
	servicesKey []string
}

// Service represents a single Home Assistant service, including its name, parameters, and description.
type Service struct {
	name             string
	camelName        string
	lowerCamelName   string
	firstLetter      string
	parameters       map[string]string
	parameterKeys    []string
	description      string
	responseRequired bool
}

// ServicesInit loads service data from a JSON file and returns a populated ServiceList struct.
func ServicesInit(servicesJSON string) ServiceList {
	var serviceList ServiceList

	serviceList.json, _ = gabs.ParseJSONFile(servicesJSON)
	serviceList.enums = make(map[string]map[string][]string)

	serviceList.setServiceNames()

	for _, name := range serviceList.serviceNames {
		d := Domain{
			name:      name,
			camelName: strcase.ToCamel(name),
			services:  make(map[string]*Service),
		}
		serviceList.setParameters(&d)
		serviceList.services = append(serviceList.services, d)
	}
	return serviceList
}

// ServicesInitJson loads service data from a JSON byte slice and returns a populated ServiceList struct.
func ServicesInitJson(data []byte) ServiceList {
	var serviceList ServiceList

	serviceList.json, _ = gabs.ParseJSON(data)
	serviceList.enums = make(map[string]map[string][]string)

	serviceList.setServiceNames()

	for _, name := range serviceList.serviceNames {
		// if name != "babybuddy" {
		//	continue
		// }
		d := Domain{
			name:      name,
			camelName: strcase.ToCamel(name),
			services:  make(map[string]*Service),
		}
		serviceList.setParameters(&d)
		serviceList.services = append(serviceList.services, d)
	}
	return serviceList
}

// setServiceNames populates the serviceNames field of the ServiceList by extracting domain names from the JSON.
// It searches for the "service_result" key in the JSON and appends each child key to the serviceNames slice.
func (sl *ServiceList) setServiceNames() {
	deviceServices := sl.json.Search("service_result")
	// this will range over each service in the json
	for k := range deviceServices.ChildrenMap() {
		sl.serviceNames = append(sl.serviceNames, k)
	}
}

// setParameters populates the parameters and metadata for all services in a given domain.
// It searches for the "service_result" key in the JSON with the given domain name, and for each service,
// it extracts the parameters, description, and response required flag, and populates the Service struct.
func (sl *ServiceList) setParameters(d *Domain) {
	deviceServices := sl.json.Search("service_result", d.name)
	// this will range over each service in the json
	for k, v := range deviceServices.ChildrenMap() {
		if strings.Contains(k, "-") {
			continue
		}
		s := &Service{
			name:             k,
			camelName:        fmt.Sprintf("%s%s", d.camelName, strcase.ToCamel(k)),
			lowerCamelName:   strcase.ToLowerCamel(fmt.Sprintf("%s%s", d.camelName, strcase.ToCamel(k))),
			firstLetter:      string(d.name[0]),
			parameters:       make(map[string]string),
			responseRequired: false,
			description:      v.Path("description").String(),
		}
		op := v.Path("response")
		if op != nil {
			sm := op.ChildrenMap()
			if b, ok := sm["optional"]; ok {
				s.responseRequired = b.Data().(bool) == false
			}
		}

		for fn, f := range v.Path("fields").ChildrenMap() {
			selector := f.Path("selector")

			if selector != nil {
				sm := selector.ChildrenMap()
				if strings.EqualFold(fn, "type") {
					fn = fmt.Sprintf("%s_%s", k, fn)
				}

				if _, ok := sm["text"]; ok {
					s.parameters[fn] = "string"
				}
				if _, ok := sm["number"]; ok {
					s.parameters[fn] = "float64"
				}
				if _, ok := sm["color_temp"]; ok {
					s.parameters[fn] = "float64"
				}
				if _, ok := sm["object"]; ok {
					if fn == "data" {
						s.parameters[fn] = "interface{}"
					}
				}
				if se, ok := sm["select"]; ok {
					options := se.ChildrenMap()["options"].Children()
					if options != nil {
						for _, o := range options {
							if _, ok := sl.enums[strcase.ToCamel(fn)]; !ok {
								sl.enums[strcase.ToCamel(fn)] = make(map[string][]string)
							}

							enumName := strings.Trim(o.Path("value").String(), "\"")
							if enumName == "null" {
								enumName = strings.Trim(o.String(), "\"")
								enumName = strings.ReplaceAll(enumName, "-", "")
								enumName = strings.ReplaceAll(enumName, " ", "")

							}
							// enumName = strcase.ToCamel(enumName)
							sl.enums[strcase.ToCamel(fn)][enumName] = append(sl.enums[strcase.ToCamel(fn)][enumName], fmt.Sprintf("%s: %s", d.name, s.name))

						}
						s.parameters[fn] = strcase.ToCamel(fn)
					}
				}

			}
		}
		for key := range s.parameters {
			s.parameterKeys = append(s.parameterKeys, key)
		}
		sort.Strings(s.parameterKeys)

		// assign service to domain services
		d.services[k] = s
	}
	for key := range d.services {
		d.servicesKey = append(d.servicesKey, key)
	}
	sort.Strings(d.servicesKey)

}
