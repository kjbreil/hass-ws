package main

import (
	"encoding/json"
	hass_ws "github.com/kjbreil/hass-ws"
	"github.com/kjbreil/hass-ws/helpers/HassWSService/servicemaker"
)

func main() {

	cfg, _ := hass_ws.ParseConfig("config.yml")

	c, _ := hass_ws.NewClient(cfg)
	err := c.Connect()
	if err != nil {
		panic(err)
	}

	ss := c.GetServices()

	data, _ := json.Marshal(ss)

	servicesList := servicemaker.ServicesInitJson(data)
	servicesFolder := "services"

	err = servicemaker.Gen(servicesFolder, servicesList)
	if err != nil {
		panic(err)
	}
}
