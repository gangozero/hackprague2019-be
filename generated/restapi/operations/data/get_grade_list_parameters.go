// Code generated by go-swagger; DO NOT EDIT.

package data

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetGradeListParams creates a new GetGradeListParams object
// no default values defined in spec.
func NewGetGradeListParams() GetGradeListParams {

	return GetGradeListParams{}
}

// GetGradeListParams contains all the bound params for the get grade list operation
// typically these are obtained from a http.Request
//
// swagger:parameters getGradeList
type GetGradeListParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  Required: true
	  In: path
	*/
	ProfileID string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetGradeListParams() beforehand.
func (o *GetGradeListParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

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
func (o *GetGradeListParams) bindProfileID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.ProfileID = raw

	return nil
}