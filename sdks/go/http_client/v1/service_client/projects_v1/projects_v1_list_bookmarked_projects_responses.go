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

package projects_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/polyaxon/polyaxon/sdks/go/http_client/v1/service_model"
)

// ProjectsV1ListBookmarkedProjectsReader is a Reader for the ProjectsV1ListBookmarkedProjects structure.
type ProjectsV1ListBookmarkedProjectsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProjectsV1ListBookmarkedProjectsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewProjectsV1ListBookmarkedProjectsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewProjectsV1ListBookmarkedProjectsNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewProjectsV1ListBookmarkedProjectsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewProjectsV1ListBookmarkedProjectsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewProjectsV1ListBookmarkedProjectsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewProjectsV1ListBookmarkedProjectsOK creates a ProjectsV1ListBookmarkedProjectsOK with default headers values
func NewProjectsV1ListBookmarkedProjectsOK() *ProjectsV1ListBookmarkedProjectsOK {
	return &ProjectsV1ListBookmarkedProjectsOK{}
}

/*ProjectsV1ListBookmarkedProjectsOK handles this case with default header values.

A successful response.
*/
type ProjectsV1ListBookmarkedProjectsOK struct {
	Payload *service_model.V1ListProjectsResponse
}

func (o *ProjectsV1ListBookmarkedProjectsOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/bookmarks/{user}/projects][%d] projectsV1ListBookmarkedProjectsOK  %+v", 200, o.Payload)
}

func (o *ProjectsV1ListBookmarkedProjectsOK) GetPayload() *service_model.V1ListProjectsResponse {
	return o.Payload
}

func (o *ProjectsV1ListBookmarkedProjectsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.V1ListProjectsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectsV1ListBookmarkedProjectsNoContent creates a ProjectsV1ListBookmarkedProjectsNoContent with default headers values
func NewProjectsV1ListBookmarkedProjectsNoContent() *ProjectsV1ListBookmarkedProjectsNoContent {
	return &ProjectsV1ListBookmarkedProjectsNoContent{}
}

/*ProjectsV1ListBookmarkedProjectsNoContent handles this case with default header values.

No content.
*/
type ProjectsV1ListBookmarkedProjectsNoContent struct {
	Payload interface{}
}

func (o *ProjectsV1ListBookmarkedProjectsNoContent) Error() string {
	return fmt.Sprintf("[GET /api/v1/bookmarks/{user}/projects][%d] projectsV1ListBookmarkedProjectsNoContent  %+v", 204, o.Payload)
}

func (o *ProjectsV1ListBookmarkedProjectsNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *ProjectsV1ListBookmarkedProjectsNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectsV1ListBookmarkedProjectsForbidden creates a ProjectsV1ListBookmarkedProjectsForbidden with default headers values
func NewProjectsV1ListBookmarkedProjectsForbidden() *ProjectsV1ListBookmarkedProjectsForbidden {
	return &ProjectsV1ListBookmarkedProjectsForbidden{}
}

/*ProjectsV1ListBookmarkedProjectsForbidden handles this case with default header values.

You don't have permission to access the resource.
*/
type ProjectsV1ListBookmarkedProjectsForbidden struct {
	Payload interface{}
}

func (o *ProjectsV1ListBookmarkedProjectsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/bookmarks/{user}/projects][%d] projectsV1ListBookmarkedProjectsForbidden  %+v", 403, o.Payload)
}

func (o *ProjectsV1ListBookmarkedProjectsForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *ProjectsV1ListBookmarkedProjectsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectsV1ListBookmarkedProjectsNotFound creates a ProjectsV1ListBookmarkedProjectsNotFound with default headers values
func NewProjectsV1ListBookmarkedProjectsNotFound() *ProjectsV1ListBookmarkedProjectsNotFound {
	return &ProjectsV1ListBookmarkedProjectsNotFound{}
}

/*ProjectsV1ListBookmarkedProjectsNotFound handles this case with default header values.

Resource does not exist.
*/
type ProjectsV1ListBookmarkedProjectsNotFound struct {
	Payload interface{}
}

func (o *ProjectsV1ListBookmarkedProjectsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/bookmarks/{user}/projects][%d] projectsV1ListBookmarkedProjectsNotFound  %+v", 404, o.Payload)
}

func (o *ProjectsV1ListBookmarkedProjectsNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *ProjectsV1ListBookmarkedProjectsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectsV1ListBookmarkedProjectsDefault creates a ProjectsV1ListBookmarkedProjectsDefault with default headers values
func NewProjectsV1ListBookmarkedProjectsDefault(code int) *ProjectsV1ListBookmarkedProjectsDefault {
	return &ProjectsV1ListBookmarkedProjectsDefault{
		_statusCode: code,
	}
}

/*ProjectsV1ListBookmarkedProjectsDefault handles this case with default header values.

An unexpected error response
*/
type ProjectsV1ListBookmarkedProjectsDefault struct {
	_statusCode int

	Payload *service_model.RuntimeError
}

// Code gets the status code for the projects v1 list bookmarked projects default response
func (o *ProjectsV1ListBookmarkedProjectsDefault) Code() int {
	return o._statusCode
}

func (o *ProjectsV1ListBookmarkedProjectsDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/bookmarks/{user}/projects][%d] ProjectsV1_ListBookmarkedProjects default  %+v", o._statusCode, o.Payload)
}

func (o *ProjectsV1ListBookmarkedProjectsDefault) GetPayload() *service_model.RuntimeError {
	return o.Payload
}

func (o *ProjectsV1ListBookmarkedProjectsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}