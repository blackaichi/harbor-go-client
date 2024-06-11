// Code generated by go-swagger; DO NOT EDIT.

package webhookjob

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

// ListWebhookJobsReader is a Reader for the ListWebhookJobs structure.
type ListWebhookJobsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListWebhookJobsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListWebhookJobsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewListWebhookJobsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewListWebhookJobsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListWebhookJobsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewListWebhookJobsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListWebhookJobsOK creates a ListWebhookJobsOK with default headers values
func NewListWebhookJobsOK() *ListWebhookJobsOK {
	return &ListWebhookJobsOK{}
}

/*
ListWebhookJobsOK describes a response with status code 200, with default header values.

List project webhook jobs successfully.
*/
type ListWebhookJobsOK struct {

	/* Link to previous page and next page
	 */
	Link string

	/* The total count of available items
	 */
	XTotalCount int64

	Payload []*models.WebhookJob
}

// IsSuccess returns true when this list webhook jobs o k response has a 2xx status code
func (o *ListWebhookJobsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list webhook jobs o k response has a 3xx status code
func (o *ListWebhookJobsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list webhook jobs o k response has a 4xx status code
func (o *ListWebhookJobsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list webhook jobs o k response has a 5xx status code
func (o *ListWebhookJobsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list webhook jobs o k response a status code equal to that given
func (o *ListWebhookJobsOK) IsCode(code int) bool {
	return code == 200
}

func (o *ListWebhookJobsOK) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name_or_id}/webhook/jobs][%d] listWebhookJobsOK  %+v", 200, o.Payload)
}

func (o *ListWebhookJobsOK) String() string {
	return fmt.Sprintf("[GET /projects/{project_name_or_id}/webhook/jobs][%d] listWebhookJobsOK  %+v", 200, o.Payload)
}

func (o *ListWebhookJobsOK) GetPayload() []*models.WebhookJob {
	return o.Payload
}

func (o *ListWebhookJobsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewListWebhookJobsBadRequest creates a ListWebhookJobsBadRequest with default headers values
func NewListWebhookJobsBadRequest() *ListWebhookJobsBadRequest {
	return &ListWebhookJobsBadRequest{}
}

/*
ListWebhookJobsBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type ListWebhookJobsBadRequest struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this list webhook jobs bad request response has a 2xx status code
func (o *ListWebhookJobsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list webhook jobs bad request response has a 3xx status code
func (o *ListWebhookJobsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list webhook jobs bad request response has a 4xx status code
func (o *ListWebhookJobsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this list webhook jobs bad request response has a 5xx status code
func (o *ListWebhookJobsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this list webhook jobs bad request response a status code equal to that given
func (o *ListWebhookJobsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *ListWebhookJobsBadRequest) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name_or_id}/webhook/jobs][%d] listWebhookJobsBadRequest  %+v", 400, o.Payload)
}

func (o *ListWebhookJobsBadRequest) String() string {
	return fmt.Sprintf("[GET /projects/{project_name_or_id}/webhook/jobs][%d] listWebhookJobsBadRequest  %+v", 400, o.Payload)
}

func (o *ListWebhookJobsBadRequest) GetPayload() *models.Errors {
	return o.Payload
}

func (o *ListWebhookJobsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewListWebhookJobsUnauthorized creates a ListWebhookJobsUnauthorized with default headers values
func NewListWebhookJobsUnauthorized() *ListWebhookJobsUnauthorized {
	return &ListWebhookJobsUnauthorized{}
}

/*
ListWebhookJobsUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type ListWebhookJobsUnauthorized struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this list webhook jobs unauthorized response has a 2xx status code
func (o *ListWebhookJobsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list webhook jobs unauthorized response has a 3xx status code
func (o *ListWebhookJobsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list webhook jobs unauthorized response has a 4xx status code
func (o *ListWebhookJobsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this list webhook jobs unauthorized response has a 5xx status code
func (o *ListWebhookJobsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this list webhook jobs unauthorized response a status code equal to that given
func (o *ListWebhookJobsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *ListWebhookJobsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name_or_id}/webhook/jobs][%d] listWebhookJobsUnauthorized  %+v", 401, o.Payload)
}

func (o *ListWebhookJobsUnauthorized) String() string {
	return fmt.Sprintf("[GET /projects/{project_name_or_id}/webhook/jobs][%d] listWebhookJobsUnauthorized  %+v", 401, o.Payload)
}

func (o *ListWebhookJobsUnauthorized) GetPayload() *models.Errors {
	return o.Payload
}

func (o *ListWebhookJobsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewListWebhookJobsForbidden creates a ListWebhookJobsForbidden with default headers values
func NewListWebhookJobsForbidden() *ListWebhookJobsForbidden {
	return &ListWebhookJobsForbidden{}
}

/*
ListWebhookJobsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type ListWebhookJobsForbidden struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this list webhook jobs forbidden response has a 2xx status code
func (o *ListWebhookJobsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list webhook jobs forbidden response has a 3xx status code
func (o *ListWebhookJobsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list webhook jobs forbidden response has a 4xx status code
func (o *ListWebhookJobsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this list webhook jobs forbidden response has a 5xx status code
func (o *ListWebhookJobsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this list webhook jobs forbidden response a status code equal to that given
func (o *ListWebhookJobsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *ListWebhookJobsForbidden) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name_or_id}/webhook/jobs][%d] listWebhookJobsForbidden  %+v", 403, o.Payload)
}

func (o *ListWebhookJobsForbidden) String() string {
	return fmt.Sprintf("[GET /projects/{project_name_or_id}/webhook/jobs][%d] listWebhookJobsForbidden  %+v", 403, o.Payload)
}

func (o *ListWebhookJobsForbidden) GetPayload() *models.Errors {
	return o.Payload
}

func (o *ListWebhookJobsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewListWebhookJobsInternalServerError creates a ListWebhookJobsInternalServerError with default headers values
func NewListWebhookJobsInternalServerError() *ListWebhookJobsInternalServerError {
	return &ListWebhookJobsInternalServerError{}
}

/*
ListWebhookJobsInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type ListWebhookJobsInternalServerError struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this list webhook jobs internal server error response has a 2xx status code
func (o *ListWebhookJobsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list webhook jobs internal server error response has a 3xx status code
func (o *ListWebhookJobsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list webhook jobs internal server error response has a 4xx status code
func (o *ListWebhookJobsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this list webhook jobs internal server error response has a 5xx status code
func (o *ListWebhookJobsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this list webhook jobs internal server error response a status code equal to that given
func (o *ListWebhookJobsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *ListWebhookJobsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name_or_id}/webhook/jobs][%d] listWebhookJobsInternalServerError  %+v", 500, o.Payload)
}

func (o *ListWebhookJobsInternalServerError) String() string {
	return fmt.Sprintf("[GET /projects/{project_name_or_id}/webhook/jobs][%d] listWebhookJobsInternalServerError  %+v", 500, o.Payload)
}

func (o *ListWebhookJobsInternalServerError) GetPayload() *models.Errors {
	return o.Payload
}

func (o *ListWebhookJobsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
