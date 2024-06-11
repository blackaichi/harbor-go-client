// Code generated by go-swagger; DO NOT EDIT.

package registry

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/blackaichi/harbor-go-client/pkg/sdk/v2.0/models"
)

// CreateRegistryReader is a Reader for the CreateRegistry structure.
type CreateRegistryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateRegistryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateRegistryCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateRegistryBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCreateRegistryUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateRegistryForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewCreateRegistryConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateRegistryInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateRegistryCreated creates a CreateRegistryCreated with default headers values
func NewCreateRegistryCreated() *CreateRegistryCreated {
	return &CreateRegistryCreated{}
}

/*
CreateRegistryCreated describes a response with status code 201, with default header values.

Created
*/
type CreateRegistryCreated struct {

	/* The location of the resource
	 */
	Location string

	/* The ID of the corresponding request for the response
	 */
	XRequestID string
}

// IsSuccess returns true when this create registry created response has a 2xx status code
func (o *CreateRegistryCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create registry created response has a 3xx status code
func (o *CreateRegistryCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create registry created response has a 4xx status code
func (o *CreateRegistryCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create registry created response has a 5xx status code
func (o *CreateRegistryCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create registry created response a status code equal to that given
func (o *CreateRegistryCreated) IsCode(code int) bool {
	return code == 201
}

func (o *CreateRegistryCreated) Error() string {
	return fmt.Sprintf("[POST /registries][%d] createRegistryCreated ", 201)
}

func (o *CreateRegistryCreated) String() string {
	return fmt.Sprintf("[POST /registries][%d] createRegistryCreated ", 201)
}

func (o *CreateRegistryCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Location
	hdrLocation := response.GetHeader("Location")

	if hdrLocation != "" {
		o.Location = hdrLocation
	}

	// hydrates response header X-Request-Id
	hdrXRequestID := response.GetHeader("X-Request-Id")

	if hdrXRequestID != "" {
		o.XRequestID = hdrXRequestID
	}

	return nil
}

// NewCreateRegistryBadRequest creates a CreateRegistryBadRequest with default headers values
func NewCreateRegistryBadRequest() *CreateRegistryBadRequest {
	return &CreateRegistryBadRequest{}
}

/*
CreateRegistryBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type CreateRegistryBadRequest struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this create registry bad request response has a 2xx status code
func (o *CreateRegistryBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create registry bad request response has a 3xx status code
func (o *CreateRegistryBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create registry bad request response has a 4xx status code
func (o *CreateRegistryBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create registry bad request response has a 5xx status code
func (o *CreateRegistryBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create registry bad request response a status code equal to that given
func (o *CreateRegistryBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *CreateRegistryBadRequest) Error() string {
	return fmt.Sprintf("[POST /registries][%d] createRegistryBadRequest  %+v", 400, o.Payload)
}

func (o *CreateRegistryBadRequest) String() string {
	return fmt.Sprintf("[POST /registries][%d] createRegistryBadRequest  %+v", 400, o.Payload)
}

func (o *CreateRegistryBadRequest) GetPayload() *models.Errors {
	return o.Payload
}

func (o *CreateRegistryBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-Request-Id
	hdrXRequestID := response.GetHeader("X-Request-Id")

	if hdrXRequestID != "" {
		o.XRequestID = hdrXRequestID
	}

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateRegistryUnauthorized creates a CreateRegistryUnauthorized with default headers values
func NewCreateRegistryUnauthorized() *CreateRegistryUnauthorized {
	return &CreateRegistryUnauthorized{}
}

/*
CreateRegistryUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type CreateRegistryUnauthorized struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this create registry unauthorized response has a 2xx status code
func (o *CreateRegistryUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create registry unauthorized response has a 3xx status code
func (o *CreateRegistryUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create registry unauthorized response has a 4xx status code
func (o *CreateRegistryUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this create registry unauthorized response has a 5xx status code
func (o *CreateRegistryUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this create registry unauthorized response a status code equal to that given
func (o *CreateRegistryUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *CreateRegistryUnauthorized) Error() string {
	return fmt.Sprintf("[POST /registries][%d] createRegistryUnauthorized  %+v", 401, o.Payload)
}

func (o *CreateRegistryUnauthorized) String() string {
	return fmt.Sprintf("[POST /registries][%d] createRegistryUnauthorized  %+v", 401, o.Payload)
}

func (o *CreateRegistryUnauthorized) GetPayload() *models.Errors {
	return o.Payload
}

func (o *CreateRegistryUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-Request-Id
	hdrXRequestID := response.GetHeader("X-Request-Id")

	if hdrXRequestID != "" {
		o.XRequestID = hdrXRequestID
	}

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateRegistryForbidden creates a CreateRegistryForbidden with default headers values
func NewCreateRegistryForbidden() *CreateRegistryForbidden {
	return &CreateRegistryForbidden{}
}

/*
CreateRegistryForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type CreateRegistryForbidden struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this create registry forbidden response has a 2xx status code
func (o *CreateRegistryForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create registry forbidden response has a 3xx status code
func (o *CreateRegistryForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create registry forbidden response has a 4xx status code
func (o *CreateRegistryForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this create registry forbidden response has a 5xx status code
func (o *CreateRegistryForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this create registry forbidden response a status code equal to that given
func (o *CreateRegistryForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *CreateRegistryForbidden) Error() string {
	return fmt.Sprintf("[POST /registries][%d] createRegistryForbidden  %+v", 403, o.Payload)
}

func (o *CreateRegistryForbidden) String() string {
	return fmt.Sprintf("[POST /registries][%d] createRegistryForbidden  %+v", 403, o.Payload)
}

func (o *CreateRegistryForbidden) GetPayload() *models.Errors {
	return o.Payload
}

func (o *CreateRegistryForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-Request-Id
	hdrXRequestID := response.GetHeader("X-Request-Id")

	if hdrXRequestID != "" {
		o.XRequestID = hdrXRequestID
	}

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateRegistryConflict creates a CreateRegistryConflict with default headers values
func NewCreateRegistryConflict() *CreateRegistryConflict {
	return &CreateRegistryConflict{}
}

/*
CreateRegistryConflict describes a response with status code 409, with default header values.

Conflict
*/
type CreateRegistryConflict struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this create registry conflict response has a 2xx status code
func (o *CreateRegistryConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create registry conflict response has a 3xx status code
func (o *CreateRegistryConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create registry conflict response has a 4xx status code
func (o *CreateRegistryConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this create registry conflict response has a 5xx status code
func (o *CreateRegistryConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this create registry conflict response a status code equal to that given
func (o *CreateRegistryConflict) IsCode(code int) bool {
	return code == 409
}

func (o *CreateRegistryConflict) Error() string {
	return fmt.Sprintf("[POST /registries][%d] createRegistryConflict  %+v", 409, o.Payload)
}

func (o *CreateRegistryConflict) String() string {
	return fmt.Sprintf("[POST /registries][%d] createRegistryConflict  %+v", 409, o.Payload)
}

func (o *CreateRegistryConflict) GetPayload() *models.Errors {
	return o.Payload
}

func (o *CreateRegistryConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-Request-Id
	hdrXRequestID := response.GetHeader("X-Request-Id")

	if hdrXRequestID != "" {
		o.XRequestID = hdrXRequestID
	}

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateRegistryInternalServerError creates a CreateRegistryInternalServerError with default headers values
func NewCreateRegistryInternalServerError() *CreateRegistryInternalServerError {
	return &CreateRegistryInternalServerError{}
}

/*
CreateRegistryInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type CreateRegistryInternalServerError struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this create registry internal server error response has a 2xx status code
func (o *CreateRegistryInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create registry internal server error response has a 3xx status code
func (o *CreateRegistryInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create registry internal server error response has a 4xx status code
func (o *CreateRegistryInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this create registry internal server error response has a 5xx status code
func (o *CreateRegistryInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this create registry internal server error response a status code equal to that given
func (o *CreateRegistryInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *CreateRegistryInternalServerError) Error() string {
	return fmt.Sprintf("[POST /registries][%d] createRegistryInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateRegistryInternalServerError) String() string {
	return fmt.Sprintf("[POST /registries][%d] createRegistryInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateRegistryInternalServerError) GetPayload() *models.Errors {
	return o.Payload
}

func (o *CreateRegistryInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-Request-Id
	hdrXRequestID := response.GetHeader("X-Request-Id")

	if hdrXRequestID != "" {
		o.XRequestID = hdrXRequestID
	}

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
