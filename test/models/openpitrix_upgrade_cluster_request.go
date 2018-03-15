// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// OpenpitrixUpgradeClusterRequest openpitrix upgrade cluster request
// swagger:model openpitrixUpgradeClusterRequest
type OpenpitrixUpgradeClusterRequest struct {

	// advanced param
	AdvancedParam []string `json:"advanced_param"`

	// app version
	AppVersion string `json:"app_version,omitempty"`

	// cluster id
	ClusterID string `json:"cluster_id,omitempty"`
}

// Validate validates this openpitrix upgrade cluster request
func (m *OpenpitrixUpgradeClusterRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAdvancedParam(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OpenpitrixUpgradeClusterRequest) validateAdvancedParam(formats strfmt.Registry) error {

	if swag.IsZero(m.AdvancedParam) { // not required
		return nil
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OpenpitrixUpgradeClusterRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenpitrixUpgradeClusterRequest) UnmarshalBinary(b []byte) error {
	var res OpenpitrixUpgradeClusterRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
