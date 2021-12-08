// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2018 Amazon.com, Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
// 	http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDescribeInstanceParams creates a new DescribeInstanceParams object
// with the default values initialized.
func NewDescribeInstanceParams() *DescribeInstanceParams {

	return &DescribeInstanceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDescribeInstanceParamsWithTimeout creates a new DescribeInstanceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDescribeInstanceParamsWithTimeout(timeout time.Duration) *DescribeInstanceParams {

	return &DescribeInstanceParams{

		timeout: timeout,
	}
}

// NewDescribeInstanceParamsWithContext creates a new DescribeInstanceParams object
// with the default values initialized, and the ability to set a context for a request
func NewDescribeInstanceParamsWithContext(ctx context.Context) *DescribeInstanceParams {

	return &DescribeInstanceParams{

		Context: ctx,
	}
}

// NewDescribeInstanceParamsWithHTTPClient creates a new DescribeInstanceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDescribeInstanceParamsWithHTTPClient(client *http.Client) *DescribeInstanceParams {

	return &DescribeInstanceParams{
		HTTPClient: client,
	}
}

/*DescribeInstanceParams contains all the parameters to send to the API endpoint
for the describe instance operation typically these are written to a http.Request
*/
type DescribeInstanceParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the describe instance params
func (o *DescribeInstanceParams) WithTimeout(timeout time.Duration) *DescribeInstanceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the describe instance params
func (o *DescribeInstanceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the describe instance params
func (o *DescribeInstanceParams) WithContext(ctx context.Context) *DescribeInstanceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the describe instance params
func (o *DescribeInstanceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the describe instance params
func (o *DescribeInstanceParams) WithHTTPClient(client *http.Client) *DescribeInstanceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the describe instance params
func (o *DescribeInstanceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *DescribeInstanceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
