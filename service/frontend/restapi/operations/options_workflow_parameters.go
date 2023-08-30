// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
)

// NewOptionsWorkflowParams creates a new OptionsWorkflowParams object
//
// There are no default values defined in the spec.
func NewOptionsWorkflowParams() OptionsWorkflowParams {

	return OptionsWorkflowParams{}
}

// OptionsWorkflowParams contains all the bound params for the options workflow operation
// typically these are obtained from a http.Request
//
// swagger:parameters optionsWorkflow
type OptionsWorkflowParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  Required: true
	  In: path
	*/
	WorkflowID string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewOptionsWorkflowParams() beforehand.
func (o *OptionsWorkflowParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rWorkflowID, rhkWorkflowID, _ := route.Params.GetOK("workflowId")
	if err := o.bindWorkflowID(rWorkflowID, rhkWorkflowID, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindWorkflowID binds and validates parameter WorkflowID from path.
func (o *OptionsWorkflowParams) bindWorkflowID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route
	o.WorkflowID = raw

	return nil
}
