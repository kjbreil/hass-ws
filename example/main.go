// Copyright 2015 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"github.com/gorilla/websocket"
	hass_ws "github.com/kjbreil/hass-ws"
	"github.com/kjbreil/hass-ws/model"
	"github.com/kjbreil/hass-ws/services"
	"log"
	"os"
	"os/signal"
)

func main() {

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)
	c, _ := hass_ws.NewClient(&hass_ws.Config{
		Host:  "192.168.1.12",
		Port:  8123,
		Token: "",
	})

	//c.AddSubscription(model.EventTypeAll)
	//c.OnType.OnClimate = func(message *model.Message, newAttrs, oldAttrs *entities.Climate) {
	//	fmt.Println(message)
	//}
	//
	//c.OnEntity["climate.kitchen"] = func(message *model.Message) {
	//	fmt.Println(message)
	//}
	c.SetOnMessage(func(message model.Message) {
		if message.Event != nil {
			//log.Println(string(message.Raw))
			if !message.Event.EventType.Valid() {
				log.Println(string(message.Raw))
			}
			//if *message.Event.EventType == model.EventTypeStateChanged {
			//	if message.Event != nil && strings.Contains(*message.Event.Data.EntityId, "binary_sensor") {
			//		log.Println(string(message.Raw))
			//	}
			//}
		}
	})

	err := c.Connect()
	if err != nil {
		log.Panicln(err)
	}

	high := 78
	low := 60
	mode := services.HvacModeHeatCool
	c.CallService(services.NewClimateSetTemperature(
		[]string{"climate.kitchen"},
		&mode,
		&high,
		&low,
		nil,
	))

	//c.GetStates()
	//c.GetServices()
	for {
		select {
		case <-interrupt:
			log.Println("interrupt")

			// Cleanly close the connection by sending a close message and then
			// waiting (with timeout) for the server to close the connection.
			err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("write close:", err)
				return
			}
			return
		}
	}
}