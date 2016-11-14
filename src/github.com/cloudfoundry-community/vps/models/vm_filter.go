package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
)

// VMFilter Vm filter
// swagger:model VmFilter
type VMFilter struct {

	// cid
	Cid int32 `json:"cid,omitempty"`

	// cpu
	CPU int32 `json:"cpu,omitempty"`

	// ip
	IP strfmt.IPv4 `json:"ip,omitempty"`

	// memory mb
	MemoryMb int32 `json:"memory_mb,omitempty"`

	// private vlan
	PrivateVlan int32 `json:"private_vlan,omitempty"`

	// public vlan
	PublicVlan int32 `json:"public_vlan,omitempty"`

	// state
	State State `json:"state,omitempty"`

	// deployment name
	DeploymentName string `json:"deploymentName,omitempty"`
}

// Validate validates this Vm filter
func (m *VMFilter) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateState(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VMFilter) validateState(formats strfmt.Registry) error {

	if swag.IsZero(m.State) { // not required
		return nil
	}

	if err := m.State.Validate(formats); err != nil {
		return err
	}

	return nil
}
