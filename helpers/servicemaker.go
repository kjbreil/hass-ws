package main

import (
	"fmt"
	"github.com/dave/jennifer/jen"
	"github.com/iancoleman/strcase"
	"sort"
)

func GenServices() error {
	servicesFolder := "services"

	servicesList := ServicesInit()

	services := make(map[string]*jen.File)
	fileList := []string{
		"types",
	}

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

			// Creation function
			services[d.name].Func().
				Id(fmt.Sprintf("New%s", s.camelName)).
				Params(
					jen.Id("target").Id("Target"),
					jen.Id(fmt.Sprintf("%sParams", s.lowerCamelName)).Id(fmt.Sprintf("%sParams", s.camelName)),
				).
				Op("*").Id(s.camelName).
				Block(
					jen.Id("serviceDomain").Op(":=").Lit(d.name),
					jen.Id("serviceType").Op(":=").Lit("call_service"),
					jen.Id("serviceService").Op(":=").Lit(k),
					jen.Id(s.firstLetter).Op(":=").Op("&").Id(s.camelName).Values(jen.Dict{
						jen.Id("ServiceBase"): jen.Id("ServiceBase").Values(jen.Dict{
							jen.Id("Id"):      jen.Id("nil"),
							jen.Id("Type"):    jen.Op("&").Id("serviceType"),
							jen.Id("Domain"):  jen.Op("&").Id("serviceDomain"),
							jen.Id("Service"): jen.Op("&").Id("serviceService"),
							jen.Id("Target"):  jen.Id("target"),
						}),
						//.Tag(map[string]string{"json": "id" + ",omitempty"}),
						jen.Id("ServiceData"): jen.Id(fmt.Sprintf("%sParams", s.lowerCamelName)),
					}),
					jen.Return(jen.Id(s.firstLetter)),
				)

			// Service Type
			services[d.name].Type().
				Id(s.camelName).
				Struct(
					jen.Id("ServiceBase"),
					jen.Id("ServiceData").Id(fmt.Sprintf("%sParams", s.camelName)).Tag(map[string]string{"json": "service_data" + ",omitempty"}),
				)

			// Service Type Params
			services[d.name].Type().
				Id(fmt.Sprintf("%sParams", s.camelName)).
				StructFunc(func(g *jen.Group) {
					for _, fn := range s.parameterKeys {
						code := s.parameters[fn]
						g.Add(jen.Id(strcase.ToCamel(fn)).Op("*").Add(code).Tag(map[string]string{"json": fn + ",omitempty"}))
					}
				})

			services[d.name].Func().Params(jen.Id(s.firstLetter).Op("*").Id(s.camelName)).Id("JSON").Params().String().Block(
				jen.List(jen.Id("data"), jen.Id("_")).Op(":=").Qual("encoding/json", "Marshal").Params(jen.Id(s.firstLetter)),
				jen.Return(jen.String().Params(jen.Id("data"))),
			)

			services[d.name].Func().Params(jen.Id(s.firstLetter).Op("*").Id(s.camelName)).Id("SetID").Params(jen.Id("id").Op("*").Int()).Block(
				jen.Id(s.firstLetter).Dot("Id").Op("=").Id("id"),
			)

		}

	}
	//// Make the types.go file
	// Generate the interface
	services["types"].Type().Id("Service").Interface(
		jen.Id("SetID").Params(jen.Id("id").Op("*").Int()),
		jen.Id("JSON").Params().String(),
	)

	// Generate the base type
	services["types"].Type().
		Id("ServiceBase").
		StructFunc(func(g *jen.Group) {
			g.Add(jen.Id("Id").Op("*").Int().Tag(map[string]string{"json": "id"}))
			g.Add(jen.Id("Type").Op("*").String().Tag(map[string]string{"json": "type"}))
			g.Add(jen.Id("Domain").Op("*").String().Tag(map[string]string{"json": "domain"}))
			g.Add(jen.Id("Service").Op("*").String().Tag(map[string]string{"json": "service"}))
			g.Add(jen.Id("Target").Id("Target").Tag(map[string]string{"json": "target" + ",omitempty"}))
		})

	services["types"].Type().Id("Target").Struct(
		jen.Id("EntityId").Index().String().Tag(map[string]string{"json": "entity_id" + ",omitempty"}),
	)

	//func EntitiesTarget(entities ...string) Target {
	//	var t Target
	//	for _, e := range entities {
	//	t.EntityId = append(t.EntityId, e)
	//}
	//	return t
	//}

	services["types"].Func().Id("Targets").Params(jen.Id("entities").Op("...").String()).Id("Target").Block(
		jen.Var().Id("t").Id("Target"),
		jen.For(jen.Id("_").Op(",").Id("e").Op(":=").Range().Id("entities")).Block(
			jen.Id("t").Dot("EntityId").Op("=").Append(jen.Id("t").Dot("EntityId").Op(",").Id("e")),
		),
		jen.Return(jen.Id("t")),
	)

	var enumKeys []string
	for k := range servicesList.enums {
		enumKeys = append(enumKeys, k)
	}
	sort.Strings(enumKeys)

	for _, k := range enumKeys {
		v := servicesList.enums[k]
		typeName := strcase.ToCamel(k)
		services["types"].Type().Id(typeName).String()
		services["types"].Const().DefsFunc(func(g *jen.Group) {

			var keysKey []string
			for kk := range v {
				keysKey = append(keysKey, kk)
			}
			sort.Strings(keysKey)

			for _, o := range keysKey {
				varName := strcase.ToCamel(o)

				g.Add(jen.Id(fmt.Sprintf("%s%s", typeName, varName)).Id(typeName).Op("=").Lit(o))
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
