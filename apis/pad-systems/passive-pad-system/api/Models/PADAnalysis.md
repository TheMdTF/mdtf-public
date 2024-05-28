# PADAnalysis
## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **PADOutcome** | **Boolean** | Whether a presentation attack was determined to be detected (True) or not detected (False).  The outcome should be calibrated to produce a BPCER of 0.01. | [default to null] |
| **PADScore** | **Double** | A score corresponding to the level of confidence that a presentation attack was detected ranging between 0 and 1. | [default to null] |
| **PADProperties** | [**List**](PADProperty.md) | Key value pairs describing presentation attack properties and their relationship to the presentation attack outcome/score.  There are no strictly defined properties.  The inclusion of descriptive properties is encouraged to provide more context.  (optional) | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

