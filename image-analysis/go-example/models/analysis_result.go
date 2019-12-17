package models

// AnalysisResult stores the AnalysisResult model for this algorithm including 3 example custom fields.
// Only Score is required for API compliance.
type AnalysisResult struct{
	Score 				float32

	Timestamp            string
	ComputationTimeMilli string
	AnalysisError        string
}
