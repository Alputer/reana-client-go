// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// InfoReader is a Reader for the Info structure.
type InfoReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *InfoReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewInfoOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewInfoInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewInfoOK creates a InfoOK with default headers values
func NewInfoOK() *InfoOK {
	return &InfoOK{}
}

/* InfoOK describes a response with status code 200, with default header values.

Request succeeded. The response contains general info about the cluster.
*/
type InfoOK struct {
	Payload *InfoOKBody
}

func (o *InfoOK) Error() string {
	return fmt.Sprintf("[GET /api/info][%d] infoOK  %+v", 200, o.Payload)
}
func (o *InfoOK) GetPayload() *InfoOKBody {
	return o.Payload
}

func (o *InfoOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(InfoOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInfoInternalServerError creates a InfoInternalServerError with default headers values
func NewInfoInternalServerError() *InfoInternalServerError {
	return &InfoInternalServerError{}
}

/* InfoInternalServerError describes a response with status code 500, with default header values.

Request failed. Internal controller error.
*/
type InfoInternalServerError struct {
	Payload *InfoInternalServerErrorBody
}

func (o *InfoInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/info][%d] infoInternalServerError  %+v", 500, o.Payload)
}
func (o *InfoInternalServerError) GetPayload() *InfoInternalServerErrorBody {
	return o.Payload
}

func (o *InfoInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(InfoInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*InfoInternalServerErrorBody info internal server error body
swagger:model InfoInternalServerErrorBody
*/
type InfoInternalServerErrorBody struct {

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this info internal server error body
func (o *InfoInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this info internal server error body based on context it is used
func (o *InfoInternalServerErrorBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *InfoInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *InfoInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res InfoInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*InfoOKBody info o k body
swagger:model InfoOKBody
*/
type InfoOKBody struct {

	// compute backends
	ComputeBackends *InfoOKBodyComputeBackends `json:"compute_backends,omitempty"`

	// default kubernetes jobs timeout
	DefaultKubernetesJobsTimeout *InfoOKBodyDefaultKubernetesJobsTimeout `json:"default_kubernetes_jobs_timeout,omitempty"`

	// default kubernetes memory limit
	DefaultKubernetesMemoryLimit *InfoOKBodyDefaultKubernetesMemoryLimit `json:"default_kubernetes_memory_limit,omitempty"`

	// default workspace
	DefaultWorkspace *InfoOKBodyDefaultWorkspace `json:"default_workspace,omitempty"`

	// kubernetes max memory limit
	KubernetesMaxMemoryLimit *InfoOKBodyKubernetesMaxMemoryLimit `json:"kubernetes_max_memory_limit,omitempty"`

	// maximum kubernetes jobs timeout
	MaximumKubernetesJobsTimeout *InfoOKBodyMaximumKubernetesJobsTimeout `json:"maximum_kubernetes_jobs_timeout,omitempty"`

	// maximum workspace retention period
	MaximumWorkspaceRetentionPeriod *InfoOKBodyMaximumWorkspaceRetentionPeriod `json:"maximum_workspace_retention_period,omitempty"`

	// workspaces available
	WorkspacesAvailable *InfoOKBodyWorkspacesAvailable `json:"workspaces_available,omitempty"`
}

// Validate validates this info o k body
func (o *InfoOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateComputeBackends(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateDefaultKubernetesJobsTimeout(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateDefaultKubernetesMemoryLimit(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateDefaultWorkspace(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateKubernetesMaxMemoryLimit(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateMaximumKubernetesJobsTimeout(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateMaximumWorkspaceRetentionPeriod(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateWorkspacesAvailable(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *InfoOKBody) validateComputeBackends(formats strfmt.Registry) error {
	if swag.IsZero(o.ComputeBackends) { // not required
		return nil
	}

	if o.ComputeBackends != nil {
		if err := o.ComputeBackends.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("infoOK" + "." + "compute_backends")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("infoOK" + "." + "compute_backends")
			}
			return err
		}
	}

	return nil
}

func (o *InfoOKBody) validateDefaultKubernetesJobsTimeout(formats strfmt.Registry) error {
	if swag.IsZero(o.DefaultKubernetesJobsTimeout) { // not required
		return nil
	}

	if o.DefaultKubernetesJobsTimeout != nil {
		if err := o.DefaultKubernetesJobsTimeout.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("infoOK" + "." + "default_kubernetes_jobs_timeout")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("infoOK" + "." + "default_kubernetes_jobs_timeout")
			}
			return err
		}
	}

	return nil
}

func (o *InfoOKBody) validateDefaultKubernetesMemoryLimit(formats strfmt.Registry) error {
	if swag.IsZero(o.DefaultKubernetesMemoryLimit) { // not required
		return nil
	}

	if o.DefaultKubernetesMemoryLimit != nil {
		if err := o.DefaultKubernetesMemoryLimit.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("infoOK" + "." + "default_kubernetes_memory_limit")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("infoOK" + "." + "default_kubernetes_memory_limit")
			}
			return err
		}
	}

	return nil
}

func (o *InfoOKBody) validateDefaultWorkspace(formats strfmt.Registry) error {
	if swag.IsZero(o.DefaultWorkspace) { // not required
		return nil
	}

	if o.DefaultWorkspace != nil {
		if err := o.DefaultWorkspace.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("infoOK" + "." + "default_workspace")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("infoOK" + "." + "default_workspace")
			}
			return err
		}
	}

	return nil
}

func (o *InfoOKBody) validateKubernetesMaxMemoryLimit(formats strfmt.Registry) error {
	if swag.IsZero(o.KubernetesMaxMemoryLimit) { // not required
		return nil
	}

	if o.KubernetesMaxMemoryLimit != nil {
		if err := o.KubernetesMaxMemoryLimit.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("infoOK" + "." + "kubernetes_max_memory_limit")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("infoOK" + "." + "kubernetes_max_memory_limit")
			}
			return err
		}
	}

	return nil
}

func (o *InfoOKBody) validateMaximumKubernetesJobsTimeout(formats strfmt.Registry) error {
	if swag.IsZero(o.MaximumKubernetesJobsTimeout) { // not required
		return nil
	}

	if o.MaximumKubernetesJobsTimeout != nil {
		if err := o.MaximumKubernetesJobsTimeout.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("infoOK" + "." + "maximum_kubernetes_jobs_timeout")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("infoOK" + "." + "maximum_kubernetes_jobs_timeout")
			}
			return err
		}
	}

	return nil
}

func (o *InfoOKBody) validateMaximumWorkspaceRetentionPeriod(formats strfmt.Registry) error {
	if swag.IsZero(o.MaximumWorkspaceRetentionPeriod) { // not required
		return nil
	}

	if o.MaximumWorkspaceRetentionPeriod != nil {
		if err := o.MaximumWorkspaceRetentionPeriod.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("infoOK" + "." + "maximum_workspace_retention_period")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("infoOK" + "." + "maximum_workspace_retention_period")
			}
			return err
		}
	}

	return nil
}

