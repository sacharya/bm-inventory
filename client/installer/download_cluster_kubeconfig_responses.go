// Code generated by go-swagger; DO NOT EDIT.

package installer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/filanov/bm-inventory/models"
)

// DownloadClusterKubeconfigReader is a Reader for the DownloadClusterKubeconfig structure.
type DownloadClusterKubeconfigReader struct {
	formats strfmt.Registry
	writer  io.Writer
}

// ReadResponse reads a server response into the received o.
func (o *DownloadClusterKubeconfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDownloadClusterKubeconfigOK(o.writer)
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewDownloadClusterKubeconfigNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewDownloadClusterKubeconfigConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDownloadClusterKubeconfigInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDownloadClusterKubeconfigOK creates a DownloadClusterKubeconfigOK with default headers values
func NewDownloadClusterKubeconfigOK(writer io.Writer) *DownloadClusterKubeconfigOK {
	return &DownloadClusterKubeconfigOK{
		Payload: writer,
	}
}

/*DownloadClusterKubeconfigOK handles this case with default header values.

Success.
*/
type DownloadClusterKubeconfigOK struct {
	Payload io.Writer
}

func (o *DownloadClusterKubeconfigOK) Error() string {
	return fmt.Sprintf("[GET /clusters/{cluster_id}/downloads/kubeconfig][%d] downloadClusterKubeconfigOK  %+v", 200, o.Payload)
}

func (o *DownloadClusterKubeconfigOK) GetPayload() io.Writer {
	return o.Payload
}

func (o *DownloadClusterKubeconfigOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDownloadClusterKubeconfigNotFound creates a DownloadClusterKubeconfigNotFound with default headers values
func NewDownloadClusterKubeconfigNotFound() *DownloadClusterKubeconfigNotFound {
	return &DownloadClusterKubeconfigNotFound{}
}

/*DownloadClusterKubeconfigNotFound handles this case with default header values.

Error.
*/
type DownloadClusterKubeconfigNotFound struct {
	Payload *models.Error
}

func (o *DownloadClusterKubeconfigNotFound) Error() string {
	return fmt.Sprintf("[GET /clusters/{cluster_id}/downloads/kubeconfig][%d] downloadClusterKubeconfigNotFound  %+v", 404, o.Payload)
}

func (o *DownloadClusterKubeconfigNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *DownloadClusterKubeconfigNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDownloadClusterKubeconfigConflict creates a DownloadClusterKubeconfigConflict with default headers values
func NewDownloadClusterKubeconfigConflict() *DownloadClusterKubeconfigConflict {
	return &DownloadClusterKubeconfigConflict{}
}

/*DownloadClusterKubeconfigConflict handles this case with default header values.

Error.
*/
type DownloadClusterKubeconfigConflict struct {
	Payload *models.Error
}

func (o *DownloadClusterKubeconfigConflict) Error() string {
	return fmt.Sprintf("[GET /clusters/{cluster_id}/downloads/kubeconfig][%d] downloadClusterKubeconfigConflict  %+v", 409, o.Payload)
}

func (o *DownloadClusterKubeconfigConflict) GetPayload() *models.Error {
	return o.Payload
}

func (o *DownloadClusterKubeconfigConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDownloadClusterKubeconfigInternalServerError creates a DownloadClusterKubeconfigInternalServerError with default headers values
func NewDownloadClusterKubeconfigInternalServerError() *DownloadClusterKubeconfigInternalServerError {
	return &DownloadClusterKubeconfigInternalServerError{}
}

/*DownloadClusterKubeconfigInternalServerError handles this case with default header values.

Error.
*/
type DownloadClusterKubeconfigInternalServerError struct {
	Payload *models.Error
}

func (o *DownloadClusterKubeconfigInternalServerError) Error() string {
	return fmt.Sprintf("[GET /clusters/{cluster_id}/downloads/kubeconfig][%d] downloadClusterKubeconfigInternalServerError  %+v", 500, o.Payload)
}

func (o *DownloadClusterKubeconfigInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *DownloadClusterKubeconfigInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
