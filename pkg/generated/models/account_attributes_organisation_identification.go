// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"log"
	"strconv"

	"github.com/form3tech-oss/go-form3/pkg/client"
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// AccountAttributesOrganisationIdentification account attributes organisation identification
// swagger:model AccountAttributesOrganisationIdentification
type AccountAttributesOrganisationIdentification struct {

	// address
	Address []string `json:"address"`

	// city
	// Max Length: 35
	// Min Length: 1
	City string `json:"city,omitempty"`

	// country
	// Pattern: ^[A-Z]{2}$
	Country string `json:"country,omitempty"`

	// name
	// Max Length: 40
	// Min Length: 1
	Name string `json:"name,omitempty"`

	// registration number
	RegistrationNumber string `json:"registration_number,omitempty"`

	// representative
	Representative *AccountAttributesOrganisationIdentificationRepresentative `json:"representative,omitempty"`

	// ISO 3166-1 code used to identify the domicile of the account
	// Pattern: ^[A-Z]{2}$
	TaxResidency string `json:"tax_residency,omitempty"`
}

func AccountAttributesOrganisationIdentificationWithDefaults(defaults client.Defaults) *AccountAttributesOrganisationIdentification {
	return &AccountAttributesOrganisationIdentification{

		Address: make([]string, 0),

		City: defaults.GetString("AccountAttributesOrganisationIdentification", "city"),

		Country: defaults.GetString("AccountAttributesOrganisationIdentification", "country"),

		Name: defaults.GetString("AccountAttributesOrganisationIdentification", "name"),

		RegistrationNumber: defaults.GetString("AccountAttributesOrganisationIdentification", "registration_number"),

		Representative: AccountAttributesOrganisationIdentificationRepresentativeWithDefaults(defaults),

		TaxResidency: defaults.GetString("AccountAttributesOrganisationIdentification", "tax_residency"),
	}
}

func (m *AccountAttributesOrganisationIdentification) WithAddress(address []string) *AccountAttributesOrganisationIdentification {

	m.Address = address

	return m
}

func (m *AccountAttributesOrganisationIdentification) WithCity(city string) *AccountAttributesOrganisationIdentification {

	m.City = city

	return m
}

func (m *AccountAttributesOrganisationIdentification) WithCountry(country string) *AccountAttributesOrganisationIdentification {

	m.Country = country

	return m
}

func (m *AccountAttributesOrganisationIdentification) WithName(name string) *AccountAttributesOrganisationIdentification {

	m.Name = name

	return m
}

func (m *AccountAttributesOrganisationIdentification) WithRegistrationNumber(registrationNumber string) *AccountAttributesOrganisationIdentification {

	m.RegistrationNumber = registrationNumber

	return m
}

func (m *AccountAttributesOrganisationIdentification) WithRepresentative(representative AccountAttributesOrganisationIdentificationRepresentative) *AccountAttributesOrganisationIdentification {

	m.Representative = &representative

	return m
}

func (m *AccountAttributesOrganisationIdentification) WithoutRepresentative() *AccountAttributesOrganisationIdentification {
	m.Representative = nil
	return m
}

func (m *AccountAttributesOrganisationIdentification) WithTaxResidency(taxResidency string) *AccountAttributesOrganisationIdentification {

	m.TaxResidency = taxResidency

	return m
}

