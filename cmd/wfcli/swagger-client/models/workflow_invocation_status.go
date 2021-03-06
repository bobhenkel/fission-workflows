package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// WorkflowInvocationStatus workflow invocation status
// swagger:model WorkflowInvocationStatus
type WorkflowInvocationStatus struct {

	// dynamic tasks
	DynamicTasks map[string]Task `json:"dynamicTasks,omitempty"`

	// error
	Error *Error `json:"error,omitempty"`

	// output
	Output *TypedValue `json:"output,omitempty"`

	// status
	Status WorkflowInvocationStatusStatus `json:"status,omitempty"`

	// tasks
	Tasks map[string]TaskInvocation `json:"tasks,omitempty"`

	// updated at
	UpdatedAt strfmt.DateTime `json:"updatedAt,omitempty"`
}

// Validate validates this workflow invocation status
func (m *WorkflowInvocationStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDynamicTasks(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateError(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateOutput(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateTasks(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WorkflowInvocationStatus) validateDynamicTasks(formats strfmt.Registry) error {

	if swag.IsZero(m.DynamicTasks) { // not required
		return nil
	}

	if err := validate.Required("dynamicTasks", "body", m.DynamicTasks); err != nil {
		return err
	}

	return nil
}

func (m *WorkflowInvocationStatus) validateError(formats strfmt.Registry) error {

	if swag.IsZero(m.Error) { // not required
		return nil
	}

	if m.Error != nil {

		if err := m.Error.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

func (m *WorkflowInvocationStatus) validateOutput(formats strfmt.Registry) error {

	if swag.IsZero(m.Output) { // not required
		return nil
	}

	if m.Output != nil {

		if err := m.Output.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

func (m *WorkflowInvocationStatus) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	if err := m.Status.Validate(formats); err != nil {
		return err
	}

	return nil
}

func (m *WorkflowInvocationStatus) validateTasks(formats strfmt.Registry) error {

	if swag.IsZero(m.Tasks) { // not required
		return nil
	}

	if err := validate.Required("tasks", "body", m.Tasks); err != nil {
		return err
	}

	return nil
}
