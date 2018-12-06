package models

type Image struct {

	// The captured image data in PNG format, encoded as a base64 string. The data string shall not exceed 1MB.
	ImageData string `json:"ImageData,omitempty"`
}
