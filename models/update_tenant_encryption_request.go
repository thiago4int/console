// Code generated by go-swagger; DO NOT EDIT.

// This file is part of MinIO Console Server
// Copyright (c) 2021 MinIO, Inc.
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.
//

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// UpdateTenantEncryptionRequest update tenant encryption request
//
// swagger:model updateTenantEncryptionRequest
type UpdateTenantEncryptionRequest struct {

	// auto cert
	AutoCert bool `json:"autoCert,omitempty"`

	// custom certificates
	CustomCertificates *UpdateTenantEncryptionRequestCustomCertificates `json:"customCertificates,omitempty"`
}

// Validate validates this update tenant encryption request
func (m *UpdateTenantEncryptionRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCustomCertificates(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpdateTenantEncryptionRequest) validateCustomCertificates(formats strfmt.Registry) error {
	if swag.IsZero(m.CustomCertificates) { // not required
		return nil
	}

	if m.CustomCertificates != nil {
		if err := m.CustomCertificates.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("customCertificates")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("customCertificates")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this update tenant encryption request based on the context it is used
func (m *UpdateTenantEncryptionRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCustomCertificates(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpdateTenantEncryptionRequest) contextValidateCustomCertificates(ctx context.Context, formats strfmt.Registry) error {

	if m.CustomCertificates != nil {
		if err := m.CustomCertificates.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("customCertificates")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("customCertificates")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *UpdateTenantEncryptionRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateTenantEncryptionRequest) UnmarshalBinary(b []byte) error {
	var res UpdateTenantEncryptionRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// UpdateTenantEncryptionRequestCustomCertificates update tenant encryption request custom certificates
//
// swagger:model UpdateTenantEncryptionRequestCustomCertificates
type UpdateTenantEncryptionRequestCustomCertificates struct {

	// minio
	Minio []*KeyPairConfiguration `json:"minio"`

	// minio c as
	MinioCAs []string `json:"minioCAs"`

	// secrets to be deleted
	SecretsToBeDeleted []string `json:"secretsToBeDeleted"`
}

// Validate validates this update tenant encryption request custom certificates
func (m *UpdateTenantEncryptionRequestCustomCertificates) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMinio(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpdateTenantEncryptionRequestCustomCertificates) validateMinio(formats strfmt.Registry) error {
	if swag.IsZero(m.Minio) { // not required
		return nil
	}

	for i := 0; i < len(m.Minio); i++ {
		if swag.IsZero(m.Minio[i]) { // not required
			continue
		}

		if m.Minio[i] != nil {
			if err := m.Minio[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("customCertificates" + "." + "minio" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("customCertificates" + "." + "minio" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this update tenant encryption request custom certificates based on the context it is used
func (m *UpdateTenantEncryptionRequestCustomCertificates) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateMinio(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpdateTenantEncryptionRequestCustomCertificates) contextValidateMinio(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Minio); i++ {

		if m.Minio[i] != nil {
			if err := m.Minio[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("customCertificates" + "." + "minio" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("customCertificates" + "." + "minio" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *UpdateTenantEncryptionRequestCustomCertificates) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateTenantEncryptionRequestCustomCertificates) UnmarshalBinary(b []byte) error {
	var res UpdateTenantEncryptionRequestCustomCertificates
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
