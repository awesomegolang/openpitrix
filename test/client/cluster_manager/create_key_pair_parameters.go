// Code generated by go-swagger; DO NOT EDIT.

package cluster_manager

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"openpitrix.io/openpitrix/test/models"
)

// NewCreateKeyPairParams creates a new CreateKeyPairParams object
// with the default values initialized.
func NewCreateKeyPairParams() *CreateKeyPairParams {
	var ()
	return &CreateKeyPairParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateKeyPairParamsWithTimeout creates a new CreateKeyPairParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateKeyPairParamsWithTimeout(timeout time.Duration) *CreateKeyPairParams {
	var ()
	return &CreateKeyPairParams{

		timeout: timeout,
	}
}

// NewCreateKeyPairParamsWithContext creates a new CreateKeyPairParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateKeyPairParamsWithContext(ctx context.Context) *CreateKeyPairParams {
	var ()
	return &CreateKeyPairParams{

		Context: ctx,
	}
}

// NewCreateKeyPairParamsWithHTTPClient creates a new CreateKeyPairParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateKeyPairParamsWithHTTPClient(client *http.Client) *CreateKeyPairParams {
	var ()
	return &CreateKeyPairParams{
		HTTPClient: client,
	}
}

/*CreateKeyPairParams contains all the parameters to send to the API endpoint
for the create key pair operation typically these are written to a http.Request
*/
type CreateKeyPairParams struct {

	/*Body*/
	Body *models.OpenpitrixCreateKeyPairRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create key pair params
func (o *CreateKeyPairParams) WithTimeout(timeout time.Duration) *CreateKeyPairParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create key pair params
func (o *CreateKeyPairParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create key pair params
func (o *CreateKeyPairParams) WithContext(ctx context.Context) *CreateKeyPairParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create key pair params
func (o *CreateKeyPairParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create key pair params
func (o *CreateKeyPairParams) WithHTTPClient(client *http.Client) *CreateKeyPairParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create key pair params
func (o *CreateKeyPairParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create key pair params
func (o *CreateKeyPairParams) WithBody(body *models.OpenpitrixCreateKeyPairRequest) *CreateKeyPairParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create key pair params
func (o *CreateKeyPairParams) SetBody(body *models.OpenpitrixCreateKeyPairRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateKeyPairParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
