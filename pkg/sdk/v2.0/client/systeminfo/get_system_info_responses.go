// Code generated by go-swagger; DO NOT EDIT.

package systeminfo

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/blackaichi/harbor-go-client/pkg/sdk/v2.0/models"
)

// GetSystemInfoReader is a Reader for the GetSystemInfo structure.
type GetSystemInfoReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSystemInfoReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSystemInfoOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewGetSystemInfoInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetSystemInfoOK creates a GetSystemInfoOK with default headers values
func NewGetSystemInfoOK() *GetSystemInfoOK {
	return &GetSystemInfoOK{}
}

/*
GetSystemInfoOK describes a response with status code 200, with default header values.

Get general info successfully.
*/
type GetSystemInfoOK struct {
	Payload *models.GeneralInfo
}

// IsSuccess returns true when this get system info o k response has a 2xx status code
func (o *GetSystemInfoOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get system info o k response has a 3xx status code
func (o *GetSystemInfoOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get system info o k response has a 4xx status code
func (o *GetSystemInfoOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get system info o k response has a 5xx status code
func (o *GetSystemInfoOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get system info o k response a status code equal to that given
func (o *GetSystemInfoOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetSystemInfoOK) Error() string {
	return fmt.Sprintf("[GET /systeminfo][%d] getSystemInfoOK  %+v", 200, o.Payload)
}

func (o *GetSystemInfoOK) String() string {
	return fmt.Sprintf("[GET /systeminfo][%d] getSystemInfoOK  %+v", 200, o.Payload)
}

func (o *GetSystemInfoOK) GetPayload() *models.GeneralInfo {
	return o.Payload
}

func (o *GetSystemInfoOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralInfo)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSystemInfoInternalServerError creates a GetSystemInfoInternalServerError with default headers values
func NewGetSystemInfoInternalServerError() *GetSystemInfoInternalServerError {
	return &GetSystemInfoInternalServerError{}
}

/*
GetSystemInfoInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GetSystemInfoInternalServerError struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this get system info internal server error response has a 2xx status code
func (o *GetSystemInfoInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get system info internal server error response has a 3xx status code
func (o *GetSystemInfoInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get system info internal server error response has a 4xx status code
func (o *GetSystemInfoInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get system info internal server error response has a 5xx status code
func (o *GetSystemInfoInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get system info internal server error response a status code equal to that given
func (o *GetSystemInfoInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetSystemInfoInternalServerError) Error() string {
	return fmt.Sprintf("[GET /systeminfo][%d] getSystemInfoInternalServerError  %+v", 500, o.Payload)
}

func (o *GetSystemInfoInternalServerError) String() string {
	return fmt.Sprintf("[GET /systeminfo][%d] getSystemInfoInternalServerError  %+v", 500, o.Payload)
}

func (o *GetSystemInfoInternalServerError) GetPayload() *models.Errors {
	return o.Payload
}

func (o *GetSystemInfoInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
