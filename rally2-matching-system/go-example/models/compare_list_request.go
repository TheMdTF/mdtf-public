package models

// CompareListRequest A single template object and a list of templates that it will be compared to.
//
// swagger:model CompareListRequest
type CompareListRequest struct {

	// single template
	SingleTemplate *Template `json:"SingleTemplate,omitempty"`

	// template list
	TemplateList []*Template `json:"TemplateList"`
}
