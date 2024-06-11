// Code generated by go-swagger; DO NOT EDIT.

package gc

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/blackaichi/harbor-go-client/pkg/sdk/v2.0/models"
)

// GetGCScheduleReader is a Reader for the GetGCSchedule structure.
type GetGCScheduleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetGCScheduleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetGCScheduleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetGCScheduleUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetGCScheduleForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetGCScheduleInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetGCScheduleOK creates a GetGCScheduleOK with default headers values
func NewGetGCScheduleOK() *GetGCScheduleOK {
	return &GetGCScheduleOK{}
}

/*
GetGCScheduleOK describes a response with status code 200, with default header values.

Get gc's schedule.
*/
type GetGCScheduleOK struct {
	Payload *models.GCHistory
}

// IsSuccess returns true when this get Gc schedule o k response has a 2xx status code
func (o *GetGCScheduleOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get Gc schedule o k response has a 3xx status code
func (o *GetGCScheduleOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get Gc schedule o k response has a 4xx status code
func (o *GetGCScheduleOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get Gc schedule o k response has a 5xx status code
func (o *GetGCScheduleOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get Gc schedule o k response a status code equal to that given
func (o *GetGCScheduleOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetGCScheduleOK) Error() string {
	return fmt.Sprintf("[GET /system/gc/schedule][%d] getGcScheduleOK  %+v", 200, o.Payload)
}

func (o *GetGCScheduleOK) String() string {
	return fmt.Sprintf("[GET /system/gc/schedule][%d] getGcScheduleOK  %+v", 200, o.Payload)
}

func (o *GetGCScheduleOK) GetPayload() *models.GCHistory {
	return o.Payload
}

func (o *GetGCScheduleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GCHistory)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGCScheduleUnauthorized creates a GetGCScheduleUnauthorized with default headers values
func NewGetGCScheduleUnauthorized() *GetGCScheduleUnauthorized {
	return &GetGCScheduleUnauthorized{}
}

/*
GetGCScheduleUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetGCScheduleUnauthorized struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this get Gc schedule unauthorized response has a 2xx status code
func (o *GetGCScheduleUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get Gc schedule unauthorized response has a 3xx status code
func (o *GetGCScheduleUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get Gc schedule unauthorized response has a 4xx status code
func (o *GetGCScheduleUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get Gc schedule unauthorized response has a 5xx status code
func (o *GetGCScheduleUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get Gc schedule unauthorized response a status code equal to that given
func (o *GetGCScheduleUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetGCScheduleUnauthorized) Error() string {
	return fmt.Sprintf("[GET /system/gc/schedule][%d] getGcScheduleUnauthorized  %+v", 401, o.Payload)
}

func (o *GetGCScheduleUnauthorized) String() string {
	return fmt.Sprintf("[GET /system/gc/schedule][%d] getGcScheduleUnauthorized  %+v", 401, o.Payload)
}

func (o *GetGCScheduleUnauthorized) GetPayload() *models.Errors {
	return o.Payload
}

func (o *GetGCScheduleUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetGCScheduleForbidden creates a GetGCScheduleForbidden with default headers values
func NewGetGCScheduleForbidden() *GetGCScheduleForbidden {
	return &GetGCScheduleForbidden{}
}

/*
GetGCScheduleForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetGCScheduleForbidden struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this get Gc schedule forbidden response has a 2xx status code
func (o *GetGCScheduleForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get Gc schedule forbidden response has a 3xx status code
func (o *GetGCScheduleForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get Gc schedule forbidden response has a 4xx status code
func (o *GetGCScheduleForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get Gc schedule forbidden response has a 5xx status code
func (o *GetGCScheduleForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get Gc schedule forbidden response a status code equal to that given
func (o *GetGCScheduleForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetGCScheduleForbidden) Error() string {
	return fmt.Sprintf("[GET /system/gc/schedule][%d] getGcScheduleForbidden  %+v", 403, o.Payload)
}

func (o *GetGCScheduleForbidden) String() string {
	return fmt.Sprintf("[GET /system/gc/schedule][%d] getGcScheduleForbidden  %+v", 403, o.Payload)
}

func (o *GetGCScheduleForbidden) GetPayload() *models.Errors {
	return o.Payload
}

func (o *GetGCScheduleForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetGCScheduleInternalServerError creates a GetGCScheduleInternalServerError with default headers values
func NewGetGCScheduleInternalServerError() *GetGCScheduleInternalServerError {
	return &GetGCScheduleInternalServerError{}
}

/*
GetGCScheduleInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GetGCScheduleInternalServerError struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this get Gc schedule internal server error response has a 2xx status code
func (o *GetGCScheduleInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get Gc schedule internal server error response has a 3xx status code
func (o *GetGCScheduleInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get Gc schedule internal server error response has a 4xx status code
func (o *GetGCScheduleInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get Gc schedule internal server error response has a 5xx status code
func (o *GetGCScheduleInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get Gc schedule internal server error response a status code equal to that given
func (o *GetGCScheduleInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetGCScheduleInternalServerError) Error() string {
	return fmt.Sprintf("[GET /system/gc/schedule][%d] getGcScheduleInternalServerError  %+v", 500, o.Payload)
}

func (o *GetGCScheduleInternalServerError) String() string {
	return fmt.Sprintf("[GET /system/gc/schedule][%d] getGcScheduleInternalServerError  %+v", 500, o.Payload)
}

func (o *GetGCScheduleInternalServerError) GetPayload() *models.Errors {
	return o.Payload
}

func (o *GetGCScheduleInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
