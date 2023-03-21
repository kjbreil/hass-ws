package main

import (
	"github.com/kjbreil/hass-ws/helpers/HassWSService/servicemaker"
)

func main() {

	err := servicemaker.GenServices()
	if err != nil {
		panic(err)
	}

	err = GenEntities()
	if err != nil {
		panic(err)
	}

}
