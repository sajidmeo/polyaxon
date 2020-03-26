// Copyright 2018-2020 Polyaxon, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by go-swagger; DO NOT EDIT.

package runs_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/polyaxon/polyaxon/sdks/go/http_client/v1/service_model"
)

// CreateRunDashboardReader is a Reader for the CreateRunDashboard structure.
type CreateRunDashboardReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateRunDashboardReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateRunDashboardOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewCreateRunDashboardNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewCreateRunDashboardForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateRunDashboardNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewCreateRunDashboardDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateRunDashboardOK creates a CreateRunDashboardOK with default headers values
func NewCreateRunDashboardOK() *CreateRunDashboardOK {
	return &CreateRunDashboardOK{}
}

/*CreateRunDashboardOK handles this case with default header values.

A successful response.
*/
type CreateRunDashboardOK struct {
	Payload *service_model.V1Dashboard
}

func (o *CreateRunDashboardOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/{project}/runs/{run}/dashboards][%d] createRunDashboardOK  %+v", 200, o.Payload)
}

func (o *CreateRunDashboardOK) GetPayload() *service_model.V1Dashboard {
	return o.Payload
}

func (o *CreateRunDashboardOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.V1Dashboard)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateRunDashboardNoContent creates a CreateRunDashboardNoContent with default headers values
func NewCreateRunDashboardNoContent() *CreateRunDashboardNoContent {
	return &CreateRunDashboardNoContent{}
}

/*CreateRunDashboardNoContent handles this case with default header values.

No content.
*/
type CreateRunDashboardNoContent struct {
	Payload interface{}
}

func (o *CreateRunDashboardNoContent) Error() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/{project}/runs/{run}/dashboards][%d] createRunDashboardNoContent  %+v", 204, o.Payload)
}

func (o *CreateRunDashboardNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *CreateRunDashboardNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateRunDashboardForbidden creates a CreateRunDashboardForbidden with default headers values
func NewCreateRunDashboardForbidden() *CreateRunDashboardForbidden {
	return &CreateRunDashboardForbidden{}
}

/*CreateRunDashboardForbidden handles this case with default header values.

You don't have permission to access the resource.
*/
type CreateRunDashboardForbidden struct {
	Payload interface{}
}

func (o *CreateRunDashboardForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/{project}/runs/{run}/dashboards][%d] createRunDashboardForbidden  %+v", 403, o.Payload)
}

func (o *CreateRunDashboardForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *CreateRunDashboardForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateRunDashboardNotFound creates a CreateRunDashboardNotFound with default headers values
func NewCreateRunDashboardNotFound() *CreateRunDashboardNotFound {
	return &CreateRunDashboardNotFound{}
}

/*CreateRunDashboardNotFound handles this case with default header values.

Resource does not exist.
*/
type CreateRunDashboardNotFound struct {
	Payload interface{}
}

func (o *CreateRunDashboardNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/{project}/runs/{run}/dashboards][%d] createRunDashboardNotFound  %+v", 404, o.Payload)
}

func (o *CreateRunDashboardNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *CreateRunDashboardNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateRunDashboardDefault creates a CreateRunDashboardDefault with default headers values
func NewCreateRunDashboardDefault(code int) *CreateRunDashboardDefault {
	return &CreateRunDashboardDefault{
		_statusCode: code,
	}
}

/*CreateRunDashboardDefault handles this case with default header values.

An unexpected error response
*/
type CreateRunDashboardDefault struct {
	_statusCode int

	Payload *service_model.RuntimeError
}

// Code gets the status code for the create run dashboard default response
func (o *CreateRunDashboardDefault) Code() int {
	return o._statusCode
}

func (o *CreateRunDashboardDefault) Error() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/{project}/runs/{run}/dashboards][%d] CreateRunDashboard default  %+v", o._statusCode, o.Payload)
}

func (o *CreateRunDashboardDefault) GetPayload() *service_model.RuntimeError {
	return o.Payload
}

func (o *CreateRunDashboardDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}