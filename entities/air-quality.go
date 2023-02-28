package entities

// //////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
// //////////////////////////////////////////////////////////////////////////////
type AirQuality struct {
	Additional               map[string]interface{} `json:"additional,omitempty"`
	AirQualityIndex          *float64               `json:"air_quality_index,omitempty"`
	CarbonDioxide            *float64               `json:"carbon_dioxide,omitempty"`
	CarbonMonoxide           *float64               `json:"carbon_monoxide,omitempty"`
	FriendlyName             *string                `json:"friendly_name,omitempty"`
	NitrogenDioxide          *float64               `json:"nitrogen_dioxide,omitempty"`
	NitrogenMonoxide         *float64               `json:"nitrogen_monoxide,omitempty"`
	NitrogenOxide            *float64               `json:"nitrogen_oxide,omitempty"`
	Ozone                    *float64               `json:"ozone,omitempty"`
	ParticulateMatter01      *float64               `json:"particulate_matter_0_1,omitempty"`
	ParticulateMatter10      *float64               `json:"particulate_matter_10,omitempty"`
	ParticulateMatter25      *float64               `json:"particulate_matter_2_5,omitempty"`
	SulphurDioxide           *float64               `json:"sulphur_dioxide,omitempty"`
	VolatileOrganicCompounds *float64               `json:"volatile_organic_compounds,omitempty"`
}

func GetAirQuality(attributes map[string]interface{}) *AirQuality {
	var a AirQuality
	FillFields(&a, attributes)
	a.Additional = attributes
	return &a
}
