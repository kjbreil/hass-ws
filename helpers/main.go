package main

import (
	"fmt"
	"github.com/dave/jennifer/jen"
	"github.com/iancoleman/strcase"
	"sort"
)

func main() {

	devices := DevicesInit()

	external := make(map[string]*jen.File)

	fileList := []string{
		"types",
	}

	outFolder := "entities"

	for _, v := range append(DeviceNames, fileList...) {
		external[v] = jen.NewFilePathName(fmt.Sprintf("./%s/%s.go", outFolder, v), outFolder)
		external[v].Comment("////////////////////////////////////////////////////////////////////////////////")
		external[v].Comment("Do not modify this file, it is automatically generated")
		external[v].Comment("////////////////////////////////////////////////////////////////////////////////")
	}

	fileList = []string{"types"}

	external["types"].Type().Id("Entity").Interface(
	// jen.UnionFunc(
	// 	func(g *jen.Group) {
	// 		for _, d := range devices {
	// 			g.Add(jen.Id(strcase.ToCamel(d.Name)))
	// 		}
	// 	},
	// ),
	//jen.Id("GetState").Params().String(),
	)

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
			st[v] = append(st[v], jen.Id(strcase.ToCamel(v)).Map(jen.String()).Interface().Tag(map[string]string{"json": v + ",omitempty"}))
		}

		sortedKeys := []string{}
		for key := range st {
			sortedKeys = append(sortedKeys, key)
		}
		sort.Strings(sortedKeys)
		external[d.Name].Type().Id(strcase.ToCamel(d.Name)).StructFunc(
			func(g *jen.Group) {
				for _, key := range sortedKeys {
					v := st[key]
					for _, val := range v {
						g.Add(val)
					}
				}
			},
		)

		external[d.Name].Func().
			Id(fmt.Sprintf("Get%s", strcase.ToCamel(d.Name))).
			Params(jen.Id("attributes").Map(jen.String()).Interface()).
			Op("*").Id(strcase.ToCamel(d.Name)).
			Block(
				jen.Var().Id(string(strcase.ToLowerCamel(d.Name)[0])).Id(strcase.ToCamel(d.Name)),
				jen.Id("FillFields").Params(jen.Op("&").Id(string(strcase.ToLowerCamel(d.Name)[0])), jen.Id("attributes")),
				jen.Id(string(strcase.ToLowerCamel(d.Name)[0])).Dot("Additional").Op("=").Id("attributes"),
				jen.Return(jen.Op("&").Id(string(strcase.ToLowerCamel(d.Name)[0]))),
			)

		//external[d.Name].Func().Params(
		//	jen.Id("d").Op("*").Id(strcase.ToCamel(d.Name)),
		//).Id("GetState").BlockFunc(
		//	func(g *jen.Group) {
		//		g.Add(
		//			jen.Op("*").Id("d").Dot("MQTT").Op("=").Id("fields"),
		//		)
		//	},
		//)
	}

	for k, v := range external {
		v.Save(fmt.Sprintf("./%s/%s.go", outFolder, k))
	}

}
