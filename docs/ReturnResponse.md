# ReturnResponse in hass-ws services

Home Assistant’s WebSocket API supports returning response data for some service actions. To request that data, clients must include the `return_response` flag when calling the service.

This library exposes that flag as the `ReturnResponse` field on `services.ServiceBase` and a helper method `SetReturnResponse(true)` on every generated service (via embedding `ServiceBase`).

## What `return_response` does
- When `return_response` is set to true, the `call_service` reply will include a `response` object in the WebSocket `result` payload for actions that support returning data.
- When the service does not support response data, the `response` value will be `null`.
- For actions that must return data, Home Assistant requires `return_response` to be included. The REST API will return HTTP 400 if you misuse it (omit it when required, or include it when not supported). The WebSocket API specifies that it must be included for actions that return data.

References:
- WebSocket API – Calling a service action: https://developers.home-assistant.io/docs/api/websocket/#calling-a-service-action
- REST API – Actions: https://developers.home-assistant.io/docs/api/rest/

## How to use it with hass-ws
Generated services default `ReturnResponse` to `false`. Enable it explicitly when you need response data from a service that supports it.

Example: request a weather forecast via `weather.get_forecast` and read the response.

```go
package main

import (
    "fmt"
    "time"
    hass_ws "github.com/kjbreil/hass-ws"
    "github.com/kjbreil/hass-ws/services"
)

func main() {
    cfg, _ := hass_ws.ParseConfig("config.yml")
    c, _ := hass_ws.NewClient(cfg)
    _ = c.Connect()

    svc := services.NewWeatherGetForecast(services.Targets("weather.home"))
    svc.GetForecastType(services.GetForecastTypehourly) // example service parameter
    svc.SetReturnResponse(true)                         // IMPORTANT: ask Home Assistant to include response data

    rsp := c.CallService(svc)
    msg := rsp.Timeout(2 * time.Second)
    if msg == nil {
        fmt.Println("timed out waiting for response")
        return
    }

    // In this library, the raw `result` object from Home Assistant is placed in Message.ServiceResult.
    // The actual response payload (if any) will be under the "response" key.
    if msg.ServiceResult != nil {
        if raw, ok := msg.ServiceResult["response"]; ok {
            fmt.Printf("service response: %#v\n", raw)
            // You can type-assert raw to the expected structure for the specific service
            // (often a map[string]any or []any depending on the service).
        } else {
            fmt.Println("service returned no response (nil)")
        }
    }
}
```

## When to set `ReturnResponse`
- Set to true for service actions that return response data (for example, `weather.get_forecast`, some conversation-related services, etc.). Consult the Home Assistant docs or Developer Tools → Services to confirm behavior for a specific service.
- Leave as false for actions that only trigger side effects (like turning on a light) and do not return response data.

## Notes and pitfalls
- WebSocket API: The response message will always include a `result` with a `response` key; it will be `null` if the service doesn’t support responses or if you didn’t request it. This library exposes that object via `model.Message.ServiceResult["response"]`.
- REST API (for reference): Use the `?return_response` query parameter when calling a service endpoint. Misuse can result in HTTP 400.
- Performance: Only enable `ReturnResponse` when you need the data; unnecessary responses add overhead.
