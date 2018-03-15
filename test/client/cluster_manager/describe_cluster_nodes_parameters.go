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
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDescribeClusterNodesParams creates a new DescribeClusterNodesParams object
// with the default values initialized.
func NewDescribeClusterNodesParams() *DescribeClusterNodesParams {
	var ()
	return &DescribeClusterNodesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDescribeClusterNodesParamsWithTimeout creates a new DescribeClusterNodesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDescribeClusterNodesParamsWithTimeout(timeout time.Duration) *DescribeClusterNodesParams {
	var ()
	return &DescribeClusterNodesParams{

		timeout: timeout,
	}
}

// NewDescribeClusterNodesParamsWithContext creates a new DescribeClusterNodesParams object
// with the default values initialized, and the ability to set a context for a request
func NewDescribeClusterNodesParamsWithContext(ctx context.Context) *DescribeClusterNodesParams {
	var ()
	return &DescribeClusterNodesParams{

		Context: ctx,
	}
}

// NewDescribeClusterNodesParamsWithHTTPClient creates a new DescribeClusterNodesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDescribeClusterNodesParamsWithHTTPClient(client *http.Client) *DescribeClusterNodesParams {
	var ()
	return &DescribeClusterNodesParams{
		HTTPClient: client,
	}
}

/*DescribeClusterNodesParams contains all the parameters to send to the API endpoint
for the describe cluster nodes operation typically these are written to a http.Request
*/
type DescribeClusterNodesParams struct {

	/*ClusterID*/
	ClusterID *string
	/*Limit
	  default is 20, max value is 200.

	*/
	Limit *int64
	/*NodeID*/
	NodeID []string
	/*Offset
	  default is 0.

	*/
	Offset *int64
	/*Status*/
	Status []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the describe cluster nodes params
func (o *DescribeClusterNodesParams) WithTimeout(timeout time.Duration) *DescribeClusterNodesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the describe cluster nodes params
func (o *DescribeClusterNodesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the describe cluster nodes params
func (o *DescribeClusterNodesParams) WithContext(ctx context.Context) *DescribeClusterNodesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the describe cluster nodes params
func (o *DescribeClusterNodesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the describe cluster nodes params
func (o *DescribeClusterNodesParams) WithHTTPClient(client *http.Client) *DescribeClusterNodesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the describe cluster nodes params
func (o *DescribeClusterNodesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusterID adds the clusterID to the describe cluster nodes params
func (o *DescribeClusterNodesParams) WithClusterID(clusterID *string) *DescribeClusterNodesParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the describe cluster nodes params
func (o *DescribeClusterNodesParams) SetClusterID(clusterID *string) {
	o.ClusterID = clusterID
}

// WithLimit adds the limit to the describe cluster nodes params
func (o *DescribeClusterNodesParams) WithLimit(limit *int64) *DescribeClusterNodesParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the describe cluster nodes params
func (o *DescribeClusterNodesParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithNodeID adds the nodeID to the describe cluster nodes params
func (o *DescribeClusterNodesParams) WithNodeID(nodeID []string) *DescribeClusterNodesParams {
	o.SetNodeID(nodeID)
	return o
}

// SetNodeID adds the nodeId to the describe cluster nodes params
func (o *DescribeClusterNodesParams) SetNodeID(nodeID []string) {
	o.NodeID = nodeID
}

// WithOffset adds the offset to the describe cluster nodes params
func (o *DescribeClusterNodesParams) WithOffset(offset *int64) *DescribeClusterNodesParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the describe cluster nodes params
func (o *DescribeClusterNodesParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithStatus adds the status to the describe cluster nodes params
func (o *DescribeClusterNodesParams) WithStatus(status []string) *DescribeClusterNodesParams {
	o.SetStatus(status)
	return o
}

// SetStatus adds the status to the describe cluster nodes params
func (o *DescribeClusterNodesParams) SetStatus(status []string) {
	o.Status = status
}

// WriteToRequest writes these params to a swagger request
func (o *DescribeClusterNodesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ClusterID != nil {

		// query param cluster_id
		var qrClusterID string
		if o.ClusterID != nil {
			qrClusterID = *o.ClusterID
		}
		qClusterID := qrClusterID
		if qClusterID != "" {
			if err := r.SetQueryParam("cluster_id", qClusterID); err != nil {
				return err
			}
		}

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

	valuesNodeID := o.NodeID

	joinedNodeID := swag.JoinByFormat(valuesNodeID, "")
	// query array param node_id
	if err := r.SetQueryParam("node_id", joinedNodeID...); err != nil {
		return err
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

	valuesStatus := o.Status

	joinedStatus := swag.JoinByFormat(valuesStatus, "")
	// query array param status
	if err := r.SetQueryParam("status", joinedStatus...); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
