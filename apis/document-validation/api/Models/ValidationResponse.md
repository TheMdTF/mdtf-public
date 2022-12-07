# ValidationResponse
## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **ValidityOutcome** | **Boolean** | Whether the document is determined to be valid (True) or invalid (False) | [default to null] |
| **ValidityScore** | **Double** | A score corresponding to the level of confidence that the document  is valid ranging between 0 and 1 (optional) | [optional] [default to null] |
| **ValidityProperties** | [**List**](ValidityProperty.md) | Key value pairs describing document properties and their relationship to the validity decision. There are no strictly defined  properties. The inclusion of descriptive properties is encouraged to provide more context. (optional) | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

