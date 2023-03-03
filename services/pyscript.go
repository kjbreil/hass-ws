package services

import "encoding/json"

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

// NewPyscriptJupyterKernelStart creates the object that can be sent to Home Assistant for domain pyscript, service jupyter_kernel_start
// "Starts a jupyter kernel for interactive use; Called by Jupyter front end and should generally not be used by users"
func NewPyscriptJupyterKernelStart(target Target, pyscriptJupyterKernelStartParams PyscriptJupyterKernelStartParams) *PyscriptJupyterKernelStart {
	serviceDomain := "pyscript"
	serviceType := "call_service"
	serviceService := "jupyter_kernel_start"
	p := &PyscriptJupyterKernelStart{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: pyscriptJupyterKernelStartParams,
	}
	return p
}

type PyscriptJupyterKernelStart struct {
	ServiceBase
	ServiceData PyscriptJupyterKernelStartParams `json:"service_data,omitempty"`
}
type PyscriptJupyterKernelStartParams struct {
	ControlPort     *float64         `json:"control_port,omitempty"`
	HbPort          *float64         `json:"hb_port,omitempty"`
	IopubPort       *float64         `json:"iopub_port,omitempty"`
	Ip              *string          `json:"ip,omitempty"`
	KernelName      *string          `json:"kernel_name,omitempty"`
	Key             *string          `json:"key,omitempty"`
	ShellPort       *float64         `json:"shell_port,omitempty"`
	SignatureScheme *SignatureScheme `json:"signature_scheme,omitempty"`
	StdinPort       *float64         `json:"stdin_port,omitempty"`
	Transport       *Transport       `json:"transport,omitempty"`
}

func (p *PyscriptJupyterKernelStart) JSON() string {
	data, _ := json.Marshal(p)
	return string(data)
}
func (p *PyscriptJupyterKernelStart) SetID(id *int) {
	p.Id = id
}

// NewPyscriptReload creates the object that can be sent to Home Assistant for domain pyscript, service reload
// "Reloads all available pyscripts and restart triggers"
func NewPyscriptReload(target Target, pyscriptReloadParams PyscriptReloadParams) *PyscriptReload {
	serviceDomain := "pyscript"
	serviceType := "call_service"
	serviceService := "reload"
	p := &PyscriptReload{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: pyscriptReloadParams,
	}
	return p
}

type PyscriptReload struct {
	ServiceBase
	ServiceData PyscriptReloadParams `json:"service_data,omitempty"`
}
type PyscriptReloadParams struct {
	GlobalCtx *string `json:"global_ctx,omitempty"`
}

func (p *PyscriptReload) JSON() string {
	data, _ := json.Marshal(p)
	return string(data)
}
func (p *PyscriptReload) SetID(id *int) {
	p.Id = id
}
