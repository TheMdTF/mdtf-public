package models

// Info Basic information describing the algorithm.
//
// swagger:model Info
type Info struct {

	// Name of algorithm
	AlgorithmName string `json:"AlgorithmName,omitempty"`

	// A string enum describing the type of biometric images the algorithm is meant to process
	// Enum: [Face Finger Iris]
	AlgorithmType string `json:"AlgorithmType,omitempty"`

	// Algorithm version identifier
	AlgorithmVersion string `json:"AlgorithmVersion,omitempty"`

	// Name of the Company which produces the algorithm
	CompanyName string `json:"CompanyName,omitempty"`

	// The recommended allocation of CPUs for the deployed docker container.
	RecommendedCpus int64 `json:"RecommendedCPUs,omitempty"`

	// The recommended allocation of memory (MB) for the deployed docker container.
	RecommendedMem int64 `json:"RecommendedMem,omitempty"`

	// The email address of an engineer or other technical resource to contact in the event of an error running your service. This field may be left blank if desired.
	TechnicalContactEmail string `json:"TechnicalContactEmail,omitempty"`
}
