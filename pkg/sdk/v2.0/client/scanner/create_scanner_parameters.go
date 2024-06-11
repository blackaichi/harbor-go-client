// Code generated by go-swagger; DO NOT EDIT.

package scanner

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

	"github.com/blackaichi/harbor-go-client/pkg/sdk/v2.0/models"
)

// NewCreateScannerParams creates a new CreateScannerParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateScannerParams() *CreateScannerParams {
	return &CreateScannerParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateScannerParamsWithTimeout creates a new CreateScannerParams object
// with the ability to set a timeout on a request.
func NewCreateScannerParamsWithTimeout(timeout time.Duration) *CreateScannerParams {
	return &CreateScannerParams{
		timeout: timeout,
	}
}

// NewCreateScannerParamsWithContext creates a new CreateScannerParams object
// with the ability to set a context for a request.
func NewCreateScannerParamsWithContext(ctx context.Context) *CreateScannerParams {
	return &CreateScannerParams{
		Context: ctx,
	}
}

// NewCreateScannerParamsWithHTTPClient creates a new CreateScannerParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateScannerParamsWithHTTPClient(client *http.Client) *CreateScannerParams {
	return &CreateScannerParams{
		HTTPClient: client,
	}
}

/*
CreateScannerParams contains all the parameters to send to the API endpoint

	for the create scanner operation.

	Typically these are written to a http.Request.
*/
type CreateScannerParams struct {

	/* XRequestID.

	   An unique ID for the request
	*/
	XRequestID *string

	/* Registration.

	   A scanner registration to be created.
	*/
	Registration *models.ScannerRegistrationReq

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create scanner params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateScannerParams) WithDefaults() *CreateScannerParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create scanner params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateScannerParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create scanner params
func (o *CreateScannerParams) WithTimeout(timeout time.Duration) *CreateScannerParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create scanner params
func (o *CreateScannerParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create scanner params
func (o *CreateScannerParams) WithContext(ctx context.Context) *CreateScannerParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create scanner params
func (o *CreateScannerParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create scanner params
func (o *CreateScannerParams) WithHTTPClient(client *http.Client) *CreateScannerParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create scanner params
func (o *CreateScannerParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the create scanner params
func (o *CreateScannerParams) WithXRequestID(xRequestID *string) *CreateScannerParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the create scanner params
func (o *CreateScannerParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithRegistration adds the registration to the create scanner params
func (o *CreateScannerParams) WithRegistration(registration *models.ScannerRegistrationReq) *CreateScannerParams {
	o.SetRegistration(registration)
	return o
}

// SetRegistration adds the registration to the create scanner params
func (o *CreateScannerParams) SetRegistration(registration *models.ScannerRegistrationReq) {
	o.Registration = registration
}

// WriteToRequest writes these params to a swagger request
func (o *CreateScannerParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.XRequestID != nil {

		// header param X-Request-Id
		if err := r.SetHeaderParam("X-Request-Id", *o.XRequestID); err != nil {
			return err
		}
	}
	if o.Registration != nil {
		if err := r.SetBodyParam(o.Registration); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
