package services

import (
	"fmt"
	gojson "github.com/goccy/go-json"
)

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

// NewLockLock creates the object that can be sent to Home Assistant for domain lock, service lock
// "Lock all or specified locks."
func NewLockLock(target Target) *LockLock {
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
		ServiceData: LockLockParams{},
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

func (l *LockLock) Code(code string) *LockLock {
	l.ServiceData.Code = &code
	return l
}
func (l *LockLock) JSON() string {
	data, _ := gojson.Marshal(l)
	return string(data)
}
func (l *LockLock) Name() string {
	return fmt.Sprintf("%s.%s", *l.Domain, *l.Service)
}
func (l *LockLock) SetID(id *int) {
	l.Id = id
}

// NewLockOpen creates the object that can be sent to Home Assistant for domain lock, service open
// "Open all or specified locks."
func NewLockOpen(target Target) *LockOpen {
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
		ServiceData: LockOpenParams{},
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

func (l *LockOpen) Code(code string) *LockOpen {
	l.ServiceData.Code = &code
	return l
}
func (l *LockOpen) JSON() string {
	data, _ := gojson.Marshal(l)
	return string(data)
}
func (l *LockOpen) Name() string {
	return fmt.Sprintf("%s.%s", *l.Domain, *l.Service)
}
func (l *LockOpen) SetID(id *int) {
	l.Id = id
}

// NewLockUnlock creates the object that can be sent to Home Assistant for domain lock, service unlock
// "Unlock all or specified locks."
func NewLockUnlock(target Target) *LockUnlock {
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
		ServiceData: LockUnlockParams{},
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

func (l *LockUnlock) Code(code string) *LockUnlock {
	l.ServiceData.Code = &code
	return l
}
func (l *LockUnlock) JSON() string {
	data, _ := gojson.Marshal(l)
	return string(data)
}
func (l *LockUnlock) Name() string {
	return fmt.Sprintf("%s.%s", *l.Domain, *l.Service)
}
func (l *LockUnlock) SetID(id *int) {
	l.Id = id
}
