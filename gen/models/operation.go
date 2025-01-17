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

// Operation Request's operation type to match in NeoFS EACL if the rule is applicable to a particular request.
//
// swagger:model Operation
type Operation string

func NewOperation(value Operation) *Operation {
	return &value
}

// Pointer returns a pointer to a freshly-allocated Operation.
func (m Operation) Pointer() *Operation {
	return &m
}

const (

	// OperationGET captures enum value "GET"
	OperationGET Operation = "GET"

	// OperationHEAD captures enum value "HEAD"
	OperationHEAD Operation = "HEAD"

	// OperationPUT captures enum value "PUT"
	OperationPUT Operation = "PUT"

	// OperationDELETE captures enum value "DELETE"
	OperationDELETE Operation = "DELETE"

	// OperationSEARCH captures enum value "SEARCH"
	OperationSEARCH Operation = "SEARCH"

	// OperationRANGE captures enum value "RANGE"
	OperationRANGE Operation = "RANGE"

	// OperationRANGEHASH captures enum value "RANGEHASH"
	OperationRANGEHASH Operation = "RANGEHASH"
)

// for schema
var operationEnum []interface{}

func init() {
	var res []Operation
	if err := json.Unmarshal([]byte(`["GET","HEAD","PUT","DELETE","SEARCH","RANGE","RANGEHASH"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		operationEnum = append(operationEnum, v)
	}
}

func (m Operation) validateOperationEnum(path, location string, value Operation) error {
	if err := validate.EnumCase(path, location, value, operationEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this operation
func (m Operation) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateOperationEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this operation based on context it is used
func (m Operation) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
