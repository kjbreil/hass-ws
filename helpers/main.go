package main

import (
	"fmt"
	"github.com/Jeffail/gabs/v2"
	"github.com/dave/jennifer/jen"
	"github.com/iancoleman/strcase"
	"log"
	"sort"
	"strings"
)

func main() {

	devices := DevicesInit()

	servicesJson, _ := gabs.ParseJSONFile("./helpers/services.json")

	entities := make(map[string]*jen.File)
	services := make(map[string]*jen.File)
	models := make(map[string]*jen.File)

	fileList := []string{
		"types",
	}

	modelsFileList := []string{
		"entity",
	}

	entitiesFolder := "entities"
	servicesFolder := "services"
	modelFolder := "model"

	for _, v := range append(DeviceNames, fileList...) {
		entities[v] = jen.NewFilePathName(fmt.Sprintf("./%s/%s.go", entitiesFolder, v), entitiesFolder)
		entities[v].Comment("////////////////////////////////////////////////////////////////////////////////")
		entities[v].Comment("Do not modify this file, it is automatically generated")
		entities[v].Comment("////////////////////////////////////////////////////////////////////////////////")
	}
	for _, v := range append(DeviceNames, fileList...) {
		services[v] = jen.NewFilePathName(fmt.Sprintf("./%s/%s.go", entitiesFolder, v), servicesFolder)
		services[v].Comment("////////////////////////////////////////////////////////////////////////////////")
		services[v].Comment("Do not modify this file, it is automatically generated")
		services[v].Comment("////////////////////////////////////////////////////////////////////////////////")
	}

	for _, v := range modelsFileList {
		models[v] = jen.NewFilePathName(fmt.Sprintf("./%s/%s.go", modelFolder, v), modelFolder)
		models[v].Comment("////////////////////////////////////////////////////////////////////////////////")
		models[v].Comment("Do not modify this file, it is automatically generated")
		models[v].Comment("////////////////////////////////////////////////////////////////////////////////")
	}

	fileList = []string{"types"}

	entities["types"].Type().Id("Entity").Interface()
	entities["types"].Type().Id("Additional").Map(jen.String()).Interface()

	for _, d := range devices {

		st := make(map[string][]*jen.Statement)
		for _, key := range d.Attributes {
			if key.DataType == "" {
				continue
			}
			st[key.Name] = append(st[key.Name], FieldAdder(key.Name, key.DataType))
		}

		for _, v := range []string{
			"friendly_name",
		} {
			st[v] = append(st[v], jen.Id(strcase.ToCamel(v)).Op("*").String().Tag(map[string]string{"json": v + ",omitempty"}))
		}

		for _, v := range []string{
			"additional",
		} {
			st[v] = append(st[v], jen.Id(strcase.ToCamel(v)).Id("Additional").Tag(map[string]string{"json": v + ",omitempty"}))
		}

		sortedKeys := []string{}
		for key := range st {
			sortedKeys = append(sortedKeys, key)
		}
		sort.Strings(sortedKeys)
		entities[d.Name].Type().Id(strcase.ToCamel(d.Name)).StructFunc(
			func(g *jen.Group) {
				for _, key := range sortedKeys {
					v := st[key]
					for _, val := range v {
						g.Add(val)
					}
				}
			},
		)

		entities[d.Name].Func().
			Id(fmt.Sprintf("Get%s", strcase.ToCamel(d.Name))).
			Params(jen.Id("attributes").Map(jen.String()).Interface()).
			Op("*").Id(strcase.ToCamel(d.Name)).
			Block(
				jen.Var().Id(string(strcase.ToLowerCamel(d.Name)[0])).Id(strcase.ToCamel(d.Name)),
				jen.Id("fillFields").Params(jen.Op("&").Id(string(strcase.ToLowerCamel(d.Name)[0])), jen.Id("attributes")),
				jen.Id(string(strcase.ToLowerCamel(d.Name)[0])).Dot("Additional").Op("=").Id("attributes"),
				jen.Return(jen.Op("&").Id(string(strcase.ToLowerCamel(d.Name)[0]))),
			)

		// Services

		if d.Name == "climate" {
			services[d.Name].Type().Id(strcase.ToCamel(d.Name)).Struct(
				jen.Id("EntityId").String(),
			)

			enumMap := make(map[string]map[string]struct{})

			//services[d.Name].Func().Id(strcase.ToCamel(d.Name))

			camelName := strcase.ToCamel(d.Name)
			firstLetter := string(camelName[0])

			deviceServices := servicesJson.Search("service_result", "climate")
			for k, v := range deviceServices.ChildrenMap() {
				services[d.Name].Func().Params(jen.Id(firstLetter).Op("*").Id(camelName)).
					Id(fmt.Sprintf(strcase.ToCamel(k))).
					ParamsFunc(func(g *jen.Group) {
						for fn, f := range v.Path("fields").ChildrenMap() {
							selector := f.Path("selector")
							if selector != nil {
								sm := selector.ChildrenMap()

								if _, ok := sm["text"]; ok {
									g.Add(jen.Id(fn).String())
								}

								if _, ok := sm["number"]; ok {
									g.Add(jen.Id(fn).Int())
								}

								if s, ok := sm["select"]; ok {
									options := s.ChildrenMap()["options"].Children()
									if options != nil {

										for _, o := range options {
											fmt.Println(o.Path("value").String())
											if _, ok := enumMap[strcase.ToCamel(fn)]; !ok {
												enumMap[strcase.ToCamel(fn)] = make(map[string]struct{})
											}
											enumMap[strcase.ToCamel(fn)][strings.Trim(o.Path("value").String(), "\"")] = struct{}{}
										}
										g.Add(jen.Id(strcase.ToLowerCamel(fn)).Id(strcase.ToCamel(fn)))
									}

									fmt.Println(options)
									//g.Add(jen.Id(fn).String())
								}

								fmt.Println(fn, f)
							}
						}
					}).
					Op("*").Id(strcase.ToCamel(d.Name)).
					Block(
						jen.Var().Id(string(strcase.ToLowerCamel(d.Name)[0])).Id(strcase.ToCamel(d.Name)),
						jen.Return(jen.Op("&").Id(string(strcase.ToLowerCamel(d.Name)[0]))),
					)
			}

			for k, v := range enumMap {
				typeName := strcase.ToCamel(k)
				services[d.Name].Type().Id(typeName).String()
				services[d.Name].Const().DefsFunc(func(g *jen.Group) {

					for o := range v {
						varName := strcase.ToCamel(o)

						g.Add(jen.Id(fmt.Sprintf("%s%s", typeName, varName)).Id(typeName).Op("=").Lit(o))
					}
				})
			}
		}

	}

	for _, v := range modelsFileList {
		sortedKeys := []string{}
		for _, key := range devices {
			sortedKeys = append(sortedKeys, key.Name)
		}
		sort.Strings(sortedKeys)

		for i := 0; i < len(sortedKeys); i++ {
			camelName := strcase.ToCamel(sortedKeys[i])
			models[v].Type().Id(fmt.Sprintf("On%sHandler", camelName)).Func().Params(jen.Id("message").Op("*").Id("Message"), jen.Id("newAttrs"), jen.Id("oldAttrs").Op("*").Qual("github.com/kjbreil/hass-ws/entities", camelName))
		}

		models[v].Type().Id("OnTypeHandlers").StructFunc(func(g *jen.Group) {
			for i := 0; i < len(sortedKeys); i++ {
				camelName := strcase.ToCamel(sortedKeys[i])
				g.Add(jen.Id(fmt.Sprintf("On%s", camelName)).Id(fmt.Sprintf("On%sHandler", camelName)))
			}
		})

		models[v].Func().Params(
			jen.Id("o").Op("*").Id("OnTypeHandlers"),
		).Id("Run").Params(
			jen.Id("message").Op("*").Id("Message"),
		).Block(
			jen.If(jen.Id("message").Dot("Event").Op("==").Id("nil").Op("||").
				Id("message").Dot("Event").Dot("Data").Op("==").Id("nil").Op("||").
				Id("message").Dot("Event").Dot("Data").Dot("EntityId").Op("==").Id("nil")).
				Block(jen.Return()),
			jen.Id("entityType").Op(":=").Qual("strings", "Split").Params(
				jen.Op("*").Id("message").Dot("Event").Dot("Data").Dot("EntityId"),
				jen.Lit("."),
			).Id("[0]"),

			jen.Switch(jen.Id("entityType")).BlockFunc(func(g *jen.Group) {
				for i := 0; i < len(sortedKeys); i++ {
					camelName := strcase.ToCamel(sortedKeys[i])
					underName := strcase.ToSnake(sortedKeys[i])
					g.Add(jen.Case(jen.Lit(underName)).Block(
						jen.If(jen.Id("o").Dot(fmt.Sprintf("On%s", camelName)).Op("==").Id("nil")).Block(jen.Return()),
						jen.Id("newAttrs").Op(":=").Qual("github.com/kjbreil/hass-ws/entities", fmt.Sprintf("Get%s", camelName)).Params(jen.Id("message").Dot("Event").Dot("Data").Dot("NewState").Dot("Attributes")),
						jen.Id("oldAttrs").Op(":=").Qual("github.com/kjbreil/hass-ws/entities", fmt.Sprintf("Get%s", camelName)).Params(jen.Id("message").Dot("Event").Dot("Data").Dot("OldState").Dot("Attributes")),
						jen.Id("o").Dot(fmt.Sprintf("On%s", camelName)).Params(jen.Id("message"), jen.Id("newAttrs"), jen.Id("oldAttrs")),
					))
					// .Params(jen.Id("message").Id("newAttrs").Id("oldAttrs"))
				}
			}),
		)

		models[v].Func().Params(
			jen.Id("o").Op("*").Id("OnTypeHandlers"),
		).Id("RunStates").Params(
			jen.Id("message").Op("*").Id("Message"),
		).Block(
			jen.If(jen.Id("message").Dot("Result").Op("==").Id("nil")).
				Block(jen.Return()),
			jen.For(jen.Id("_").Op(",").Id("r").Op(":=").Range().Id("message").Dot("Result")).Block(

				jen.Id("entityType").Op(":=").Qual("strings", "Split").Params(
					jen.Op("*").Id("r").Dot("EntityId"),
					jen.Lit("."),
				).Id("[0]"),
				jen.Switch(jen.Id("entityType")).BlockFunc(func(g *jen.Group) {
					for i := 0; i < len(sortedKeys); i++ {
						camelName := strcase.ToCamel(sortedKeys[i])
						underName := strcase.ToSnake(sortedKeys[i])
						g.Add(jen.Case(jen.Lit(underName)).Block(
							jen.If(jen.Id("o").Dot(fmt.Sprintf("On%s", camelName)).Op("==").Id("nil")).Block(jen.Continue()),
							jen.Id("newAttrs").Op(":=").Qual("github.com/kjbreil/hass-ws/entities", fmt.Sprintf("Get%s", camelName)).Params(jen.Id("r").Dot("Attributes")),
							jen.Id("newMsg").Op(":=").Op("&").Id("Message").Values(jen.Dict{

								jen.Id("ID"):   jen.Id("message").Dot("ID"),
								jen.Id("Type"): jen.Id("MessageTypeEvent"),
								jen.Id("Event"): jen.Op("&").Id("Event").Values(
									jen.Dict{
										jen.Id("Data"): jen.Op("&").Id("Data").Values(jen.Dict{
											jen.Id("EntityId"): jen.Id("r").Dot("EntityId"),
											jen.Id("NewState"): jen.Op("&").Id("State").Values(
												jen.Dict{
													jen.Id("EntityId"):    jen.Id("r").Dot("EntityId"),
													jen.Id("LastChanged"): jen.Id("r").Dot("LastChanged"),
													jen.Id("State"):       jen.Id("r").Dot("State"),
													jen.Id("Attributes"):  jen.Id("r").Dot("Attributes"),
													jen.Id("LastUpdated"): jen.Id("r").Dot("LastUpdated"),
													jen.Id("Context"):     jen.Id("r").Dot("Context"),
												},
											),
											jen.Id("OldState"): jen.Id("nil"),
										}),
										jen.Id("Context"):   jen.Id("r").Dot("Context"),
										jen.Id("EventType"): jen.Id("EventTypeState"),
									},
								),
							},
							),

							jen.Id("o").Dot(fmt.Sprintf("On%s", camelName)).Params(jen.Id("newMsg"), jen.Id("newAttrs"), jen.Id("nil")),
						))
						// .Params(jen.Id("message").Id("newAttrs").Id("oldAttrs"))
					}
				}),
			),
		)

		models[v].Type().Id("OnEntityHandlers").Map(jen.String()).Func().Params(jen.Id("message").Op("*").Id("Message"))

		models[v].Func().Params(
			jen.Id("o").Id("OnEntityHandlers"),
		).Id("Run").Params(
			jen.Id("message").Op("*").Id("Message"),
		).Block(
			jen.If(jen.Id("message").Dot("Event").Op("==").Id("nil").Op("||").
				Id("message").Dot("Event").Dot("Data").Op("==").Id("nil").Op("||").
				Id("message").Dot("Event").Dot("Data").Dot("EntityId").Op("==").Id("nil")).
				Block(jen.Return()),
			jen.For(jen.Id("k").Op(",").Id("v").Op(":=").Range().Id("o")).Block(
				jen.If(jen.Id("k").Op("==").Op("*").Id("message").Dot("Event").Dot("Data").Dot("EntityId").Block(
					jen.Id("v").Call(jen.Id("message")),
				)),
			),
		)

		models[v].Func().Params(
			jen.Id("o").Id("OnEntityHandlers"),
		).Id("RunStates").Params(
			jen.Id("message").Op("*").Id("Message"),
		).Block(
			jen.If(jen.Id("message").Dot("Result").Op("==").Id("nil")).
				Block(jen.Return()),
			jen.For(jen.Id("_").Op(",").Id("r").Op(":=").Range().Id("message").Dot("Result")).Block(
				jen.For(jen.Id("k").Op(",").Id("v").Op(":=").Range().Id("o")).Block(

					jen.If(jen.Id("k").Op("==").Op("*").Id("r").Dot("EntityId").Block(

						jen.Id("newMsg").Op(":=").Op("&").Id("Message").Values(jen.Dict{

							jen.Id("ID"):   jen.Id("message").Dot("ID"),
							jen.Id("Type"): jen.Id("MessageTypeEvent"),
							jen.Id("Event"): jen.Op("&").Id("Event").Values(
								jen.Dict{
									jen.Id("Data"): jen.Op("&").Id("Data").Values(jen.Dict{
										jen.Id("EntityId"): jen.Id("r").Dot("EntityId"),
										jen.Id("NewState"): jen.Op("&").Id("State").Values(
											jen.Dict{
												jen.Id("EntityId"):    jen.Id("r").Dot("EntityId"),
												jen.Id("LastChanged"): jen.Id("r").Dot("LastChanged"),
												jen.Id("State"):       jen.Id("r").Dot("State"),
												jen.Id("Attributes"):  jen.Id("r").Dot("Attributes"),
												jen.Id("LastUpdated"): jen.Id("r").Dot("LastUpdated"),
												jen.Id("Context"):     jen.Id("r").Dot("Context"),
											},
										),
										jen.Id("OldState"): jen.Id("nil"),
									}),
									jen.Id("Context"):   jen.Id("r").Dot("Context"),
									jen.Id("EventType"): jen.Id("EventTypeState"),
								},
							),
						},
						),

						jen.Id("v").Call(jen.Id("newMsg")),
					)),
				),
			),
		)
	}

	for k, v := range entities {
		v.Save(fmt.Sprintf("./%s/%s.go", entitiesFolder, k))
	}

	for k, v := range services {
		err := v.Save(fmt.Sprintf("./%s/%s.go", servicesFolder, k))
		if err != nil {
			log.Panicln(err)
		}
	}

	for k, v := range models {
		err := v.Save(fmt.Sprintf("./%s/%s.go", modelFolder, k))
		if err != nil {
			log.Panicln(err)
		}
	}

}
