// Code generated by go-swagger; DO NOT EDIT.

package scanner

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/blackaichi/harbor-go-client/pkg/sdk/v2.0/models"
)

// ListScannersReader is a Reader for the ListScanners structure.
type ListScannersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListScannersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListScannersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewListScannersBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewListScannersUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListScannersForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewListScannersInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListScannersOK creates a ListScannersOK with default headers values
func NewListScannersOK() *ListScannersOK {
	return &ListScannersOK{}
}

/*
ListScannersOK describes a response with status code 200, with default header values.

A list of scanner registrations.
*/
type ListScannersOK struct {

	/* Link to previous page and next page
	 */
	Link string

	/* The total count of available items
	 */
	XTotalCount int64

	Payload []*models.ScannerRegistration
}

// IsSuccess returns true when this list scanners o k response has a 2xx status code
func (o *ListScannersOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list scanners o k response has a 3xx status code
func (o *ListScannersOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list scanners o k response has a 4xx status code
func (o *ListScannersOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list scanners o k response has a 5xx status code
func (o *ListScannersOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list scanners o k response a status code equal to that given
func (o *ListScannersOK) IsCode(code int) bool {
	return code == 200
}

func (o *ListScannersOK) Error() string {
	return fmt.Sprintf("[GET /scanners][%d] listScannersOK  %+v", 200, o.Payload)
}

func (o *ListScannersOK) String() string {
	return fmt.Sprintf("[GET /scanners][%d] listScannersOK  %+v", 200, o.Payload)
}

func (o *ListScannersOK) GetPayload() []*models.ScannerRegistration {
	return o.Payload
}

func (o *ListScannersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Link
	hdrLink := response.GetHeader("Link")

	if hdrLink != "" {
		o.Link = hdrLink
	}

	// hydrates response header X-Total-Count
	hdrXTotalCount := response.GetHeader("X-Total-Count")

	if hdrXTotalCount != "" {
		valxTotalCount, err := swag.ConvertInt64(hdrXTotalCount)
		if err != nil {
			return errors.InvalidType("X-Total-Count", "header", "int64", hdrXTotalCount)
		}
		o.XTotalCount = valxTotalCount
	}

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListScannersBadRequest creates a ListScannersBadRequest with default headers values
func NewListScannersBadRequest() *ListScannersBadRequest {
	return &ListScannersBadRequest{}
}

/*
ListScannersBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type ListScannersBadRequest struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this list scanners bad request response has a 2xx status code
func (o *ListScannersBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list scanners bad request response has a 3xx status code
func (o *ListScannersBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list scanners bad request response has a 4xx status code
func (o *ListScannersBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this list scanners bad request response has a 5xx status code
func (o *ListScannersBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this list scanners bad request response a status code equal to that given
func (o *ListScannersBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *ListScannersBadRequest) Error() string {
	return fmt.Sprintf("[GET /scanners][%d] listScannersBadRequest  %+v", 400, o.Payload)
}

func (o *ListScannersBadRequest) String() string {
	return fmt.Sprintf("[GET /scanners][%d] listScannersBadRequest  %+v", 400, o.Payload)
}

func (o *ListScannersBadRequest) GetPayload() *models.Errors {
	return o.Payload
}

func (o *ListScannersBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewListScannersUnauthorized creates a ListScannersUnauthorized with default headers values
func NewListScannersUnauthorized() *ListScannersUnauthorized {
	return &ListScannersUnauthorized{}
}

/*
ListScannersUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type ListScannersUnauthorized struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this list scanners unauthorized response has a 2xx status code
func (o *ListScannersUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list scanners unauthorized response has a 3xx status code
func (o *ListScannersUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list scanners unauthorized response has a 4xx status code
func (o *ListScannersUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this list scanners unauthorized response has a 5xx status code
func (o *ListScannersUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this list scanners unauthorized response a status code equal to that given
func (o *ListScannersUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *ListScannersUnauthorized) Error() string {
	return fmt.Sprintf("[GET /scanners][%d] listScannersUnauthorized  %+v", 401, o.Payload)
}

func (o *ListScannersUnauthorized) String() string {
	return fmt.Sprintf("[GET /scanners][%d] listScannersUnauthorized  %+v", 401, o.Payload)
}

func (o *ListScannersUnauthorized) GetPayload() *models.Errors {
	return o.Payload
}

func (o *ListScannersUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewListScannersForbidden creates a ListScannersForbidden with default headers values
func NewListScannersForbidden() *ListScannersForbidden {
	return &ListScannersForbidden{}
}

/*
ListScannersForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type ListScannersForbidden struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this list scanners forbidden response has a 2xx status code
func (o *ListScannersForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list scanners forbidden response has a 3xx status code
func (o *ListScannersForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list scanners forbidden response has a 4xx status code
func (o *ListScannersForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this list scanners forbidden response has a 5xx status code
func (o *ListScannersForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this list scanners forbidden response a status code equal to that given
func (o *ListScannersForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *ListScannersForbidden) Error() string {
	return fmt.Sprintf("[GET /scanners][%d] listScannersForbidden  %+v", 403, o.Payload)
}

func (o *ListScannersForbidden) String() string {
	return fmt.Sprintf("[GET /scanners][%d] listScannersForbidden  %+v", 403, o.Payload)
}

func (o *ListScannersForbidden) GetPayload() *models.Errors {
	return o.Payload
}

func (o *ListScannersForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewListScannersInternalServerError creates a ListScannersInternalServerError with default headers values
func NewListScannersInternalServerError() *ListScannersInternalServerError {
	return &ListScannersInternalServerError{}
}

/*
ListScannersInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type ListScannersInternalServerError struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this list scanners internal server error response has a 2xx status code
func (o *ListScannersInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list scanners internal server error response has a 3xx status code
func (o *ListScannersInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list scanners internal server error response has a 4xx status code
func (o *ListScannersInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this list scanners internal server error response has a 5xx status code
func (o *ListScannersInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this list scanners internal server error response a status code equal to that given
func (o *ListScannersInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *ListScannersInternalServerError) Error() string {
	return fmt.Sprintf("[GET /scanners][%d] listScannersInternalServerError  %+v", 500, o.Payload)
}

func (o *ListScannersInternalServerError) String() string {
	return fmt.Sprintf("[GET /scanners][%d] listScannersInternalServerError  %+v", 500, o.Payload)
}

func (o *ListScannersInternalServerError) GetPayload() *models.Errors {
	return o.Payload
}

func (o *ListScannersInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
