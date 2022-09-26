// Code generated by go-swagger; DO NOT EDIT.

// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
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
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/firecracker-microvm/firecracker-go-sdk/client/models"
)

// PutGuestBootSourceReader is a Reader for the PutGuestBootSource structure.
type PutGuestBootSourceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutGuestBootSourceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPutGuestBootSourceNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutGuestBootSourceBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewPutGuestBootSourceDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPutGuestBootSourceNoContent creates a PutGuestBootSourceNoContent with default headers values
func NewPutGuestBootSourceNoContent() *PutGuestBootSourceNoContent {
	return &PutGuestBootSourceNoContent{}
}

/*PutGuestBootSourceNoContent handles this case with default header values.

Boot source created/updated
*/
type PutGuestBootSourceNoContent struct {
}

func (o *PutGuestBootSourceNoContent) Error() string {
	return fmt.Sprintf("[PUT /boot-source][%d] putGuestBootSourceNoContent ", 204)
}

func (o *PutGuestBootSourceNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutGuestBootSourceBadRequest creates a PutGuestBootSourceBadRequest with default headers values
func NewPutGuestBootSourceBadRequest() *PutGuestBootSourceBadRequest {
	return &PutGuestBootSourceBadRequest{}
}

/*PutGuestBootSourceBadRequest handles this case with default header values.

Boot source cannot be created due to bad input
*/
type PutGuestBootSourceBadRequest struct {
	Payload *models.Error
}

func (o *PutGuestBootSourceBadRequest) Error() string {
	return fmt.Sprintf("[PUT /boot-source][%d] putGuestBootSourceBadRequest  %+v", 400, o.Payload)
}

func (o *PutGuestBootSourceBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutGuestBootSourceBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutGuestBootSourceDefault creates a PutGuestBootSourceDefault with default headers values
func NewPutGuestBootSourceDefault(code int) *PutGuestBootSourceDefault {
	return &PutGuestBootSourceDefault{
		_statusCode: code,
	}
}

/*PutGuestBootSourceDefault handles this case with default header values.

Internal server error
*/
type PutGuestBootSourceDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the put guest boot source default response
func (o *PutGuestBootSourceDefault) Code() int {
	return o._statusCode
}

func (o *PutGuestBootSourceDefault) Error() string {
	return fmt.Sprintf("[PUT /boot-source][%d] putGuestBootSource default  %+v", o._statusCode, o.Payload)
}

func (o *PutGuestBootSourceDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutGuestBootSourceDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
