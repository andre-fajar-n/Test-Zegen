// Code generated by go-swagger; DO NOT EDIT.

package authentication

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"zegen/gen/models"
)

// RegisterOKCode is the HTTP code returned for type RegisterOK
const RegisterOKCode int = 200

/*
RegisterOK Success register

swagger:response registerOK
*/
type RegisterOK struct {

	/*
	  In: Body
	*/
	Payload *models.SuccessRegister `json:"body,omitempty"`
}

// NewRegisterOK creates RegisterOK with default headers values
func NewRegisterOK() *RegisterOK {

	return &RegisterOK{}
}

// WithPayload adds the payload to the register o k response
func (o *RegisterOK) WithPayload(payload *models.SuccessRegister) *RegisterOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the register o k response
func (o *RegisterOK) SetPayload(payload *models.SuccessRegister) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RegisterOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*
RegisterDefault Server Error

swagger:response registerDefault
*/
type RegisterDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewRegisterDefault creates RegisterDefault with default headers values
func NewRegisterDefault(code int) *RegisterDefault {
	if code <= 0 {
		code = 500
	}

	return &RegisterDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the register default response
func (o *RegisterDefault) WithStatusCode(code int) *RegisterDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the register default response
func (o *RegisterDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the register default response
func (o *RegisterDefault) WithPayload(payload *models.Error) *RegisterDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the register default response
func (o *RegisterDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RegisterDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
