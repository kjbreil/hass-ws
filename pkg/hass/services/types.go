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
