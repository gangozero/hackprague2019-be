// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// Sample sample
// swagger:model sample
type Sample struct {

	// grade
	Grade int64 `json:"grade,omitempty"`

	// Lat
	Lat float64 `json:"lat,omitempty"`

	// Lon
	Lon float64 `json:"lon,omitempty"`

	// ts
	Ts int64 `json:"ts,omitempty"`

	// User ID
	UserID string `json:"user_id,omitempty"`
}

// Validate validates this sample
func (m *Sample) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Sample) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Sample) UnmarshalBinary(b []byte) error {
	var res Sample
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
