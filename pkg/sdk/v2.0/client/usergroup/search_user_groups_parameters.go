// Code generated by go-swagger; DO NOT EDIT.

package usergroup

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
)

// NewSearchUserGroupsParams creates a new SearchUserGroupsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSearchUserGroupsParams() *SearchUserGroupsParams {
	return &SearchUserGroupsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSearchUserGroupsParamsWithTimeout creates a new SearchUserGroupsParams object
// with the ability to set a timeout on a request.
func NewSearchUserGroupsParamsWithTimeout(timeout time.Duration) *SearchUserGroupsParams {
	return &SearchUserGroupsParams{
		timeout: timeout,
	}
}

// NewSearchUserGroupsParamsWithContext creates a new SearchUserGroupsParams object
// with the ability to set a context for a request.
func NewSearchUserGroupsParamsWithContext(ctx context.Context) *SearchUserGroupsParams {
	return &SearchUserGroupsParams{
		Context: ctx,
	}
}

// NewSearchUserGroupsParamsWithHTTPClient creates a new SearchUserGroupsParams object
// with the ability to set a custom HTTPClient for a request.
func NewSearchUserGroupsParamsWithHTTPClient(client *http.Client) *SearchUserGroupsParams {
	return &SearchUserGroupsParams{
		HTTPClient: client,
	}
}

/*
SearchUserGroupsParams contains all the parameters to send to the API endpoint

	for the search user groups operation.

	Typically these are written to a http.Request.
*/
type SearchUserGroupsParams struct {

	/* XRequestID.

	   An unique ID for the request
	*/
	XRequestID *string

	/* Groupname.

	   Group name for filtering results.
	*/
	Groupname string

	/* Page.

	   The page number

	   Format: int64
	   Default: 1
	*/
	Page *int64

	/* PageSize.

	   The size of per page

	   Format: int64
	   Default: 10
	*/
	PageSize *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the search user groups params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SearchUserGroupsParams) WithDefaults() *SearchUserGroupsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the search user groups params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SearchUserGroupsParams) SetDefaults() {
	var (
		pageDefault = int64(1)

		pageSizeDefault = int64(10)
	)

	val := SearchUserGroupsParams{
		Page:     &pageDefault,
		PageSize: &pageSizeDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the search user groups params
func (o *SearchUserGroupsParams) WithTimeout(timeout time.Duration) *SearchUserGroupsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the search user groups params
func (o *SearchUserGroupsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the search user groups params
func (o *SearchUserGroupsParams) WithContext(ctx context.Context) *SearchUserGroupsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the search user groups params
func (o *SearchUserGroupsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the search user groups params
func (o *SearchUserGroupsParams) WithHTTPClient(client *http.Client) *SearchUserGroupsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the search user groups params
func (o *SearchUserGroupsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the search user groups params
func (o *SearchUserGroupsParams) WithXRequestID(xRequestID *string) *SearchUserGroupsParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the search user groups params
func (o *SearchUserGroupsParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithGroupname adds the groupname to the search user groups params
func (o *SearchUserGroupsParams) WithGroupname(groupname string) *SearchUserGroupsParams {
	o.SetGroupname(groupname)
	return o
}

// SetGroupname adds the groupname to the search user groups params
func (o *SearchUserGroupsParams) SetGroupname(groupname string) {
	o.Groupname = groupname
}

// WithPage adds the page to the search user groups params
func (o *SearchUserGroupsParams) WithPage(page *int64) *SearchUserGroupsParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the search user groups params
func (o *SearchUserGroupsParams) SetPage(page *int64) {
	o.Page = page
}

// WithPageSize adds the pageSize to the search user groups params
func (o *SearchUserGroupsParams) WithPageSize(pageSize *int64) *SearchUserGroupsParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the search user groups params
func (o *SearchUserGroupsParams) SetPageSize(pageSize *int64) {
	o.PageSize = pageSize
}

// WriteToRequest writes these params to a swagger request
func (o *SearchUserGroupsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// query param groupname
	qrGroupname := o.Groupname
	qGroupname := qrGroupname
	if qGroupname != "" {

		if err := r.SetQueryParam("groupname", qGroupname); err != nil {
			return err
		}
	}

	if o.Page != nil {

		// query param page
		var qrPage int64

		if o.Page != nil {
			qrPage = *o.Page
		}
		qPage := swag.FormatInt64(qrPage)
		if qPage != "" {

			if err := r.SetQueryParam("page", qPage); err != nil {
				return err
			}
		}
	}

	if o.PageSize != nil {

		// query param page_size
		var qrPageSize int64

		if o.PageSize != nil {
			qrPageSize = *o.PageSize
		}
		qPageSize := swag.FormatInt64(qrPageSize)
		if qPageSize != "" {

			if err := r.SetQueryParam("page_size", qPageSize); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
