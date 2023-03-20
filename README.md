# HASS-WS

Connect to Home Assistant Websocket and subscribe to events

Work in progress. API is very likely to change. Many thanks to https://github.com/W-Floyd/ha-mqtt-iot for the base idea and start generation code.

The client itself subscribes to all events. 
client.OnType is a struct of handlers for each type, the data you get within the handler is the whole message and new and old attributes in a specific type for the entity type. The type is generated from Home Assistant configs and only contains attributes defined there, extra attributes are put into a map[string]interface{} called Additional.  
client.OnEntity is a map of entities (containing the domain so "climate.living" not "living") of functions to run when that entity state changes. Only the message is available with attributes as a map[string]interface{}.

## Entities
For each entity there is a OnXXX function that can be run. The new and old attributes will have entity specific types returned.

## Services
Each service has a generator that is passed to client.CallService. If a service has parameters then it will have a type associated. Right now there is no validation that what will be sent will be valid in home asssistant, the error callback will need to be inspected.