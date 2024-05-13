package main

import (
	"flag"
	"github.com/goccy/go-json"
	hass_ws "github.com/kjbreil/hass-ws"
	"github.com/kjbreil/hass-ws/helpers/HassWSService/servicemaker"
	"os"
)

var (
	configLocation   = flag.String("c", "config.yml", "location of hass-ws config")
	packageName      = flag.String("p", "services", "folder and package name for the generated code")
	makeServicesJson = flag.Bool("j", false, "make the services.json")
	onlyServices     = flag.String("os", "", "single service to generate")
)

func main() {

	flag.Parse()

	cfg, _ := hass_ws.ParseConfig(*configLocation)

	c, _ := hass_ws.NewClient(cfg)
	err := c.Connect()
	if err != nil {
		panic(err)
	}

	ss := c.GetServices()

	if *onlyServices != "" {
		for serviceName, service := range ss.ServiceResult {
			if serviceName == *onlyServices {
				ss.ServiceResult = map[string]any{serviceName: service}
				break
			}
		}
	}

	data, _ := json.MarshalIndent(ss, "", "  ")

	if *makeServicesJson {
		os.WriteFile("services.json", data, os.ModePerm)
	}

	servicesList := servicemaker.ServicesInitJson(data)
	servicesFolder := *packageName

	err = servicemaker.Gen(servicesFolder, servicesList)
	if err != nil {
		panic(err)
	}
}
