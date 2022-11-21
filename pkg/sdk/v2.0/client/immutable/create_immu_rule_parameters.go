// Code generated by go-swagger; DO NOT EDIT.

package immutable

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/goharbor/go-client/pkg/sdk/v2.0/models"
)

// NewCreateImmuRuleParams creates a new CreateImmuRuleParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateImmuRuleParams() *CreateImmuRuleParams {
	return &CreateImmuRuleParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateImmuRuleParamsWithTimeout creates a new CreateImmuRuleParams object
// with the ability to set a timeout on a request.
func NewCreateImmuRuleParamsWithTimeout(timeout time.Duration) *CreateImmuRuleParams {
	return &CreateImmuRuleParams{
		timeout: timeout,
	}
}

// NewCreateImmuRuleParamsWithContext creates a new CreateImmuRuleParams object
// with the ability to set a context for a request.
func NewCreateImmuRuleParamsWithContext(ctx context.Context) *CreateImmuRuleParams {
	return &CreateImmuRuleParams{
		Context: ctx,
	}
}

// NewCreateImmuRuleParamsWithHTTPClient creates a new CreateImmuRuleParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateImmuRuleParamsWithHTTPClient(client *http.Client) *CreateImmuRuleParams {
	return &CreateImmuRuleParams{
		HTTPClient: client,
	}
}

/*
CreateImmuRuleParams contains all the parameters to send to the API endpoint

	for the create immu rule operation.

	Typically these are written to a http.Request.
*/
type CreateImmuRuleParams struct {

	// ImmutableRule.
	ImmutableRule *models.ImmutableRule

	/* XIsResourceName.

	   The flag to indicate whether the parameter which supports both name and id in the path is the name of the resource. When the X-Is-Resource-Name is false and the parameter can be converted to an integer, the parameter will be as an id, otherwise, it will be as a name.
	*/
	XIsResourceName *bool

	/* XRequestID.

	   An unique ID for the request
	*/
	XRequestID *string

	/* ProjectNameOrID.

	   The name or id of the project
	*/
	ProjectNameOrID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create immu rule params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateImmuRuleParams) WithDefaults() *CreateImmuRuleParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create immu rule params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateImmuRuleParams) SetDefaults() {
	var (
		xIsResourceNameDefault = bool(false)
	)

	val := CreateImmuRuleParams{
		XIsResourceName: &xIsResourceNameDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the create immu rule params
func (o *CreateImmuRuleParams) WithTimeout(timeout time.Duration) *CreateImmuRuleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create immu rule params
func (o *CreateImmuRuleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create immu rule params
func (o *CreateImmuRuleParams) WithContext(ctx context.Context) *CreateImmuRuleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create immu rule params
func (o *CreateImmuRuleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create immu rule params
func (o *CreateImmuRuleParams) WithHTTPClient(client *http.Client) *CreateImmuRuleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create immu rule params
func (o *CreateImmuRuleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithImmutableRule adds the immutableRule to the create immu rule params
func (o *CreateImmuRuleParams) WithImmutableRule(immutableRule *models.ImmutableRule) *CreateImmuRuleParams {
	o.SetImmutableRule(immutableRule)
	return o
}

// SetImmutableRule adds the immutableRule to the create immu rule params
func (o *CreateImmuRuleParams) SetImmutableRule(immutableRule *models.ImmutableRule) {
	o.ImmutableRule = immutableRule
}

// WithXIsResourceName adds the xIsResourceName to the create immu rule params
func (o *CreateImmuRuleParams) WithXIsResourceName(xIsResourceName *bool) *CreateImmuRuleParams {
	o.SetXIsResourceName(xIsResourceName)
	return o
}

// SetXIsResourceName adds the xIsResourceName to the create immu rule params
func (o *CreateImmuRuleParams) SetXIsResourceName(xIsResourceName *bool) {
	o.XIsResourceName = xIsResourceName
}

// WithXRequestID adds the xRequestID to the create immu rule params
func (o *CreateImmuRuleParams) WithXRequestID(xRequestID *string) *CreateImmuRuleParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the create immu rule params
func (o *CreateImmuRuleParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithProjectNameOrID adds the projectNameOrID to the create immu rule params
func (o *CreateImmuRuleParams) WithProjectNameOrID(projectNameOrID string) *CreateImmuRuleParams {
	o.SetProjectNameOrID(projectNameOrID)
	return o
}

// SetProjectNameOrID adds the projectNameOrId to the create immu rule params
func (o *CreateImmuRuleParams) SetProjectNameOrID(projectNameOrID string) {
	o.ProjectNameOrID = projectNameOrID
}

// WriteToRequest writes these params to a swagger request
func (o *CreateImmuRuleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.ImmutableRule != nil {
		if err := r.SetBodyParam(o.ImmutableRule); err != nil {
			return err
		}
	}

	if o.XIsResourceName != nil {

		// header param X-Is-Resource-Name
		if err := r.SetHeaderParam("X-Is-Resource-Name", swag.FormatBool(*o.XIsResourceName)); err != nil {
			return err
		}
	}

	if o.XRequestID != nil {

		// header param X-Request-Id
		if err := r.SetHeaderParam("X-Request-Id", *o.XRequestID); err != nil {
			return err
		}
	}

	// path param project_name_or_id
	if err := r.SetPathParam("project_name_or_id", o.ProjectNameOrID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
