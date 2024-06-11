// Code generated by go-swagger; DO NOT EDIT.

package robotv1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/blackaichi/harbor-go-client/pkg/sdk/v2.0/models"
)

// GetRobotByIDV1Reader is a Reader for the GetRobotByIDV1 structure.
type GetRobotByIDV1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRobotByIDV1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRobotByIDV1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetRobotByIDV1Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetRobotByIDV1Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetRobotByIDV1NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetRobotByIDV1InternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetRobotByIDV1OK creates a GetRobotByIDV1OK with default headers values
func NewGetRobotByIDV1OK() *GetRobotByIDV1OK {
	return &GetRobotByIDV1OK{}
}

/*
GetRobotByIDV1OK describes a response with status code 200, with default header values.

Return matched robot information.
*/
type GetRobotByIDV1OK struct {
	Payload *models.Robot
}

// IsSuccess returns true when this get robot by Id v1 o k response has a 2xx status code
func (o *GetRobotByIDV1OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get robot by Id v1 o k response has a 3xx status code
func (o *GetRobotByIDV1OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get robot by Id v1 o k response has a 4xx status code
func (o *GetRobotByIDV1OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get robot by Id v1 o k response has a 5xx status code
func (o *GetRobotByIDV1OK) IsServerError() bool {
	return false
}

// IsCode returns true when this get robot by Id v1 o k response a status code equal to that given
func (o *GetRobotByIDV1OK) IsCode(code int) bool {
	return code == 200
}

func (o *GetRobotByIDV1OK) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name_or_id}/robots/{robot_id}][%d] getRobotByIdV1OK  %+v", 200, o.Payload)
}

func (o *GetRobotByIDV1OK) String() string {
	return fmt.Sprintf("[GET /projects/{project_name_or_id}/robots/{robot_id}][%d] getRobotByIdV1OK  %+v", 200, o.Payload)
}

func (o *GetRobotByIDV1OK) GetPayload() *models.Robot {
	return o.Payload
}

func (o *GetRobotByIDV1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Robot)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRobotByIDV1Unauthorized creates a GetRobotByIDV1Unauthorized with default headers values
func NewGetRobotByIDV1Unauthorized() *GetRobotByIDV1Unauthorized {
	return &GetRobotByIDV1Unauthorized{}
}

