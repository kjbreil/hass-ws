package entities

// //////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
// //////////////////////////////////////////////////////////////////////////////
type WaterHeater struct {
	Additional   map[string]interface{} `json:"additional,omitempty"`
	FriendlyName *string                `json:"friendly_name,omitempty"`
}

func GetWaterHeater(attributes map[string]interface{}) *WaterHeater {
	var w WaterHeater
	fillFields(&w, attributes)
	w.Additional = attributes
	return &w
}
