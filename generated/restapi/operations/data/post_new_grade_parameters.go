// Code generated by go-swagger; DO NOT EDIT.

package data

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/gangozero/hackprague2019-be/generated/models"
)

// NewPostNewGradeParams creates a new PostNewGradeParams object
// no default values defined in spec.
func NewPostNewGradeParams() PostNewGradeParams {

	return PostNewGradeParams{}
}

// PostNewGradeParams contains all the bound params for the post new grade operation
// typically these are obtained from a http.Request
//
// swagger:parameters postNewGrade
type PostNewGradeParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  Required: true
	  In: body
	*/
	Data *models.Sample
	/*
	  Required: true
	  In: path
	*/
	ProfileID string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewPostNewGradeParams() beforehand.
func (o *PostNewGradeParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.Sample
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("data", "body"))
			} else {
				res = append(res, errors.NewParseError("data", "body", "", err))
			}
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.Data = &body
			}
		}
	} else {
		res = append(res, errors.Required("data", "body"))
	}
	rProfileID, rhkProfileID, _ := route.Params.GetOK("profile_id")
	if err := o.bindProfileID(rProfileID, rhkProfileID, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindProfileID binds and validates parameter ProfileID from path.
func (o *PostNewGradeParams) bindProfileID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.ProfileID = raw

	return nil
}
