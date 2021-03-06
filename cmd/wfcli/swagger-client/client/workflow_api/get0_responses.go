package workflow_api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/fission/fission-workflows/cmd/wfcli/swagger-client/models"
)

// Get0Reader is a Reader for the Get0 structure.
type Get0Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *Get0Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGet0OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGet0OK creates a Get0OK with default headers values
func NewGet0OK() *Get0OK {
	return &Get0OK{}
}

/*Get0OK handles this case with default header values.

Get0OK get0 o k
*/
type Get0OK struct {
	Payload *models.Workflow
}

func (o *Get0OK) Error() string {
	return fmt.Sprintf("[GET /workflow/{id}][%d] get0OK  %+v", 200, o.Payload)
}

func (o *Get0OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Workflow)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
