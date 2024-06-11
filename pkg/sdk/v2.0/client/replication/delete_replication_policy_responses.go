// Code generated by go-swagger; DO NOT EDIT.

package replication

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/blackaichi/harbor-go-client/pkg/sdk/v2.0/models"
)

// DeleteReplicationPolicyReader is a Reader for the DeleteReplicationPolicy structure.
type DeleteReplicationPolicyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteReplicationPolicyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteReplicationPolicyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteReplicationPolicyUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteReplicationPolicyForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteReplicationPolicyNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 412:
		result := NewDeleteReplicationPolicyPreconditionFailed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteReplicationPolicyInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteReplicationPolicyOK creates a DeleteReplicationPolicyOK with default headers values
func NewDeleteReplicationPolicyOK() *DeleteReplicationPolicyOK {
	return &DeleteReplicationPolicyOK{}
}

/*
DeleteReplicationPolicyOK describes a response with status code 200, with default header values.

Success
*/
type DeleteReplicationPolicyOK struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string
}

// IsSuccess returns true when this delete replication policy o k response has a 2xx status code
func (o *DeleteReplicationPolicyOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete replication policy o k response has a 3xx status code
func (o *DeleteReplicationPolicyOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete replication policy o k response has a 4xx status code
func (o *DeleteReplicationPolicyOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete replication policy o k response has a 5xx status code
func (o *DeleteReplicationPolicyOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete replication policy o k response a status code equal to that given
func (o *DeleteReplicationPolicyOK) IsCode(code int) bool {
	return code == 200
}

func (o *DeleteReplicationPolicyOK) Error() string {
	return fmt.Sprintf("[DELETE /replication/policies/{id}][%d] deleteReplicationPolicyOK ", 200)
}

func (o *DeleteReplicationPolicyOK) String() string {
	return fmt.Sprintf("[DELETE /replication/policies/{id}][%d] deleteReplicationPolicyOK ", 200)
}

func (o *DeleteReplicationPolicyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-Request-Id
	hdrXRequestID := response.GetHeader("X-Request-Id")

	if hdrXRequestID != "" {
		o.XRequestID = hdrXRequestID
	}

	return nil
}

// NewDeleteReplicationPolicyUnauthorized creates a DeleteReplicationPolicyUnauthorized with default headers values
func NewDeleteReplicationPolicyUnauthorized() *DeleteReplicationPolicyUnauthorized {
	return &DeleteReplicationPolicyUnauthorized{}
}

/*
DeleteReplicationPolicyUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type DeleteReplicationPolicyUnauthorized struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this delete replication policy unauthorized response has a 2xx status code
func (o *DeleteReplicationPolicyUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete replication policy unauthorized response has a 3xx status code
func (o *DeleteReplicationPolicyUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete replication policy unauthorized response has a 4xx status code
func (o *DeleteReplicationPolicyUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete replication policy unauthorized response has a 5xx status code
func (o *DeleteReplicationPolicyUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this delete replication policy unauthorized response a status code equal to that given
func (o *DeleteReplicationPolicyUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *DeleteReplicationPolicyUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /replication/policies/{id}][%d] deleteReplicationPolicyUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteReplicationPolicyUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /replication/policies/{id}][%d] deleteReplicationPolicyUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteReplicationPolicyUnauthorized) GetPayload() *models.Errors {
	return o.Payload
}

func (o *DeleteReplicationPolicyUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteReplicationPolicyForbidden creates a DeleteReplicationPolicyForbidden with default headers values
func NewDeleteReplicationPolicyForbidden() *DeleteReplicationPolicyForbidden {
	return &DeleteReplicationPolicyForbidden{}
}

/*
DeleteReplicationPolicyForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type DeleteReplicationPolicyForbidden struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this delete replication policy forbidden response has a 2xx status code
func (o *DeleteReplicationPolicyForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete replication policy forbidden response has a 3xx status code
func (o *DeleteReplicationPolicyForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete replication policy forbidden response has a 4xx status code
func (o *DeleteReplicationPolicyForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete replication policy forbidden response has a 5xx status code
func (o *DeleteReplicationPolicyForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete replication policy forbidden response a status code equal to that given
func (o *DeleteReplicationPolicyForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *DeleteReplicationPolicyForbidden) Error() string {
	return fmt.Sprintf("[DELETE /replication/policies/{id}][%d] deleteReplicationPolicyForbidden  %+v", 403, o.Payload)
}

func (o *DeleteReplicationPolicyForbidden) String() string {
	return fmt.Sprintf("[DELETE /replication/policies/{id}][%d] deleteReplicationPolicyForbidden  %+v", 403, o.Payload)
}

func (o *DeleteReplicationPolicyForbidden) GetPayload() *models.Errors {
	return o.Payload
}

func (o *DeleteReplicationPolicyForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteReplicationPolicyNotFound creates a DeleteReplicationPolicyNotFound with default headers values
func NewDeleteReplicationPolicyNotFound() *DeleteReplicationPolicyNotFound {
	return &DeleteReplicationPolicyNotFound{}
}

/*
DeleteReplicationPolicyNotFound describes a response with status code 404, with default header values.

Not found
*/
type DeleteReplicationPolicyNotFound struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this delete replication policy not found response has a 2xx status code
func (o *DeleteReplicationPolicyNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete replication policy not found response has a 3xx status code
func (o *DeleteReplicationPolicyNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete replication policy not found response has a 4xx status code
func (o *DeleteReplicationPolicyNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete replication policy not found response has a 5xx status code
func (o *DeleteReplicationPolicyNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete replication policy not found response a status code equal to that given
func (o *DeleteReplicationPolicyNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *DeleteReplicationPolicyNotFound) Error() string {
	return fmt.Sprintf("[DELETE /replication/policies/{id}][%d] deleteReplicationPolicyNotFound  %+v", 404, o.Payload)
}

func (o *DeleteReplicationPolicyNotFound) String() string {
	return fmt.Sprintf("[DELETE /replication/policies/{id}][%d] deleteReplicationPolicyNotFound  %+v", 404, o.Payload)
}

func (o *DeleteReplicationPolicyNotFound) GetPayload() *models.Errors {
	return o.Payload
}

func (o *DeleteReplicationPolicyNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteReplicationPolicyPreconditionFailed creates a DeleteReplicationPolicyPreconditionFailed with default headers values
func NewDeleteReplicationPolicyPreconditionFailed() *DeleteReplicationPolicyPreconditionFailed {
	return &DeleteReplicationPolicyPreconditionFailed{}
}

/*
DeleteReplicationPolicyPreconditionFailed describes a response with status code 412, with default header values.

Precondition failed
*/
type DeleteReplicationPolicyPreconditionFailed struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this delete replication policy precondition failed response has a 2xx status code
func (o *DeleteReplicationPolicyPreconditionFailed) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete replication policy precondition failed response has a 3xx status code
func (o *DeleteReplicationPolicyPreconditionFailed) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete replication policy precondition failed response has a 4xx status code
func (o *DeleteReplicationPolicyPreconditionFailed) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete replication policy precondition failed response has a 5xx status code
func (o *DeleteReplicationPolicyPreconditionFailed) IsServerError() bool {
	return false
}

// IsCode returns true when this delete replication policy precondition failed response a status code equal to that given
func (o *DeleteReplicationPolicyPreconditionFailed) IsCode(code int) bool {
	return code == 412
}

func (o *DeleteReplicationPolicyPreconditionFailed) Error() string {
	return fmt.Sprintf("[DELETE /replication/policies/{id}][%d] deleteReplicationPolicyPreconditionFailed  %+v", 412, o.Payload)
}

func (o *DeleteReplicationPolicyPreconditionFailed) String() string {
	return fmt.Sprintf("[DELETE /replication/policies/{id}][%d] deleteReplicationPolicyPreconditionFailed  %+v", 412, o.Payload)
}

func (o *DeleteReplicationPolicyPreconditionFailed) GetPayload() *models.Errors {
	return o.Payload
}

func (o *DeleteReplicationPolicyPreconditionFailed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteReplicationPolicyInternalServerError creates a DeleteReplicationPolicyInternalServerError with default headers values
func NewDeleteReplicationPolicyInternalServerError() *DeleteReplicationPolicyInternalServerError {
	return &DeleteReplicationPolicyInternalServerError{}
}

/*
DeleteReplicationPolicyInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type DeleteReplicationPolicyInternalServerError struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this delete replication policy internal server error response has a 2xx status code
func (o *DeleteReplicationPolicyInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete replication policy internal server error response has a 3xx status code
func (o *DeleteReplicationPolicyInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete replication policy internal server error response has a 4xx status code
func (o *DeleteReplicationPolicyInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete replication policy internal server error response has a 5xx status code
func (o *DeleteReplicationPolicyInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this delete replication policy internal server error response a status code equal to that given
func (o *DeleteReplicationPolicyInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *DeleteReplicationPolicyInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /replication/policies/{id}][%d] deleteReplicationPolicyInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteReplicationPolicyInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /replication/policies/{id}][%d] deleteReplicationPolicyInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteReplicationPolicyInternalServerError) GetPayload() *models.Errors {
	return o.Payload
}

func (o *DeleteReplicationPolicyInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