// Validate validates this account attributes organisation identification
func (m *AccountAttributesOrganisationIdentification) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCountry(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRepresentative(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTaxResidency(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AccountAttributesOrganisationIdentification) validateAddress(formats strfmt.Registry) error {

	if swag.IsZero(m.Address) { // not required
		return nil
	}

	for i := 0; i < len(m.Address); i++ {

		if err := validate.MinLength("address"+"."+strconv.Itoa(i), "body", string(m.Address[i]), 1); err != nil {
			return err
		}

		if err := validate.MaxLength("address"+"."+strconv.Itoa(i), "body", string(m.Address[i]), 140); err != nil {
			return err
		}

	}

	return nil
}

func (m *AccountAttributesOrganisationIdentification) validateCity(formats strfmt.Registry) error {

	if swag.IsZero(m.City) { // not required
		return nil
	}

	if err := validate.MinLength("city", "body", string(m.City), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("city", "body", string(m.City), 35); err != nil {
		return err
	}

	return nil
}

func (m *AccountAttributesOrganisationIdentification) validateCountry(formats strfmt.Registry) error {

	if swag.IsZero(m.Country) { // not required
		return nil
	}

	if err := validate.Pattern("country", "body", string(m.Country), `^[A-Z]{2}$`); err != nil {
		return err
	}

	return nil
}

func (m *AccountAttributesOrganisationIdentification) validateName(formats strfmt.Registry) error {

	if swag.IsZero(m.Name) { // not required
		return nil
	}

	if err := validate.MinLength("name", "body", string(m.Name), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", string(m.Name), 40); err != nil {
		return err
	}

	return nil
}

func (m *AccountAttributesOrganisationIdentification) validateRepresentative(formats strfmt.Registry) error {

	if swag.IsZero(m.Representative) { // not required
		return nil
	}

	if m.Representative != nil {
		if err := m.Representative.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("representative")
			}
			return err
		}
	}

	return nil
}

func (m *AccountAttributesOrganisationIdentification) validateTaxResidency(formats strfmt.Registry) error {

	if swag.IsZero(m.TaxResidency) { // not required
		return nil
	}

	if err := validate.Pattern("tax_residency", "body", string(m.TaxResidency), `^[A-Z]{2}$`); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AccountAttributesOrganisationIdentification) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AccountAttributesOrganisationIdentification) UnmarshalBinary(b []byte) error {
	var res AccountAttributesOrganisationIdentification
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *AccountAttributesOrganisationIdentification) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// AccountAttributesOrganisationIdentificationRepresentative account attributes organisation identification representative
// swagger:model AccountAttributesOrganisationIdentificationRepresentative
type AccountAttributesOrganisationIdentificationRepresentative struct {

	// birth date
	// Format: date
	BirthDate *strfmt.Date `json:"birth_date,omitempty"`

	// name
	// Max Length: 40
	// Min Length: 1
	Name string `json:"name,omitempty"`

	// residency
	Residency string `json:"residency,omitempty"`
}

func AccountAttributesOrganisationIdentificationRepresentativeWithDefaults(defaults client.Defaults) *AccountAttributesOrganisationIdentificationRepresentative {
	return &AccountAttributesOrganisationIdentificationRepresentative{

		BirthDate: defaults.GetStrfmtDatePtr("AccountAttributesOrganisationIdentificationRepresentative", "birth_date"),

		Name: defaults.GetString("AccountAttributesOrganisationIdentificationRepresentative", "name"),

		Residency: defaults.GetString("AccountAttributesOrganisationIdentificationRepresentative", "residency"),
	}
}

func (m *AccountAttributesOrganisationIdentificationRepresentative) WithBirthDate(birthDate strfmt.Date) *AccountAttributesOrganisationIdentificationRepresentative {

	m.BirthDate = &birthDate

	return m
}

func (m *AccountAttributesOrganisationIdentificationRepresentative) WithoutBirthDate() *AccountAttributesOrganisationIdentificationRepresentative {
	m.BirthDate = nil
	return m
}

func (m *AccountAttributesOrganisationIdentificationRepresentative) WithName(name string) *AccountAttributesOrganisationIdentificationRepresentative {

	m.Name = name

	return m
}

func (m *AccountAttributesOrganisationIdentificationRepresentative) WithResidency(residency string) *AccountAttributesOrganisationIdentificationRepresentative {

	m.Residency = residency

	return m
}

// Validate validates this account attributes organisation identification representative
func (m *AccountAttributesOrganisationIdentificationRepresentative) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBirthDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AccountAttributesOrganisationIdentificationRepresentative) validateBirthDate(formats strfmt.Registry) error {

	if swag.IsZero(m.BirthDate) { // not required
		return nil
	}

	if err := validate.FormatOf("representative"+"."+"birth_date", "body", "date", m.BirthDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *AccountAttributesOrganisationIdentificationRepresentative) validateName(formats strfmt.Registry) error {

	if swag.IsZero(m.Name) { // not required
		return nil
	}

	if err := validate.MinLength("representative"+"."+"name", "body", string(m.Name), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("representative"+"."+"name", "body", string(m.Name), 40); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AccountAttributesOrganisationIdentificationRepresentative) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AccountAttributesOrganisationIdentificationRepresentative) UnmarshalBinary(b []byte) error {
	var res AccountAttributesOrganisationIdentificationRepresentative
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *AccountAttributesOrganisationIdentificationRepresentative) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
