package main

import (
	"flag"
	"github.com/goccy/go-json"
	hass_ws "github.com/kjbreil/hass-ws"
	"github.com/kjbreil/hass-ws/helpers/HassWSService/servicemaker"
	"os"
)

// Command-line flags for configuration
var (
	configLocation   = flag.String("c", "config.yml", "location of hass-ws config")
	packageName      = flag.String("p", "services", "folder and package name for the generated code")
	makeServicesJson = flag.Bool("j", false, "make the services.json")
	onlyServices     = flag.String("os", "", "single service to generate")
)

// main is the entry point for the service code generator. It connects to Home Assistant using hass-ws,
// retrieves the available services, and generates Go code for them using the servicemaker package.
// Optionally, it can output a services.json file or generate code for a single service.
func main() {
	flag.Parse()

	// Parse the Home Assistant configuration
	cfg, _ := hass_ws.ParseConfig(*configLocation)

	// Create and connect the client
	c, _ := hass_ws.NewClient(cfg)
	err := c.Connect()
	if err != nil {
		panic(err) // Panic if connection fails
	}

	// Retrieve available Home Assistant services
	ss := c.GetServices()

	// If only a single service should be generated, filter the results
	if *onlyServices != "" {
		for serviceName, service := range ss.ServiceResult {
			if serviceName == *onlyServices {
				ss.ServiceResult = map[string]any{serviceName: service}
				break
			}
		}
	}

	// Marshal the services to JSON
	data, _ := json.MarshalIndent(ss, "", "  ")

	// Optionally write the services.json file
	if *makeServicesJson {
		os.WriteFile("services.json", data, os.ModePerm)
	}

	// Generate Go code for the services using servicemaker
	servicesList := servicemaker.ServicesInitJson(data)
	servicesFolder := *packageName

	err = servicemaker.Gen(servicesFolder, servicesList)
	if err != nil {
		panic(err) // Panic if code generation fails
	}
}
