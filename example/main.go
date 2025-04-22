// Copyright 2015 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	hass_ws "github.com/kjbreil/hass-ws"
	"github.com/kjbreil/hass-ws/helpers/HassWSService/services"
	"github.com/kjbreil/hass-ws/model"
	"github.com/kjbreil/hass-ws/rest"
	"log"
	"net/url"
	"os"
	"os/signal"
	"time"
)

func main() {

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	cfg, _ := hass_ws.ParseConfig("config.yml")

	rc := rest.New(cfg.Host, cfg.Port, cfg.Token, cfg.SSL)

	values := url.Values{}
	values.Set("type", "call_service")

	api, err := rc.GetState("light.light_main_floor_bathroom")
	if err != nil {
		return
	}

	fmt.Println(api)

	c, _ := hass_ws.NewClient(cfg)

	// c.Logger(setupLogging())

	// Subscribe to all events
	c.AddSubscription(model.EventTypeCallService)

	// Setup a handler for sensor updates
	// use message.FriendlyName() rather than newAttr.FriendlyName and need to check for nil
	// c.OnType.OnSensor = func(message *model.Message, newAttrs, oldAttrs *entities.Sensor) {
	//	fmt.Printf("Sensor: %s - %s - %s\n", message.EntityID(), message.FriendlyName(), message.State())
	// }

	// c.OnType.OnLight = func(message *model.Message, newAttrs, oldAttrs *entities.Light) {
	//	c.Logger().Info(fmt.Sprintf("Sensor: %s - %s - %s", message.EntityID(), message.FriendlyName(), message.State()))
	// }

	// // setup handler for single entity updates
	// c.OnEntity["climate.kitchen"] = func(message *model.Message) {
	//	c.Logger().Info(fmt.Sprintf("Sensor: %s - %s", message.FriendlyName(), message.State()))
	//
	// }

	// Set a message handler to run on every message even if hit by other message handlers
	c.OnMessage = func(message model.Message) {
		// some processing code
		if message.Type == model.MessageTypeEvent && message.Event != nil && message.Event.EventType == model.EventTypeCallService {

			fmt.Printf("Message: %s\n", message.Raw)

		}
	}

	// Set a message handler to run when both OnType and OnEntity did not run
	c.OnUnhandled = func(message model.Message) {
		// some processing code
	}
	c.InitStates = true

	err = c.Connect()
	if err != nil {
		log.Panicln(err)
	}

	// service := services.NewClimateSetTemperature(services.Targets("climate.kitchen")).
	// 	HvacMode(services.HvacModeheat_cool).
	// 	TargetTempLow(60).
	// 	TargetTempHigh(78)
	// c.CallService(service)

	service := services.NewWeatherGetForecast(services.Targets("weather.forecast_merritt_ave_se"))
	service.GetForecastType(services.GetForecastTypehourly)

	rsp := c.CallService(service)

	msg := rsp.Timeout(time.Second)
	fmt.Println(msg)

	// service := services.NewCalendarListEvents(services.Targets("calendar.rainier"))
	// c.CallService(service)

	// Get all states, they ar ether run through the OnType and OnEntity handlers but not OnMessage or OnUnhandled
	// c.GetStates()

	// Get all the services, in ServiceResults its a map[string]interface{} and not easy to work with
	// ss := c.GetServices()
	// for sn := range ss.ServiceResult {
	//	fmt.Printf("Service %s returned\n", sn)
	// }

	// areas := c.GetDeviceRegistry()

	// fmt.Println(areas)
	// for sn := range areas.ServiceResult {
	//	fmt.Printf("Service %s returned\n", sn)
	// }

	// _, err = c.GetHistory(time.Now().Add(-24*time.Hour), time.Now(), "climate.kids_room")

	if err != nil {
		log.Panicln(err)
	}
	for {
		select {
		case <-interrupt:

			err := c.Close()
			if err != nil {
				log.Println("write close:", err)
				return
			}
			return
		}
	}
}
