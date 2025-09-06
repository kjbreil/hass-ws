//go:generate stringer -type=Format -trimprefix=Format
//go:generate stringer -type=GetForecastsType -trimprefix=GetForecastsType
//go:generate stringer -type=Level -trimprefix=Level
//go:generate stringer -type=Mode -trimprefix=Mode
//go:generate stringer -type=Period -trimprefix=Period
//go:generate stringer -type=Qos -trimprefix=Qos
//go:generate stringer -type=Types -trimprefix=Types
//go:generate jsonenums -type=Format
//go:generate jsonenums -type=GetForecastsType
//go:generate jsonenums -type=Level
//go:generate jsonenums -type=Mode
//go:generate jsonenums -type=Period
//go:generate jsonenums -type=Qos
//go:generate jsonenums -type=Types

package services

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

type Service interface {
	SetID(id *int)
	JSON() string
	Name() string
	Targets() []string
	SetReturnResponse(b bool)
}
type ServiceBase struct {
	Id             *int    `json:"id"`
	Type           *string `json:"type"`
	Domain         *string `json:"domain"`
	Service        *string `json:"service"`
	Target         Target  `json:"target,omitempty"`
	ReturnResponse bool    `json:"return_response"`
}

func (s *ServiceBase) SetID(id *int) {
	s.Id = id
}
func (s *ServiceBase) SetReturnResponse(b bool) {
	s.ReturnResponse = b
}

type Target struct {
	EntityId []string `json:"entity_id,omitempty"`
}

func Targets(entities ...string) Target {
	var t Target
	for _, e := range entities {
		t.EntityId = append(t.EntityId, e)
	}
	return t
}

type Types int

const (
	Typeschange Types = iota
	Typeslast_reset
	Typesmax
	Typesmean
	Typesmin
	Typesstate
	Typessum
)
