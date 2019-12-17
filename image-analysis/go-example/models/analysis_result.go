package models

// AnalysisResult stores the AnalysisResult model for this algorithm including 4 custom fields.
type AnalysisResult struct{
	Score 				string
	Timestamp            string
	ComputationTimeMilli string
	AnalysisError        string
}
