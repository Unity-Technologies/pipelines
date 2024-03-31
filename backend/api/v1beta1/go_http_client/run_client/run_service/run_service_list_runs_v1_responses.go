// Code generated by go-swagger; DO NOT EDIT.

package run_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	run_model "github.com/kubeflow/pipelines/backend/api/v1beta1/go_http_client/run_model"
)

// RunServiceListRunsV1Reader is a Reader for the RunServiceListRunsV1 structure.
type RunServiceListRunsV1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RunServiceListRunsV1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewRunServiceListRunsV1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewRunServiceListRunsV1Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewRunServiceListRunsV1OK creates a RunServiceListRunsV1OK with default headers values
func NewRunServiceListRunsV1OK() *RunServiceListRunsV1OK {
	return &RunServiceListRunsV1OK{}
}

/*RunServiceListRunsV1OK handles this case with default header values.

A successful response.
*/
type RunServiceListRunsV1OK struct {
	Payload *run_model.APIListRunsResponse
}

func (o *RunServiceListRunsV1OK) Error() string {
	return fmt.Sprintf("[GET /apis/v1beta1/runs][%d] runServiceListRunsV1OK  %+v", 200, o.Payload)
}

func (o *RunServiceListRunsV1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(run_model.APIListRunsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRunServiceListRunsV1Default creates a RunServiceListRunsV1Default with default headers values
func NewRunServiceListRunsV1Default(code int) *RunServiceListRunsV1Default {
	return &RunServiceListRunsV1Default{
		_statusCode: code,
	}
}

/*RunServiceListRunsV1Default handles this case with default header values.

An unexpected error response.
*/
type RunServiceListRunsV1Default struct {
	_statusCode int

	Payload *run_model.GatewayruntimeError
}

// Code gets the status code for the run service list runs v1 default response
func (o *RunServiceListRunsV1Default) Code() int {
	return o._statusCode
}

func (o *RunServiceListRunsV1Default) Error() string {
	return fmt.Sprintf("[GET /apis/v1beta1/runs][%d] RunService_ListRunsV1 default  %+v", o._statusCode, o.Payload)
}

func (o *RunServiceListRunsV1Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(run_model.GatewayruntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}