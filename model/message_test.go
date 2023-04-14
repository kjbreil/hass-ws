package model

import (
	"testing"
)

func intPointer(i int) *int {
	return &i
}

// Tests don't test for correct output only for if error is returned
// and then if the methods work properly
func TestParseMessage(t *testing.T) {
	type args struct {
		msgJson []byte
	}
	type methodResults struct {
		domainEntity string
		domain       string
		entityID     string
		state        string
		oldState     string
		stateFloat   float64
		friendlyName string
	}
	tests := []struct {
		name    string
		args    args
		methods methodResults
		wantErr bool
	}{
		{
			name: "Climate",
			args: args{msgJson: []byte("{\"id\":1,\"type\":\"event\",\"event\":{\"event_type\":\"state_changed\",\"data\":{\"entity_id\":\"climate.kitchen\",\"old_state\":{\"entity_id\":\"climate.kitchen\",\"state\":\"heat_cool\",\"attributes\":{\"hvac_modes\":[\"off\",\"heat\",\"cool\",\"heat_cool\"],\"min_temp\":50,\"max_temp\":90,\"current_temperature\":69,\"temperature\":null,\"target_temp_high\":76,\"target_temp_low\":69,\"current_humidity\":34,\"hvac_action\":\"idle\",\"friendly_name\":\"Kitchen Thermostat\",\"supported_features\":3},\"last_changed\":\"2023-02-22T13:41:42.681237+00:00\",\"last_updated\":\"2023-02-28T15:58:39.026248+00:00\",\"context\":{\"id\":\"01GTCBYJXJ83KQT53DDVTXJ23C\",\"parent_id\":null,\"user_id\":null}},\"new_state\":{\"entity_id\":\"climate.kitchen\",\"state\":\"heat_cool\",\"attributes\":{\"hvac_modes\":[\"off\",\"heat\",\"cool\",\"heat_cool\"],\"min_temp\":50,\"max_temp\":90,\"current_temperature\":69,\"temperature\":null,\"target_temp_high\":77,\"target_temp_low\":71,\"current_humidity\":34,\"hvac_action\":\"idle\",\"friendly_name\":\"Kitchen Thermostat\",\"supported_features\":3},\"last_changed\":\"2023-02-22T13:41:42.681237+00:00\",\"last_updated\":\"2023-02-28T15:59:03.464870+00:00\",\"context\":{\"id\":\"01GTCBZ9E3RGMPEF1J73X6DPY6\",\"parent_id\":\"01GTCBZ9DR44XXMRPT4ZX9MF5Y\",\"user_id\":null}}},\"origin\":\"LOCAL\",\"time_fired\":\"2023-02-28T15:59:03.464870+00:00\",\"context\":{\"id\":\"01GTCBZ9E3RGMPEF1J73X6DPY6\",\"parent_id\":\"01GTCBZ9DR44XXMRPT4ZX9MF5Y\",\"user_id\":null}}}")},
			methods: methodResults{
				domainEntity: "climate.kitchen",
				domain:       "climate",
				entityID:     "kitchen",
				state:        "heat_cool",
				oldState:     "heat_cool",
				stateFloat:   0,
				friendlyName: "Kitchen Thermostat",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			msg, err := ParseMessage(tt.args.msgJson)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseMessage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if got := msg.DomainEntity(); got != tt.methods.domainEntity {
				t.Errorf("DomainEntity() = %v, want %v", got, tt.methods.domainEntity)
			}

			if got := msg.Domain(); got != tt.methods.domain {
				t.Errorf("Domain() = %v, want %v", got, tt.methods.domainEntity)
			}

			if got := msg.EntityID(); got != tt.methods.entityID {
				t.Errorf("EntityID() = %v, want %v", got, tt.methods.entityID)
			}

			if got := msg.State(); got != tt.methods.state {
				t.Errorf("State() = %v, want %v", got, tt.methods.state)
			}

			if got := msg.OldState(); *got != tt.methods.oldState {
				t.Errorf("OldState() = %v, want %v", *got, tt.methods.oldState)
			}

			if got, _ := msg.StateFloat(); got != tt.methods.stateFloat {
				t.Errorf("StateFloat() = %v, want %v", got, tt.methods.stateFloat)
			}

			if got := msg.Attributes(); got == nil {
				t.Errorf("Attributes() nil")
			}

			if got := msg.OldAttributes(); got == nil {
				t.Errorf("OldAttributes() nil")
			}

			if got := msg.FriendlyName(); got != tt.methods.friendlyName {
				t.Errorf("FriendlyName() = %v, want %v", got, tt.methods.friendlyName)
			}
		})
	}
}
