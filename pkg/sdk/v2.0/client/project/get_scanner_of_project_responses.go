// Code generated by go-swagger; DO NOT EDIT.

package project

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/blackaichi/harbor-go-client/pkg/sdk/v2.0/models"
)

// GetScannerOfProjectReader is a Reader for the GetScannerOfProject structure.
type GetScannerOfProjectReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetScannerOfProjectReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetScannerOfProjectOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetScannerOfProjectBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetScannerOfProjectUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetScannerOfProjectForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetScannerOfProjectNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetScannerOfProjectInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetScannerOfProjectOK creates a GetScannerOfProjectOK with default headers values
func NewGetScannerOfProjectOK() *GetScannerOfProjectOK {
	return &GetScannerOfProjectOK{}
}

/*
GetScannerOfProjectOK describes a response with status code 200, with default header values.

The details of the scanner registration.
*/
type GetScannerOfProjectOK struct {
	Payload *models.ScannerRegistration
}

// IsSuccess returns true when this get scanner of project o k response has a 2xx status code
func (o *GetScannerOfProjectOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get scanner of project o k response has a 3xx status code
func (o *GetScannerOfProjectOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get scanner of project o k response has a 4xx status code
func (o *GetScannerOfProjectOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get scanner of project o k response has a 5xx status code
func (o *GetScannerOfProjectOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get scanner of project o k response a status code equal to that given
func (o *GetScannerOfProjectOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetScannerOfProjectOK) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name_or_id}/scanner][%d] getScannerOfProjectOK  %+v", 200, o.Payload)
}

func (o *GetScannerOfProjectOK) String() string {
	return fmt.Sprintf("[GET /projects/{project_name_or_id}/scanner][%d] getScannerOfProjectOK  %+v", 200, o.Payload)
}

func (o *GetScannerOfProjectOK) GetPayload() *models.ScannerRegistration {
	return o.Payload
}

func (o *GetScannerOfProjectOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ScannerRegistration)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScannerOfProjectBadRequest creates a GetScannerOfProjectBadRequest with default headers values
func NewGetScannerOfProjectBadRequest() *GetScannerOfProjectBadRequest {
	return &GetScannerOfProjectBadRequest{}
}

/*
GetScannerOfProjectBadRequest describes a response with status code 400, with default header values.

Bad project ID
*/
type GetScannerOfProjectBadRequest struct {
}

// IsSuccess returns true when this get scanner of project bad request response has a 2xx status code
func (o *GetScannerOfProjectBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get scanner of project bad request response has a 3xx status code
func (o *GetScannerOfProjectBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get scanner of project bad request response has a 4xx status code
func (o *GetScannerOfProjectBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get scanner of project bad request response has a 5xx status code
func (o *GetScannerOfProjectBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get scanner of project bad request response a status code equal to that given
func (o *GetScannerOfProjectBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetScannerOfProjectBadRequest) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name_or_id}/scanner][%d] getScannerOfProjectBadRequest ", 400)
}

func (o *GetScannerOfProjectBadRequest) String() string {
	return fmt.Sprintf("[GET /projects/{project_name_or_id}/scanner][%d] getScannerOfProjectBadRequest ", 400)
}

func (o *GetScannerOfProjectBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetScannerOfProjectUnauthorized creates a GetScannerOfProjectUnauthorized with default headers values
func NewGetScannerOfProjectUnauthorized() *GetScannerOfProjectUnauthorized {
	return &GetScannerOfProjectUnauthorized{}
}

/*
GetScannerOfProjectUnauthorized describes a response with status code 401, with default header values.

Unauthorized request
*/
type GetScannerOfProjectUnauthorized struct {
}

// IsSuccess returns true when this get scanner of project unauthorized response has a 2xx status code
func (o *GetScannerOfProjectUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get scanner of project unauthorized response has a 3xx status code
func (o *GetScannerOfProjectUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get scanner of project unauthorized response has a 4xx status code
func (o *GetScannerOfProjectUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get scanner of project unauthorized response has a 5xx status code
func (o *GetScannerOfProjectUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get scanner of project unauthorized response a status code equal to that given
func (o *GetScannerOfProjectUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetScannerOfProjectUnauthorized) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name_or_id}/scanner][%d] getScannerOfProjectUnauthorized ", 401)
}

func (o *GetScannerOfProjectUnauthorized) String() string {
	return fmt.Sprintf("[GET /projects/{project_name_or_id}/scanner][%d] getScannerOfProjectUnauthorized ", 401)
}

func (o *GetScannerOfProjectUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetScannerOfProjectForbidden creates a GetScannerOfProjectForbidden with default headers values
func NewGetScannerOfProjectForbidden() *GetScannerOfProjectForbidden {
	return &GetScannerOfProjectForbidden{}
}

/*
GetScannerOfProjectForbidden describes a response with status code 403, with default header values.

Request is not allowed
*/
type GetScannerOfProjectForbidden struct {
}

// IsSuccess returns true when this get scanner of project forbidden response has a 2xx status code
func (o *GetScannerOfProjectForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get scanner of project forbidden response has a 3xx status code
func (o *GetScannerOfProjectForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get scanner of project forbidden response has a 4xx status code
func (o *GetScannerOfProjectForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get scanner of project forbidden response has a 5xx status code
func (o *GetScannerOfProjectForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get scanner of project forbidden response a status code equal to that given
func (o *GetScannerOfProjectForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetScannerOfProjectForbidden) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name_or_id}/scanner][%d] getScannerOfProjectForbidden ", 403)
}

func (o *GetScannerOfProjectForbidden) String() string {
	return fmt.Sprintf("[GET /projects/{project_name_or_id}/scanner][%d] getScannerOfProjectForbidden ", 403)
}

func (o *GetScannerOfProjectForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetScannerOfProjectNotFound creates a GetScannerOfProjectNotFound with default headers values
func NewGetScannerOfProjectNotFound() *GetScannerOfProjectNotFound {
	return &GetScannerOfProjectNotFound{}
}

/*
GetScannerOfProjectNotFound describes a response with status code 404, with default header values.

The requested object is not found
*/
type GetScannerOfProjectNotFound struct {
}

// IsSuccess returns true when this get scanner of project not found response has a 2xx status code
func (o *GetScannerOfProjectNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get scanner of project not found response has a 3xx status code
func (o *GetScannerOfProjectNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get scanner of project not found response has a 4xx status code
func (o *GetScannerOfProjectNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get scanner of project not found response has a 5xx status code
func (o *GetScannerOfProjectNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get scanner of project not found response a status code equal to that given
func (o *GetScannerOfProjectNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetScannerOfProjectNotFound) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name_or_id}/scanner][%d] getScannerOfProjectNotFound ", 404)
}

func (o *GetScannerOfProjectNotFound) String() string {
	return fmt.Sprintf("[GET /projects/{project_name_or_id}/scanner][%d] getScannerOfProjectNotFound ", 404)
}

func (o *GetScannerOfProjectNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetScannerOfProjectInternalServerError creates a GetScannerOfProjectInternalServerError with default headers values
func NewGetScannerOfProjectInternalServerError() *GetScannerOfProjectInternalServerError {
	return &GetScannerOfProjectInternalServerError{}
}

/*
GetScannerOfProjectInternalServerError describes a response with status code 500, with default header values.

Internal server error happened
*/
type GetScannerOfProjectInternalServerError struct {
}

// IsSuccess returns true when this get scanner of project internal server error response has a 2xx status code
func (o *GetScannerOfProjectInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get scanner of project internal server error response has a 3xx status code
func (o *GetScannerOfProjectInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get scanner of project internal server error response has a 4xx status code
func (o *GetScannerOfProjectInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get scanner of project internal server error response has a 5xx status code
func (o *GetScannerOfProjectInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get scanner of project internal server error response a status code equal to that given
func (o *GetScannerOfProjectInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetScannerOfProjectInternalServerError) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name_or_id}/scanner][%d] getScannerOfProjectInternalServerError ", 500)
}

func (o *GetScannerOfProjectInternalServerError) String() string {
	return fmt.Sprintf("[GET /projects/{project_name_or_id}/scanner][%d] getScannerOfProjectInternalServerError ", 500)
}

func (o *GetScannerOfProjectInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
