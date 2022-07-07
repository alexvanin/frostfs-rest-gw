// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// TokenResponse token response
// Example: [{"token":"sometoken-todo-add","type":"object"},{"token":"ChCpanIBJCpJuJz42KOmGMSnEhsKGTWquaX2Lq6GhhO4faOYkLD0f9WkXuYJlq4aBAhnGAMiIQJgFcIEghQB5lq3AJZOVswInwc1IGhlQ7NCUh4DFO3UATIECAEQAQ==","type":"container"}]
//
// swagger:model TokenResponse
type TokenResponse struct {

	// name
	Name string `json:"name,omitempty"`

	// token
	// Required: true
	Token *string `json:"token"`

	// type
	// Required: true
	Type *TokenType `json:"type"`
}

// Validate validates this token response
func (m *TokenResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateToken(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TokenResponse) validateToken(formats strfmt.Registry) error {

	if err := validate.Required("token", "body", m.Token); err != nil {
		return err
	}

	return nil
}

func (m *TokenResponse) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	if m.Type != nil {
		if err := m.Type.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("type")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("type")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this token response based on the context it is used
func (m *TokenResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TokenResponse) contextValidateType(ctx context.Context, formats strfmt.Registry) error {

	if m.Type != nil {
		if err := m.Type.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("type")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("type")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TokenResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TokenResponse) UnmarshalBinary(b []byte) error {
	var res TokenResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
