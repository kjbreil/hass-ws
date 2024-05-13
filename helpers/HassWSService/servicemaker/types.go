package servicemaker

import (
	"fmt"
	"github.com/dave/jennifer/jen"
)

func makeTypesFile(enumKeys []string, services map[string]*jen.File) {
	// // Make the types.go file
	// make the generate comment for stringer

	for _, k := range enumKeys {
		services["types"].HeaderComment(fmt.Sprintf("//go:generate stringer -type=%s -trimprefix=%s", k, k))
	}
	for _, k := range enumKeys {
		services["types"].HeaderComment(fmt.Sprintf("//go:generate jsonenums -type=%s", k))
	}

	// Generate the interface
	services["types"].Type().Id("Service").Interface(
		jen.Id("SetID").Params(jen.Id("id").Op("*").Int()),
		jen.Id("JSON").Params().String(),
		jen.Id("Name").Params().String(),
		jen.Id("Targets").Params().Index().String(),
		jen.Id("SetReturnResponse").Params(jen.Id("b").Bool()),
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
			g.Add(jen.Id("ReturnResponse").Bool().Tag(map[string]string{"json": "return_response"}))
		})

	services["types"].Func().Params(jen.Id("s").Op("*").Id("ServiceBase")).Id("SetID").Params(jen.Id("id").Op("*").Int()).Block(
		jen.Id("s").Dot("Id").Op("=").Id("id"),
	)
	services["types"].Func().Params(jen.Id("s").Op("*").Id("ServiceBase")).Id("SetReturnResponse").Params(jen.Id("b").Bool()).Block(
		jen.Id("s").Dot("ReturnResponse").Op("=").Id("b"),
	)

	services["types"].Type().Id("Target").Struct(
		jen.Id("EntityId").Index().String().Tag(map[string]string{"json": "entity_id" + ",omitempty"}),
	)

	services["types"].Func().Id("Targets").Params(jen.Id("entities").Op("...").String()).Id("Target").Block(
		jen.Var().Id("t").Id("Target"),
		jen.For(jen.Id("_").Op(",").Id("e").Op(":=").Range().Id("entities")).Block(
			jen.Id("t").Dot("EntityId").Op("=").Append(jen.Id("t").Dot("EntityId").Op(",").Id("e")),
		),
		jen.Return(jen.Id("t")),
	)
}
