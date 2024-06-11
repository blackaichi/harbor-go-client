// Code generated by go-swagger; DO NOT EDIT.

package preheat

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/blackaichi/harbor-go-client/pkg/sdk/v2.0/models"
)

// GetInstanceReader is a Reader for the GetInstance structure.
type GetInstanceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetInstanceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetInstanceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetInstanceBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetInstanceUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetInstanceForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetInstanceNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetInstanceInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetInstanceOK creates a GetInstanceOK with default headers values
func NewGetInstanceOK() *GetInstanceOK {
	return &GetInstanceOK{}
}

/*
GetInstanceOK describes a response with status code 200, with default header values.

Success
*/
type GetInstanceOK struct {
	Payload *models.Instance
}

// IsSuccess returns true when this get instance o k response has a 2xx status code
func (o *GetInstanceOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get instance o k response has a 3xx status code
func (o *GetInstanceOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get instance o k response has a 4xx status code
func (o *GetInstanceOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get instance o k response has a 5xx status code
func (o *GetInstanceOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get instance o k response a status code equal to that given
func (o *GetInstanceOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetInstanceOK) Error() string {
	return fmt.Sprintf("[GET /p2p/preheat/instances/{preheat_instance_name}][%d] getInstanceOK  %+v", 200, o.Payload)
}

func (o *GetInstanceOK) String() string {
	return fmt.Sprintf("[GET /p2p/preheat/instances/{preheat_instance_name}][%d] getInstanceOK  %+v", 200, o.Payload)
}

func (o *GetInstanceOK) GetPayload() *models.Instance {
	return o.Payload
}

func (o *GetInstanceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Instance)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInstanceBadRequest creates a GetInstanceBadRequest with default headers values
func NewGetInstanceBadRequest() *GetInstanceBadRequest {
	return &GetInstanceBadRequest{}
}

/*
GetInstanceBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetInstanceBadRequest struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this get instance bad request response has a 2xx status code
func (o *GetInstanceBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get instance bad request response has a 3xx status code
func (o *GetInstanceBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get instance bad request response has a 4xx status code
func (o *GetInstanceBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get instance bad request response has a 5xx status code
func (o *GetInstanceBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get instance bad request response a status code equal to that given
func (o *GetInstanceBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetInstanceBadRequest) Error() string {
	return fmt.Sprintf("[GET /p2p/preheat/instances/{preheat_instance_name}][%d] getInstanceBadRequest  %+v", 400, o.Payload)
}

func (o *GetInstanceBadRequest) String() string {
	return fmt.Sprintf("[GET /p2p/preheat/instances/{preheat_instance_name}][%d] getInstanceBadRequest  %+v", 400, o.Payload)
}

func (o *GetInstanceBadRequest) GetPayload() *models.Errors {
	return o.Payload
}

func (o *GetInstanceBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetInstanceUnauthorized creates a GetInstanceUnauthorized with default headers values
func NewGetInstanceUnauthorized() *GetInstanceUnauthorized {
	return &GetInstanceUnauthorized{}
}

/*
GetInstanceUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetInstanceUnauthorized struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this get instance unauthorized response has a 2xx status code
func (o *GetInstanceUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get instance unauthorized response has a 3xx status code
func (o *GetInstanceUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get instance unauthorized response has a 4xx status code
func (o *GetInstanceUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get instance unauthorized response has a 5xx status code
func (o *GetInstanceUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get instance unauthorized response a status code equal to that given
func (o *GetInstanceUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetInstanceUnauthorized) Error() string {
	return fmt.Sprintf("[GET /p2p/preheat/instances/{preheat_instance_name}][%d] getInstanceUnauthorized  %+v", 401, o.Payload)
}

func (o *GetInstanceUnauthorized) String() string {
	return fmt.Sprintf("[GET /p2p/preheat/instances/{preheat_instance_name}][%d] getInstanceUnauthorized  %+v", 401, o.Payload)
}

func (o *GetInstanceUnauthorized) GetPayload() *models.Errors {
	return o.Payload
}

func (o *GetInstanceUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetInstanceForbidden creates a GetInstanceForbidden with default headers values
func NewGetInstanceForbidden() *GetInstanceForbidden {
	return &GetInstanceForbidden{}
}

/*
GetInstanceForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetInstanceForbidden struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this get instance forbidden response has a 2xx status code
func (o *GetInstanceForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get instance forbidden response has a 3xx status code
func (o *GetInstanceForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get instance forbidden response has a 4xx status code
func (o *GetInstanceForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get instance forbidden response has a 5xx status code
func (o *GetInstanceForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get instance forbidden response a status code equal to that given
func (o *GetInstanceForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetInstanceForbidden) Error() string {
	return fmt.Sprintf("[GET /p2p/preheat/instances/{preheat_instance_name}][%d] getInstanceForbidden  %+v", 403, o.Payload)
}

func (o *GetInstanceForbidden) String() string {
	return fmt.Sprintf("[GET /p2p/preheat/instances/{preheat_instance_name}][%d] getInstanceForbidden  %+v", 403, o.Payload)
}

func (o *GetInstanceForbidden) GetPayload() *models.Errors {
	return o.Payload
}

func (o *GetInstanceForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetInstanceNotFound creates a GetInstanceNotFound with default headers values
func NewGetInstanceNotFound() *GetInstanceNotFound {
	return &GetInstanceNotFound{}
}

/*
GetInstanceNotFound describes a response with status code 404, with default header values.

Not found
*/
type GetInstanceNotFound struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this get instance not found response has a 2xx status code
func (o *GetInstanceNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get instance not found response has a 3xx status code
func (o *GetInstanceNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get instance not found response has a 4xx status code
func (o *GetInstanceNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get instance not found response has a 5xx status code
func (o *GetInstanceNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get instance not found response a status code equal to that given
func (o *GetInstanceNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetInstanceNotFound) Error() string {
	return fmt.Sprintf("[GET /p2p/preheat/instances/{preheat_instance_name}][%d] getInstanceNotFound  %+v", 404, o.Payload)
}

func (o *GetInstanceNotFound) String() string {
	return fmt.Sprintf("[GET /p2p/preheat/instances/{preheat_instance_name}][%d] getInstanceNotFound  %+v", 404, o.Payload)
}

func (o *GetInstanceNotFound) GetPayload() *models.Errors {
	return o.Payload
}

func (o *GetInstanceNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetInstanceInternalServerError creates a GetInstanceInternalServerError with default headers values
func NewGetInstanceInternalServerError() *GetInstanceInternalServerError {
	return &GetInstanceInternalServerError{}
}

/*
GetInstanceInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GetInstanceInternalServerError struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this get instance internal server error response has a 2xx status code
func (o *GetInstanceInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get instance internal server error response has a 3xx status code
func (o *GetInstanceInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get instance internal server error response has a 4xx status code
func (o *GetInstanceInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get instance internal server error response has a 5xx status code
func (o *GetInstanceInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get instance internal server error response a status code equal to that given
func (o *GetInstanceInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetInstanceInternalServerError) Error() string {
	return fmt.Sprintf("[GET /p2p/preheat/instances/{preheat_instance_name}][%d] getInstanceInternalServerError  %+v", 500, o.Payload)
}

func (o *GetInstanceInternalServerError) String() string {
	return fmt.Sprintf("[GET /p2p/preheat/instances/{preheat_instance_name}][%d] getInstanceInternalServerError  %+v", 500, o.Payload)
}

func (o *GetInstanceInternalServerError) GetPayload() *models.Errors {
	return o.Payload
}

func (o *GetInstanceInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
