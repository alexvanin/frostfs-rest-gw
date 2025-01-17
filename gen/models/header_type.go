// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// HeaderType Enumeration of possible sources of Headers to apply filters in NeoFS EACL.
//
// swagger:model HeaderType
type HeaderType string

func NewHeaderType(value HeaderType) *HeaderType {
	return &value
}

// Pointer returns a pointer to a freshly-allocated HeaderType.
func (m HeaderType) Pointer() *HeaderType {
	return &m
}

const (

	// HeaderTypeREQUEST captures enum value "REQUEST"
	HeaderTypeREQUEST HeaderType = "REQUEST"

	// HeaderTypeOBJECT captures enum value "OBJECT"
	HeaderTypeOBJECT HeaderType = "OBJECT"

	// HeaderTypeSERVICE captures enum value "SERVICE"
	HeaderTypeSERVICE HeaderType = "SERVICE"
)

// for schema
var headerTypeEnum []interface{}

func init() {
	var res []HeaderType
	if err := json.Unmarshal([]byte(`["REQUEST","OBJECT","SERVICE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		headerTypeEnum = append(headerTypeEnum, v)
	}
}

func (m HeaderType) validateHeaderTypeEnum(path, location string, value HeaderType) error {
	if err := validate.EnumCase(path, location, value, headerTypeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this header type
func (m HeaderType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateHeaderTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this header type based on context it is used
func (m HeaderType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
