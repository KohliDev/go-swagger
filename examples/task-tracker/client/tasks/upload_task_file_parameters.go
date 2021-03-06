package tasks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"os"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewUploadTaskFileParams creates a new UploadTaskFileParams object
// with the default values initialized.
func NewUploadTaskFileParams() *UploadTaskFileParams {
	var ()
	return &UploadTaskFileParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUploadTaskFileParamsWithTimeout creates a new UploadTaskFileParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUploadTaskFileParamsWithTimeout(timeout time.Duration) *UploadTaskFileParams {
	var ()
	return &UploadTaskFileParams{

		timeout: timeout,
	}
}

/*UploadTaskFileParams contains all the parameters to send to the API endpoint
for the upload task file operation typically these are written to a http.Request
*/
type UploadTaskFileParams struct {

	/*Description
	  Extra information describing the file

	*/
	Description *string
	/*File
	  The file to upload

	*/
	File *os.File
	/*ID
	  The id of the item

	*/
	ID int64

	timeout time.Duration
}

// WithDescription adds the description to the upload task file params
func (o *UploadTaskFileParams) WithDescription(description *string) *UploadTaskFileParams {
	o.Description = description
	return o
}

// WithFile adds the file to the upload task file params
func (o *UploadTaskFileParams) WithFile(file *os.File) *UploadTaskFileParams {
	o.File = file
	return o
}

// WithID adds the id to the upload task file params
func (o *UploadTaskFileParams) WithID(id int64) *UploadTaskFileParams {
	o.ID = id
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *UploadTaskFileParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if o.Description != nil {

		// form param description
		var frDescription string
		if o.Description != nil {
			frDescription = *o.Description
		}
		fDescription := frDescription
		if fDescription != "" {
			if err := r.SetFormParam("description", fDescription); err != nil {
				return err
			}
		}

	}

	if o.File != nil {

		if o.File != nil {

			// form file param file
			if err := r.SetFileParam("file", o.File); err != nil {
				return err
			}

		}

	}

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
