// Code generated by go-swagger; DO NOT EDIT.

package pipeline_service

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

	pipeline_model "github.com/kubeflow/pipelines/backend/api/v2beta1/go_http_client/pipeline_model"
)

// NewPipelineServiceCreatePipelineAndVersionParams creates a new PipelineServiceCreatePipelineAndVersionParams object
// with the default values initialized.
func NewPipelineServiceCreatePipelineAndVersionParams() *PipelineServiceCreatePipelineAndVersionParams {
	var ()
	return &PipelineServiceCreatePipelineAndVersionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPipelineServiceCreatePipelineAndVersionParamsWithTimeout creates a new PipelineServiceCreatePipelineAndVersionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPipelineServiceCreatePipelineAndVersionParamsWithTimeout(timeout time.Duration) *PipelineServiceCreatePipelineAndVersionParams {
	var ()
	return &PipelineServiceCreatePipelineAndVersionParams{

		timeout: timeout,
	}
}

// NewPipelineServiceCreatePipelineAndVersionParamsWithContext creates a new PipelineServiceCreatePipelineAndVersionParams object
// with the default values initialized, and the ability to set a context for a request
func NewPipelineServiceCreatePipelineAndVersionParamsWithContext(ctx context.Context) *PipelineServiceCreatePipelineAndVersionParams {
	var ()
	return &PipelineServiceCreatePipelineAndVersionParams{

		Context: ctx,
	}
}

// NewPipelineServiceCreatePipelineAndVersionParamsWithHTTPClient creates a new PipelineServiceCreatePipelineAndVersionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPipelineServiceCreatePipelineAndVersionParamsWithHTTPClient(client *http.Client) *PipelineServiceCreatePipelineAndVersionParams {
	var ()
	return &PipelineServiceCreatePipelineAndVersionParams{
		HTTPClient: client,
	}
}

/*PipelineServiceCreatePipelineAndVersionParams contains all the parameters to send to the API endpoint
for the pipeline service create pipeline and version operation typically these are written to a http.Request
*/
type PipelineServiceCreatePipelineAndVersionParams struct {

	/*Body*/
	Body *pipeline_model.V2beta1CreatePipelineAndVersionRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the pipeline service create pipeline and version params
func (o *PipelineServiceCreatePipelineAndVersionParams) WithTimeout(timeout time.Duration) *PipelineServiceCreatePipelineAndVersionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the pipeline service create pipeline and version params
func (o *PipelineServiceCreatePipelineAndVersionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the pipeline service create pipeline and version params
func (o *PipelineServiceCreatePipelineAndVersionParams) WithContext(ctx context.Context) *PipelineServiceCreatePipelineAndVersionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the pipeline service create pipeline and version params
func (o *PipelineServiceCreatePipelineAndVersionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the pipeline service create pipeline and version params
func (o *PipelineServiceCreatePipelineAndVersionParams) WithHTTPClient(client *http.Client) *PipelineServiceCreatePipelineAndVersionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the pipeline service create pipeline and version params
func (o *PipelineServiceCreatePipelineAndVersionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the pipeline service create pipeline and version params
func (o *PipelineServiceCreatePipelineAndVersionParams) WithBody(body *pipeline_model.V2beta1CreatePipelineAndVersionRequest) *PipelineServiceCreatePipelineAndVersionParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the pipeline service create pipeline and version params
func (o *PipelineServiceCreatePipelineAndVersionParams) SetBody(body *pipeline_model.V2beta1CreatePipelineAndVersionRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PipelineServiceCreatePipelineAndVersionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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