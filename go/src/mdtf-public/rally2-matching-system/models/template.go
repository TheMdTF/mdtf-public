package models

// Template Data transfer object for a template.
//
// swagger:model Template
type Template struct {

	// A base64 encoded string that contains template data.
	// Format: byte
	Template string `json:"Template,omitempty"`
}
