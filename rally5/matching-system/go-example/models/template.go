package models

// Template Data transfer object for a template.
//
// swagger:model Template
type Template struct {

	// The template data, encoded as a base64 string.  The data string shall not exceed 1 MB.
	Template string
}
