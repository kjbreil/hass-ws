package entities

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

type Humidifier struct {
	Additional        Additional `json:"additional,omitempty"`
	AvailableModes    *[]string  `json:"available_modes,omitempty"`
	DeviceClass       *string    `json:"device_class,omitempty"`
	FriendlyName      *string    `json:"friendly_name,omitempty"`
	IsOn              *bool      `json:"is_on,omitempty"`
	MaxHumidity       *int       `json:"max_humidity,omitempty"`
	MinHumidity       *int       `json:"min_humidity,omitempty"`
	Mode              *string    `json:"mode,omitempty"`
	SupportedFeatures *int       `json:"supported_features,omitempty"`
	TargetHumidity    *int       `json:"target_humidity,omitempty"`
}

func GetHumidifier(attributes map[string]interface{}) *Humidifier {
	var h Humidifier
	fillFields(&h, attributes)
	h.Additional = attributes
	return &h
}
