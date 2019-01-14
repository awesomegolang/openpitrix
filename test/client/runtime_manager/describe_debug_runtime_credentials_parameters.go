// Code generated by go-swagger; DO NOT EDIT.

package runtime_manager

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDescribeDebugRuntimeCredentialsParams creates a new DescribeDebugRuntimeCredentialsParams object
// with the default values initialized.
func NewDescribeDebugRuntimeCredentialsParams() *DescribeDebugRuntimeCredentialsParams {
	var ()
	return &DescribeDebugRuntimeCredentialsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDescribeDebugRuntimeCredentialsParamsWithTimeout creates a new DescribeDebugRuntimeCredentialsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDescribeDebugRuntimeCredentialsParamsWithTimeout(timeout time.Duration) *DescribeDebugRuntimeCredentialsParams {
	var ()
	return &DescribeDebugRuntimeCredentialsParams{

		timeout: timeout,
	}
}

// NewDescribeDebugRuntimeCredentialsParamsWithContext creates a new DescribeDebugRuntimeCredentialsParams object
// with the default values initialized, and the ability to set a context for a request
func NewDescribeDebugRuntimeCredentialsParamsWithContext(ctx context.Context) *DescribeDebugRuntimeCredentialsParams {
	var ()
	return &DescribeDebugRuntimeCredentialsParams{

		Context: ctx,
	}
}

// NewDescribeDebugRuntimeCredentialsParamsWithHTTPClient creates a new DescribeDebugRuntimeCredentialsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDescribeDebugRuntimeCredentialsParamsWithHTTPClient(client *http.Client) *DescribeDebugRuntimeCredentialsParams {
	var ()
	return &DescribeDebugRuntimeCredentialsParams{
		HTTPClient: client,
	}
}

