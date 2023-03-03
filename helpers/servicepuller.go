package main

import (
	"fmt"
	"github.com/Jeffail/gabs/v2"
	"github.com/dave/jennifer/jen"
	"github.com/iancoleman/strcase"
	"sort"
	"strings"
)

type ServiceList struct {
	services     []Domain
	serviceNames []string
	json         *gabs.Container
	enums        map[string]map[string][]string
}

type Domain struct {
	name        string
	camelName   string
	services    map[string]*Service
	servicesKey []string
}

type Service struct {
	name           string
	camelName      string
	lowerCamelName string
	firstLetter    string
	parameters     map[string]*jen.Statement
	parameterKeys  []string
	description    string
}

func ServicesInit() ServiceList {
	var serviceList ServiceList

	serviceList.json, _ = gabs.ParseJSONFile("./helpers/services.json")
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

func (sl *ServiceList) setServiceNames() {
	deviceServices := sl.json.Search("service_result")
	// this will range over each service in the json
	for k := range deviceServices.ChildrenMap() {
		sl.serviceNames = append(sl.serviceNames, k)
	}

}

func (sl *ServiceList) setParameters(d *Domain) {
	deviceServices := sl.json.Search("service_result", d.name)
	// this will range over each service in the json
	for k, v := range deviceServices.ChildrenMap() {
		s := &Service{
			name:           k,
			camelName:      fmt.Sprintf("%s%s", d.camelName, strcase.ToCamel(k)),
			lowerCamelName: strcase.ToLowerCamel(fmt.Sprintf("%s%s", d.camelName, strcase.ToCamel(k))),
			firstLetter:    string(d.name[0]),
			parameters:     make(map[string]*jen.Statement),
			description:    v.Path("description").String(),
		}
		for fn, f := range v.Path("fields").ChildrenMap() {
			selector := f.Path("selector")

			if selector != nil {
				sm := selector.ChildrenMap()

				if _, ok := sm["text"]; ok {
					s.parameters[fn] = jen.String()
				}
				if _, ok := sm["number"]; ok {
					s.parameters[fn] = jen.Float64()
				}
				if _, ok := sm["color_temp"]; ok {
					s.parameters[fn] = jen.Float64()
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
							}
							sl.enums[strcase.ToCamel(fn)][enumName] = append(sl.enums[strcase.ToCamel(fn)][enumName], fmt.Sprintf("%s: %s", d.name, s.name))

						}
						s.parameters[fn] = jen.Id(strcase.ToCamel(fn))
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
