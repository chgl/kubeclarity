// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Component component
//
// swagger:model Component
type Component struct {

	// author
	// Max Length: 255
	// Min Length: 0
	Author *string `json:"author,omitempty"`

	// blake2b 256
	// Pattern: ^[0-9a-f]{64}$
	Blake2b256 string `json:"blake2b_256,omitempty"`

	// blake2b 384
	// Pattern: ^[0-9a-f]{96}$
	Blake2b384 string `json:"blake2b_384,omitempty"`

	// blake2b 512
	// Pattern: ^[0-9a-f]{128}$
	Blake2b512 string `json:"blake2b_512,omitempty"`

	// blake3
	// Pattern: ^[A-Fa-f0-9]*$
	Blake3 string `json:"blake3,omitempty"`

	// bom ref
	BomRef string `json:"bomRef,omitempty"`

	// children
	Children []*Component `json:"children"`

	// classifier
	// Enum: [APPLICATION FRAMEWORK LIBRARY CONTAINER OPERATING_SYSTEM DEVICE FIRMWARE FILE]
	Classifier string `json:"classifier,omitempty"`

	// copyright
	// Max Length: 1024
	// Min Length: 0
	Copyright *string `json:"copyright,omitempty"`

	// cpe
	// Max Length: 255
	// Min Length: 0
	// Pattern: (cpe:2\.3:[aho\*\-](:(((\?*|\*?)([a-zA-Z0-9\-\._]|(\\[\\\*\?!"#$$%&'\(\)\+,/:;<=>@\[\]\^`\{\|}~]))+(\?*|\*?))|[\*\-])){5}(:(([a-zA-Z]{2,3}(-([a-zA-Z]{2}|[0-9]{3}))?)|[\*\-]))(:(((\?*|\*?)([a-zA-Z0-9\-\._]|(\\[\\\*\?!"#$$%&'\(\)\+,/:;<=>@\[\]\^`\{\|}~]))+(\?*|\*?))|[\*\-])){4})|([c][pP][eE]:/[AHOaho]?(:[A-Za-z0-9\._\-~%]*){0,6})
	Cpe *string `json:"cpe,omitempty"`

	// description
	// Max Length: 1024
	// Min Length: 0
	Description *string `json:"description,omitempty"`

	// direct dependencies
	DirectDependencies string `json:"directDependencies,omitempty"`

	// extension
	// Max Length: 255
	// Min Length: 0
	Extension *string `json:"extension,omitempty"`

	// filename
	// Max Length: 255
	// Min Length: 0
	Filename *string `json:"filename,omitempty"`

	// group
	// Max Length: 255
	// Min Length: 0
	Group *string `json:"group,omitempty"`

	// is internal
	IsInternal bool `json:"isInternal,omitempty"`

	// last inherited risk score
	LastInheritedRiskScore float64 `json:"lastInheritedRiskScore,omitempty"`

	// license
	// Max Length: 255
	// Min Length: 0
	License *string `json:"license,omitempty"`

	// md5
	// Pattern: ^[0-9a-f]{32}$
	Md5 string `json:"md5,omitempty"`

	// metrics
	Metrics *DependencyMetrics `json:"metrics,omitempty"`

	// name
	// Max Length: 255
	// Min Length: 1
	Name string `json:"name,omitempty"`

	// notes
	Notes string `json:"notes,omitempty"`

	// parent
	Parent *Component `json:"parent,omitempty"`

	// project
	// Required: true
	Project *Project `json:"project"`

	// publisher
	// Max Length: 255
	// Min Length: 0
	Publisher *string `json:"publisher,omitempty"`

	// purl
	Purl string `json:"purl,omitempty"`

	// purl coordinates
	PurlCoordinates string `json:"purlCoordinates,omitempty"`

	// repository meta
	RepositoryMeta *RepositoryMetaComponent `json:"repositoryMeta,omitempty"`

	// resolved license
	ResolvedLicense *License `json:"resolvedLicense,omitempty"`

	// sha1
	// Pattern: ^[0-9a-f]{40}$
	Sha1 string `json:"sha1,omitempty"`

	// sha256
	// Pattern: ^[0-9a-f]{64}$
	Sha256 string `json:"sha256,omitempty"`

	// sha384
	// Pattern: ^[0-9a-f]{96}$
	Sha384 string `json:"sha384,omitempty"`

	// sha3 256
	// Pattern: ^[0-9a-f]{64}$
	Sha3256 string `json:"sha3_256,omitempty"`

	// sha3 384
	// Pattern: ^[0-9a-f]{96}$
	Sha3384 string `json:"sha3_384,omitempty"`

	// sha3 512
	// Pattern: ^[0-9a-f]{128}$
	Sha3512 string `json:"sha3_512,omitempty"`

	// sha512
	// Pattern: ^[0-9a-f]{128}$
	Sha512 string `json:"sha512,omitempty"`

	// swid tag Id
	// Max Length: 255
	// Min Length: 0
	SwidTagID *string `json:"swidTagId,omitempty"`

	// used by
	UsedBy int32 `json:"usedBy,omitempty"`

	// uuid
	// Required: true
	// Format: uuid
	UUID *strfmt.UUID `json:"uuid"`

	// version
	// Max Length: 255
	// Min Length: 0
	Version *string `json:"version,omitempty"`

	// vulnerabilities
	Vulnerabilities []*Vulnerability `json:"vulnerabilities"`
}

