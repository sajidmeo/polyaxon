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

package queues_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/polyaxon/polyaxon/sdks/go/http_client/v1/service_model"
)

// QueuesV1CreateQueueReader is a Reader for the QueuesV1CreateQueue structure.
type QueuesV1CreateQueueReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *QueuesV1CreateQueueReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewQueuesV1CreateQueueOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewQueuesV1CreateQueueNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewQueuesV1CreateQueueForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewQueuesV1CreateQueueNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewQueuesV1CreateQueueDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewQueuesV1CreateQueueOK creates a QueuesV1CreateQueueOK with default headers values
func NewQueuesV1CreateQueueOK() *QueuesV1CreateQueueOK {
	return &QueuesV1CreateQueueOK{}
}

/*QueuesV1CreateQueueOK handles this case with default header values.

A successful response.
*/
type QueuesV1CreateQueueOK struct {
	Payload *service_model.V1Agent
}

func (o *QueuesV1CreateQueueOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/orgs/{owner}/agents/{agent}/queues][%d] queuesV1CreateQueueOK  %+v", 200, o.Payload)
}

func (o *QueuesV1CreateQueueOK) GetPayload() *service_model.V1Agent {
	return o.Payload
}

func (o *QueuesV1CreateQueueOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.V1Agent)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueuesV1CreateQueueNoContent creates a QueuesV1CreateQueueNoContent with default headers values
func NewQueuesV1CreateQueueNoContent() *QueuesV1CreateQueueNoContent {
	return &QueuesV1CreateQueueNoContent{}
}

/*QueuesV1CreateQueueNoContent handles this case with default header values.

No content.
*/
type QueuesV1CreateQueueNoContent struct {
	Payload interface{}
}

func (o *QueuesV1CreateQueueNoContent) Error() string {
	return fmt.Sprintf("[POST /api/v1/orgs/{owner}/agents/{agent}/queues][%d] queuesV1CreateQueueNoContent  %+v", 204, o.Payload)
}

func (o *QueuesV1CreateQueueNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *QueuesV1CreateQueueNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueuesV1CreateQueueForbidden creates a QueuesV1CreateQueueForbidden with default headers values
func NewQueuesV1CreateQueueForbidden() *QueuesV1CreateQueueForbidden {
	return &QueuesV1CreateQueueForbidden{}
}

/*QueuesV1CreateQueueForbidden handles this case with default header values.

You don't have permission to access the resource.
*/
type QueuesV1CreateQueueForbidden struct {
	Payload interface{}
}

func (o *QueuesV1CreateQueueForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v1/orgs/{owner}/agents/{agent}/queues][%d] queuesV1CreateQueueForbidden  %+v", 403, o.Payload)
}

func (o *QueuesV1CreateQueueForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *QueuesV1CreateQueueForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueuesV1CreateQueueNotFound creates a QueuesV1CreateQueueNotFound with default headers values
func NewQueuesV1CreateQueueNotFound() *QueuesV1CreateQueueNotFound {
	return &QueuesV1CreateQueueNotFound{}
}

/*QueuesV1CreateQueueNotFound handles this case with default header values.

Resource does not exist.
*/
type QueuesV1CreateQueueNotFound struct {
	Payload interface{}
}

func (o *QueuesV1CreateQueueNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v1/orgs/{owner}/agents/{agent}/queues][%d] queuesV1CreateQueueNotFound  %+v", 404, o.Payload)
}

func (o *QueuesV1CreateQueueNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *QueuesV1CreateQueueNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueuesV1CreateQueueDefault creates a QueuesV1CreateQueueDefault with default headers values
func NewQueuesV1CreateQueueDefault(code int) *QueuesV1CreateQueueDefault {
	return &QueuesV1CreateQueueDefault{
		_statusCode: code,
	}
}

/*QueuesV1CreateQueueDefault handles this case with default header values.

An unexpected error response
*/
type QueuesV1CreateQueueDefault struct {
	_statusCode int

	Payload *service_model.RuntimeError
}

// Code gets the status code for the queues v1 create queue default response
func (o *QueuesV1CreateQueueDefault) Code() int {
	return o._statusCode
}

func (o *QueuesV1CreateQueueDefault) Error() string {
	return fmt.Sprintf("[POST /api/v1/orgs/{owner}/agents/{agent}/queues][%d] QueuesV1_CreateQueue default  %+v", o._statusCode, o.Payload)
}

func (o *QueuesV1CreateQueueDefault) GetPayload() *service_model.RuntimeError {
	return o.Payload
}

func (o *QueuesV1CreateQueueDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}