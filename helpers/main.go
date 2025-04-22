package main

import (
	"github.com/kjbreil/hass-ws/helpers/HassWSService/servicemaker"
)

// main is the entry point for the helpers tool. It generates Home Assistant services and entities by calling
// external generators. If any generation step fails, the program panics and prints the error.
func main() {
	// Generate Home Assistant services using the servicemaker package
	err := servicemaker.GenServices()
	if err != nil {
		panic(err) // Panic if service generation fails
	}

	// Generate Home Assistant entities
	err = GenEntities()
	if err != nil {
		panic(err) // Panic if entity generation fails
	}
}