/*
GetRobotByIDV1Unauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetRobotByIDV1Unauthorized struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this get robot by Id v1 unauthorized response has a 2xx status code
func (o *GetRobotByIDV1Unauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get robot by Id v1 unauthorized response has a 3xx status code
func (o *GetRobotByIDV1Unauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get robot by Id v1 unauthorized response has a 4xx status code
func (o *GetRobotByIDV1Unauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get robot by Id v1 unauthorized response has a 5xx status code
func (o *GetRobotByIDV1Unauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get robot by Id v1 unauthorized response a status code equal to that given
func (o *GetRobotByIDV1Unauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetRobotByIDV1Unauthorized) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name_or_id}/robots/{robot_id}][%d] getRobotByIdV1Unauthorized  %+v", 401, o.Payload)
}

func (o *GetRobotByIDV1Unauthorized) String() string {
	return fmt.Sprintf("[GET /projects/{project_name_or_id}/robots/{robot_id}][%d] getRobotByIdV1Unauthorized  %+v", 401, o.Payload)
}

func (o *GetRobotByIDV1Unauthorized) GetPayload() *models.Errors {
	return o.Payload
}

func (o *GetRobotByIDV1Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetRobotByIDV1Forbidden creates a GetRobotByIDV1Forbidden with default headers values
func NewGetRobotByIDV1Forbidden() *GetRobotByIDV1Forbidden {
	return &GetRobotByIDV1Forbidden{}
}

/*
GetRobotByIDV1Forbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetRobotByIDV1Forbidden struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this get robot by Id v1 forbidden response has a 2xx status code
func (o *GetRobotByIDV1Forbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get robot by Id v1 forbidden response has a 3xx status code
func (o *GetRobotByIDV1Forbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get robot by Id v1 forbidden response has a 4xx status code
func (o *GetRobotByIDV1Forbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get robot by Id v1 forbidden response has a 5xx status code
func (o *GetRobotByIDV1Forbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get robot by Id v1 forbidden response a status code equal to that given
func (o *GetRobotByIDV1Forbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetRobotByIDV1Forbidden) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name_or_id}/robots/{robot_id}][%d] getRobotByIdV1Forbidden  %+v", 403, o.Payload)
}

func (o *GetRobotByIDV1Forbidden) String() string {
	return fmt.Sprintf("[GET /projects/{project_name_or_id}/robots/{robot_id}][%d] getRobotByIdV1Forbidden  %+v", 403, o.Payload)
}

func (o *GetRobotByIDV1Forbidden) GetPayload() *models.Errors {
	return o.Payload
}

func (o *GetRobotByIDV1Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetRobotByIDV1NotFound creates a GetRobotByIDV1NotFound with default headers values
func NewGetRobotByIDV1NotFound() *GetRobotByIDV1NotFound {
	return &GetRobotByIDV1NotFound{}
}

/*
GetRobotByIDV1NotFound describes a response with status code 404, with default header values.

Not found
*/
type GetRobotByIDV1NotFound struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this get robot by Id v1 not found response has a 2xx status code
func (o *GetRobotByIDV1NotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get robot by Id v1 not found response has a 3xx status code
func (o *GetRobotByIDV1NotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get robot by Id v1 not found response has a 4xx status code
func (o *GetRobotByIDV1NotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get robot by Id v1 not found response has a 5xx status code
func (o *GetRobotByIDV1NotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get robot by Id v1 not found response a status code equal to that given
func (o *GetRobotByIDV1NotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetRobotByIDV1NotFound) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name_or_id}/robots/{robot_id}][%d] getRobotByIdV1NotFound  %+v", 404, o.Payload)
}

func (o *GetRobotByIDV1NotFound) String() string {
	return fmt.Sprintf("[GET /projects/{project_name_or_id}/robots/{robot_id}][%d] getRobotByIdV1NotFound  %+v", 404, o.Payload)
}

func (o *GetRobotByIDV1NotFound) GetPayload() *models.Errors {
	return o.Payload
}

func (o *GetRobotByIDV1NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetRobotByIDV1InternalServerError creates a GetRobotByIDV1InternalServerError with default headers values
func NewGetRobotByIDV1InternalServerError() *GetRobotByIDV1InternalServerError {
	return &GetRobotByIDV1InternalServerError{}
}

/*
GetRobotByIDV1InternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GetRobotByIDV1InternalServerError struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this get robot by Id v1 internal server error response has a 2xx status code
func (o *GetRobotByIDV1InternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get robot by Id v1 internal server error response has a 3xx status code
func (o *GetRobotByIDV1InternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get robot by Id v1 internal server error response has a 4xx status code
func (o *GetRobotByIDV1InternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get robot by Id v1 internal server error response has a 5xx status code
func (o *GetRobotByIDV1InternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get robot by Id v1 internal server error response a status code equal to that given
func (o *GetRobotByIDV1InternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetRobotByIDV1InternalServerError) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name_or_id}/robots/{robot_id}][%d] getRobotByIdV1InternalServerError  %+v", 500, o.Payload)
}

func (o *GetRobotByIDV1InternalServerError) String() string {
	return fmt.Sprintf("[GET /projects/{project_name_or_id}/robots/{robot_id}][%d] getRobotByIdV1InternalServerError  %+v", 500, o.Payload)
}

func (o *GetRobotByIDV1InternalServerError) GetPayload() *models.Errors {
	return o.Payload
}

func (o *GetRobotByIDV1InternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
