// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/nspcc-dev/neofs-rest-gw/gen/models"
)

// AuthOKCode is the HTTP code returned for type AuthOK
const AuthOKCode int = 200

/*AuthOK Base64 encoded stable binary marshaled bearer token bodies.

swagger:response authOK
*/
type AuthOK struct {
	/*

	 */
	AccessControlAllowOrigin string `json:"Access-Control-Allow-Origin"`

	/*
	  In: Body
	*/
	Payload []*models.TokenResponse `json:"body,omitempty"`
}

// NewAuthOK creates AuthOK with default headers values
func NewAuthOK() *AuthOK {

	return &AuthOK{}
}

// WithAccessControlAllowOrigin adds the accessControlAllowOrigin to the auth o k response
func (o *AuthOK) WithAccessControlAllowOrigin(accessControlAllowOrigin string) *AuthOK {
	o.AccessControlAllowOrigin = accessControlAllowOrigin
	return o
}

// SetAccessControlAllowOrigin sets the accessControlAllowOrigin to the auth o k response
func (o *AuthOK) SetAccessControlAllowOrigin(accessControlAllowOrigin string) {
	o.AccessControlAllowOrigin = accessControlAllowOrigin
}

// WithPayload adds the payload to the auth o k response
func (o *AuthOK) WithPayload(payload []*models.TokenResponse) *AuthOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the auth o k response
func (o *AuthOK) SetPayload(payload []*models.TokenResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AuthOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Access-Control-Allow-Origin

	accessControlAllowOrigin := o.AccessControlAllowOrigin
	if accessControlAllowOrigin != "" {
		rw.Header().Set("Access-Control-Allow-Origin", accessControlAllowOrigin)
	}

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.TokenResponse, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// AuthBadRequestCode is the HTTP code returned for type AuthBadRequest
const AuthBadRequestCode int = 400

/*AuthBadRequest Bad request

swagger:response authBadRequest
*/
type AuthBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewAuthBadRequest creates AuthBadRequest with default headers values
func NewAuthBadRequest() *AuthBadRequest {

	return &AuthBadRequest{}
}

// WithPayload adds the payload to the auth bad request response
func (o *AuthBadRequest) WithPayload(payload *models.ErrorResponse) *AuthBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the auth bad request response
func (o *AuthBadRequest) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AuthBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
