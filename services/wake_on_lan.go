package services

import "encoding/json"

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

// NewWakeOnLanSendMagicPacket creates the object that can be sent to Home Assistant for domain wake_on_lan, service send_magic_packet
// "Send a 'magic packet' to wake up a device with 'Wake-On-LAN' capabilities."
func NewWakeOnLanSendMagicPacket(entities []string, broadcastAddress *string, broadcastPort *float64, mac *string) *WakeOnLanSendMagicPacket {
	serviceDomain := "wake_on_lan"
	serviceType := "call_service"
	serviceService := "send_magic_packet"
	w := &WakeOnLanSendMagicPacket{
		Domain:  &serviceDomain,
		Id:      nil,
		Service: &serviceService,
		ServiceData: struct {
			BroadcastAddress *string  `json:"broadcast_address,omitempty"`
			BroadcastPort    *float64 `json:"broadcast_port,omitempty"`
			Mac              *string  `json:"mac,omitempty"`
		}{
			BroadcastAddress: broadcastAddress,
			BroadcastPort:    broadcastPort,
			Mac:              mac,
		},
		Target: struct {
			EntityId []string `json:"entity_id,omitempty"`
		}{EntityId: entities},
		Type: &serviceType,
	}
	return w
}

type WakeOnLanSendMagicPacket struct {
	Id          *int    `json:"id"`
	Type        *string `json:"type"`
	Domain      *string `json:"domain"`
	Service     *string `json:"service"`
	ServiceData struct {
		BroadcastAddress *string  `json:"broadcast_address,omitempty"`
		BroadcastPort    *float64 `json:"broadcast_port,omitempty"`
		Mac              *string  `json:"mac,omitempty"`
	} `json:"service_data,omitempty"`
	Target struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (w *WakeOnLanSendMagicPacket) JSON() string {
	data, _ := json.Marshal(w)
	return string(data)
}
func (w *WakeOnLanSendMagicPacket) SetID(id *int) {
	w.Id = id
}
