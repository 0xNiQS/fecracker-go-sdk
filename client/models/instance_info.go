// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2018 Amazon.com, Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
// 	http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// InstanceInfo Describes MicroVM instance information.
// swagger:model InstanceInfo
type InstanceInfo struct {

	// MicroVM / instance ID.
	// Required: true
	ID *string `json:"id"`

	// The current detailed state of the Firecracker instance. This value is read-only for the control-plane.
	// Required: true
	// Enum: [Uninitialized Starting Running Halting Halted]
	State *string `json:"state"`

	// MicroVM hypervisor build version.
	// Required: true
	VmmVersion *string `json:"vmm_version"`
}

// Validate validates this instance info
func (m *InstanceInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVmmVersion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InstanceInfo) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

var instanceInfoTypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Uninitialized","Starting","Running","Halting","Halted"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		instanceInfoTypeStatePropEnum = append(instanceInfoTypeStatePropEnum, v)
	}
}

const (

	// InstanceInfoStateUninitialized captures enum value "Uninitialized"
	InstanceInfoStateUninitialized string = "Uninitialized"

	// InstanceInfoStateStarting captures enum value "Starting"
	InstanceInfoStateStarting string = "Starting"

	// InstanceInfoStateRunning captures enum value "Running"
	InstanceInfoStateRunning string = "Running"

	// InstanceInfoStateHalting captures enum value "Halting"
	InstanceInfoStateHalting string = "Halting"

	// InstanceInfoStateHalted captures enum value "Halted"
	InstanceInfoStateHalted string = "Halted"
)

// prop value enum
func (m *InstanceInfo) validateStateEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, instanceInfoTypeStatePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *InstanceInfo) validateState(formats strfmt.Registry) error {

	if err := validate.Required("state", "body", m.State); err != nil {
		return err
	}

	// value enum
	if err := m.validateStateEnum("state", "body", *m.State); err != nil {
		return err
	}

	return nil
}

func (m *InstanceInfo) validateVmmVersion(formats strfmt.Registry) error {

	if err := validate.Required("vmm_version", "body", m.VmmVersion); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *InstanceInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InstanceInfo) UnmarshalBinary(b []byte) error {
	var res InstanceInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
