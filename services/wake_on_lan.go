package services

import (
	"encoding/json"
	"fmt"
)

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

// NewWakeOnLanSendMagicPacket creates the object that can be sent to Home Assistant for domain wake_on_lan, service send_magic_packet
// "Send a 'magic packet' to wake up a device with 'Wake-On-LAN' capabilities."
func NewWakeOnLanSendMagicPacket(target Target) *WakeOnLanSendMagicPacket {
	serviceDomain := "wake_on_lan"
	serviceType := "call_service"
	serviceService := "send_magic_packet"
	w := &WakeOnLanSendMagicPacket{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: WakeOnLanSendMagicPacketParams{},
	}
	return w
}

type WakeOnLanSendMagicPacket struct {
	ServiceBase
	ServiceData WakeOnLanSendMagicPacketParams `json:"service_data,omitempty"`
}
type WakeOnLanSendMagicPacketParams struct {
	BroadcastAddress *string  `json:"broadcast_address,omitempty"`
	BroadcastPort    *float64 `json:"broadcast_port,omitempty"`
	Mac              *string  `json:"mac,omitempty"`
}

func (w *WakeOnLanSendMagicPacket) BroadcastAddress(broadcastAddress string) *WakeOnLanSendMagicPacket {
	w.ServiceData.BroadcastAddress = &broadcastAddress
	return w
}
func (w *WakeOnLanSendMagicPacket) BroadcastPort(broadcastPort float64) *WakeOnLanSendMagicPacket {
	w.ServiceData.BroadcastPort = &broadcastPort
	return w
}
func (w *WakeOnLanSendMagicPacket) Mac(mac string) *WakeOnLanSendMagicPacket {
	w.ServiceData.Mac = &mac
	return w
}
func (w *WakeOnLanSendMagicPacket) JSON() string {
	data, _ := json.Marshal(w)
	return string(data)
}
func (w *WakeOnLanSendMagicPacket) Name() string {
	return fmt.Sprintf("%s.%s", *w.Domain, *w.Service)
}
func (w *WakeOnLanSendMagicPacket) SetID(id *int) {
	w.Id = id
}
