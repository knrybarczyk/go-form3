// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"log"

	"github.com/form3tech-oss/go-form3/pkg/client"
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ClaimSubmission claim submission
// swagger:model ClaimSubmission
type ClaimSubmission struct {

	// attributes
	Attributes *ClaimSubmissionAttributes `json:"attributes,omitempty"`

	// id
	// Required: true
	// Format: uuid
	ID *strfmt.UUID `json:"id"`

	// organisation id
	// Required: true
	// Format: uuid
	OrganisationID *strfmt.UUID `json:"organisation_id"`

	// relationships
	Relationships *ClaimSubmissionRelationships `json:"relationships,omitempty"`

	// type
	// Pattern: ^[A-Za-z_]*$
	Type string `json:"type,omitempty"`

	// version
	// Minimum: 0
	Version *int64 `json:"version,omitempty"`
}

func ClaimSubmissionWithDefaults(defaults client.Defaults) *ClaimSubmission {
	return &ClaimSubmission{

		Attributes: ClaimSubmissionAttributesWithDefaults(defaults),

		ID: defaults.GetStrfmtUUIDPtr("ClaimSubmission", "id"),

		OrganisationID: defaults.GetStrfmtUUIDPtr("ClaimSubmission", "organisation_id"),

		Relationships: ClaimSubmissionRelationshipsWithDefaults(defaults),

		Type: defaults.GetString("ClaimSubmission", "type"),

		Version: defaults.GetInt64Ptr("ClaimSubmission", "version"),
	}
}

func (m *ClaimSubmission) WithAttributes(attributes ClaimSubmissionAttributes) *ClaimSubmission {

	m.Attributes = &attributes

	return m
}

func (m *ClaimSubmission) WithoutAttributes() *ClaimSubmission {
	m.Attributes = nil
	return m
}

func (m *ClaimSubmission) WithID(id strfmt.UUID) *ClaimSubmission {

	m.ID = &id

	return m
}

func (m *ClaimSubmission) WithoutID() *ClaimSubmission {
	m.ID = nil
	return m
}

func (m *ClaimSubmission) WithOrganisationID(organisationID strfmt.UUID) *ClaimSubmission {

	m.OrganisationID = &organisationID

	return m
}

func (m *ClaimSubmission) WithoutOrganisationID() *ClaimSubmission {
	m.OrganisationID = nil
	return m
}

func (m *ClaimSubmission) WithRelationships(relationships ClaimSubmissionRelationships) *ClaimSubmission {

	m.Relationships = &relationships

	return m
}

func (m *ClaimSubmission) WithoutRelationships() *ClaimSubmission {
	m.Relationships = nil
	return m
}

func (m *ClaimSubmission) WithType(typeVar string) *ClaimSubmission {

	m.Type = typeVar

	return m
}

func (m *ClaimSubmission) WithVersion(version int64) *ClaimSubmission {

	m.Version = &version

	return m
}

func (m *ClaimSubmission) WithoutVersion() *ClaimSubmission {
	m.Version = nil
	return m
}

// Validate validates this claim submission
func (m *ClaimSubmission) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAttributes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrganisationID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRelationships(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVersion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClaimSubmission) validateAttributes(formats strfmt.Registry) error {

	if swag.IsZero(m.Attributes) { // not required
		return nil
	}

	if m.Attributes != nil {
		if err := m.Attributes.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("attributes")
			}
			return err
		}
	}

	return nil
}

func (m *ClaimSubmission) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ClaimSubmission) validateOrganisationID(formats strfmt.Registry) error {

	if err := validate.Required("organisation_id", "body", m.OrganisationID); err != nil {
		return err
	}

	if err := validate.FormatOf("organisation_id", "body", "uuid", m.OrganisationID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ClaimSubmission) validateRelationships(formats strfmt.Registry) error {

	if swag.IsZero(m.Relationships) { // not required
		return nil
	}

	if m.Relationships != nil {
		if err := m.Relationships.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("relationships")
			}
			return err
		}
	}

	return nil
}

func (m *ClaimSubmission) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	if err := validate.Pattern("type", "body", string(m.Type), `^[A-Za-z_]*$`); err != nil {
		return err
	}

	return nil
}

func (m *ClaimSubmission) validateVersion(formats strfmt.Registry) error {

	if swag.IsZero(m.Version) { // not required
		return nil
	}

	if err := validate.MinimumInt("version", "body", int64(*m.Version), 0, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ClaimSubmission) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClaimSubmission) UnmarshalBinary(b []byte) error {
	var res ClaimSubmission
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *ClaimSubmission) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// ClaimSubmissionAttributes claim submission attributes
// swagger:model ClaimSubmissionAttributes
type ClaimSubmissionAttributes struct {

	// scheme message id
	SchemeMessageID string `json:"scheme_message_id,omitempty"`

	// status
	Status ClaimSubmissionStatus `json:"status,omitempty"`

	// status reason
	StatusReason string `json:"status_reason,omitempty"`

	// submission datetime
	// Format: date-time
	SubmissionDatetime strfmt.DateTime `json:"submission_datetime,omitempty"`
}

func ClaimSubmissionAttributesWithDefaults(defaults client.Defaults) *ClaimSubmissionAttributes {
	return &ClaimSubmissionAttributes{

		SchemeMessageID: defaults.GetString("ClaimSubmissionAttributes", "scheme_message_id"),

		// TODO Status: ClaimSubmissionStatus,

		StatusReason: defaults.GetString("ClaimSubmissionAttributes", "status_reason"),

		SubmissionDatetime: defaults.GetStrfmtDateTime("ClaimSubmissionAttributes", "submission_datetime"),
	}
}

func (m *ClaimSubmissionAttributes) WithSchemeMessageID(schemeMessageID string) *ClaimSubmissionAttributes {

	m.SchemeMessageID = schemeMessageID

	return m
}

func (m *ClaimSubmissionAttributes) WithStatus(status ClaimSubmissionStatus) *ClaimSubmissionAttributes {

	m.Status = status

	return m
}

func (m *ClaimSubmissionAttributes) WithStatusReason(statusReason string) *ClaimSubmissionAttributes {

	m.StatusReason = statusReason

	return m
}

func (m *ClaimSubmissionAttributes) WithSubmissionDatetime(submissionDatetime strfmt.DateTime) *ClaimSubmissionAttributes {

	m.SubmissionDatetime = submissionDatetime

	return m
}

// Validate validates this claim submission attributes
func (m *ClaimSubmissionAttributes) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubmissionDatetime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClaimSubmissionAttributes) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	if err := m.Status.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("attributes" + "." + "status")
		}
		return err
	}

	return nil
}

func (m *ClaimSubmissionAttributes) validateSubmissionDatetime(formats strfmt.Registry) error {

	if swag.IsZero(m.SubmissionDatetime) { // not required
		return nil
	}

	if err := validate.FormatOf("attributes"+"."+"submission_datetime", "body", "date-time", m.SubmissionDatetime.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ClaimSubmissionAttributes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClaimSubmissionAttributes) UnmarshalBinary(b []byte) error {
	var res ClaimSubmissionAttributes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *ClaimSubmissionAttributes) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}