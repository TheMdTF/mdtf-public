package models

// Info Basic information describing the algorithm.
type Info struct {

	// Name of algorithm
	AlgorithmName string

	// A string enum describing the type of biometric images the algorithm is meant to process
	// Enum: [Face Finger Iris]
	AlgorithmType string

	// Algorithm version identifier
	AlgorithmVersion string

	// Name of the Company which produces the algorithm
	CompanyName string

	// The recommended allocation of CPUs for the deployed docker container.
	RecommendedCPUs int64

	// The recommended allocation of memory (MB) for the deployed docker container.
	RecommendedMem int64

	// The email address of an engineer or other technical resource to contact in the event of an error running your service.
	// This field may be left blank if desired.
	TechnicalContactEmail string

	// The list of threshold scores for various false match rates (FMRs) for your submitted matching system.
	// Systems are required to provide at least thresholds for achieving FMRs for 1:500, 1:1e3,  1:1e4,  1:1e5,  1:1e6
	Thresholds map[string] float64
}
