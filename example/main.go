// Copyright 2015 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	hass_ws "github.com/kjbreil/hass-ws"
	"github.com/kjbreil/hass-ws/entities"
	"github.com/kjbreil/hass-ws/model"
	"log"
	"os"
	"os/signal"
)

func main() {

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	cfg, _ := hass_ws.ParseConfig("config.yml")

	c, _ := hass_ws.NewClient(cfg)

	//c.Logger(setupLogging())

	// Subscribe to all events
	c.AddSubscription(model.EventTypeAll)

	// Setup a handler for sensor updates
	// use message.FriendlyName() rather than newAttr.FriendlyName and need to check for nil
	//c.OnType.OnSensor = func(message *model.Message, newAttrs, oldAttrs *entities.Sensor) {
	//	fmt.Printf("Sensor: %s - %s - %s\n", message.EntityID(), message.FriendlyName(), message.State())
	//}

	c.OnType.OnLight = func(message *model.Message, newAttrs, oldAttrs *entities.Light) {
		c.Logger().Info(fmt.Sprintf("Sensor: %s - %s - %s", message.EntityID(), message.FriendlyName(), message.State()))
	}

	// setup handler for single entity updates
	c.OnEntity["climate.kitchen"] = func(message *model.Message) {
		c.Logger().Info(fmt.Sprintf("Sensor: %s - %s", message.FriendlyName(), message.State()))

	}

	// Set a message handler to run on every message even if hit by other message handlers
	c.OnMessage = func(message model.Message) {
		// some processing code
	}

	// Set a message handler to run when both OnType and OnEntity did not run
	c.OnUnhandled = func(message model.Message) {
		// some processing code
	}
	c.InitStates = true

	err := c.Connect()
	if err != nil {
		log.Panicln(err)
	}
	//
	//// Call a service
	//high := 78.0
	//low := 60.0
	//mode := services.HvacModeheat_cool
	//c.CallService(services.NewClimateSetTemperature(services.Targets("climate.kitchen"), &services.ClimateSetTemperatureParams{
	//	HvacMode:       &mode,
	//	TargetTempHigh: &high,
	//	TargetTempLow:  &low,
	//	Temperature:    nil,
	//}))
	//
	//// Get all states, they ar ether run through the OnType and OnEntity handlers but not OnMessage or OnUnhandled
	//c.GetStates()
	//
	//// Get all the services, in ServiceResults its a map[string]interface{} and not easy to work with
	//ss := c.GetServices()
	//for sn := range ss.ServiceResult {
	//	fmt.Printf("Service %s returned\n", sn)
	//}

	//areas := c.GetDeviceRegistry()

	//fmt.Println(areas)
	//for sn := range areas.ServiceResult {
	//	fmt.Printf("Service %s returned\n", sn)
	//}

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
