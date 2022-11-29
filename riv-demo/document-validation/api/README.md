# Documentation for The Maryland Test Facility Document Validation Interface - OpenAPI 3.0

<a name="documentation-for-api-endpoints"></a>
## Documentation for API Endpoints

All URIs are relative to *http://localhost*

| Class | Method | HTTP request | Description |
|------------ | ------------- | ------------- | -------------|
| *InfoApi* | [**info**](Apis/InfoApi.md#info) | **GET** /v1/info | Return information about the document validation algorithm |
| *ValidateApi* | [**v1ValidatePost**](Apis/ValidateApi.md#v1validatepost) | **POST** /v1/validate | Validate an identity document based on images |


<a name="documentation-for-models"></a>
## Documentation for Models

 - [CapturedDocument](./Models/CapturedDocument.md)
 - [Info](./Models/Info.md)
 - [ValidationErrorResponse](./Models/ValidationErrorResponse.md)
 - [ValidationItem](./Models/ValidationItem.md)
 - [ValidationResponse](./Models/ValidationResponse.md)
 - [ValidationResponse_vector_inner](./Models/ValidationResponse_vector_inner.md)


<a name="documentation-for-authorization"></a>
## Documentation for Authorization

All endpoints do not require authorization.
