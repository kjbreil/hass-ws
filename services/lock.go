package services

import "encoding/json"

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

// NewLockLock creates the object that can be sent to Home Assistant for domain lock, service lock
// "Lock all or specified locks."
func NewLockLock(target Target, lockLockParams *LockLockParams) *LockLock {
	serviceDomain := "lock"
	serviceType := "call_service"
	serviceService := "lock"
	l := &LockLock{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: *lockLockParams,
	}
	return l
}

type LockLock struct {
	ServiceBase
	ServiceData LockLockParams `json:"service_data,omitempty"`
}
type LockLockParams struct {
	Code *string `json:"code,omitempty"`
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
func NewLockOpen(target Target, lockOpenParams *LockOpenParams) *LockOpen {
	serviceDomain := "lock"
	serviceType := "call_service"
	serviceService := "open"
	l := &LockOpen{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: *lockOpenParams,
	}
	return l
}

type LockOpen struct {
	ServiceBase
	ServiceData LockOpenParams `json:"service_data,omitempty"`
}
type LockOpenParams struct {
	Code *string `json:"code,omitempty"`
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
func NewLockUnlock(target Target, lockUnlockParams *LockUnlockParams) *LockUnlock {
	serviceDomain := "lock"
	serviceType := "call_service"
	serviceService := "unlock"
	l := &LockUnlock{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: *lockUnlockParams,
	}
	return l
}

type LockUnlock struct {
	ServiceBase
	ServiceData LockUnlockParams `json:"service_data,omitempty"`
}
type LockUnlockParams struct {
	Code *string `json:"code,omitempty"`
}

func (l *LockUnlock) JSON() string {
	data, _ := json.Marshal(l)
	return string(data)
}
func (l *LockUnlock) SetID(id *int) {
	l.Id = id
}
