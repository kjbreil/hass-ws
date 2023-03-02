package model

import (
	entities "github.com/kjbreil/hass-ws/entities"
	"strings"
)

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

type OnAirQualityHandler func(message *Message, newAttrs, oldAttrs *entities.AirQuality)
type OnAlarmControlPanelHandler func(message *Message, newAttrs, oldAttrs *entities.AlarmControlPanel)
type OnBinarySensorHandler func(message *Message, newAttrs, oldAttrs *entities.BinarySensor)
type OnButtonHandler func(message *Message, newAttrs, oldAttrs *entities.Button)
type OnCameraHandler func(message *Message, newAttrs, oldAttrs *entities.Camera)
type OnClimateHandler func(message *Message, newAttrs, oldAttrs *entities.Climate)
type OnCoverHandler func(message *Message, newAttrs, oldAttrs *entities.Cover)
type OnDeviceTrackerHandler func(message *Message, newAttrs, oldAttrs *entities.DeviceTracker)
type OnFanHandler func(message *Message, newAttrs, oldAttrs *entities.Fan)
type OnHumidifierHandler func(message *Message, newAttrs, oldAttrs *entities.Humidifier)
type OnIntroductionHandler func(message *Message, newAttrs, oldAttrs *entities.Introduction)
type OnLightHandler func(message *Message, newAttrs, oldAttrs *entities.Light)
type OnLockHandler func(message *Message, newAttrs, oldAttrs *entities.Lock)
type OnMediaPlayerHandler func(message *Message, newAttrs, oldAttrs *entities.MediaPlayer)
type OnNumberHandler func(message *Message, newAttrs, oldAttrs *entities.Number)
type OnRemoteHandler func(message *Message, newAttrs, oldAttrs *entities.Remote)
type OnSelectHandler func(message *Message, newAttrs, oldAttrs *entities.Select)
type OnSensorHandler func(message *Message, newAttrs, oldAttrs *entities.Sensor)
type OnSirenHandler func(message *Message, newAttrs, oldAttrs *entities.Siren)
type OnSwitchHandler func(message *Message, newAttrs, oldAttrs *entities.Switch)
type OnTextHandler func(message *Message, newAttrs, oldAttrs *entities.Text)
type OnUpdateHandler func(message *Message, newAttrs, oldAttrs *entities.Update)
type OnVacuumHandler func(message *Message, newAttrs, oldAttrs *entities.Vacuum)
type OnWaterHeaterHandler func(message *Message, newAttrs, oldAttrs *entities.WaterHeater)
type OnWeatherHandler func(message *Message, newAttrs, oldAttrs *entities.Weather)
type OnTypeHandlers struct {
	OnAirQuality        OnAirQualityHandler
	OnAlarmControlPanel OnAlarmControlPanelHandler
	OnBinarySensor      OnBinarySensorHandler
	OnButton            OnButtonHandler
	OnCamera            OnCameraHandler
	OnClimate           OnClimateHandler
	OnCover             OnCoverHandler
	OnDeviceTracker     OnDeviceTrackerHandler
	OnFan               OnFanHandler
	OnHumidifier        OnHumidifierHandler
	OnIntroduction      OnIntroductionHandler
	OnLight             OnLightHandler
	OnLock              OnLockHandler
	OnMediaPlayer       OnMediaPlayerHandler
	OnNumber            OnNumberHandler
	OnRemote            OnRemoteHandler
	OnSelect            OnSelectHandler
	OnSensor            OnSensorHandler
	OnSiren             OnSirenHandler
	OnSwitch            OnSwitchHandler
	OnText              OnTextHandler
	OnUpdate            OnUpdateHandler
	OnVacuum            OnVacuumHandler
	OnWaterHeater       OnWaterHeaterHandler
	OnWeather           OnWeatherHandler
}