func (o *InfoOKBody) validateWorkspacesAvailable(formats strfmt.Registry) error {
	if swag.IsZero(o.WorkspacesAvailable) { // not required
		return nil
	}

	if o.WorkspacesAvailable != nil {
		if err := o.WorkspacesAvailable.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("infoOK" + "." + "workspaces_available")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("infoOK" + "." + "workspaces_available")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this info o k body based on the context it is used
func (o *InfoOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateComputeBackends(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateDefaultKubernetesJobsTimeout(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateDefaultKubernetesMemoryLimit(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateDefaultWorkspace(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateKubernetesMaxMemoryLimit(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateMaximumKubernetesJobsTimeout(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateMaximumWorkspaceRetentionPeriod(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateWorkspacesAvailable(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *InfoOKBody) contextValidateComputeBackends(ctx context.Context, formats strfmt.Registry) error {

	if o.ComputeBackends != nil {
		if err := o.ComputeBackends.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("infoOK" + "." + "compute_backends")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("infoOK" + "." + "compute_backends")
			}
			return err
		}
	}

	return nil
}

func (o *InfoOKBody) contextValidateDefaultKubernetesJobsTimeout(ctx context.Context, formats strfmt.Registry) error {

	if o.DefaultKubernetesJobsTimeout != nil {
		if err := o.DefaultKubernetesJobsTimeout.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("infoOK" + "." + "default_kubernetes_jobs_timeout")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("infoOK" + "." + "default_kubernetes_jobs_timeout")
			}
			return err
		}
	}

	return nil
}

func (o *InfoOKBody) contextValidateDefaultKubernetesMemoryLimit(ctx context.Context, formats strfmt.Registry) error {

	if o.DefaultKubernetesMemoryLimit != nil {
		if err := o.DefaultKubernetesMemoryLimit.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("infoOK" + "." + "default_kubernetes_memory_limit")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("infoOK" + "." + "default_kubernetes_memory_limit")
			}
			return err
		}
	}

	return nil
}

func (o *InfoOKBody) contextValidateDefaultWorkspace(ctx context.Context, formats strfmt.Registry) error {

	if o.DefaultWorkspace != nil {
		if err := o.DefaultWorkspace.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("infoOK" + "." + "default_workspace")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("infoOK" + "." + "default_workspace")
			}
			return err
		}
	}

	return nil
}

func (o *InfoOKBody) contextValidateKubernetesMaxMemoryLimit(ctx context.Context, formats strfmt.Registry) error {

	if o.KubernetesMaxMemoryLimit != nil {
		if err := o.KubernetesMaxMemoryLimit.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("infoOK" + "." + "kubernetes_max_memory_limit")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("infoOK" + "." + "kubernetes_max_memory_limit")
			}
			return err
		}
	}

	return nil
}

