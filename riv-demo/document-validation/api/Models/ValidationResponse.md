# ValidationResponse
## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **validityOutcome** | **Boolean** | Whether the document is determined to be valid (True) or invalid (False) | [default to null] |
| **validityScore** | **Double** | A score corresponding to the level of confidence that the document  is valid ranging between 0 and 1 (optional) | [optional] [default to null] |
| **validityProperties** | [**List**](ValidationResponse_validityProperties_inner.md) | Key value pairs describing document properties and their relationship to the validity decision. There are no strictly defined  properties. The inclusion of descriptive properties is encouraged to provide more context. (optional) | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

