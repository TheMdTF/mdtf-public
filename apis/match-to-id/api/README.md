# Documentation for The Maryland Test Facility Match-to-ID Interface

<a name="documentation-for-api-endpoints"></a>
## Documentation for API Endpoints

All URIs are relative to *http://localhost*

| Class | Method | HTTP request | Description |
|------------ | ------------- | ------------- | -------------|
| *AlgorithmInformationApi* | [**getInfo**](Apis/AlgorithmInformationApi.md#getinfo) | **GET** /v1/info | Returns basic information for the algorithm. |
| *BiometricOperationsApi* | [**compareTemplateList**](Apis/BiometricOperationsApi.md#comparetemplatelist) | **POST** /v1/compare-list | Compare a single template to a list of target templates. |
*BiometricOperationsApi* | [**generateTemplate**](Apis/BiometricOperationsApi.md#generatetemplate) | **POST** /v1/create-template | Generate a template from the provided facial biometric sample or the identity document image. |


<a name="documentation-for-models"></a>
## Documentation for Models

 - [CompareListRequest](./Models/CompareListRequest.md)
 - [Comparison](./Models/Comparison.md)
 - [Image](./Models/Image.md)
 - [Info](./Models/Info.md)
 - [Info_Thresholds](./Models/Info_Thresholds.md)
 - [Template](./Models/Template.md)


<a name="documentation-for-authorization"></a>
## Documentation for Authorization

All endpoints do not require authorization.