func (o *InfoOKBody) contextValidateMaximumKubernetesJobsTimeout(ctx context.Context, formats strfmt.Registry) error {

	if o.MaximumKubernetesJobsTimeout != nil {
		if err := o.MaximumKubernetesJobsTimeout.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("infoOK" + "." + "maximum_kubernetes_jobs_timeout")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("infoOK" + "." + "maximum_kubernetes_jobs_timeout")
			}
			return err
		}
	}

	return nil
}

func (o *InfoOKBody) contextValidateMaximumWorkspaceRetentionPeriod(ctx context.Context, formats strfmt.Registry) error {

	if o.MaximumWorkspaceRetentionPeriod != nil {
		if err := o.MaximumWorkspaceRetentionPeriod.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("infoOK" + "." + "maximum_workspace_retention_period")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("infoOK" + "." + "maximum_workspace_retention_period")
			}
			return err
		}
	}

	return nil
}

func (o *InfoOKBody) contextValidateWorkspacesAvailable(ctx context.Context, formats strfmt.Registry) error {

	if o.WorkspacesAvailable != nil {
		if err := o.WorkspacesAvailable.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("infoOK" + "." + "workspaces_available")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("infoOK" + "." + "workspaces_available")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *InfoOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *InfoOKBody) UnmarshalBinary(b []byte) error {
	var res InfoOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*InfoOKBodyComputeBackends info o k body compute backends
swagger:model InfoOKBodyComputeBackends
*/
type InfoOKBodyComputeBackends struct {

	// title
	Title string `json:"title,omitempty"`

	// value
	Value []string `json:"value"`
}

// Validate validates this info o k body compute backends
func (o *InfoOKBodyComputeBackends) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this info o k body compute backends based on context it is used
func (o *InfoOKBodyComputeBackends) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *InfoOKBodyComputeBackends) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *InfoOKBodyComputeBackends) UnmarshalBinary(b []byte) error {
	var res InfoOKBodyComputeBackends
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*InfoOKBodyDefaultKubernetesJobsTimeout info o k body default kubernetes jobs timeout
swagger:model InfoOKBodyDefaultKubernetesJobsTimeout
*/
type InfoOKBodyDefaultKubernetesJobsTimeout struct {

	// title
	Title string `json:"title,omitempty"`

	// value
	Value string `json:"value,omitempty"`
}

// Validate validates this info o k body default kubernetes jobs timeout
func (o *InfoOKBodyDefaultKubernetesJobsTimeout) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this info o k body default kubernetes jobs timeout based on context it is used
func (o *InfoOKBodyDefaultKubernetesJobsTimeout) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *InfoOKBodyDefaultKubernetesJobsTimeout) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *InfoOKBodyDefaultKubernetesJobsTimeout) UnmarshalBinary(b []byte) error {
	var res InfoOKBodyDefaultKubernetesJobsTimeout
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*InfoOKBodyDefaultKubernetesMemoryLimit info o k body default kubernetes memory limit
swagger:model InfoOKBodyDefaultKubernetesMemoryLimit
*/
type InfoOKBodyDefaultKubernetesMemoryLimit struct {

	// title
	Title string `json:"title,omitempty"`

	// value
	Value string `json:"value,omitempty"`
}

// Validate validates this info o k body default kubernetes memory limit
func (o *InfoOKBodyDefaultKubernetesMemoryLimit) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this info o k body default kubernetes memory limit based on context it is used
func (o *InfoOKBodyDefaultKubernetesMemoryLimit) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *InfoOKBodyDefaultKubernetesMemoryLimit) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *InfoOKBodyDefaultKubernetesMemoryLimit) UnmarshalBinary(b []byte) error {
	var res InfoOKBodyDefaultKubernetesMemoryLimit
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*InfoOKBodyDefaultWorkspace info o k body default workspace
swagger:model InfoOKBodyDefaultWorkspace
*/
type InfoOKBodyDefaultWorkspace struct {

	// title
	Title string `json:"title,omitempty"`

	// value
	Value string `json:"value,omitempty"`
}

// Validate validates this info o k body default workspace
func (o *InfoOKBodyDefaultWorkspace) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this info o k body default workspace based on context it is used
func (o *InfoOKBodyDefaultWorkspace) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *InfoOKBodyDefaultWorkspace) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *InfoOKBodyDefaultWorkspace) UnmarshalBinary(b []byte) error {
	var res InfoOKBodyDefaultWorkspace
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*InfoOKBodyKubernetesMaxMemoryLimit info o k body kubernetes max memory limit
swagger:model InfoOKBodyKubernetesMaxMemoryLimit
*/
type InfoOKBodyKubernetesMaxMemoryLimit struct {

	// title
	Title string `json:"title,omitempty"`

	// value
	Value *string `json:"value,omitempty"`
}

// Validate validates this info o k body kubernetes max memory limit
func (o *InfoOKBodyKubernetesMaxMemoryLimit) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this info o k body kubernetes max memory limit based on context it is used
func (o *InfoOKBodyKubernetesMaxMemoryLimit) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *InfoOKBodyKubernetesMaxMemoryLimit) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *InfoOKBodyKubernetesMaxMemoryLimit) UnmarshalBinary(b []byte) error {
	var res InfoOKBodyKubernetesMaxMemoryLimit
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*InfoOKBodyMaximumKubernetesJobsTimeout info o k body maximum kubernetes jobs timeout
swagger:model InfoOKBodyMaximumKubernetesJobsTimeout
*/
type InfoOKBodyMaximumKubernetesJobsTimeout struct {

	// title
	Title string `json:"title,omitempty"`

	// value
	Value string `json:"value,omitempty"`
}

// Validate validates this info o k body maximum kubernetes jobs timeout
func (o *InfoOKBodyMaximumKubernetesJobsTimeout) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this info o k body maximum kubernetes jobs timeout based on context it is used
func (o *InfoOKBodyMaximumKubernetesJobsTimeout) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *InfoOKBodyMaximumKubernetesJobsTimeout) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *InfoOKBodyMaximumKubernetesJobsTimeout) UnmarshalBinary(b []byte) error {
	var res InfoOKBodyMaximumKubernetesJobsTimeout
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*InfoOKBodyMaximumWorkspaceRetentionPeriod info o k body maximum workspace retention period
swagger:model InfoOKBodyMaximumWorkspaceRetentionPeriod
*/
type InfoOKBodyMaximumWorkspaceRetentionPeriod struct {

	// title
	Title string `json:"title,omitempty"`

	// value
	Value *string `json:"value,omitempty"`
}

// Validate validates this info o k body maximum workspace retention period
func (o *InfoOKBodyMaximumWorkspaceRetentionPeriod) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this info o k body maximum workspace retention period based on context it is used
func (o *InfoOKBodyMaximumWorkspaceRetentionPeriod) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *InfoOKBodyMaximumWorkspaceRetentionPeriod) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *InfoOKBodyMaximumWorkspaceRetentionPeriod) UnmarshalBinary(b []byte) error {
	var res InfoOKBodyMaximumWorkspaceRetentionPeriod
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*InfoOKBodyWorkspacesAvailable info o k body workspaces available
swagger:model InfoOKBodyWorkspacesAvailable
*/
type InfoOKBodyWorkspacesAvailable struct {

	// title
	Title string `json:"title,omitempty"`

	// value
	Value []string `json:"value"`
}

// Validate validates this info o k body workspaces available
func (o *InfoOKBodyWorkspacesAvailable) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this info o k body workspaces available based on context it is used
func (o *InfoOKBodyWorkspacesAvailable) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *InfoOKBodyWorkspacesAvailable) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *InfoOKBodyWorkspacesAvailable) UnmarshalBinary(b []byte) error {
	var res InfoOKBodyWorkspacesAvailable
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
