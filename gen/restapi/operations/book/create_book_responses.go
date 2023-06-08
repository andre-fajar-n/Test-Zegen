// Code generated by go-swagger; DO NOT EDIT.

package book

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"zegen/gen/models"
)

// CreateBookCreatedCode is the HTTP code returned for type CreateBookCreated
const CreateBookCreatedCode int = 201

/*
CreateBookCreated Success create

swagger:response createBookCreated
*/
type CreateBookCreated struct {

	/*
	  In: Body
	*/
	Payload *models.SuccessCreateBook `json:"body,omitempty"`
}

// NewCreateBookCreated creates CreateBookCreated with default headers values
func NewCreateBookCreated() *CreateBookCreated {

	return &CreateBookCreated{}
}

// WithPayload adds the payload to the create book created response
func (o *CreateBookCreated) WithPayload(payload *models.SuccessCreateBook) *CreateBookCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create book created response
func (o *CreateBookCreated) SetPayload(payload *models.SuccessCreateBook) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateBookCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*
CreateBookDefault Server Error

swagger:response createBookDefault
*/
type CreateBookDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCreateBookDefault creates CreateBookDefault with default headers values
func NewCreateBookDefault(code int) *CreateBookDefault {
	if code <= 0 {
		code = 500
	}

	return &CreateBookDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the create book default response
func (o *CreateBookDefault) WithStatusCode(code int) *CreateBookDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the create book default response
func (o *CreateBookDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the create book default response
func (o *CreateBookDefault) WithPayload(payload *models.Error) *CreateBookDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create book default response
func (o *CreateBookDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateBookDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