/*DescribeDebugRuntimeCredentialsParams contains all the parameters to send to the API endpoint
for the describe debug runtime credentials operation typically these are written to a http.Request
*/
type DescribeDebugRuntimeCredentialsParams struct {

	/*DisplayColumns*/
	DisplayColumns []string
	/*Limit
	  default is 20, max value is 200.

	*/
	Limit *int64
	/*Offset
	  default is 0.

	*/
	Offset *int64
	/*OwnerPath*/
	OwnerPath []string
	/*Provider*/
	Provider []string
	/*Reverse*/
	Reverse *bool
	/*RuntimeCredentialID*/
	RuntimeCredentialID []string
	/*SearchWord*/
	SearchWord *string
	/*SortKey*/
	SortKey *string
	/*Status*/
	Status []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the describe debug runtime credentials params
func (o *DescribeDebugRuntimeCredentialsParams) WithTimeout(timeout time.Duration) *DescribeDebugRuntimeCredentialsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the describe debug runtime credentials params
func (o *DescribeDebugRuntimeCredentialsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the describe debug runtime credentials params
func (o *DescribeDebugRuntimeCredentialsParams) WithContext(ctx context.Context) *DescribeDebugRuntimeCredentialsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the describe debug runtime credentials params
func (o *DescribeDebugRuntimeCredentialsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the describe debug runtime credentials params
func (o *DescribeDebugRuntimeCredentialsParams) WithHTTPClient(client *http.Client) *DescribeDebugRuntimeCredentialsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the describe debug runtime credentials params
func (o *DescribeDebugRuntimeCredentialsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDisplayColumns adds the displayColumns to the describe debug runtime credentials params
func (o *DescribeDebugRuntimeCredentialsParams) WithDisplayColumns(displayColumns []string) *DescribeDebugRuntimeCredentialsParams {
	o.SetDisplayColumns(displayColumns)
	return o
}

// SetDisplayColumns adds the displayColumns to the describe debug runtime credentials params
func (o *DescribeDebugRuntimeCredentialsParams) SetDisplayColumns(displayColumns []string) {
	o.DisplayColumns = displayColumns
}

// WithLimit adds the limit to the describe debug runtime credentials params
func (o *DescribeDebugRuntimeCredentialsParams) WithLimit(limit *int64) *DescribeDebugRuntimeCredentialsParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the describe debug runtime credentials params
func (o *DescribeDebugRuntimeCredentialsParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithOffset adds the offset to the describe debug runtime credentials params
func (o *DescribeDebugRuntimeCredentialsParams) WithOffset(offset *int64) *DescribeDebugRuntimeCredentialsParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the describe debug runtime credentials params
func (o *DescribeDebugRuntimeCredentialsParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithOwnerPath adds the ownerPath to the describe debug runtime credentials params
func (o *DescribeDebugRuntimeCredentialsParams) WithOwnerPath(ownerPath []string) *DescribeDebugRuntimeCredentialsParams {
	o.SetOwnerPath(ownerPath)
	return o
}

// SetOwnerPath adds the ownerPath to the describe debug runtime credentials params
func (o *DescribeDebugRuntimeCredentialsParams) SetOwnerPath(ownerPath []string) {
	o.OwnerPath = ownerPath
}

// WithProvider adds the provider to the describe debug runtime credentials params
func (o *DescribeDebugRuntimeCredentialsParams) WithProvider(provider []string) *DescribeDebugRuntimeCredentialsParams {
	o.SetProvider(provider)
	return o
}

// SetProvider adds the provider to the describe debug runtime credentials params
func (o *DescribeDebugRuntimeCredentialsParams) SetProvider(provider []string) {
	o.Provider = provider
}

// WithReverse adds the reverse to the describe debug runtime credentials params
func (o *DescribeDebugRuntimeCredentialsParams) WithReverse(reverse *bool) *DescribeDebugRuntimeCredentialsParams {
	o.SetReverse(reverse)
	return o
}

// SetReverse adds the reverse to the describe debug runtime credentials params
func (o *DescribeDebugRuntimeCredentialsParams) SetReverse(reverse *bool) {
	o.Reverse = reverse
}

// WithRuntimeCredentialID adds the runtimeCredentialID to the describe debug runtime credentials params
func (o *DescribeDebugRuntimeCredentialsParams) WithRuntimeCredentialID(runtimeCredentialID []string) *DescribeDebugRuntimeCredentialsParams {
	o.SetRuntimeCredentialID(runtimeCredentialID)
	return o
}

// SetRuntimeCredentialID adds the runtimeCredentialId to the describe debug runtime credentials params
func (o *DescribeDebugRuntimeCredentialsParams) SetRuntimeCredentialID(runtimeCredentialID []string) {
	o.RuntimeCredentialID = runtimeCredentialID
}

// WithSearchWord adds the searchWord to the describe debug runtime credentials params
func (o *DescribeDebugRuntimeCredentialsParams) WithSearchWord(searchWord *string) *DescribeDebugRuntimeCredentialsParams {
	o.SetSearchWord(searchWord)
	return o
}

// SetSearchWord adds the searchWord to the describe debug runtime credentials params
func (o *DescribeDebugRuntimeCredentialsParams) SetSearchWord(searchWord *string) {
	o.SearchWord = searchWord
}

// WithSortKey adds the sortKey to the describe debug runtime credentials params
func (o *DescribeDebugRuntimeCredentialsParams) WithSortKey(sortKey *string) *DescribeDebugRuntimeCredentialsParams {
	o.SetSortKey(sortKey)
	return o
}

// SetSortKey adds the sortKey to the describe debug runtime credentials params
func (o *DescribeDebugRuntimeCredentialsParams) SetSortKey(sortKey *string) {
	o.SortKey = sortKey
}

// WithStatus adds the status to the describe debug runtime credentials params
func (o *DescribeDebugRuntimeCredentialsParams) WithStatus(status []string) *DescribeDebugRuntimeCredentialsParams {
	o.SetStatus(status)
	return o
}

// SetStatus adds the status to the describe debug runtime credentials params
func (o *DescribeDebugRuntimeCredentialsParams) SetStatus(status []string) {
	o.Status = status
}

// WriteToRequest writes these params to a swagger request
func (o *DescribeDebugRuntimeCredentialsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	valuesDisplayColumns := o.DisplayColumns

	joinedDisplayColumns := swag.JoinByFormat(valuesDisplayColumns, "multi")
	// query array param display_columns
	if err := r.SetQueryParam("display_columns", joinedDisplayColumns...); err != nil {
		return err
	}

	if o.Limit != nil {

		// query param limit
		var qrLimit int64
		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {
			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}

	}

	if o.Offset != nil {

		// query param offset
		var qrOffset int64
		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatInt64(qrOffset)
		if qOffset != "" {
			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}

	}

	valuesOwnerPath := o.OwnerPath

	joinedOwnerPath := swag.JoinByFormat(valuesOwnerPath, "multi")
	// query array param owner_path
	if err := r.SetQueryParam("owner_path", joinedOwnerPath...); err != nil {
		return err
	}

	valuesProvider := o.Provider

	joinedProvider := swag.JoinByFormat(valuesProvider, "multi")
	// query array param provider
	if err := r.SetQueryParam("provider", joinedProvider...); err != nil {
		return err
	}

	if o.Reverse != nil {

		// query param reverse
		var qrReverse bool
		if o.Reverse != nil {
			qrReverse = *o.Reverse
		}
		qReverse := swag.FormatBool(qrReverse)
		if qReverse != "" {
			if err := r.SetQueryParam("reverse", qReverse); err != nil {
				return err
			}
		}

	}

	valuesRuntimeCredentialID := o.RuntimeCredentialID

	joinedRuntimeCredentialID := swag.JoinByFormat(valuesRuntimeCredentialID, "multi")
	// query array param runtime_credential_id
	if err := r.SetQueryParam("runtime_credential_id", joinedRuntimeCredentialID...); err != nil {
		return err
	}

	if o.SearchWord != nil {

		// query param search_word
		var qrSearchWord string
		if o.SearchWord != nil {
			qrSearchWord = *o.SearchWord
		}
		qSearchWord := qrSearchWord
		if qSearchWord != "" {
			if err := r.SetQueryParam("search_word", qSearchWord); err != nil {
				return err
			}
		}

	}

	if o.SortKey != nil {

		// query param sort_key
		var qrSortKey string
		if o.SortKey != nil {
			qrSortKey = *o.SortKey
		}
		qSortKey := qrSortKey
		if qSortKey != "" {
			if err := r.SetQueryParam("sort_key", qSortKey); err != nil {
				return err
			}
		}

	}

	valuesStatus := o.Status

	joinedStatus := swag.JoinByFormat(valuesStatus, "multi")
	// query array param status
	if err := r.SetQueryParam("status", joinedStatus...); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
