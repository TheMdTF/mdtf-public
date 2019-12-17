package models

// AnalysisResult stores the AnalysisResult model for this algorithm including 3 example custom fields.
// Only Score is a required field and up to 7 custom string fields may be provided.
type AnalysisResult struct{
	Score 				float32

	Timestamp            string
	ComputationTimeMilli string
	AnalysisError        string
}