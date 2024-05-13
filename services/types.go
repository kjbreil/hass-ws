//go:generate stringer -type=ColorName -trimprefix=ColorName
//go:generate stringer -type=CommandType -trimprefix=CommandType
//go:generate stringer -type=Direction -trimprefix=Direction
//go:generate stringer -type=Enqueue -trimprefix=Enqueue
//go:generate stringer -type=Flash -trimprefix=Flash
//go:generate stringer -type=Format -trimprefix=Format
//go:generate stringer -type=HvacMode -trimprefix=HvacMode
//go:generate stringer -type=Level -trimprefix=Level
//go:generate stringer -type=Mode -trimprefix=Mode
//go:generate stringer -type=Qos -trimprefix=Qos
//go:generate stringer -type=Repeat -trimprefix=Repeat
//go:generate stringer -type=RepeatType -trimprefix=RepeatType
//go:generate jsonenums -type=ColorName
//go:generate jsonenums -type=CommandType
//go:generate jsonenums -type=Direction
//go:generate jsonenums -type=Enqueue
//go:generate jsonenums -type=Flash
//go:generate jsonenums -type=Format
//go:generate jsonenums -type=HvacMode
//go:generate jsonenums -type=Level
//go:generate jsonenums -type=Mode
//go:generate jsonenums -type=Qos
//go:generate jsonenums -type=Repeat
//go:generate jsonenums -type=RepeatType

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
