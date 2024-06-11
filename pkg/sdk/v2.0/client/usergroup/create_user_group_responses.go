// Code generated by go-swagger; DO NOT EDIT.

package usergroup

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/blackaichi/harbor-go-client/pkg/sdk/v2.0/models"
)

// CreateUserGroupReader is a Reader for the CreateUserGroup structure.
type CreateUserGroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateUserGroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateUserGroupCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateUserGroupBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCreateUserGroupUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateUserGroupForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewCreateUserGroupConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateUserGroupInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateUserGroupCreated creates a CreateUserGroupCreated with default headers values
func NewCreateUserGroupCreated() *CreateUserGroupCreated {
	return &CreateUserGroupCreated{}
}

/*
CreateUserGroupCreated describes a response with status code 201, with default header values.

User group created successfully.
*/
type CreateUserGroupCreated struct {

	/* The URL of the created resource
	 */
	Location string
}

// IsSuccess returns true when this create user group created response has a 2xx status code
func (o *CreateUserGroupCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create user group created response has a 3xx status code
func (o *CreateUserGroupCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create user group created response has a 4xx status code
func (o *CreateUserGroupCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create user group created response has a 5xx status code
func (o *CreateUserGroupCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create user group created response a status code equal to that given
func (o *CreateUserGroupCreated) IsCode(code int) bool {
	return code == 201
}

func (o *CreateUserGroupCreated) Error() string {
	return fmt.Sprintf("[POST /usergroups][%d] createUserGroupCreated ", 201)
}

func (o *CreateUserGroupCreated) String() string {
	return fmt.Sprintf("[POST /usergroups][%d] createUserGroupCreated ", 201)
}

func (o *CreateUserGroupCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Location
	hdrLocation := response.GetHeader("Location")

	if hdrLocation != "" {
		o.Location = hdrLocation
	}

	return nil
}

// NewCreateUserGroupBadRequest creates a CreateUserGroupBadRequest with default headers values
func NewCreateUserGroupBadRequest() *CreateUserGroupBadRequest {
	return &CreateUserGroupBadRequest{}
}

/*
CreateUserGroupBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type CreateUserGroupBadRequest struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this create user group bad request response has a 2xx status code
func (o *CreateUserGroupBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create user group bad request response has a 3xx status code
func (o *CreateUserGroupBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create user group bad request response has a 4xx status code
func (o *CreateUserGroupBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create user group bad request response has a 5xx status code
func (o *CreateUserGroupBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create user group bad request response a status code equal to that given
func (o *CreateUserGroupBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *CreateUserGroupBadRequest) Error() string {
	return fmt.Sprintf("[POST /usergroups][%d] createUserGroupBadRequest  %+v", 400, o.Payload)
}

func (o *CreateUserGroupBadRequest) String() string {
	return fmt.Sprintf("[POST /usergroups][%d] createUserGroupBadRequest  %+v", 400, o.Payload)
}

func (o *CreateUserGroupBadRequest) GetPayload() *models.Errors {
	return o.Payload
}

func (o *CreateUserGroupBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateUserGroupUnauthorized creates a CreateUserGroupUnauthorized with default headers values
func NewCreateUserGroupUnauthorized() *CreateUserGroupUnauthorized {
	return &CreateUserGroupUnauthorized{}
}

/*
CreateUserGroupUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type CreateUserGroupUnauthorized struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this create user group unauthorized response has a 2xx status code
func (o *CreateUserGroupUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create user group unauthorized response has a 3xx status code
func (o *CreateUserGroupUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create user group unauthorized response has a 4xx status code
func (o *CreateUserGroupUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this create user group unauthorized response has a 5xx status code
func (o *CreateUserGroupUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this create user group unauthorized response a status code equal to that given
func (o *CreateUserGroupUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *CreateUserGroupUnauthorized) Error() string {
	return fmt.Sprintf("[POST /usergroups][%d] createUserGroupUnauthorized  %+v", 401, o.Payload)
}

func (o *CreateUserGroupUnauthorized) String() string {
	return fmt.Sprintf("[POST /usergroups][%d] createUserGroupUnauthorized  %+v", 401, o.Payload)
}

func (o *CreateUserGroupUnauthorized) GetPayload() *models.Errors {
	return o.Payload
}

func (o *CreateUserGroupUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateUserGroupForbidden creates a CreateUserGroupForbidden with default headers values
func NewCreateUserGroupForbidden() *CreateUserGroupForbidden {
	return &CreateUserGroupForbidden{}
}

/*
CreateUserGroupForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type CreateUserGroupForbidden struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this create user group forbidden response has a 2xx status code
func (o *CreateUserGroupForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create user group forbidden response has a 3xx status code
func (o *CreateUserGroupForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create user group forbidden response has a 4xx status code
func (o *CreateUserGroupForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this create user group forbidden response has a 5xx status code
func (o *CreateUserGroupForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this create user group forbidden response a status code equal to that given
func (o *CreateUserGroupForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *CreateUserGroupForbidden) Error() string {
	return fmt.Sprintf("[POST /usergroups][%d] createUserGroupForbidden  %+v", 403, o.Payload)
}

func (o *CreateUserGroupForbidden) String() string {
	return fmt.Sprintf("[POST /usergroups][%d] createUserGroupForbidden  %+v", 403, o.Payload)
}

func (o *CreateUserGroupForbidden) GetPayload() *models.Errors {
	return o.Payload
}

func (o *CreateUserGroupForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateUserGroupConflict creates a CreateUserGroupConflict with default headers values
func NewCreateUserGroupConflict() *CreateUserGroupConflict {
	return &CreateUserGroupConflict{}
}

/*
CreateUserGroupConflict describes a response with status code 409, with default header values.

Conflict
*/
type CreateUserGroupConflict struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this create user group conflict response has a 2xx status code
func (o *CreateUserGroupConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create user group conflict response has a 3xx status code
func (o *CreateUserGroupConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create user group conflict response has a 4xx status code
func (o *CreateUserGroupConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this create user group conflict response has a 5xx status code
func (o *CreateUserGroupConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this create user group conflict response a status code equal to that given
func (o *CreateUserGroupConflict) IsCode(code int) bool {
	return code == 409
}

func (o *CreateUserGroupConflict) Error() string {
	return fmt.Sprintf("[POST /usergroups][%d] createUserGroupConflict  %+v", 409, o.Payload)
}

func (o *CreateUserGroupConflict) String() string {
	return fmt.Sprintf("[POST /usergroups][%d] createUserGroupConflict  %+v", 409, o.Payload)
}

func (o *CreateUserGroupConflict) GetPayload() *models.Errors {
	return o.Payload
}

func (o *CreateUserGroupConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateUserGroupInternalServerError creates a CreateUserGroupInternalServerError with default headers values
func NewCreateUserGroupInternalServerError() *CreateUserGroupInternalServerError {
	return &CreateUserGroupInternalServerError{}
}

/*
CreateUserGroupInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type CreateUserGroupInternalServerError struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this create user group internal server error response has a 2xx status code
func (o *CreateUserGroupInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create user group internal server error response has a 3xx status code
func (o *CreateUserGroupInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create user group internal server error response has a 4xx status code
func (o *CreateUserGroupInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this create user group internal server error response has a 5xx status code
func (o *CreateUserGroupInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this create user group internal server error response a status code equal to that given
func (o *CreateUserGroupInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *CreateUserGroupInternalServerError) Error() string {
	return fmt.Sprintf("[POST /usergroups][%d] createUserGroupInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateUserGroupInternalServerError) String() string {
	return fmt.Sprintf("[POST /usergroups][%d] createUserGroupInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateUserGroupInternalServerError) GetPayload() *models.Errors {
	return o.Payload
}

func (o *CreateUserGroupInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
