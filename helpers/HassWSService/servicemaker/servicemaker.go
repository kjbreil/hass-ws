package servicemaker

import (
	"fmt"
	"github.com/dave/jennifer/jen"
	"github.com/iancoleman/strcase"
	"os"
	"sort"
	"strings"
)

func GenServices() error {
	servicesFolder := "services"

	servicesList := ServicesInit("./helpers/services.json")
	err := Gen(servicesFolder, servicesList)
	return err
}

func Gen(servicesFolder string, servicesList ServiceList) error {
	err := os.MkdirAll(servicesFolder, os.ModePerm)
	if err != nil {
		return err
	}

	services := make(map[string]*jen.File)
	fileList := []string{
		"types",
	}

	for _, s := range servicesList.serviceNames {
		fileList = append(fileList, fmt.Sprintf("%s_test", s))
	}

	var enumKeys []string
	for k := range servicesList.enums {
		enumKeys = append(enumKeys, k)
		fileList = append(fileList, strings.ToLower(k))
	}
	sort.Strings(enumKeys)
	sort.Strings(fileList)

	// Generate the file headers
	for _, v := range append(servicesList.serviceNames, fileList...) {
		services[v] = jen.NewFilePathName(fmt.Sprintf("./%s/%s.go", servicesFolder, v), servicesFolder)
		services[v].Comment("////////////////////////////////////////////////////////////////////////////////")
		services[v].Comment("Do not modify this file, it is automatically generated")
		services[v].Comment("////////////////////////////////////////////////////////////////////////////////")
		services[v].Line()
	}

	for _, d := range servicesList.services {
		for _, k := range d.servicesKey {
			s := d.services[k]

			services[d.name].Comment(fmt.Sprintf("New%s creates the object that can be sent to Home Assistant for domain %s, service %s", s.camelName, d.name, s.name))
			services[d.name].Comment(fmt.Sprintf("%s ", s.description))

			testFileName := fmt.Sprintf("%s_test", d.name)

			// Creation function
			services[d.name].Func().
				Id(fmt.Sprintf("New%s", s.camelName)).
				ParamsFunc(func(g *jen.Group) {
					g.Add(jen.Id("target").Id("Target"))
					// if len(s.parameterKeys) > 0 {
					//	g.Add(jen.Id(fmt.Sprintf("%sParams", s.lowerCamelName)).Op("*").Id(fmt.Sprintf("%sParams", s.camelName)))
					// }
				}).
				Op("*").Id(s.camelName).
				Block(
					jen.Id("serviceDomain").Op(":=").Lit(d.name),
					jen.Id("serviceType").Op(":=").Lit("call_service"),
					jen.Id("serviceService").Op(":=").Lit(k),
					jen.Id(s.firstLetter).Op(":=").Op("&").Id(s.camelName).Values(
						jen.DictFunc(func(d jen.Dict) {
							d[jen.Id("ServiceBase")] = jen.Id("ServiceBase").Values(jen.Dict{
								jen.Id("Id"):             jen.Id("nil"),
								jen.Id("Type"):           jen.Op("&").Id("serviceType"),
								jen.Id("Domain"):         jen.Op("&").Id("serviceDomain"),
								jen.Id("Service"):        jen.Op("&").Id("serviceService"),
								jen.Id("Target"):         jen.Id("target"),
								jen.Id("ReturnResponse"): jen.Lit(s.responseRequired),
							})
							if len(s.parameters) > 0 {
								// d[jen.Id("ServiceData")] = jen.Op("*").Id(fmt.Sprintf("%sParams", s.lowerCamelName))
								d[jen.Id("ServiceData")] = jen.Id(fmt.Sprintf("%sParams", s.camelName)).Values()
							} else {
								d[jen.Id("ServiceData")] = jen.Id("nil")

							}
						}),
					),
					jen.Return(jen.Id(s.firstLetter)),
				)

			// Service Type
			services[d.name].Type().
				Id(s.camelName).
				StructFunc(func(g *jen.Group) {
					g.Add(jen.Id("ServiceBase"))
					if len(s.parameters) > 0 {
						g.Add(jen.Id("ServiceData").Id(fmt.Sprintf("%sParams", s.camelName)).Tag(map[string]string{"json": "service_data" + ",omitempty"}))
					} else {
						g.Add(jen.Id("ServiceData").Interface().Tag(map[string]string{"json": "service_data" + ",omitempty"}))

					}
				})

			// Service Type Params
			if len(s.parameters) > 0 {

				services[d.name].Type().
					Id(fmt.Sprintf("%sParams", s.camelName)).
					StructFunc(func(g *jen.Group) {
						for _, fn := range s.parameterKeys {
							code := s.parameters[fn]
							jsonTag := fn
							if strings.HasPrefix(fn, s.name) {
								jsonTag = strings.Replace(jsonTag, s.name+"_", "", 1)
							}
							if code == "interface{}" {
								g.Add(jen.Id(strcase.ToCamel(fn)).Add(codeGetter(code)).Tag(map[string]string{"json": jsonTag + ",omitempty"}))
							} else {
								g.Add(jen.Id(strcase.ToCamel(fn)).Op("*").Add(codeGetter(code)).Tag(map[string]string{"json": jsonTag + ",omitempty"}))
							}
						}
					})
			}

			if len(s.parameterKeys) > 0 {
				for _, fn := range s.parameterKeys {

					// firstLetter := string(fn[0])
					code := s.parameters[fn]

					finalCamelName := strcase.ToCamel(fn)
					switch fn {
					case "name", "domain":
						fn = s.camelName + strcase.ToCamel(fn)
					}
					camelName := strcase.ToCamel(fn)
					lowerCamelName := strcase.ToLowerCamel(fn)
					services[d.name].Func().
						Params(jen.Id(s.firstLetter).Op("*").Id(s.camelName)).
						Id(fmt.Sprintf("%s", camelName)).
						Params(jen.Id(lowerCamelName).Add(codeGetter(code))).
						Op("*").Id(s.camelName).
						Block(
							jen.Id(s.firstLetter).Dot("ServiceData").Dot(finalCamelName).Op("=").Op("&").Id(lowerCamelName),
							jen.Return(jen.Id(s.firstLetter)),
						)
				}
			}

			services[d.name].Func().Params(jen.Id(s.firstLetter).Op("*").Id(s.camelName)).Id("JSON").Params().String().Block(
				jen.List(jen.Id("data"), jen.Id("_")).Op(":=").Qual("github.com/goccy/go-json", "Marshal").Params(jen.Id(s.firstLetter)),
				jen.Return(jen.String().Params(jen.Id("data"))),
			)

			services[d.name].Func().Params(jen.Id(s.firstLetter).Op("*").Id(s.camelName)).Id("Targets").Params().Index().String().Block(
				jen.Return(jen.Id(s.firstLetter).Dot("Target").Dot("EntityId")),
			)

			services[d.name].Func().Params(jen.Id(s.firstLetter).Op("*").Id(s.camelName)).Id("Name").Params().String().Block(
				jen.Return(jen.Qual("fmt", "Sprintf").Params(jen.Lit("%s.%s").Op(",").Op("*").Id(s.firstLetter).Dot("Domain").Op(",").Op("*").Id(s.firstLetter).Dot("Service"))),
			)

			// Tests
			services[testFileName].Func().Id(fmt.Sprintf("Test%s_JSON", s.camelName)).Params(jen.Id("t").Op("*").Qual("testing", "T")).BlockFunc(func(g *jen.Group) {
				for _, fn := range s.parameterKeys {

					code := s.parameters[fn]
					switch fn {
					case "name", "domain":
						fn = s.camelName + strcase.ToCamel(fn)
					}
					fn = strcase.ToLowerCamel(fn)
					g.Add(servicesList.codeAssigner(fn, code))

				}
				g.Add(jen.Line())

				g.Add(
					jen.Id("tests").Op(":=").Index().Struct(
						jen.Id("name").String(),
						jen.Id("fields").Op("*").Id(s.camelName),
						jen.Id("want").String(),
					).ValuesFunc(func(g *jen.Group) {
						add := jen.Add()
						if len(s.parameters) > 0 {
							for _, fn := range s.parameterKeys {

								switch fn {
								case "name", "domain":
									fn = s.camelName + strcase.ToCamel(fn)
								}
								camelName := strcase.ToCamel(fn)
								fn = strcase.ToLowerCamel(fn)
								add.Add(jen.Dot(camelName).Params(jen.Id(fn)))
							}
						}

						g.Add(jen.Values(jen.Dict{
							jen.Id("name"): jen.Lit("base"),
							jen.Id("fields"): jen.Id(fmt.Sprintf("New%s", s.camelName)).
								CallFunc(func(g *jen.Group) {
									g.Add(jen.Id("Targets").Call(jen.Lit("climate.kitchen")))
									// if len(s.parameters) > 0 {
									//	g.Add(jen.Op("&").Id(fmt.Sprintf("%sParams", s.camelName)).Values(jen.DictFunc(func(d jen.Dict) {
									//		for _, fn := range s.parameterKeys {
									//			d[jen.Id(strcase.ToCamel(fn))] = jen.Op("&").Id(strcase.ToLowerCamel(fn))
									//
									//		}
									//
									//	})))
									// }
								}).Add(add),

							jen.Id("want"): jen.LitFunc(func() interface{} {

								var serviceData []string

								for _, fn := range s.parameterKeys {
									code := s.parameters[fn]
									fn = strcase.ToSnake(fn)
									left := strcase.ToSnake(code)

									switch code {
									case "string":
										serviceData = append(serviceData, fmt.Sprintf(`"%s":"%s"`, fn, "data"))
									case "[]byte":
										serviceData = append(serviceData, fmt.Sprintf(`"%s":"%s"`, fn, "data"))
									case "interface{}":
										serviceData = append(serviceData, fmt.Sprintf(`"%s":"%s"`, fn, "data"))
									case "float64":
										serviceData = append(serviceData, fmt.Sprintf(`"%s":%.1f`, fn, 1.2))
									default:

										v := servicesList.enums[code]

										var keysKey []string
										for kk := range v {
											keysKey = append(keysKey, kk)
										}
										sort.Strings(keysKey)

										right := keysKey[0]

										serviceData = append(serviceData, fmt.Sprintf(`"%s":"%s"`, left, right))

									}

								}
								finalServiceData := ""
								if len(s.parameters) > 0 {
									finalServiceData = fmt.Sprintf(`,"service_data":{%s}`, strings.Join(serviceData, ","))
								}
								return fmt.Sprintf(`{"id":null,"type":"call_service","domain":"%s","service":"%s","target":{"entity_id":["climate.kitchen"]}%s}`, d.name, k, finalServiceData)

							},
							// fmt.Sprintf("{\"id\":null,\"type\":\"call_service\",\"domain\":\"%s\",\"service\":\"%s\",\"target\":{\"entity_id\":[\"climate.kitchen\"]},\"service_data\":{}}", d, k),
							),
						}))
					}))
				g.Add(
					jen.For(jen.Id("_").Op(",").Id("tt").Op(":=").Range().Id("tests")).Block(
						jen.Id("t").Dot("Run").Call(jen.Id("tt").Dot("name").Op(",").Func().Params(jen.Id("t").Op("*").Qual("testing", "T")).Block(

							jen.Id("d").Op(":=").Id("tt").Dot("fields"),
							jen.If(
								jen.Id("got").Op(":=").Id("d").Dot("JSON").Call(),
								jen.Id("got").Op("!=").Id("tt").Dot("want"),
							).Block(
								jen.Id("t").Dot("Errorf").Call(jen.Lit("JSON() = %v, want %v"), jen.Id("got"), jen.Id("tt").Dot("want"))),
						)),
					))
			})

		}
		// Tests
		// func TestDeviceTrackerSee_JSON(t *testing.T) {
		//	tests := []struct {
		//		name   string
		//		fields *DeviceTrackerSee
		//		want   string
		//	}{
		//		{
		//			name: "base",
		//			fields: NewDeviceTrackerSee(
		//				Targets("climate.kitchen"),
		//				&DeviceTrackerSeeParams{
		//					Battery:      nil,
		//					DevId:        nil,
		//					GpsAccuracy:  nil,
		//					HostName:     nil,
		//					LocationName: nil,
		//					Mac:          nil,
		//				},
		//			),
		//			want: "{\"id\":null,\"type\":\"call_service\",\"domain\":\"device_tracker\",\"service\":\"see\",\"target\":{\"entity_id\":[\"climate.kitchen\"]},\"service_data\":{}}",
		//		},
		//	}
		//	for _, tt := range tests {
		//		t.Run(tt.name, func(t *testing.T) {
		//			d := tt.fields
		//			if got := d.JSON(); got != tt.want {
		//				t.Errorf("JSON() = %v, want %v", got, tt.want)
		//			}
		//		})
		//	}
		// }

	}
	makeTypesFile(enumKeys, services)

	for _, k := range enumKeys {
		v := servicesList.enums[k]
		typeName := strcase.ToCamel(k)
		services[strings.ToLower(k)].Type().Id(typeName).Int()
		services[strings.ToLower(k)].Const().DefsFunc(func(g *jen.Group) {

			var keysKey []string
			for kk := range v {
				keysKey = append(keysKey, kk)
			}
			sort.Strings(keysKey)

			for i, o := range keysKey {

				// varName := strcase.ToCamel(o)

				if i == 0 {
					g.Add(jen.Id(fmt.Sprintf("%s%s", typeName, o)).Id(typeName).Op("=").Id("iota"))
				} else {
					g.Add(jen.Id(fmt.Sprintf("%s%s", typeName, o)))
				}
			}
		})
	}

	for k, v := range services {
		err := v.Save(fmt.Sprintf("./%s/%s.go", servicesFolder, k))
		if err != nil {
			return err
		}
	}

	return nil
}

func codeGetter(s string) jen.Code {
	switch s {
	case "string":
		return jen.String()
	case "float64":
		return jen.Float64()
	default:
		return jen.Id(s)
	}
}

func (sl *ServiceList) codeAssigner(fn string, code string) jen.Code {
	fn = strcase.ToLowerCamel(fn)
	switch code {
	case "string":
		return jen.Id(fn).Op(":=").Lit("data")
	case "float64":
		return jen.Id(fn).Op(":=").Lit(1.2)
	case "[]byte":
		return jen.Id(fn).Op(":=").Lit("data")
	case "interface{}":
		return jen.Id(fn).Op(":=").Lit("data")
	default:

		v := sl.enums[code]
		typeName := strcase.ToCamel(code)
		var keysKey []string
		for kk := range v {
			keysKey = append(keysKey, kk)
		}
		sort.Strings(keysKey)
		firstEnum := keysKey[0]

		return jen.Id(fn).Op(":=").Id(fmt.Sprintf("%s%s", typeName, firstEnum))
	}
}
