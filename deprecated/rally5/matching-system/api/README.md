# Documentation for The Maryland Test Facility Matching System Interface

<a name="documentation-for-api-endpoints"></a>
## Documentation for API Endpoints

All URIs are relative to *http://localhost:8080*

| Class | Method | HTTP request | Description |
|------------ | ------------- | ------------- | -------------|
| *AlgorithmInformationApi* | [**getInfo**](Apis/AlgorithmInformationApi.md#getinfo) | **GET** /v1/info | Returns basic information for the algorithm. |
| *BiometricOperationsApi* | [**compareTemplateList**](Apis/BiometricOperationsApi.md#comparetemplatelist) | **POST** /v1/compare-list | Compare a single probe template to a list of target templates. |
*BiometricOperationsApi* | [**generateTemplate**](Apis/BiometricOperationsApi.md#generatetemplate) | **POST** /v1/create-template | Generate a template from the provided biometric image. |


<a name="documentation-for-models"></a>
## Documentation for Models

 - [CompareListRequest](Models/CompareListRequest.md)
 - [Comparison](Models/Comparison.md)
 - [Image](Models/Image.md)
 - [Info](Models/Info.md)
 - [Info_Thresholds](Models/Info_Thresholds.md)
 - [Template](Models/Template.md)

<a name="documentation-for-authorization"></a>
## Documentation for Authorization

All endpoints do not require authorization.
