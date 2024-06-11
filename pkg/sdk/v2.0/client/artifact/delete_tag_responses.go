// Code generated by go-swagger; DO NOT EDIT.

package artifact

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/blackaichi/harbor-go-client/pkg/sdk/v2.0/models"
)

// DeleteTagReader is a Reader for the DeleteTag structure.
type DeleteTagReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteTagReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteTagOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteTagUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteTagForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteTagNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteTagInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteTagOK creates a DeleteTagOK with default headers values
func NewDeleteTagOK() *DeleteTagOK {
	return &DeleteTagOK{}
}

/*
DeleteTagOK describes a response with status code 200, with default header values.

Success
*/
type DeleteTagOK struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string
}

// IsSuccess returns true when this delete tag o k response has a 2xx status code
func (o *DeleteTagOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete tag o k response has a 3xx status code
func (o *DeleteTagOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete tag o k response has a 4xx status code
func (o *DeleteTagOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete tag o k response has a 5xx status code
func (o *DeleteTagOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete tag o k response a status code equal to that given
func (o *DeleteTagOK) IsCode(code int) bool {
	return code == 200
}

func (o *DeleteTagOK) Error() string {
	return fmt.Sprintf("[DELETE /projects/{project_name}/repositories/{repository_name}/artifacts/{reference}/tags/{tag_name}][%d] deleteTagOK ", 200)
}

func (o *DeleteTagOK) String() string {
	return fmt.Sprintf("[DELETE /projects/{project_name}/repositories/{repository_name}/artifacts/{reference}/tags/{tag_name}][%d] deleteTagOK ", 200)
}

func (o *DeleteTagOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-Request-Id
	hdrXRequestID := response.GetHeader("X-Request-Id")

	if hdrXRequestID != "" {
		o.XRequestID = hdrXRequestID
	}

	return nil
}

// NewDeleteTagUnauthorized creates a DeleteTagUnauthorized with default headers values
func NewDeleteTagUnauthorized() *DeleteTagUnauthorized {
	return &DeleteTagUnauthorized{}
}

/*
DeleteTagUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type DeleteTagUnauthorized struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this delete tag unauthorized response has a 2xx status code
func (o *DeleteTagUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete tag unauthorized response has a 3xx status code
func (o *DeleteTagUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete tag unauthorized response has a 4xx status code
func (o *DeleteTagUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete tag unauthorized response has a 5xx status code
func (o *DeleteTagUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this delete tag unauthorized response a status code equal to that given
func (o *DeleteTagUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *DeleteTagUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /projects/{project_name}/repositories/{repository_name}/artifacts/{reference}/tags/{tag_name}][%d] deleteTagUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteTagUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /projects/{project_name}/repositories/{repository_name}/artifacts/{reference}/tags/{tag_name}][%d] deleteTagUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteTagUnauthorized) GetPayload() *models.Errors {
	return o.Payload
}

func (o *DeleteTagUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteTagForbidden creates a DeleteTagForbidden with default headers values
func NewDeleteTagForbidden() *DeleteTagForbidden {
	return &DeleteTagForbidden{}
}

/*
DeleteTagForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type DeleteTagForbidden struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this delete tag forbidden response has a 2xx status code
func (o *DeleteTagForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete tag forbidden response has a 3xx status code
func (o *DeleteTagForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete tag forbidden response has a 4xx status code
func (o *DeleteTagForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete tag forbidden response has a 5xx status code
func (o *DeleteTagForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete tag forbidden response a status code equal to that given
func (o *DeleteTagForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *DeleteTagForbidden) Error() string {
	return fmt.Sprintf("[DELETE /projects/{project_name}/repositories/{repository_name}/artifacts/{reference}/tags/{tag_name}][%d] deleteTagForbidden  %+v", 403, o.Payload)
}

func (o *DeleteTagForbidden) String() string {
	return fmt.Sprintf("[DELETE /projects/{project_name}/repositories/{repository_name}/artifacts/{reference}/tags/{tag_name}][%d] deleteTagForbidden  %+v", 403, o.Payload)
}

func (o *DeleteTagForbidden) GetPayload() *models.Errors {
	return o.Payload
}

func (o *DeleteTagForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteTagNotFound creates a DeleteTagNotFound with default headers values
func NewDeleteTagNotFound() *DeleteTagNotFound {
	return &DeleteTagNotFound{}
}

/*
DeleteTagNotFound describes a response with status code 404, with default header values.

Not found
*/
type DeleteTagNotFound struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this delete tag not found response has a 2xx status code
func (o *DeleteTagNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete tag not found response has a 3xx status code
func (o *DeleteTagNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete tag not found response has a 4xx status code
func (o *DeleteTagNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete tag not found response has a 5xx status code
func (o *DeleteTagNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete tag not found response a status code equal to that given
func (o *DeleteTagNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *DeleteTagNotFound) Error() string {
	return fmt.Sprintf("[DELETE /projects/{project_name}/repositories/{repository_name}/artifacts/{reference}/tags/{tag_name}][%d] deleteTagNotFound  %+v", 404, o.Payload)
}

func (o *DeleteTagNotFound) String() string {
	return fmt.Sprintf("[DELETE /projects/{project_name}/repositories/{repository_name}/artifacts/{reference}/tags/{tag_name}][%d] deleteTagNotFound  %+v", 404, o.Payload)
}

func (o *DeleteTagNotFound) GetPayload() *models.Errors {
	return o.Payload
}

func (o *DeleteTagNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteTagInternalServerError creates a DeleteTagInternalServerError with default headers values
func NewDeleteTagInternalServerError() *DeleteTagInternalServerError {
	return &DeleteTagInternalServerError{}
}

/*
DeleteTagInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type DeleteTagInternalServerError struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this delete tag internal server error response has a 2xx status code
func (o *DeleteTagInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete tag internal server error response has a 3xx status code
func (o *DeleteTagInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete tag internal server error response has a 4xx status code
func (o *DeleteTagInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete tag internal server error response has a 5xx status code
func (o *DeleteTagInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this delete tag internal server error response a status code equal to that given
func (o *DeleteTagInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *DeleteTagInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /projects/{project_name}/repositories/{repository_name}/artifacts/{reference}/tags/{tag_name}][%d] deleteTagInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteTagInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /projects/{project_name}/repositories/{repository_name}/artifacts/{reference}/tags/{tag_name}][%d] deleteTagInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteTagInternalServerError) GetPayload() *models.Errors {
	return o.Payload
}

func (o *DeleteTagInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
