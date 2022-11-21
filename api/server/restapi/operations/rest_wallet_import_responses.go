// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/massalabs/thyra-plugin-massa-core/api/server/models"
)

// RestWalletImportNoContentCode is the HTTP code returned for type RestWalletImportNoContent
const RestWalletImportNoContentCode int = 204

/*RestWalletImportNoContent Wallet imported.

swagger:response restWalletImportNoContent
*/
type RestWalletImportNoContent struct {
}

// NewRestWalletImportNoContent creates RestWalletImportNoContent with default headers values
func NewRestWalletImportNoContent() *RestWalletImportNoContent {

	return &RestWalletImportNoContent{}
}

// WriteResponse to the client
func (o *RestWalletImportNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

// RestWalletImportBadRequestCode is the HTTP code returned for type RestWalletImportBadRequest
const RestWalletImportBadRequestCode int = 400

/*RestWalletImportBadRequest Bad request.

swagger:response restWalletImportBadRequest
*/
type RestWalletImportBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewRestWalletImportBadRequest creates RestWalletImportBadRequest with default headers values
func NewRestWalletImportBadRequest() *RestWalletImportBadRequest {

	return &RestWalletImportBadRequest{}
}

// WithPayload adds the payload to the rest wallet import bad request response
func (o *RestWalletImportBadRequest) WithPayload(payload *models.Error) *RestWalletImportBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the rest wallet import bad request response
func (o *RestWalletImportBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RestWalletImportBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// RestWalletImportUnprocessableEntityCode is the HTTP code returned for type RestWalletImportUnprocessableEntity
const RestWalletImportUnprocessableEntityCode int = 422

/*RestWalletImportUnprocessableEntity Unprocessable Entity - syntax is correct, but the server was unable to process the contained instructions.

swagger:response restWalletImportUnprocessableEntity
*/
type RestWalletImportUnprocessableEntity struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewRestWalletImportUnprocessableEntity creates RestWalletImportUnprocessableEntity with default headers values
func NewRestWalletImportUnprocessableEntity() *RestWalletImportUnprocessableEntity {

	return &RestWalletImportUnprocessableEntity{}
}

// WithPayload adds the payload to the rest wallet import unprocessable entity response
func (o *RestWalletImportUnprocessableEntity) WithPayload(payload *models.Error) *RestWalletImportUnprocessableEntity {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the rest wallet import unprocessable entity response
func (o *RestWalletImportUnprocessableEntity) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RestWalletImportUnprocessableEntity) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(422)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// RestWalletImportInternalServerErrorCode is the HTTP code returned for type RestWalletImportInternalServerError
const RestWalletImportInternalServerErrorCode int = 500

/*RestWalletImportInternalServerError Internal Server Error - The server has encountered a situation it does not know how to handle.

swagger:response restWalletImportInternalServerError
*/
type RestWalletImportInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewRestWalletImportInternalServerError creates RestWalletImportInternalServerError with default headers values
func NewRestWalletImportInternalServerError() *RestWalletImportInternalServerError {

	return &RestWalletImportInternalServerError{}
}

// WithPayload adds the payload to the rest wallet import internal server error response
func (o *RestWalletImportInternalServerError) WithPayload(payload *models.Error) *RestWalletImportInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the rest wallet import internal server error response
func (o *RestWalletImportInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RestWalletImportInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}