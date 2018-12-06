package models

import (
	strfmt "github.com/go-openapi/strfmt"
)

// Template Data transfer object for a template.
//
// swagger:model Template
type Template struct {

	// A base64 encoded string that contains template data.
	// Format: byte
	Template strfmt.Base64 `json:"Template,omitempty"`
}
