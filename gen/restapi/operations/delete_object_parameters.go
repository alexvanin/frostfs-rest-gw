// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// NewDeleteObjectParams creates a new DeleteObjectParams object
// with the default values initialized.
func NewDeleteObjectParams() DeleteObjectParams {

	var (
		// initialize parameters with default values

		walletConnectDefault = bool(false)
	)

	return DeleteObjectParams{
		WalletConnect: &walletConnectDefault,
	}
}

// DeleteObjectParams contains all the bound params for the delete object operation
// typically these are obtained from a http.Request
//
// swagger:parameters deleteObject
type DeleteObjectParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Base64 encoded signature for bearer token
	  Required: true
	  In: header
	*/
	XBearerSignature string
	/*Hex encoded the public part of the key that signed the bearer token
	  Required: true
	  In: header
	*/
	XBearerSignatureKey string
	/*Base58 encoded container id
	  Required: true
	  In: path
	*/
	ContainerID string
	/*Base58 encoded object id
	  Required: true
	  In: path
	*/
	ObjectID string
	/*Use wallect connect signature scheme or not
	  In: query
	  Default: false
	*/
	WalletConnect *bool
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewDeleteObjectParams() beforehand.
func (o *DeleteObjectParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	if err := o.bindXBearerSignature(r.Header[http.CanonicalHeaderKey("X-Bearer-Signature")], true, route.Formats); err != nil {
		res = append(res, err)
	}

	if err := o.bindXBearerSignatureKey(r.Header[http.CanonicalHeaderKey("X-Bearer-Signature-Key")], true, route.Formats); err != nil {
		res = append(res, err)
	}

	rContainerID, rhkContainerID, _ := route.Params.GetOK("containerId")
	if err := o.bindContainerID(rContainerID, rhkContainerID, route.Formats); err != nil {
		res = append(res, err)
	}

	rObjectID, rhkObjectID, _ := route.Params.GetOK("objectId")
	if err := o.bindObjectID(rObjectID, rhkObjectID, route.Formats); err != nil {
		res = append(res, err)
	}

	qWalletConnect, qhkWalletConnect, _ := qs.GetOK("walletConnect")
	if err := o.bindWalletConnect(qWalletConnect, qhkWalletConnect, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindXBearerSignature binds and validates parameter XBearerSignature from header.
func (o *DeleteObjectParams) bindXBearerSignature(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("X-Bearer-Signature", "header", rawData)
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true

	if err := validate.RequiredString("X-Bearer-Signature", "header", raw); err != nil {
		return err
	}
	o.XBearerSignature = raw

	return nil
}

// bindXBearerSignatureKey binds and validates parameter XBearerSignatureKey from header.
func (o *DeleteObjectParams) bindXBearerSignatureKey(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("X-Bearer-Signature-Key", "header", rawData)
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true

	if err := validate.RequiredString("X-Bearer-Signature-Key", "header", raw); err != nil {
		return err
	}
	o.XBearerSignatureKey = raw

	return nil
}

// bindContainerID binds and validates parameter ContainerID from path.
func (o *DeleteObjectParams) bindContainerID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route
	o.ContainerID = raw

	return nil
}

// bindObjectID binds and validates parameter ObjectID from path.
func (o *DeleteObjectParams) bindObjectID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route
	o.ObjectID = raw

	return nil
}

// bindWalletConnect binds and validates parameter WalletConnect from query.
func (o *DeleteObjectParams) bindWalletConnect(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		// Default values have been previously initialized by NewDeleteObjectParams()
		return nil
	}

	value, err := swag.ConvertBool(raw)
	if err != nil {
		return errors.InvalidType("walletConnect", "query", "bool", raw)
	}
	o.WalletConnect = &value

	return nil
}