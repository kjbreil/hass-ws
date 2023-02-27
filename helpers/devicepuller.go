package main

import (
	"errors"
	"github.com/dave/jennifer/jen"
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/ast"
	"github.com/iancoleman/strcase"
	"io"
	"io/fs"
	"log"
	"net/http"
	"os"
)

var DeviceNames = []string{
	"introduction",
	"air-quality",
	"alarm-control-panel",
	"binary-sensor",
	"button",
	"camera",
	"climate",
	"cover",
	"device-tracker",
	"fan",
	"humidifier",
	"light",
	"lock",
	"media-player",
	"number",
	"remote",
	"select",
	"sensor",
	"siren",
	"switch",
	"text",
	"update",
	"vacuum",
	"water-heater",
	"weather",
}

type Device struct {
	Name       string
	Attributes map[string]attributes
}

var PullNew = false

func DevicesInit() (retval []Device) {
	for _, name := range DeviceNames {
		d := Device{
			Name: name,
		}
		d.init()
		retval = append(retval, d)
	}
	return retval
}

func (dev *Device) init() {
	dev.Attributes, _ = splitDocument(dev.Name)
}

type attributes struct {
	Name     string
	DataType string
	Required bool
}

func splitDocument(devicename string) (map[string]attributes, error) {

	data, err := fetchDocument(devicename)
	if err != nil {
		return nil, err
	}
	attrs := make(map[string]attributes)
	tags := markdown.Parse(data, nil)
	for _, c := range tags.GetChildren() {
		switch c.(type) {
		case *ast.Table:
			var header []string
			var body [][]string
			for _, t := range c.GetChildren() {
				switch t.(type) {
				case *ast.TableHeader:
					for _, r := range t.GetChildren() {
						switch r.(type) {
						case *ast.TableRow:
							for _, ce := range r.GetChildren() {
								switch ce.(type) {
								case *ast.TableCell:
									for _, te := range ce.GetChildren() {
										switch te.(type) {
										case *ast.Text:
											l := te.AsLeaf()
											header = append(header, string(l.Literal))
										}
									}
								}
							}
						}
					}
				case *ast.TableBody:
					for _, r := range t.GetChildren() {
						switch r.(type) {
						case *ast.TableRow:
							var line []string
							for _, ce := range r.GetChildren() {
								switch ce.(type) {
								case *ast.TableCell:
									for _, te := range ce.GetChildren() {
										switch te.(type) {
										case *ast.Text:
											l := te.AsLeaf()
											line = append(line, string(l.Literal))
										}
									}
								}
							}
							body = append(body, line)

						}
					}
				}
			}

			if len(header) == 4 && header[0] == "Name" && header[1] == "Type" && header[2] == "Default" && header[3] == "Description" {
				for _, b := range body {
					atr := attributes{
						Name:     b[0],
						DataType: b[1],
					}
					if b[2] == "**Required**" {
						atr.Required = true
					}
					attrs[atr.Name] = atr
				}
			}
		}
	}

	return attrs, nil

}

func exists(path string) bool {
	_, err := os.Stat(path)
	return !errors.Is(err, os.ErrNotExist)
}

func fetchDocument(devicename string) ([]byte, error) {

	url := "https://raw.githubusercontent.com/home-assistant/developers.home-assistant/master/docs/core/entity/" + devicename + ".md"

	targetFile := "./helpers/cache/" + devicename + ".md"

	if exists(targetFile) && PullNew == false {

		data, err := os.ReadFile(targetFile)
		if err == nil {
			return data, nil
		}

		log.Println("Loading " + targetFile + " failed")

	}

	log.Println("Fetching " + devicename)

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = os.WriteFile(targetFile, bodyBytes, fs.FileMode(0644))
	if err != nil {
		return nil, err
	}

	return bodyBytes, nil

}

func Unquote(s string) string {
	if len(s) > 0 && s[0] == '"' {
		s = s[1:]
	}
	if len(s) > 0 && s[len(s)-1] == '"' {
		s = s[:len(s)-1]
	}
	return s
}
func FieldAdder(key string, t string) *jen.Statement {

	return TypeTranslator(t, jen.Id(strcase.ToCamel(
		func(key string) string {
			var s string
			if key == "topic" {
				s = "state_topic"
			} else {
				s = key
			}
			return s
		}(key),
	))).Tag(map[string]string{"json": key + ",omitempty"})
}

func TypeTranslator(t string, s *jen.Statement) *jen.Statement {

	v := Unquote(t)
	switch v {
	case "string", "String", "template", "icon", "device_class", "str", "SourceType", "float or string":
		return s.Op("*").String()
	case "integer", "int", "int (bitwise)", "bool, int":
		return s.Op("*").Int()
	case "boolean", "bool":
		return s.Op("*").Bool()
	case "list", "[\"list\"]", "[\"string\",\"list\"]", "set", "HVACMode", "list or dict", "array":
		return s.Op("*").Index().String()
	case "tuple":
		return s.Op("*").Index().Float64()
	case "map":
		return s.Op("*").Map(jen.String()).String()
	case "float":
		return s.Op("*").Float64()
	default:
		log.Fatalln("Unknown type " + t)
		return nil
	}
}
