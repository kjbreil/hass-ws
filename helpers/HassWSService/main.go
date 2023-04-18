package main

import (
	"flag"
	"github.com/goccy/go-json"
	hass_ws "github.com/kjbreil/hass-ws"
	"github.com/kjbreil/hass-ws/helpers/HassWSService/servicemaker"
)

var (
	configLocation = flag.String("c", "config.yml", "location of hass-ws config")
	packageName    = flag.String("p", "services", "folder and package name for the generated code")
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

	data, _ := json.Marshal(ss)

	servicesList := servicemaker.ServicesInitJson(data)
	servicesFolder := *packageName

	err = servicemaker.Gen(servicesFolder, servicesList)
	if err != nil {
		panic(err)
	}
}
