package main

import (
	"fmt"
	"github.com/Jeffail/gabs/v2"
	"github.com/dave/jennifer/jen"
	"github.com/iancoleman/strcase"
	"strings"
)

var ServiceNames = append([]string{
	"automation",
	"backup",
	"counter",
	"frontend",
	"group",
	"homeassistant",
	"input_boolean",
	"input_button",
	"input_datetime",
	"input_number",
	"input_select",
	"input_text",
	"logbook",
	"logger",
	"min_max",
	"mqtt",
	"notify",
	"persistent_notification",
	"person",
	"recorder",
	"scene",
	"schedule",
	"script",
	"shell_command",
	"system_log",
	"template",
	"timer",
	"tts",
	"wake_on_lan",
	"zone",
}, DeviceNames...)

type ServiceList struct {
	services []Domain
	json     *gabs.Container
	enums    map[string]map[string]struct{}
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
}

func ServicesInit() ServiceList {
	var serviceList ServiceList

	serviceList.json, _ = gabs.ParseJSONFile("./helpers/services.json")
	serviceList.enums = make(map[string]map[string]struct{})

	for _, name := range ServiceNames {
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

func (sl *ServiceList) setParameters(d *Domain) {
	deviceServices := sl.json.Search("service_result", d.name)
	// this will range over each service in the json
	for k, v := range deviceServices.ChildrenMap() {
		s := &Service{
			name:           k,
			camelName:      fmt.Sprintf("%s%s", d.camelName, strcase.ToCamel(k)),
			lowerCamelName: strcase.ToLowerCamel(fmt.Sprintf("%s%s", d.camelName, strcase.ToCamel(k))),
			firstLetter:    string(k[0]),
			parameters:     make(map[string]*jen.Statement),
		}
		for fn, f := range v.Path("fields").ChildrenMap() {
			selector := f.Path("selector")

			if selector != nil {
				sm := selector.ChildrenMap()

				if _, ok := sm["text"]; ok {
					s.parameters[fn] = jen.String()
				}
				if _, ok := sm["number"]; ok {
					s.parameters[fn] = jen.Int()
				}
				if se, ok := sm["select"]; ok {
					options := se.ChildrenMap()["options"].Children()
					if options != nil {
						for _, o := range options {
							if _, ok := sl.enums[strcase.ToCamel(fn)]; !ok {
								sl.enums[strcase.ToCamel(fn)] = make(map[string]struct{})
							}
							sl.enums[strcase.ToCamel(fn)][strings.Trim(o.Path("value").String(), "\"")] = struct{}{}
						}
						s.parameters[fn] = jen.Id(strcase.ToCamel(fn))
					}
				}
			}
		}
		for key := range s.parameters {
			s.parameterKeys = append(s.parameterKeys, key)
		}
		// assign service to domain services
		d.services[k] = s
	}
	for key := range d.services {
		d.servicesKey = append(d.servicesKey, key)
	}

}
