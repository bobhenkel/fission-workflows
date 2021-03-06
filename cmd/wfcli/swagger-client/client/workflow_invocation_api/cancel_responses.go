package workflow_invocation_api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/fission/fission-workflows/cmd/wfcli/swagger-client/models"
)

// CancelReader is a Reader for the Cancel structure.
type CancelReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CancelReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCancelOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCancelOK creates a CancelOK with default headers values
func NewCancelOK() *CancelOK {
	return &CancelOK{}
}

/*CancelOK handles this case with default header values.

CancelOK cancel o k
*/
type CancelOK struct {
	Payload models.ProtobufEmpty
}

func (o *CancelOK) Error() string {
	return fmt.Sprintf("[DELETE /invocation/{id}][%d] cancelOK  %+v", 200, o.Payload)
}

func (o *CancelOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
