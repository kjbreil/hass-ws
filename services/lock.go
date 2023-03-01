package services

import "encoding/json"

// //////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
// //////////////////////////////////////////////////////////////////////////////
//
// NewLockLock creates the object that can be sent to Home Assistant for domain lock, service lock
// "Lock all or specified locks."
func NewLockLock(entities []string, code *string) *LockLock {
	serviceDomain := "lock"
	serviceType := "call_service"
	serviceService := "lock"
	l := &LockLock{
		Domain:  &serviceDomain,
		Id:      nil,
		Service: &serviceService,
		ServiceData: struct {
			Code *string `json:"code,omitempty"`
		}{Code: code},
		Target: struct {
			EntityId []string `json:"entity_id,omitempty"`
		}{EntityId: entities},
		Type: &serviceType,
	}
	return l
}

type LockLock struct {
	Id          *int    `json:"id"`
	Type        *string `json:"type"`
	Domain      *string `json:"domain"`
	Service     *string `json:"service"`
	ServiceData struct {
		Code *string `json:"code,omitempty"`
	} `json:"service_data,omitempty"`
	Target struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (l *LockLock) JSON() string {
	data, _ := json.Marshal(l)
	return string(data)
}
func (l *LockLock) SetID(id *int) {
	l.Id = id
}

// NewLockOpen creates the object that can be sent to Home Assistant for domain lock, service open
// "Open all or specified locks."
func NewLockOpen(entities []string, code *string) *LockOpen {
	serviceDomain := "lock"
	serviceType := "call_service"
	serviceService := "open"
	l := &LockOpen{
		Domain:  &serviceDomain,
		Id:      nil,
		Service: &serviceService,
		ServiceData: struct {
			Code *string `json:"code,omitempty"`
		}{Code: code},
		Target: struct {
			EntityId []string `json:"entity_id,omitempty"`
		}{EntityId: entities},
		Type: &serviceType,
	}
	return l
}

type LockOpen struct {
	Id          *int    `json:"id"`
	Type        *string `json:"type"`
	Domain      *string `json:"domain"`
	Service     *string `json:"service"`
	ServiceData struct {
		Code *string `json:"code,omitempty"`
	} `json:"service_data,omitempty"`
	Target struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (l *LockOpen) JSON() string {
	data, _ := json.Marshal(l)
	return string(data)
}
func (l *LockOpen) SetID(id *int) {
	l.Id = id
}

// NewLockUnlock creates the object that can be sent to Home Assistant for domain lock, service unlock
// "Unlock all or specified locks."
func NewLockUnlock(entities []string, code *string) *LockUnlock {
	serviceDomain := "lock"
	serviceType := "call_service"
	serviceService := "unlock"
	l := &LockUnlock{
		Domain:  &serviceDomain,
		Id:      nil,
		Service: &serviceService,
		ServiceData: struct {
			Code *string `json:"code,omitempty"`
		}{Code: code},
		Target: struct {
			EntityId []string `json:"entity_id,omitempty"`
		}{EntityId: entities},
		Type: &serviceType,
	}
	return l
}

type LockUnlock struct {
	Id          *int    `json:"id"`
	Type        *string `json:"type"`
	Domain      *string `json:"domain"`
	Service     *string `json:"service"`
	ServiceData struct {
		Code *string `json:"code,omitempty"`
	} `json:"service_data,omitempty"`
	Target struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (l *LockUnlock) JSON() string {
	data, _ := json.Marshal(l)
	return string(data)
}
func (l *LockUnlock) SetID(id *int) {
	l.Id = id
}