// Validate validates this component
func (m *Component) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAuthor(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBlake2b256(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBlake2b384(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBlake2b512(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBlake3(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateChildren(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClassifier(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCopyright(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCpe(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExtension(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFilename(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGroup(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLicense(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMd5(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMetrics(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateParent(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProject(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePublisher(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRepositoryMeta(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResolvedLicense(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSha1(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSha256(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSha384(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSha3256(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSha3384(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSha3512(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSha512(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSwidTagID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUUID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVersion(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVulnerabilities(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Component) validateAuthor(formats strfmt.Registry) error {
	if swag.IsZero(m.Author) { // not required
		return nil
	}

	if err := validate.MinLength("author", "body", *m.Author, 0); err != nil {
		return err
	}

	if err := validate.MaxLength("author", "body", *m.Author, 255); err != nil {
		return err
	}

	return nil
}

func (m *Component) validateBlake2b256(formats strfmt.Registry) error {
	if swag.IsZero(m.Blake2b256) { // not required
		return nil
	}

	if err := validate.Pattern("blake2b_256", "body", m.Blake2b256, `^[0-9a-f]{64}$`); err != nil {
		return err
	}

	return nil
}

func (m *Component) validateBlake2b384(formats strfmt.Registry) error {
	if swag.IsZero(m.Blake2b384) { // not required
		return nil
	}

	if err := validate.Pattern("blake2b_384", "body", m.Blake2b384, `^[0-9a-f]{96}$`); err != nil {
		return err
	}

	return nil
}

func (m *Component) validateBlake2b512(formats strfmt.Registry) error {
	if swag.IsZero(m.Blake2b512) { // not required
		return nil
	}

	if err := validate.Pattern("blake2b_512", "body", m.Blake2b512, `^[0-9a-f]{128}$`); err != nil {
		return err
	}

	return nil
}

func (m *Component) validateBlake3(formats strfmt.Registry) error {
	if swag.IsZero(m.Blake3) { // not required
		return nil
	}

	if err := validate.Pattern("blake3", "body", m.Blake3, `^[A-Fa-f0-9]*$`); err != nil {
		return err
	}

	return nil
}

func (m *Component) validateChildren(formats strfmt.Registry) error {
	if swag.IsZero(m.Children) { // not required
		return nil
	}

	for i := 0; i < len(m.Children); i++ {
		if swag.IsZero(m.Children[i]) { // not required
			continue
		}

		if m.Children[i] != nil {
			if err := m.Children[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("children" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var componentTypeClassifierPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["APPLICATION","FRAMEWORK","LIBRARY","CONTAINER","OPERATING_SYSTEM","DEVICE","FIRMWARE","FILE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		componentTypeClassifierPropEnum = append(componentTypeClassifierPropEnum, v)
	}
}

const (

	// ComponentClassifierAPPLICATION captures enum value "APPLICATION"
	ComponentClassifierAPPLICATION string = "APPLICATION"

	// ComponentClassifierFRAMEWORK captures enum value "FRAMEWORK"
	ComponentClassifierFRAMEWORK string = "FRAMEWORK"

	// ComponentClassifierLIBRARY captures enum value "LIBRARY"
	ComponentClassifierLIBRARY string = "LIBRARY"

	// ComponentClassifierCONTAINER captures enum value "CONTAINER"
	ComponentClassifierCONTAINER string = "CONTAINER"

	// ComponentClassifierOPERATINGSYSTEM captures enum value "OPERATING_SYSTEM"
	ComponentClassifierOPERATINGSYSTEM string = "OPERATING_SYSTEM"

	// ComponentClassifierDEVICE captures enum value "DEVICE"
	ComponentClassifierDEVICE string = "DEVICE"

	// ComponentClassifierFIRMWARE captures enum value "FIRMWARE"
	ComponentClassifierFIRMWARE string = "FIRMWARE"

	// ComponentClassifierFILE captures enum value "FILE"
	ComponentClassifierFILE string = "FILE"
)

// prop value enum
func (m *Component) validateClassifierEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, componentTypeClassifierPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Component) validateClassifier(formats strfmt.Registry) error {
	if swag.IsZero(m.Classifier) { // not required
		return nil
	}

	// value enum
	if err := m.validateClassifierEnum("classifier", "body", m.Classifier); err != nil {
		return err
	}

	return nil
}

func (m *Component) validateCopyright(formats strfmt.Registry) error {
	if swag.IsZero(m.Copyright) { // not required
		return nil
	}

	if err := validate.MinLength("copyright", "body", *m.Copyright, 0); err != nil {
		return err
	}

	if err := validate.MaxLength("copyright", "body", *m.Copyright, 1024); err != nil {
		return err
	}

	return nil
}

func (m *Component) validateCpe(formats strfmt.Registry) error {
	if swag.IsZero(m.Cpe) { // not required
		return nil
	}

	if err := validate.MinLength("cpe", "body", *m.Cpe, 0); err != nil {
		return err
	}

	if err := validate.MaxLength("cpe", "body", *m.Cpe, 255); err != nil {
		return err
	}

	if err := validate.Pattern("cpe", "body", *m.Cpe, `(cpe:2\.3:[aho\*\-](:(((\?*|\*?)([a-zA-Z0-9\-\._]|(\\[\\\*\?!"#$$%&'\(\)\+,/:;<=>@\[\]\^`+"`"+`\{\|}~]))+(\?*|\*?))|[\*\-])){5}(:(([a-zA-Z]{2,3}(-([a-zA-Z]{2}|[0-9]{3}))?)|[\*\-]))(:(((\?*|\*?)([a-zA-Z0-9\-\._]|(\\[\\\*\?!"#$$%&'\(\)\+,/:;<=>@\[\]\^`+"`"+`\{\|}~]))+(\?*|\*?))|[\*\-])){4})|([c][pP][eE]:/[AHOaho]?(:[A-Za-z0-9\._\-~%]*){0,6})`); err != nil {
		return err
	}

	return nil
}

func (m *Component) validateDescription(formats strfmt.Registry) error {
	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if err := validate.MinLength("description", "body", *m.Description, 0); err != nil {
		return err
	}

	if err := validate.MaxLength("description", "body", *m.Description, 1024); err != nil {
		return err
	}

	return nil
}

func (m *Component) validateExtension(formats strfmt.Registry) error {
	if swag.IsZero(m.Extension) { // not required
		return nil
	}

	if err := validate.MinLength("extension", "body", *m.Extension, 0); err != nil {
		return err
	}

	if err := validate.MaxLength("extension", "body", *m.Extension, 255); err != nil {
		return err
	}

	return nil
}

func (m *Component) validateFilename(formats strfmt.Registry) error {
	if swag.IsZero(m.Filename) { // not required
		return nil
	}

	if err := validate.MinLength("filename", "body", *m.Filename, 0); err != nil {
		return err
	}

	if err := validate.MaxLength("filename", "body", *m.Filename, 255); err != nil {
		return err
	}

	return nil
}

func (m *Component) validateGroup(formats strfmt.Registry) error {
	if swag.IsZero(m.Group) { // not required
		return nil
	}

	if err := validate.MinLength("group", "body", *m.Group, 0); err != nil {
		return err
	}

	if err := validate.MaxLength("group", "body", *m.Group, 255); err != nil {
		return err
	}

	return nil
}

func (m *Component) validateLicense(formats strfmt.Registry) error {
	if swag.IsZero(m.License) { // not required
		return nil
	}

	if err := validate.MinLength("license", "body", *m.License, 0); err != nil {
		return err
	}

	if err := validate.MaxLength("license", "body", *m.License, 255); err != nil {
		return err
	}

	return nil
}

func (m *Component) validateMd5(formats strfmt.Registry) error {
	if swag.IsZero(m.Md5) { // not required
		return nil
	}

	if err := validate.Pattern("md5", "body", m.Md5, `^[0-9a-f]{32}$`); err != nil {
		return err
	}

	return nil
}

func (m *Component) validateMetrics(formats strfmt.Registry) error {
	if swag.IsZero(m.Metrics) { // not required
		return nil
	}

	if m.Metrics != nil {
		if err := m.Metrics.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("metrics")
			}
			return err
		}
	}

	return nil
}

func (m *Component) validateName(formats strfmt.Registry) error {
	if swag.IsZero(m.Name) { // not required
		return nil
	}

	if err := validate.MinLength("name", "body", m.Name, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", m.Name, 255); err != nil {
		return err
	}

	return nil
}

func (m *Component) validateParent(formats strfmt.Registry) error {
	if swag.IsZero(m.Parent) { // not required
		return nil
	}

	if m.Parent != nil {
		if err := m.Parent.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("parent")
			}
			return err
		}
	}

	return nil
}

func (m *Component) validateProject(formats strfmt.Registry) error {

	if err := validate.Required("project", "body", m.Project); err != nil {
		return err
	}

	if m.Project != nil {
		if err := m.Project.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("project")
			}
			return err
		}
	}

	return nil
}

func (m *Component) validatePublisher(formats strfmt.Registry) error {
	if swag.IsZero(m.Publisher) { // not required
		return nil
	}

	if err := validate.MinLength("publisher", "body", *m.Publisher, 0); err != nil {
		return err
	}

	if err := validate.MaxLength("publisher", "body", *m.Publisher, 255); err != nil {
		return err
	}

	return nil
}

func (m *Component) validateRepositoryMeta(formats strfmt.Registry) error {
	if swag.IsZero(m.RepositoryMeta) { // not required
		return nil
	}

	if m.RepositoryMeta != nil {
		if err := m.RepositoryMeta.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("repositoryMeta")
			}
			return err
		}
	}

	return nil
}

func (m *Component) validateResolvedLicense(formats strfmt.Registry) error {
	if swag.IsZero(m.ResolvedLicense) { // not required
		return nil
	}

	if m.ResolvedLicense != nil {
		if err := m.ResolvedLicense.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("resolvedLicense")
			}
			return err
		}
	}

	return nil
}

func (m *Component) validateSha1(formats strfmt.Registry) error {
	if swag.IsZero(m.Sha1) { // not required
		return nil
	}

	if err := validate.Pattern("sha1", "body", m.Sha1, `^[0-9a-f]{40}$`); err != nil {
		return err
	}

	return nil
}

func (m *Component) validateSha256(formats strfmt.Registry) error {
	if swag.IsZero(m.Sha256) { // not required
		return nil
	}

	if err := validate.Pattern("sha256", "body", m.Sha256, `^[0-9a-f]{64}$`); err != nil {
		return err
	}

	return nil
}

func (m *Component) validateSha384(formats strfmt.Registry) error {
	if swag.IsZero(m.Sha384) { // not required
		return nil
	}

	if err := validate.Pattern("sha384", "body", m.Sha384, `^[0-9a-f]{96}$`); err != nil {
		return err
	}

	return nil
}

func (m *Component) validateSha3256(formats strfmt.Registry) error {
	if swag.IsZero(m.Sha3256) { // not required
		return nil
	}

	if err := validate.Pattern("sha3_256", "body", m.Sha3256, `^[0-9a-f]{64}$`); err != nil {
		return err
	}

	return nil
}

func (m *Component) validateSha3384(formats strfmt.Registry) error {
	if swag.IsZero(m.Sha3384) { // not required
		return nil
	}

	if err := validate.Pattern("sha3_384", "body", m.Sha3384, `^[0-9a-f]{96}$`); err != nil {
		return err
	}

	return nil
}

func (m *Component) validateSha3512(formats strfmt.Registry) error {
	if swag.IsZero(m.Sha3512) { // not required
		return nil
	}

	if err := validate.Pattern("sha3_512", "body", m.Sha3512, `^[0-9a-f]{128}$`); err != nil {
		return err
	}

	return nil
}

func (m *Component) validateSha512(formats strfmt.Registry) error {
	if swag.IsZero(m.Sha512) { // not required
		return nil
	}

	if err := validate.Pattern("sha512", "body", m.Sha512, `^[0-9a-f]{128}$`); err != nil {
		return err
	}

	return nil
}

func (m *Component) validateSwidTagID(formats strfmt.Registry) error {
	if swag.IsZero(m.SwidTagID) { // not required
		return nil
	}

	if err := validate.MinLength("swidTagId", "body", *m.SwidTagID, 0); err != nil {
		return err
	}

	if err := validate.MaxLength("swidTagId", "body", *m.SwidTagID, 255); err != nil {
		return err
	}

	return nil
}

func (m *Component) validateUUID(formats strfmt.Registry) error {

	if err := validate.Required("uuid", "body", m.UUID); err != nil {
		return err
	}

	if err := validate.FormatOf("uuid", "body", "uuid", m.UUID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Component) validateVersion(formats strfmt.Registry) error {
	if swag.IsZero(m.Version) { // not required
		return nil
	}

	if err := validate.MinLength("version", "body", *m.Version, 0); err != nil {
		return err
	}

	if err := validate.MaxLength("version", "body", *m.Version, 255); err != nil {
		return err
	}

	return nil
}

func (m *Component) validateVulnerabilities(formats strfmt.Registry) error {
	if swag.IsZero(m.Vulnerabilities) { // not required
		return nil
	}

	for i := 0; i < len(m.Vulnerabilities); i++ {
		if swag.IsZero(m.Vulnerabilities[i]) { // not required
			continue
		}

		if m.Vulnerabilities[i] != nil {
			if err := m.Vulnerabilities[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("vulnerabilities" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this component based on the context it is used
func (m *Component) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateChildren(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMetrics(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateParent(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateProject(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRepositoryMeta(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateResolvedLicense(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVulnerabilities(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Component) contextValidateChildren(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Children); i++ {

		if m.Children[i] != nil {
			if err := m.Children[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("children" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Component) contextValidateMetrics(ctx context.Context, formats strfmt.Registry) error {

	if m.Metrics != nil {
		if err := m.Metrics.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("metrics")
			}
			return err
		}
	}

	return nil
}

func (m *Component) contextValidateParent(ctx context.Context, formats strfmt.Registry) error {

	if m.Parent != nil {
		if err := m.Parent.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("parent")
			}
			return err
		}
	}

	return nil
}

func (m *Component) contextValidateProject(ctx context.Context, formats strfmt.Registry) error {

	if m.Project != nil {
		if err := m.Project.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("project")
			}
			return err
		}
	}

	return nil
}

func (m *Component) contextValidateRepositoryMeta(ctx context.Context, formats strfmt.Registry) error {

	if m.RepositoryMeta != nil {
		if err := m.RepositoryMeta.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("repositoryMeta")
			}
			return err
		}
	}

	return nil
}

func (m *Component) contextValidateResolvedLicense(ctx context.Context, formats strfmt.Registry) error {

	if m.ResolvedLicense != nil {
		if err := m.ResolvedLicense.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("resolvedLicense")
			}
			return err
		}
	}

	return nil
}

func (m *Component) contextValidateVulnerabilities(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Vulnerabilities); i++ {

		if m.Vulnerabilities[i] != nil {
			if err := m.Vulnerabilities[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("vulnerabilities" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Component) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Component) UnmarshalBinary(b []byte) error {
	var res Component
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}