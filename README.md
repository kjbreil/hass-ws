# HASS-WS

Connect to Home Assistant Websocket and subscribe to events

Work in progress. API is very likely to change. Many thanks to https://github.com/W-Floyd/ha-mqtt-iot for the base idea and start generation code.

The client itself subscribes to all events. 
client.OnType is a struct of handlers for each type, the data you get within the handler is the whole message and new and old attributes in a specific type for the entity type. The type is generated from Home Assistant configs and only contains attributes defined there, extra attributes are put into a map[string]interface{} called Additional.  
client.OnEntity is a map of entities (containing the domain so "climate.living" not "living") of functions to run when that entity state changes. Only the message is available with attributes as a map[string]interface{}.

## Entities
For each entity there is a OnXXX function that can be run. The new and old attributes will have entity specific types returned.

## Services
Each service has a generator that is passed to client.CallService. Methods exist on the service type to add parameters to the service.

 This package contains a default list of services however this is most likely not a complete list of services available in your Home Assistant installation since the service list is dynamic based on integrations installed. Generating your own service definitions is needed to interact properly with all your specific integrations. 

There is a HassWSService application that can be used to generate a services package based on your Home Assistant installation.
```bash
go install github.com/kjbreil/hass-ws/helpers/HassWSService@latest
go install github.com/campoy/jsonenums@latest
go install golang.org/x/tools/cmd/stringer@latest
HassWSService
go generate ./...
```

### Calling a Service
To call a service you create a service and then add options to the service. I recommend using Home Assistant Developer Tools -> Services page to get a better understanding of what is needed for each call each service. There is no reporting of requirements in the service definitions so be warned, some parameters are required and others are not, it is also conditional at times. For example `ClimateSetTemperature{}` needs `TargetTempHigh(float64)` and `TargetTempLow(float64)` when the mode is Heat/Cool however if the mode is Heat or the mode is Cool then `Temperature(float64)` is needed and both `TargetTempHigh(float64)` and `TargetTempLow(float64)` are ignored.
```go
service := services.NewClimateSetTemperature(services.Targets("climate.kitchen")).
		HvacMode(services.HvacModeheat_cool).
		TargetTempHigh(75).
		TargetTempLow(65)
gs.ServiceChan <- service
```