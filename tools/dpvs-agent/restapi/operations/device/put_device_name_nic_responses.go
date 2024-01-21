// Code generated by go-swagger; DO NOT EDIT.

package device

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// PutDeviceNameNicOKCode is the HTTP code returned for type PutDeviceNameNicOK
const PutDeviceNameNicOKCode int = 200

/*
PutDeviceNameNicOK Success

swagger:response putDeviceNameNicOK
*/
type PutDeviceNameNicOK struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewPutDeviceNameNicOK creates PutDeviceNameNicOK with default headers values
func NewPutDeviceNameNicOK() *PutDeviceNameNicOK {

	return &PutDeviceNameNicOK{}
}

// WithPayload adds the payload to the put device name nic o k response
func (o *PutDeviceNameNicOK) WithPayload(payload string) *PutDeviceNameNicOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the put device name nic o k response
func (o *PutDeviceNameNicOK) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PutDeviceNameNicOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// PutDeviceNameNicInternalServerErrorCode is the HTTP code returned for type PutDeviceNameNicInternalServerError
const PutDeviceNameNicInternalServerErrorCode int = 500

/*
PutDeviceNameNicInternalServerError Failure

swagger:response putDeviceNameNicInternalServerError
*/
type PutDeviceNameNicInternalServerError struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewPutDeviceNameNicInternalServerError creates PutDeviceNameNicInternalServerError with default headers values
func NewPutDeviceNameNicInternalServerError() *PutDeviceNameNicInternalServerError {

	return &PutDeviceNameNicInternalServerError{}
}

// WithPayload adds the payload to the put device name nic internal server error response
func (o *PutDeviceNameNicInternalServerError) WithPayload(payload string) *PutDeviceNameNicInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the put device name nic internal server error response
func (o *PutDeviceNameNicInternalServerError) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PutDeviceNameNicInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}
