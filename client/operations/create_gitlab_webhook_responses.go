// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CreateGitlabWebhookReader is a Reader for the CreateGitlabWebhook structure.
type CreateGitlabWebhookReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateGitlabWebhookReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateGitlabWebhookCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewCreateGitlabWebhookForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateGitlabWebhookInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateGitlabWebhookCreated creates a CreateGitlabWebhookCreated with default headers values
func NewCreateGitlabWebhookCreated() *CreateGitlabWebhookCreated {
	return &CreateGitlabWebhookCreated{}
}

/* CreateGitlabWebhookCreated describes a response with status code 201, with default header values.

The webhook was created.
*/
type CreateGitlabWebhookCreated struct {
}

func (o *CreateGitlabWebhookCreated) Error() string {
	return fmt.Sprintf("[POST /api/gitlab/webhook][%d] createGitlabWebhookCreated ", 201)
}

func (o *CreateGitlabWebhookCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateGitlabWebhookForbidden creates a CreateGitlabWebhookForbidden with default headers values
func NewCreateGitlabWebhookForbidden() *CreateGitlabWebhookForbidden {
	return &CreateGitlabWebhookForbidden{}
}

/* CreateGitlabWebhookForbidden describes a response with status code 403, with default header values.

Request failed. User token not valid.
*/
type CreateGitlabWebhookForbidden struct {
}

func (o *CreateGitlabWebhookForbidden) Error() string {
	return fmt.Sprintf("[POST /api/gitlab/webhook][%d] createGitlabWebhookForbidden ", 403)
}

func (o *CreateGitlabWebhookForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateGitlabWebhookInternalServerError creates a CreateGitlabWebhookInternalServerError with default headers values
func NewCreateGitlabWebhookInternalServerError() *CreateGitlabWebhookInternalServerError {
	return &CreateGitlabWebhookInternalServerError{}
}

/* CreateGitlabWebhookInternalServerError describes a response with status code 500, with default header values.

Request failed. Internal controller error.
*/
type CreateGitlabWebhookInternalServerError struct {
}

func (o *CreateGitlabWebhookInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/gitlab/webhook][%d] createGitlabWebhookInternalServerError ", 500)
}

func (o *CreateGitlabWebhookInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*CreateGitlabWebhookBody create gitlab webhook body
swagger:model CreateGitlabWebhookBody
*/
type CreateGitlabWebhookBody struct {

	// The GitLab project id.
	// Required: true
	ProjectID *string `json:"project_id"`
}

// Validate validates this create gitlab webhook body
func (o *CreateGitlabWebhookBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateProjectID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateGitlabWebhookBody) validateProjectID(formats strfmt.Registry) error {

	if err := validate.Required("data"+"."+"project_id", "body", o.ProjectID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this create gitlab webhook body based on context it is used
func (o *CreateGitlabWebhookBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *CreateGitlabWebhookBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateGitlabWebhookBody) UnmarshalBinary(b []byte) error {
	var res CreateGitlabWebhookBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
