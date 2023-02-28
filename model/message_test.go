package model

import (
	"reflect"
	"testing"
)

// {"id":1,"type":"event","event":{"event_type":"state_changed","data":{"entity_id":"climate.kitchen","old_state":{"entity_id":"climate.kitchen","state":"heat_cool","attributes":{"hvac_modes":["off","heat","cool","heat_cool"],"min_temp":50,"max_temp":90,"current_temperature":69,"temperature":null,"target_temp_high":76,"target_temp_low":69,"current_humidity":34,"hvac_action":"idle","friendly_name":"Kitchen Thermostat","supported_features":3},"last_changed":"2023-02-22T13:41:42.681237+00:00","last_updated":"2023-02-28T15:58:39.026248+00:00","context":{"id":"01GTCBYJXJ83KQT53DDVTXJ23C","parent_id":null,"user_id":null}},"new_state":{"entity_id":"climate.kitchen","state":"heat_cool","attributes":{"hvac_modes":["off","heat","cool","heat_cool"],"min_temp":50,"max_temp":90,"current_temperature":69,"temperature":null,"target_temp_high":77,"target_temp_low":71,"current_humidity":34,"hvac_action":"idle","friendly_name":"Kitchen Thermostat","supported_features":3},"last_changed":"2023-02-22T13:41:42.681237+00:00","last_updated":"2023-02-28T15:59:03.464870+00:00","context":{"id":"01GTCBZ9E3RGMPEF1J73X6DPY6","parent_id":"01GTCBZ9DR44XXMRPT4ZX9MF5Y","user_id":null}}},"origin":"LOCAL","time_fired":"2023-02-28T15:59:03.464870+00:00","context":{"id":"01GTCBZ9E3RGMPEF1J73X6DPY6","parent_id":"01GTCBZ9DR44XXMRPT4ZX9MF5Y","user_id":null}}}
func TestParseMessage(t *testing.T) {
	type args struct {
		msgJson []byte
	}
	tests := []struct {
		name    string
		args    args
		want    Message
		wantErr bool
	}{
		{
			name: "Climate",
			args: args{msgJson: []byte("{\"id\":1,\"type\":\"event\",\"event\":{\"event_type\":\"state_changed\",\"data\":{\"entity_id\":\"climate.kitchen\",\"old_state\":{\"entity_id\":\"climate.kitchen\",\"state\":\"heat_cool\",\"attributes\":{\"hvac_modes\":[\"off\",\"heat\",\"cool\",\"heat_cool\"],\"min_temp\":50,\"max_temp\":90,\"current_temperature\":69,\"temperature\":null,\"target_temp_high\":76,\"target_temp_low\":69,\"current_humidity\":34,\"hvac_action\":\"idle\",\"friendly_name\":\"Kitchen Thermostat\",\"supported_features\":3},\"last_changed\":\"2023-02-22T13:41:42.681237+00:00\",\"last_updated\":\"2023-02-28T15:58:39.026248+00:00\",\"context\":{\"id\":\"01GTCBYJXJ83KQT53DDVTXJ23C\",\"parent_id\":null,\"user_id\":null}},\"new_state\":{\"entity_id\":\"climate.kitchen\",\"state\":\"heat_cool\",\"attributes\":{\"hvac_modes\":[\"off\",\"heat\",\"cool\",\"heat_cool\"],\"min_temp\":50,\"max_temp\":90,\"current_temperature\":69,\"temperature\":null,\"target_temp_high\":77,\"target_temp_low\":71,\"current_humidity\":34,\"hvac_action\":\"idle\",\"friendly_name\":\"Kitchen Thermostat\",\"supported_features\":3},\"last_changed\":\"2023-02-22T13:41:42.681237+00:00\",\"last_updated\":\"2023-02-28T15:59:03.464870+00:00\",\"context\":{\"id\":\"01GTCBZ9E3RGMPEF1J73X6DPY6\",\"parent_id\":\"01GTCBZ9DR44XXMRPT4ZX9MF5Y\",\"user_id\":null}}},\"origin\":\"LOCAL\",\"time_fired\":\"2023-02-28T15:59:03.464870+00:00\",\"context\":{\"id\":\"01GTCBZ9E3RGMPEF1J73X6DPY6\",\"parent_id\":\"01GTCBZ9DR44XXMRPT4ZX9MF5Y\",\"user_id\":null}}}")},
			want: Message{
				ID:          nil,
				Type:        "",
				AccessToken: nil,
				Success:     nil,
				EventType:   nil,
				Error:       nil,
				Event:       nil,
				Raw:         nil,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseMessage(tt.args.msgJson)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseMessage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseMessage() got = %v, want %v", got, tt.want)
			}
		})
	}
}
