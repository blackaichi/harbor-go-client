// Code generated by go-swagger; DO NOT EDIT.

package preheat

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/goharbor/go-client/pkg/sdk/v2.0/models"
)

// DeletePolicyReader is a Reader for the DeletePolicy structure.
type DeletePolicyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeletePolicyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeletePolicyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeletePolicyBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeletePolicyUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeletePolicyForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeletePolicyNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeletePolicyInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeletePolicyOK creates a DeletePolicyOK with default headers values
func NewDeletePolicyOK() *DeletePolicyOK {
	return &DeletePolicyOK{}
}

/*
DeletePolicyOK describes a response with status code 200, with default header values.

Success
*/
type DeletePolicyOK struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string
}

// IsSuccess returns true when this delete policy o k response has a 2xx status code
func (o *DeletePolicyOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete policy o k response has a 3xx status code
func (o *DeletePolicyOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete policy o k response has a 4xx status code
func (o *DeletePolicyOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete policy o k response has a 5xx status code
func (o *DeletePolicyOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete policy o k response a status code equal to that given
func (o *DeletePolicyOK) IsCode(code int) bool {
	return code == 200
}

func (o *DeletePolicyOK) Error() string {
	return fmt.Sprintf("[DELETE /projects/{project_name}/preheat/policies/{preheat_policy_name}][%d] deletePolicyOK ", 200)
}

func (o *DeletePolicyOK) String() string {
	return fmt.Sprintf("[DELETE /projects/{project_name}/preheat/policies/{preheat_policy_name}][%d] deletePolicyOK ", 200)
}

func (o *DeletePolicyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-Request-Id
	hdrXRequestID := response.GetHeader("X-Request-Id")

	if hdrXRequestID != "" {
		o.XRequestID = hdrXRequestID
	}

	return nil
}

// NewDeletePolicyBadRequest creates a DeletePolicyBadRequest with default headers values
func NewDeletePolicyBadRequest() *DeletePolicyBadRequest {
	return &DeletePolicyBadRequest{}
}

/*
DeletePolicyBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type DeletePolicyBadRequest struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this delete policy bad request response has a 2xx status code
func (o *DeletePolicyBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete policy bad request response has a 3xx status code
func (o *DeletePolicyBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete policy bad request response has a 4xx status code
func (o *DeletePolicyBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete policy bad request response has a 5xx status code
func (o *DeletePolicyBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this delete policy bad request response a status code equal to that given
func (o *DeletePolicyBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *DeletePolicyBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /projects/{project_name}/preheat/policies/{preheat_policy_name}][%d] deletePolicyBadRequest  %+v", 400, o.Payload)
}

func (o *DeletePolicyBadRequest) String() string {
	return fmt.Sprintf("[DELETE /projects/{project_name}/preheat/policies/{preheat_policy_name}][%d] deletePolicyBadRequest  %+v", 400, o.Payload)
}

func (o *DeletePolicyBadRequest) GetPayload() *models.Errors {
	return o.Payload
}

func (o *DeletePolicyBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeletePolicyUnauthorized creates a DeletePolicyUnauthorized with default headers values
func NewDeletePolicyUnauthorized() *DeletePolicyUnauthorized {
	return &DeletePolicyUnauthorized{}
}

/*
DeletePolicyUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type DeletePolicyUnauthorized struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this delete policy unauthorized response has a 2xx status code
func (o *DeletePolicyUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete policy unauthorized response has a 3xx status code
func (o *DeletePolicyUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete policy unauthorized response has a 4xx status code
func (o *DeletePolicyUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete policy unauthorized response has a 5xx status code
func (o *DeletePolicyUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this delete policy unauthorized response a status code equal to that given
func (o *DeletePolicyUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *DeletePolicyUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /projects/{project_name}/preheat/policies/{preheat_policy_name}][%d] deletePolicyUnauthorized  %+v", 401, o.Payload)
}

func (o *DeletePolicyUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /projects/{project_name}/preheat/policies/{preheat_policy_name}][%d] deletePolicyUnauthorized  %+v", 401, o.Payload)
}

func (o *DeletePolicyUnauthorized) GetPayload() *models.Errors {
	return o.Payload
}

func (o *DeletePolicyUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeletePolicyForbidden creates a DeletePolicyForbidden with default headers values
func NewDeletePolicyForbidden() *DeletePolicyForbidden {
	return &DeletePolicyForbidden{}
}

/*
DeletePolicyForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type DeletePolicyForbidden struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this delete policy forbidden response has a 2xx status code
func (o *DeletePolicyForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete policy forbidden response has a 3xx status code
func (o *DeletePolicyForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete policy forbidden response has a 4xx status code
func (o *DeletePolicyForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete policy forbidden response has a 5xx status code
func (o *DeletePolicyForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete policy forbidden response a status code equal to that given
func (o *DeletePolicyForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *DeletePolicyForbidden) Error() string {
	return fmt.Sprintf("[DELETE /projects/{project_name}/preheat/policies/{preheat_policy_name}][%d] deletePolicyForbidden  %+v", 403, o.Payload)
}

func (o *DeletePolicyForbidden) String() string {
	return fmt.Sprintf("[DELETE /projects/{project_name}/preheat/policies/{preheat_policy_name}][%d] deletePolicyForbidden  %+v", 403, o.Payload)
}

func (o *DeletePolicyForbidden) GetPayload() *models.Errors {
	return o.Payload
}

func (o *DeletePolicyForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeletePolicyNotFound creates a DeletePolicyNotFound with default headers values
func NewDeletePolicyNotFound() *DeletePolicyNotFound {
	return &DeletePolicyNotFound{}
}

/*
DeletePolicyNotFound describes a response with status code 404, with default header values.

Not found
*/
type DeletePolicyNotFound struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this delete policy not found response has a 2xx status code
func (o *DeletePolicyNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete policy not found response has a 3xx status code
func (o *DeletePolicyNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete policy not found response has a 4xx status code
func (o *DeletePolicyNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete policy not found response has a 5xx status code
func (o *DeletePolicyNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete policy not found response a status code equal to that given
func (o *DeletePolicyNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *DeletePolicyNotFound) Error() string {
	return fmt.Sprintf("[DELETE /projects/{project_name}/preheat/policies/{preheat_policy_name}][%d] deletePolicyNotFound  %+v", 404, o.Payload)
}

func (o *DeletePolicyNotFound) String() string {
	return fmt.Sprintf("[DELETE /projects/{project_name}/preheat/policies/{preheat_policy_name}][%d] deletePolicyNotFound  %+v", 404, o.Payload)
}

func (o *DeletePolicyNotFound) GetPayload() *models.Errors {
	return o.Payload
}

func (o *DeletePolicyNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeletePolicyInternalServerError creates a DeletePolicyInternalServerError with default headers values
func NewDeletePolicyInternalServerError() *DeletePolicyInternalServerError {
	return &DeletePolicyInternalServerError{}
}

/*
DeletePolicyInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type DeletePolicyInternalServerError struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this delete policy internal server error response has a 2xx status code
func (o *DeletePolicyInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete policy internal server error response has a 3xx status code
func (o *DeletePolicyInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete policy internal server error response has a 4xx status code
func (o *DeletePolicyInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete policy internal server error response has a 5xx status code
func (o *DeletePolicyInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this delete policy internal server error response a status code equal to that given
func (o *DeletePolicyInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *DeletePolicyInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /projects/{project_name}/preheat/policies/{preheat_policy_name}][%d] deletePolicyInternalServerError  %+v", 500, o.Payload)
}

func (o *DeletePolicyInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /projects/{project_name}/preheat/policies/{preheat_policy_name}][%d] deletePolicyInternalServerError  %+v", 500, o.Payload)
}

func (o *DeletePolicyInternalServerError) GetPayload() *models.Errors {
	return o.Payload
}

func (o *DeletePolicyInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
