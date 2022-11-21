// Code generated by go-swagger; DO NOT EDIT.

package scanner

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/goharbor/go-client/pkg/sdk/v2.0/models"
)

// GetScannerMetadataReader is a Reader for the GetScannerMetadata structure.
type GetScannerMetadataReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetScannerMetadataReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetScannerMetadataOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetScannerMetadataUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetScannerMetadataForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetScannerMetadataInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetScannerMetadataOK creates a GetScannerMetadataOK with default headers values
func NewGetScannerMetadataOK() *GetScannerMetadataOK {
	return &GetScannerMetadataOK{}
}

/*
GetScannerMetadataOK describes a response with status code 200, with default header values.

The metadata of the specified scanner adapter
*/
type GetScannerMetadataOK struct {
	Payload *models.ScannerAdapterMetadata
}

// IsSuccess returns true when this get scanner metadata o k response has a 2xx status code
func (o *GetScannerMetadataOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get scanner metadata o k response has a 3xx status code
func (o *GetScannerMetadataOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get scanner metadata o k response has a 4xx status code
func (o *GetScannerMetadataOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get scanner metadata o k response has a 5xx status code
func (o *GetScannerMetadataOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get scanner metadata o k response a status code equal to that given
func (o *GetScannerMetadataOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetScannerMetadataOK) Error() string {
	return fmt.Sprintf("[GET /scanners/{registration_id}/metadata][%d] getScannerMetadataOK  %+v", 200, o.Payload)
}

func (o *GetScannerMetadataOK) String() string {
	return fmt.Sprintf("[GET /scanners/{registration_id}/metadata][%d] getScannerMetadataOK  %+v", 200, o.Payload)
}

func (o *GetScannerMetadataOK) GetPayload() *models.ScannerAdapterMetadata {
	return o.Payload
}

func (o *GetScannerMetadataOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ScannerAdapterMetadata)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScannerMetadataUnauthorized creates a GetScannerMetadataUnauthorized with default headers values
func NewGetScannerMetadataUnauthorized() *GetScannerMetadataUnauthorized {
	return &GetScannerMetadataUnauthorized{}
}

/*
GetScannerMetadataUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetScannerMetadataUnauthorized struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this get scanner metadata unauthorized response has a 2xx status code
func (o *GetScannerMetadataUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get scanner metadata unauthorized response has a 3xx status code
func (o *GetScannerMetadataUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get scanner metadata unauthorized response has a 4xx status code
func (o *GetScannerMetadataUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get scanner metadata unauthorized response has a 5xx status code
func (o *GetScannerMetadataUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get scanner metadata unauthorized response a status code equal to that given
func (o *GetScannerMetadataUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetScannerMetadataUnauthorized) Error() string {
	return fmt.Sprintf("[GET /scanners/{registration_id}/metadata][%d] getScannerMetadataUnauthorized  %+v", 401, o.Payload)
}

func (o *GetScannerMetadataUnauthorized) String() string {
	return fmt.Sprintf("[GET /scanners/{registration_id}/metadata][%d] getScannerMetadataUnauthorized  %+v", 401, o.Payload)
}

func (o *GetScannerMetadataUnauthorized) GetPayload() *models.Errors {
	return o.Payload
}

func (o *GetScannerMetadataUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetScannerMetadataForbidden creates a GetScannerMetadataForbidden with default headers values
func NewGetScannerMetadataForbidden() *GetScannerMetadataForbidden {
	return &GetScannerMetadataForbidden{}
}

/*
GetScannerMetadataForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetScannerMetadataForbidden struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this get scanner metadata forbidden response has a 2xx status code
func (o *GetScannerMetadataForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get scanner metadata forbidden response has a 3xx status code
func (o *GetScannerMetadataForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get scanner metadata forbidden response has a 4xx status code
func (o *GetScannerMetadataForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get scanner metadata forbidden response has a 5xx status code
func (o *GetScannerMetadataForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get scanner metadata forbidden response a status code equal to that given
func (o *GetScannerMetadataForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetScannerMetadataForbidden) Error() string {
	return fmt.Sprintf("[GET /scanners/{registration_id}/metadata][%d] getScannerMetadataForbidden  %+v", 403, o.Payload)
}

func (o *GetScannerMetadataForbidden) String() string {
	return fmt.Sprintf("[GET /scanners/{registration_id}/metadata][%d] getScannerMetadataForbidden  %+v", 403, o.Payload)
}

func (o *GetScannerMetadataForbidden) GetPayload() *models.Errors {
	return o.Payload
}

func (o *GetScannerMetadataForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetScannerMetadataInternalServerError creates a GetScannerMetadataInternalServerError with default headers values
func NewGetScannerMetadataInternalServerError() *GetScannerMetadataInternalServerError {
	return &GetScannerMetadataInternalServerError{}
}

/*
GetScannerMetadataInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GetScannerMetadataInternalServerError struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this get scanner metadata internal server error response has a 2xx status code
func (o *GetScannerMetadataInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get scanner metadata internal server error response has a 3xx status code
func (o *GetScannerMetadataInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get scanner metadata internal server error response has a 4xx status code
func (o *GetScannerMetadataInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get scanner metadata internal server error response has a 5xx status code
func (o *GetScannerMetadataInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get scanner metadata internal server error response a status code equal to that given
func (o *GetScannerMetadataInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetScannerMetadataInternalServerError) Error() string {
	return fmt.Sprintf("[GET /scanners/{registration_id}/metadata][%d] getScannerMetadataInternalServerError  %+v", 500, o.Payload)
}

func (o *GetScannerMetadataInternalServerError) String() string {
	return fmt.Sprintf("[GET /scanners/{registration_id}/metadata][%d] getScannerMetadataInternalServerError  %+v", 500, o.Payload)
}

func (o *GetScannerMetadataInternalServerError) GetPayload() *models.Errors {
	return o.Payload
}

func (o *GetScannerMetadataInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
