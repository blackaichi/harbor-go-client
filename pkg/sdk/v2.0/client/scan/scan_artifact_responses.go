// Code generated by go-swagger; DO NOT EDIT.

package scan

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/blackaichi/harbor-go-client/pkg/sdk/v2.0/models"
)

// ScanArtifactReader is a Reader for the ScanArtifact structure.
type ScanArtifactReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ScanArtifactReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewScanArtifactAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewScanArtifactBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewScanArtifactUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewScanArtifactForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewScanArtifactNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewScanArtifactInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewScanArtifactAccepted creates a ScanArtifactAccepted with default headers values
func NewScanArtifactAccepted() *ScanArtifactAccepted {
	return &ScanArtifactAccepted{}
}

/*
ScanArtifactAccepted describes a response with status code 202, with default header values.

Accepted
*/
type ScanArtifactAccepted struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string
}

// IsSuccess returns true when this scan artifact accepted response has a 2xx status code
func (o *ScanArtifactAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this scan artifact accepted response has a 3xx status code
func (o *ScanArtifactAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this scan artifact accepted response has a 4xx status code
func (o *ScanArtifactAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this scan artifact accepted response has a 5xx status code
func (o *ScanArtifactAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this scan artifact accepted response a status code equal to that given
func (o *ScanArtifactAccepted) IsCode(code int) bool {
	return code == 202
}

func (o *ScanArtifactAccepted) Error() string {
	return fmt.Sprintf("[POST /projects/{project_name}/repositories/{repository_name}/artifacts/{reference}/scan][%d] scanArtifactAccepted ", 202)
}

func (o *ScanArtifactAccepted) String() string {
	return fmt.Sprintf("[POST /projects/{project_name}/repositories/{repository_name}/artifacts/{reference}/scan][%d] scanArtifactAccepted ", 202)
}

func (o *ScanArtifactAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-Request-Id
	hdrXRequestID := response.GetHeader("X-Request-Id")

	if hdrXRequestID != "" {
		o.XRequestID = hdrXRequestID
	}

	return nil
}

// NewScanArtifactBadRequest creates a ScanArtifactBadRequest with default headers values
func NewScanArtifactBadRequest() *ScanArtifactBadRequest {
	return &ScanArtifactBadRequest{}
}

/*
ScanArtifactBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type ScanArtifactBadRequest struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this scan artifact bad request response has a 2xx status code
func (o *ScanArtifactBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this scan artifact bad request response has a 3xx status code
func (o *ScanArtifactBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this scan artifact bad request response has a 4xx status code
func (o *ScanArtifactBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this scan artifact bad request response has a 5xx status code
func (o *ScanArtifactBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this scan artifact bad request response a status code equal to that given
func (o *ScanArtifactBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *ScanArtifactBadRequest) Error() string {
	return fmt.Sprintf("[POST /projects/{project_name}/repositories/{repository_name}/artifacts/{reference}/scan][%d] scanArtifactBadRequest  %+v", 400, o.Payload)
}

func (o *ScanArtifactBadRequest) String() string {
	return fmt.Sprintf("[POST /projects/{project_name}/repositories/{repository_name}/artifacts/{reference}/scan][%d] scanArtifactBadRequest  %+v", 400, o.Payload)
}

func (o *ScanArtifactBadRequest) GetPayload() *models.Errors {
	return o.Payload
}

func (o *ScanArtifactBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewScanArtifactUnauthorized creates a ScanArtifactUnauthorized with default headers values
func NewScanArtifactUnauthorized() *ScanArtifactUnauthorized {
	return &ScanArtifactUnauthorized{}
}

/*
ScanArtifactUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type ScanArtifactUnauthorized struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this scan artifact unauthorized response has a 2xx status code
func (o *ScanArtifactUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this scan artifact unauthorized response has a 3xx status code
func (o *ScanArtifactUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this scan artifact unauthorized response has a 4xx status code
func (o *ScanArtifactUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this scan artifact unauthorized response has a 5xx status code
func (o *ScanArtifactUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this scan artifact unauthorized response a status code equal to that given
func (o *ScanArtifactUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *ScanArtifactUnauthorized) Error() string {
	return fmt.Sprintf("[POST /projects/{project_name}/repositories/{repository_name}/artifacts/{reference}/scan][%d] scanArtifactUnauthorized  %+v", 401, o.Payload)
}

func (o *ScanArtifactUnauthorized) String() string {
	return fmt.Sprintf("[POST /projects/{project_name}/repositories/{repository_name}/artifacts/{reference}/scan][%d] scanArtifactUnauthorized  %+v", 401, o.Payload)
}

func (o *ScanArtifactUnauthorized) GetPayload() *models.Errors {
	return o.Payload
}

func (o *ScanArtifactUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewScanArtifactForbidden creates a ScanArtifactForbidden with default headers values
func NewScanArtifactForbidden() *ScanArtifactForbidden {
	return &ScanArtifactForbidden{}
}

/*
ScanArtifactForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type ScanArtifactForbidden struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this scan artifact forbidden response has a 2xx status code
func (o *ScanArtifactForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this scan artifact forbidden response has a 3xx status code
func (o *ScanArtifactForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this scan artifact forbidden response has a 4xx status code
func (o *ScanArtifactForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this scan artifact forbidden response has a 5xx status code
func (o *ScanArtifactForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this scan artifact forbidden response a status code equal to that given
func (o *ScanArtifactForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *ScanArtifactForbidden) Error() string {
	return fmt.Sprintf("[POST /projects/{project_name}/repositories/{repository_name}/artifacts/{reference}/scan][%d] scanArtifactForbidden  %+v", 403, o.Payload)
}

func (o *ScanArtifactForbidden) String() string {
	return fmt.Sprintf("[POST /projects/{project_name}/repositories/{repository_name}/artifacts/{reference}/scan][%d] scanArtifactForbidden  %+v", 403, o.Payload)
}

func (o *ScanArtifactForbidden) GetPayload() *models.Errors {
	return o.Payload
}

func (o *ScanArtifactForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewScanArtifactNotFound creates a ScanArtifactNotFound with default headers values
func NewScanArtifactNotFound() *ScanArtifactNotFound {
	return &ScanArtifactNotFound{}
}

/*
ScanArtifactNotFound describes a response with status code 404, with default header values.

Not found
*/
type ScanArtifactNotFound struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this scan artifact not found response has a 2xx status code
func (o *ScanArtifactNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this scan artifact not found response has a 3xx status code
func (o *ScanArtifactNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this scan artifact not found response has a 4xx status code
func (o *ScanArtifactNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this scan artifact not found response has a 5xx status code
func (o *ScanArtifactNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this scan artifact not found response a status code equal to that given
func (o *ScanArtifactNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *ScanArtifactNotFound) Error() string {
	return fmt.Sprintf("[POST /projects/{project_name}/repositories/{repository_name}/artifacts/{reference}/scan][%d] scanArtifactNotFound  %+v", 404, o.Payload)
}

func (o *ScanArtifactNotFound) String() string {
	return fmt.Sprintf("[POST /projects/{project_name}/repositories/{repository_name}/artifacts/{reference}/scan][%d] scanArtifactNotFound  %+v", 404, o.Payload)
}

func (o *ScanArtifactNotFound) GetPayload() *models.Errors {
	return o.Payload
}

func (o *ScanArtifactNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewScanArtifactInternalServerError creates a ScanArtifactInternalServerError with default headers values
func NewScanArtifactInternalServerError() *ScanArtifactInternalServerError {
	return &ScanArtifactInternalServerError{}
}

/*
ScanArtifactInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type ScanArtifactInternalServerError struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this scan artifact internal server error response has a 2xx status code
func (o *ScanArtifactInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this scan artifact internal server error response has a 3xx status code
func (o *ScanArtifactInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this scan artifact internal server error response has a 4xx status code
func (o *ScanArtifactInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this scan artifact internal server error response has a 5xx status code
func (o *ScanArtifactInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this scan artifact internal server error response a status code equal to that given
func (o *ScanArtifactInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *ScanArtifactInternalServerError) Error() string {
	return fmt.Sprintf("[POST /projects/{project_name}/repositories/{repository_name}/artifacts/{reference}/scan][%d] scanArtifactInternalServerError  %+v", 500, o.Payload)
}

func (o *ScanArtifactInternalServerError) String() string {
	return fmt.Sprintf("[POST /projects/{project_name}/repositories/{repository_name}/artifacts/{reference}/scan][%d] scanArtifactInternalServerError  %+v", 500, o.Payload)
}

func (o *ScanArtifactInternalServerError) GetPayload() *models.Errors {
	return o.Payload
}

func (o *ScanArtifactInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
