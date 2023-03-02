package services

import "encoding/json"

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

// NewPyscriptJupyterKernelStart creates the object that can be sent to Home Assistant for domain pyscript, service jupyter_kernel_start
// "Starts a jupyter kernel for interactive use; Called by Jupyter front end and should generally not be used by users"
func NewPyscriptJupyterKernelStart(entities []string, controlPort *int, hbPort *int, iopubPort *int, ip *string, kernelName *string, key *string, shellPort *int, signatureScheme *SignatureScheme, stdinPort *int, transport *Transport) *PyscriptJupyterKernelStart {
	serviceDomain := "pyscript"
	serviceType := "call_service"
	serviceService := "jupyter_kernel_start"
	p := &PyscriptJupyterKernelStart{
		Domain:  &serviceDomain,
		Id:      nil,
		Service: &serviceService,
		ServiceData: struct {
			ControlPort     *int             `json:"control_port,omitempty"`
			HbPort          *int             `json:"hb_port,omitempty"`
			IopubPort       *int             `json:"iopub_port,omitempty"`
			Ip              *string          `json:"ip,omitempty"`
			KernelName      *string          `json:"kernel_name,omitempty"`
			Key             *string          `json:"key,omitempty"`
			ShellPort       *int             `json:"shell_port,omitempty"`
			SignatureScheme *SignatureScheme `json:"signature_scheme,omitempty"`
			StdinPort       *int             `json:"stdin_port,omitempty"`
			Transport       *Transport       `json:"transport,omitempty"`
		}{
			ControlPort:     controlPort,
			HbPort:          hbPort,
			IopubPort:       iopubPort,
			Ip:              ip,
			KernelName:      kernelName,
			Key:             key,
			ShellPort:       shellPort,
			SignatureScheme: signatureScheme,
			StdinPort:       stdinPort,
			Transport:       transport,
		},
		Target: struct {
			EntityId []string `json:"entity_id,omitempty"`
		}{EntityId: entities},
		Type: &serviceType,
	}
	return p
}

type PyscriptJupyterKernelStart struct {
	Id          *int    `json:"id"`
	Type        *string `json:"type"`
	Domain      *string `json:"domain"`
	Service     *string `json:"service"`
	ServiceData struct {
		ControlPort     *int             `json:"control_port,omitempty"`
		HbPort          *int             `json:"hb_port,omitempty"`
		IopubPort       *int             `json:"iopub_port,omitempty"`
		Ip              *string          `json:"ip,omitempty"`
		KernelName      *string          `json:"kernel_name,omitempty"`
		Key             *string          `json:"key,omitempty"`
		ShellPort       *int             `json:"shell_port,omitempty"`
		SignatureScheme *SignatureScheme `json:"signature_scheme,omitempty"`
		StdinPort       *int             `json:"stdin_port,omitempty"`
		Transport       *Transport       `json:"transport,omitempty"`
	} `json:"service_data,omitempty"`
	Target struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
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
func NewPyscriptReload(entities []string, globalCtx *string) *PyscriptReload {
	serviceDomain := "pyscript"
	serviceType := "call_service"
	serviceService := "reload"
	p := &PyscriptReload{
		Domain:  &serviceDomain,
		Id:      nil,
		Service: &serviceService,
		ServiceData: struct {
			GlobalCtx *string `json:"global_ctx,omitempty"`
		}{GlobalCtx: globalCtx},
		Target: struct {
			EntityId []string `json:"entity_id,omitempty"`
		}{EntityId: entities},
		Type: &serviceType,
	}
	return p
}

type PyscriptReload struct {
	Id          *int    `json:"id"`
	Type        *string `json:"type"`
	Domain      *string `json:"domain"`
	Service     *string `json:"service"`
	ServiceData struct {
		GlobalCtx *string `json:"global_ctx,omitempty"`
	} `json:"service_data,omitempty"`
	Target struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (p *PyscriptReload) JSON() string {
	data, _ := json.Marshal(p)
	return string(data)
}
func (p *PyscriptReload) SetID(id *int) {
	p.Id = id
}
