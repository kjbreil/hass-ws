// Copyright 2015 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"github.com/gorilla/websocket"
	hass_ws "github.com/kjbreil/hass-ws"
	"github.com/kjbreil/hass-ws/model"
	"log"
	"os"
	"os/signal"
	"strings"
)

func main() {

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)
	c, _ := hass_ws.NewClient(&hass_ws.Config{
		Host:  "192.168.1.12",
		Port:  8123,
		Token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJiNWU0OGM1Mjk1MjI0ZDMzOTgwZTczZWRlNDlhN2VhZSIsImlhdCI6MTY3NzUyMjU4OCwiZXhwIjoxOTkyODgyNTg4fQ.gQyMv28xc8IlZJURTqbkWTaEbM4risgg5wMx_x9c8iQ",
	})

	c.AddSubscription(model.EventTypeAll)
	c.SetOnMessage(func(message model.Message) {
		if message.Event != nil && message.Event.EventType != nil {
			if !message.Event.EventType.Valid() {
				log.Println(string(message.Raw))
			}
			if *message.Event.EventType == model.EventTypeStateChanged {
				if message.Event != nil && strings.Contains(*message.Event.Data.EntityId, "climate") {
					log.Println(string(message.Raw))
				}
			}
		}
	})

	err := c.Connect()
	if err != nil {
		log.Panicln(err)
	}
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