func (o *OnTypeHandlers) Run(message *Message) bool {
	if message.Event == nil || message.Event.Data == nil || message.Event.Data.EntityId == nil {
		return false
	}
	entityType := strings.Split(*message.Event.Data.EntityId, ".")[0]
	switch entityType {
	case "air_quality":
		if o.OnAirQuality == nil {
			return true
		}
		newAttrs := entities.GetAirQuality(message.Event.Data.NewState.Attributes)
		oldAttrs := entities.GetAirQuality(message.Event.Data.OldState.Attributes)
		o.OnAirQuality(message, newAttrs, oldAttrs)
	case "alarm_control_panel":
		if o.OnAlarmControlPanel == nil {
			return true
		}
		newAttrs := entities.GetAlarmControlPanel(message.Event.Data.NewState.Attributes)
		oldAttrs := entities.GetAlarmControlPanel(message.Event.Data.OldState.Attributes)
		o.OnAlarmControlPanel(message, newAttrs, oldAttrs)
	case "binary_sensor":
		if o.OnBinarySensor == nil {
			return true
		}
		newAttrs := entities.GetBinarySensor(message.Event.Data.NewState.Attributes)
		oldAttrs := entities.GetBinarySensor(message.Event.Data.OldState.Attributes)
		o.OnBinarySensor(message, newAttrs, oldAttrs)
	case "button":
		if o.OnButton == nil {
			return true
		}
		newAttrs := entities.GetButton(message.Event.Data.NewState.Attributes)
		oldAttrs := entities.GetButton(message.Event.Data.OldState.Attributes)
		o.OnButton(message, newAttrs, oldAttrs)
	case "camera":
		if o.OnCamera == nil {
			return true
		}
		newAttrs := entities.GetCamera(message.Event.Data.NewState.Attributes)
		oldAttrs := entities.GetCamera(message.Event.Data.OldState.Attributes)
		o.OnCamera(message, newAttrs, oldAttrs)
	case "climate":
		if o.OnClimate == nil {
			return true
		}
		newAttrs := entities.GetClimate(message.Event.Data.NewState.Attributes)
		oldAttrs := entities.GetClimate(message.Event.Data.OldState.Attributes)
		o.OnClimate(message, newAttrs, oldAttrs)
	case "cover":
		if o.OnCover == nil {
			return true
		}
		newAttrs := entities.GetCover(message.Event.Data.NewState.Attributes)
		oldAttrs := entities.GetCover(message.Event.Data.OldState.Attributes)
		o.OnCover(message, newAttrs, oldAttrs)
	case "device_tracker":
		if o.OnDeviceTracker == nil {
			return true
		}
		newAttrs := entities.GetDeviceTracker(message.Event.Data.NewState.Attributes)
		oldAttrs := entities.GetDeviceTracker(message.Event.Data.OldState.Attributes)
		o.OnDeviceTracker(message, newAttrs, oldAttrs)
	case "fan":
		if o.OnFan == nil {
			return true
		}
		newAttrs := entities.GetFan(message.Event.Data.NewState.Attributes)
		oldAttrs := entities.GetFan(message.Event.Data.OldState.Attributes)
		o.OnFan(message, newAttrs, oldAttrs)
	case "humidifier":
		if o.OnHumidifier == nil {
			return true
		}
		newAttrs := entities.GetHumidifier(message.Event.Data.NewState.Attributes)
		oldAttrs := entities.GetHumidifier(message.Event.Data.OldState.Attributes)
		o.OnHumidifier(message, newAttrs, oldAttrs)
	case "introduction":
		if o.OnIntroduction == nil {
			return true
		}
		newAttrs := entities.GetIntroduction(message.Event.Data.NewState.Attributes)
		oldAttrs := entities.GetIntroduction(message.Event.Data.OldState.Attributes)
		o.OnIntroduction(message, newAttrs, oldAttrs)
	case "light":
		if o.OnLight == nil {
			return true
		}
		newAttrs := entities.GetLight(message.Event.Data.NewState.Attributes)
		oldAttrs := entities.GetLight(message.Event.Data.OldState.Attributes)
		o.OnLight(message, newAttrs, oldAttrs)
	case "lock":
		if o.OnLock == nil {
			return true
		}
		newAttrs := entities.GetLock(message.Event.Data.NewState.Attributes)
		oldAttrs := entities.GetLock(message.Event.Data.OldState.Attributes)
		o.OnLock(message, newAttrs, oldAttrs)
	case "media_player":
		if o.OnMediaPlayer == nil {
			return true
		}
		newAttrs := entities.GetMediaPlayer(message.Event.Data.NewState.Attributes)
		oldAttrs := entities.GetMediaPlayer(message.Event.Data.OldState.Attributes)
		o.OnMediaPlayer(message, newAttrs, oldAttrs)
	case "number":
		if o.OnNumber == nil {
			return true
		}
		newAttrs := entities.GetNumber(message.Event.Data.NewState.Attributes)
		oldAttrs := entities.GetNumber(message.Event.Data.OldState.Attributes)
		o.OnNumber(message, newAttrs, oldAttrs)
	case "remote":
		if o.OnRemote == nil {
			return true
		}
		newAttrs := entities.GetRemote(message.Event.Data.NewState.Attributes)
		oldAttrs := entities.GetRemote(message.Event.Data.OldState.Attributes)
		o.OnRemote(message, newAttrs, oldAttrs)
	case "select":
		if o.OnSelect == nil {
			return true
		}
		newAttrs := entities.GetSelect(message.Event.Data.NewState.Attributes)
		oldAttrs := entities.GetSelect(message.Event.Data.OldState.Attributes)
		o.OnSelect(message, newAttrs, oldAttrs)
	case "sensor":
		if o.OnSensor == nil {
			return true
		}
		newAttrs := entities.GetSensor(message.Event.Data.NewState.Attributes)
		oldAttrs := entities.GetSensor(message.Event.Data.OldState.Attributes)
		o.OnSensor(message, newAttrs, oldAttrs)
	case "siren":
		if o.OnSiren == nil {
			return true
		}
		newAttrs := entities.GetSiren(message.Event.Data.NewState.Attributes)
		oldAttrs := entities.GetSiren(message.Event.Data.OldState.Attributes)
		o.OnSiren(message, newAttrs, oldAttrs)
	case "switch":
		if o.OnSwitch == nil {
			return true
		}
		newAttrs := entities.GetSwitch(message.Event.Data.NewState.Attributes)
		oldAttrs := entities.GetSwitch(message.Event.Data.OldState.Attributes)
		o.OnSwitch(message, newAttrs, oldAttrs)
	case "text":
		if o.OnText == nil {
			return true
		}
		newAttrs := entities.GetText(message.Event.Data.NewState.Attributes)
		oldAttrs := entities.GetText(message.Event.Data.OldState.Attributes)
		o.OnText(message, newAttrs, oldAttrs)
	case "update":
		if o.OnUpdate == nil {
			return true
		}
		newAttrs := entities.GetUpdate(message.Event.Data.NewState.Attributes)
		oldAttrs := entities.GetUpdate(message.Event.Data.OldState.Attributes)
		o.OnUpdate(message, newAttrs, oldAttrs)
	case "vacuum":
		if o.OnVacuum == nil {
			return true
		}
		newAttrs := entities.GetVacuum(message.Event.Data.NewState.Attributes)
		oldAttrs := entities.GetVacuum(message.Event.Data.OldState.Attributes)
		o.OnVacuum(message, newAttrs, oldAttrs)
	case "water_heater":
		if o.OnWaterHeater == nil {
			return true
		}
		newAttrs := entities.GetWaterHeater(message.Event.Data.NewState.Attributes)
		oldAttrs := entities.GetWaterHeater(message.Event.Data.OldState.Attributes)
		o.OnWaterHeater(message, newAttrs, oldAttrs)
	case "weather":
		if o.OnWeather == nil {
			return true
		}
		newAttrs := entities.GetWeather(message.Event.Data.NewState.Attributes)
		oldAttrs := entities.GetWeather(message.Event.Data.OldState.Attributes)
		o.OnWeather(message, newAttrs, oldAttrs)
	}
	return false
}
func (o *OnTypeHandlers) RunStates(message *Message) {
	if message.Result == nil {
		return
	}
	for _, r := range message.Result {
		entityType := strings.Split(*r.EntityId, ".")[0]
		switch entityType {
		case "air_quality":
			if o.OnAirQuality == nil {
				continue
			}
			newAttrs := entities.GetAirQuality(r.Attributes)
			newMsg := &Message{
				Event: &Event{
					Context: r.Context,
					Data: &Data{
						EntityId: r.EntityId,
						NewState: &State{
							Attributes:  r.Attributes,
							Context:     r.Context,
							EntityId:    r.EntityId,
							LastChanged: r.LastChanged,
							LastUpdated: r.LastUpdated,
							State:       r.State,
						},
						OldState: nil,
					},
					EventType: EventTypeState,
				},
				ID:   message.ID,
				Type: MessageTypeEvent,
			}
			o.OnAirQuality(newMsg, newAttrs, nil)
		case "alarm_control_panel":
			if o.OnAlarmControlPanel == nil {
				continue
			}
			newAttrs := entities.GetAlarmControlPanel(r.Attributes)
			newMsg := &Message{
				Event: &Event{
					Context: r.Context,
					Data: &Data{
						EntityId: r.EntityId,
						NewState: &State{
							Attributes:  r.Attributes,
							Context:     r.Context,
							EntityId:    r.EntityId,
							LastChanged: r.LastChanged,
							LastUpdated: r.LastUpdated,
							State:       r.State,
						},
						OldState: nil,
					},
					EventType: EventTypeState,
				},
				ID:   message.ID,
				Type: MessageTypeEvent,
			}
			o.OnAlarmControlPanel(newMsg, newAttrs, nil)
		case "binary_sensor":
			if o.OnBinarySensor == nil {
				continue
			}
			newAttrs := entities.GetBinarySensor(r.Attributes)
			newMsg := &Message{
				Event: &Event{
					Context: r.Context,
					Data: &Data{
						EntityId: r.EntityId,
						NewState: &State{
							Attributes:  r.Attributes,
							Context:     r.Context,
							EntityId:    r.EntityId,
							LastChanged: r.LastChanged,
							LastUpdated: r.LastUpdated,
							State:       r.State,
						},
						OldState: nil,
					},
					EventType: EventTypeState,
				},
				ID:   message.ID,
				Type: MessageTypeEvent,
			}
			o.OnBinarySensor(newMsg, newAttrs, nil)
		case "button":
			if o.OnButton == nil {
				continue
			}
			newAttrs := entities.GetButton(r.Attributes)
			newMsg := &Message{
				Event: &Event{
					Context: r.Context,
					Data: &Data{
						EntityId: r.EntityId,
						NewState: &State{
							Attributes:  r.Attributes,
							Context:     r.Context,
							EntityId:    r.EntityId,
							LastChanged: r.LastChanged,
							LastUpdated: r.LastUpdated,
							State:       r.State,
						},
						OldState: nil,
					},
					EventType: EventTypeState,
				},
				ID:   message.ID,
				Type: MessageTypeEvent,
			}
			o.OnButton(newMsg, newAttrs, nil)
		case "camera":
			if o.OnCamera == nil {
				continue
			}
			newAttrs := entities.GetCamera(r.Attributes)
			newMsg := &Message{
				Event: &Event{
					Context: r.Context,
					Data: &Data{
						EntityId: r.EntityId,
						NewState: &State{
							Attributes:  r.Attributes,
							Context:     r.Context,
							EntityId:    r.EntityId,
							LastChanged: r.LastChanged,
							LastUpdated: r.LastUpdated,
							State:       r.State,
						},
						OldState: nil,
					},
					EventType: EventTypeState,
				},
				ID:   message.ID,
				Type: MessageTypeEvent,
			}
			o.OnCamera(newMsg, newAttrs, nil)
		case "climate":
			if o.OnClimate == nil {
				continue
			}
			newAttrs := entities.GetClimate(r.Attributes)
			newMsg := &Message{
				Event: &Event{
					Context: r.Context,
					Data: &Data{
						EntityId: r.EntityId,
						NewState: &State{
							Attributes:  r.Attributes,
							Context:     r.Context,
							EntityId:    r.EntityId,
							LastChanged: r.LastChanged,
							LastUpdated: r.LastUpdated,
							State:       r.State,
						},
						OldState: nil,
					},
					EventType: EventTypeState,
				},
				ID:   message.ID,
				Type: MessageTypeEvent,
			}
			o.OnClimate(newMsg, newAttrs, nil)
		case "cover":
			if o.OnCover == nil {
				continue
			}
			newAttrs := entities.GetCover(r.Attributes)
			newMsg := &Message{
				Event: &Event{
					Context: r.Context,
					Data: &Data{
						EntityId: r.EntityId,
						NewState: &State{
							Attributes:  r.Attributes,
							Context:     r.Context,
							EntityId:    r.EntityId,
							LastChanged: r.LastChanged,
							LastUpdated: r.LastUpdated,
							State:       r.State,
						},
						OldState: nil,
					},
					EventType: EventTypeState,
				},
				ID:   message.ID,
				Type: MessageTypeEvent,
			}
			o.OnCover(newMsg, newAttrs, nil)
		case "device_tracker":
			if o.OnDeviceTracker == nil {
				continue
			}
			newAttrs := entities.GetDeviceTracker(r.Attributes)
			newMsg := &Message{
				Event: &Event{
					Context: r.Context,
					Data: &Data{
						EntityId: r.EntityId,
						NewState: &State{
							Attributes:  r.Attributes,
							Context:     r.Context,
							EntityId:    r.EntityId,
							LastChanged: r.LastChanged,
							LastUpdated: r.LastUpdated,
							State:       r.State,
						},
						OldState: nil,
					},
					EventType: EventTypeState,
				},
				ID:   message.ID,
				Type: MessageTypeEvent,
			}
			o.OnDeviceTracker(newMsg, newAttrs, nil)
		case "fan":
			if o.OnFan == nil {
				continue
			}
			newAttrs := entities.GetFan(r.Attributes)
			newMsg := &Message{
				Event: &Event{
					Context: r.Context,
					Data: &Data{
						EntityId: r.EntityId,
						NewState: &State{
							Attributes:  r.Attributes,
							Context:     r.Context,
							EntityId:    r.EntityId,
							LastChanged: r.LastChanged,
							LastUpdated: r.LastUpdated,
							State:       r.State,
						},
						OldState: nil,
					},
					EventType: EventTypeState,
				},
				ID:   message.ID,
				Type: MessageTypeEvent,
			}
			o.OnFan(newMsg, newAttrs, nil)
		case "humidifier":
			if o.OnHumidifier == nil {
				continue
			}
			newAttrs := entities.GetHumidifier(r.Attributes)
			newMsg := &Message{
				Event: &Event{
					Context: r.Context,
					Data: &Data{
						EntityId: r.EntityId,
						NewState: &State{
							Attributes:  r.Attributes,
							Context:     r.Context,
							EntityId:    r.EntityId,
							LastChanged: r.LastChanged,
							LastUpdated: r.LastUpdated,
							State:       r.State,
						},
						OldState: nil,
					},
					EventType: EventTypeState,
				},
				ID:   message.ID,
				Type: MessageTypeEvent,
			}
			o.OnHumidifier(newMsg, newAttrs, nil)
		case "introduction":
			if o.OnIntroduction == nil {
				continue
			}
			newAttrs := entities.GetIntroduction(r.Attributes)
			newMsg := &Message{
				Event: &Event{
					Context: r.Context,
					Data: &Data{
						EntityId: r.EntityId,
						NewState: &State{
							Attributes:  r.Attributes,
							Context:     r.Context,
							EntityId:    r.EntityId,
							LastChanged: r.LastChanged,
							LastUpdated: r.LastUpdated,
							State:       r.State,
						},
						OldState: nil,
					},
					EventType: EventTypeState,
				},
				ID:   message.ID,
				Type: MessageTypeEvent,
			}
			o.OnIntroduction(newMsg, newAttrs, nil)
		case "light":
			if o.OnLight == nil {
				continue
			}
			newAttrs := entities.GetLight(r.Attributes)
			newMsg := &Message{
				Event: &Event{
					Context: r.Context,
					Data: &Data{
						EntityId: r.EntityId,
						NewState: &State{
							Attributes:  r.Attributes,
							Context:     r.Context,
							EntityId:    r.EntityId,
							LastChanged: r.LastChanged,
							LastUpdated: r.LastUpdated,
							State:       r.State,
						},
						OldState: nil,
					},
					EventType: EventTypeState,
				},
				ID:   message.ID,
				Type: MessageTypeEvent,
			}
			o.OnLight(newMsg, newAttrs, nil)
		case "lock":
			if o.OnLock == nil {
				continue
			}
			newAttrs := entities.GetLock(r.Attributes)
			newMsg := &Message{
				Event: &Event{
					Context: r.Context,
					Data: &Data{
						EntityId: r.EntityId,
						NewState: &State{
							Attributes:  r.Attributes,
							Context:     r.Context,
							EntityId:    r.EntityId,
							LastChanged: r.LastChanged,
							LastUpdated: r.LastUpdated,
							State:       r.State,
						},
						OldState: nil,
					},
					EventType: EventTypeState,
				},
				ID:   message.ID,
				Type: MessageTypeEvent,
			}
			o.OnLock(newMsg, newAttrs, nil)
		case "media_player":
			if o.OnMediaPlayer == nil {
				continue
			}
			newAttrs := entities.GetMediaPlayer(r.Attributes)
			newMsg := &Message{
				Event: &Event{
					Context: r.Context,
					Data: &Data{
						EntityId: r.EntityId,
						NewState: &State{
							Attributes:  r.Attributes,
							Context:     r.Context,
							EntityId:    r.EntityId,
							LastChanged: r.LastChanged,
							LastUpdated: r.LastUpdated,
							State:       r.State,
						},
						OldState: nil,
					},
					EventType: EventTypeState,
				},
				ID:   message.ID,
				Type: MessageTypeEvent,
			}
			o.OnMediaPlayer(newMsg, newAttrs, nil)
		case "number":
			if o.OnNumber == nil {
				continue
			}
			newAttrs := entities.GetNumber(r.Attributes)
			newMsg := &Message{
				Event: &Event{
					Context: r.Context,
					Data: &Data{
						EntityId: r.EntityId,
						NewState: &State{
							Attributes:  r.Attributes,
							Context:     r.Context,
							EntityId:    r.EntityId,
							LastChanged: r.LastChanged,
							LastUpdated: r.LastUpdated,
							State:       r.State,
						},
						OldState: nil,
					},
					EventType: EventTypeState,
				},
				ID:   message.ID,
				Type: MessageTypeEvent,
			}
			o.OnNumber(newMsg, newAttrs, nil)
		case "remote":
			if o.OnRemote == nil {
				continue
			}
			newAttrs := entities.GetRemote(r.Attributes)
			newMsg := &Message{
				Event: &Event{
					Context: r.Context,
					Data: &Data{
						EntityId: r.EntityId,
						NewState: &State{
							Attributes:  r.Attributes,
							Context:     r.Context,
							EntityId:    r.EntityId,
							LastChanged: r.LastChanged,
							LastUpdated: r.LastUpdated,
							State:       r.State,
						},
						OldState: nil,
					},
					EventType: EventTypeState,
				},
				ID:   message.ID,
				Type: MessageTypeEvent,
			}
			o.OnRemote(newMsg, newAttrs, nil)
		case "select":
			if o.OnSelect == nil {
				continue
			}
			newAttrs := entities.GetSelect(r.Attributes)
			newMsg := &Message{
				Event: &Event{
					Context: r.Context,
					Data: &Data{
						EntityId: r.EntityId,
						NewState: &State{
							Attributes:  r.Attributes,
							Context:     r.Context,
							EntityId:    r.EntityId,
							LastChanged: r.LastChanged,
							LastUpdated: r.LastUpdated,
							State:       r.State,
						},
						OldState: nil,
					},
					EventType: EventTypeState,
				},
				ID:   message.ID,
				Type: MessageTypeEvent,
			}
			o.OnSelect(newMsg, newAttrs, nil)
		case "sensor":
			if o.OnSensor == nil {
				continue
			}
			newAttrs := entities.GetSensor(r.Attributes)
			newMsg := &Message{
				Event: &Event{
					Context: r.Context,
					Data: &Data{
						EntityId: r.EntityId,
						NewState: &State{
							Attributes:  r.Attributes,
							Context:     r.Context,
							EntityId:    r.EntityId,
							LastChanged: r.LastChanged,
							LastUpdated: r.LastUpdated,
							State:       r.State,
						},
						OldState: nil,
					},
					EventType: EventTypeState,
				},
				ID:   message.ID,
				Type: MessageTypeEvent,
			}
			o.OnSensor(newMsg, newAttrs, nil)
		case "siren":
			if o.OnSiren == nil {
				continue
			}
			newAttrs := entities.GetSiren(r.Attributes)
			newMsg := &Message{
				Event: &Event{
					Context: r.Context,
					Data: &Data{
						EntityId: r.EntityId,
						NewState: &State{
							Attributes:  r.Attributes,
							Context:     r.Context,
							EntityId:    r.EntityId,
							LastChanged: r.LastChanged,
							LastUpdated: r.LastUpdated,
							State:       r.State,
						},
						OldState: nil,
					},
					EventType: EventTypeState,
				},
				ID:   message.ID,
				Type: MessageTypeEvent,
			}
			o.OnSiren(newMsg, newAttrs, nil)
		case "switch":
			if o.OnSwitch == nil {
				continue
			}
			newAttrs := entities.GetSwitch(r.Attributes)
			newMsg := &Message{
				Event: &Event{
					Context: r.Context,
					Data: &Data{
						EntityId: r.EntityId,
						NewState: &State{
							Attributes:  r.Attributes,
							Context:     r.Context,
							EntityId:    r.EntityId,
							LastChanged: r.LastChanged,
							LastUpdated: r.LastUpdated,
							State:       r.State,
						},
						OldState: nil,
					},
					EventType: EventTypeState,
				},
				ID:   message.ID,
				Type: MessageTypeEvent,
			}
			o.OnSwitch(newMsg, newAttrs, nil)
		case "text":
			if o.OnText == nil {
				continue
			}
			newAttrs := entities.GetText(r.Attributes)
			newMsg := &Message{
				Event: &Event{
					Context: r.Context,
					Data: &Data{
						EntityId: r.EntityId,
						NewState: &State{
							Attributes:  r.Attributes,
							Context:     r.Context,
							EntityId:    r.EntityId,
							LastChanged: r.LastChanged,
							LastUpdated: r.LastUpdated,
							State:       r.State,
						},
						OldState: nil,
					},
					EventType: EventTypeState,
				},
				ID:   message.ID,
				Type: MessageTypeEvent,
			}
			o.OnText(newMsg, newAttrs, nil)
		case "update":
			if o.OnUpdate == nil {
				continue
			}
			newAttrs := entities.GetUpdate(r.Attributes)
			newMsg := &Message{
				Event: &Event{
					Context: r.Context,
					Data: &Data{
						EntityId: r.EntityId,
						NewState: &State{
							Attributes:  r.Attributes,
							Context:     r.Context,
							EntityId:    r.EntityId,
							LastChanged: r.LastChanged,
							LastUpdated: r.LastUpdated,
							State:       r.State,
						},
						OldState: nil,
					},
					EventType: EventTypeState,
				},
				ID:   message.ID,
				Type: MessageTypeEvent,
			}
			o.OnUpdate(newMsg, newAttrs, nil)
		case "vacuum":
			if o.OnVacuum == nil {
				continue
			}
			newAttrs := entities.GetVacuum(r.Attributes)
			newMsg := &Message{
				Event: &Event{
					Context: r.Context,
					Data: &Data{
						EntityId: r.EntityId,
						NewState: &State{
							Attributes:  r.Attributes,
							Context:     r.Context,
							EntityId:    r.EntityId,
							LastChanged: r.LastChanged,
							LastUpdated: r.LastUpdated,
							State:       r.State,
						},
						OldState: nil,
					},
					EventType: EventTypeState,
				},
				ID:   message.ID,
				Type: MessageTypeEvent,
			}
			o.OnVacuum(newMsg, newAttrs, nil)
		case "water_heater":
			if o.OnWaterHeater == nil {
				continue
			}
			newAttrs := entities.GetWaterHeater(r.Attributes)
			newMsg := &Message{
				Event: &Event{
					Context: r.Context,
					Data: &Data{
						EntityId: r.EntityId,
						NewState: &State{
							Attributes:  r.Attributes,
							Context:     r.Context,
							EntityId:    r.EntityId,
							LastChanged: r.LastChanged,
							LastUpdated: r.LastUpdated,
							State:       r.State,
						},
						OldState: nil,
					},
					EventType: EventTypeState,
				},
				ID:   message.ID,
				Type: MessageTypeEvent,
			}
			o.OnWaterHeater(newMsg, newAttrs, nil)
		case "weather":
			if o.OnWeather == nil {
				continue
			}
			newAttrs := entities.GetWeather(r.Attributes)
			newMsg := &Message{
				Event: &Event{
					Context: r.Context,
					Data: &Data{
						EntityId: r.EntityId,
						NewState: &State{
							Attributes:  r.Attributes,
							Context:     r.Context,
							EntityId:    r.EntityId,
							LastChanged: r.LastChanged,
							LastUpdated: r.LastUpdated,
							State:       r.State,
						},
						OldState: nil,
					},
					EventType: EventTypeState,
				},
				ID:   message.ID,
				Type: MessageTypeEvent,
			}
			o.OnWeather(newMsg, newAttrs, nil)
		}
	}
}

type OnEntityHandlers map[string]func(message *Message)

func (o OnEntityHandlers) Run(message *Message) bool {
	if message.Event == nil || message.Event.Data == nil || message.Event.Data.EntityId == nil {
		return false
	}
	for k, v := range o {
		if k == *message.Event.Data.EntityId {
			v(message)
			return true
		}
	}
	return false
}
func (o OnEntityHandlers) RunStates(message *Message) {
	if message.Result == nil {
		return
	}
	for _, r := range message.Result {
		for k, v := range o {
			if k == *r.EntityId {
				newMsg := &Message{
					Event: &Event{
						Context: r.Context,
						Data: &Data{
							EntityId: r.EntityId,
							NewState: &State{
								Attributes:  r.Attributes,
								Context:     r.Context,
								EntityId:    r.EntityId,
								LastChanged: r.LastChanged,
								LastUpdated: r.LastUpdated,
								State:       r.State,
							},
							OldState: nil,
						},
						EventType: EventTypeState,
					},
					ID:   message.ID,
					Type: MessageTypeEvent,
				}
				v(newMsg)
			}
		}
	}
}
