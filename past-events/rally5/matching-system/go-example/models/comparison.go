package models

// Comparison A similarity score for a single (1:1) template comparison operation.
//
// swagger:model Comparison
type Comparison struct {
	// A similarity score, as produced by the algorithm
	Score float32
}
