package workflow_invocation_api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewInvokeSync0Params creates a new InvokeSync0Params object
// with the default values initialized.
func NewInvokeSync0Params() *InvokeSync0Params {
	var ()
	return &InvokeSync0Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewInvokeSync0ParamsWithTimeout creates a new InvokeSync0Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewInvokeSync0ParamsWithTimeout(timeout time.Duration) *InvokeSync0Params {
	var ()
	return &InvokeSync0Params{

		timeout: timeout,
	}
}

// NewInvokeSync0ParamsWithContext creates a new InvokeSync0Params object
// with the default values initialized, and the ability to set a context for a request
func NewInvokeSync0ParamsWithContext(ctx context.Context) *InvokeSync0Params {
	var ()
	return &InvokeSync0Params{

		Context: ctx,
	}
}

/*InvokeSync0Params contains all the parameters to send to the API endpoint
for the invoke sync 0 operation typically these are written to a http.Request
*/
type InvokeSync0Params struct {

	/*WorkflowID*/
	WorkflowID *string

	timeout time.Duration
	Context context.Context
}

// WithTimeout adds the timeout to the invoke sync 0 params
func (o *InvokeSync0Params) WithTimeout(timeout time.Duration) *InvokeSync0Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the invoke sync 0 params
func (o *InvokeSync0Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the invoke sync 0 params
func (o *InvokeSync0Params) WithContext(ctx context.Context) *InvokeSync0Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the invoke sync 0 params
func (o *InvokeSync0Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithWorkflowID adds the workflowID to the invoke sync 0 params
func (o *InvokeSync0Params) WithWorkflowID(workflowID *string) *InvokeSync0Params {
	o.SetWorkflowID(workflowID)
	return o
}

// SetWorkflowID adds the workflowId to the invoke sync 0 params
func (o *InvokeSync0Params) SetWorkflowID(workflowID *string) {
	o.WorkflowID = workflowID
}

// WriteToRequest writes these params to a swagger request
func (o *InvokeSync0Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if o.WorkflowID != nil {

		// query param workflowId
		var qrWorkflowID string
		if o.WorkflowID != nil {
			qrWorkflowID = *o.WorkflowID
		}
		qWorkflowID := qrWorkflowID
		if qWorkflowID != "" {
			if err := r.SetQueryParam("workflowId", qWorkflowID); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
