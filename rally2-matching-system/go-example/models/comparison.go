package models

// Comparison A similarity score for a single (1:1) template comparison operation.
//
// swagger:model Comparison
type Comparison struct {

	// Similarity score between 0 and 1, with 1 being the highest score the algorithm can produce
	NormalizedScore float32

	// An un-normalized similarity score, as produced by the algorithm
	Score float32
}
