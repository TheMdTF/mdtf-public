# Documentation for The Maryland Test Facility Document Validation Interface

<a name="documentation-for-api-endpoints"></a>
## Documentation for API Endpoints

All URIs are relative to *http://localhost*

| Class | Method | HTTP request | Description |
|------------ | ------------- | ------------- | -------------|
| *InfoApi* | [**info**](Apis/InfoApi.md#info) | **GET** /v1/info | Return information about the document validation algorithm |
| *ValidateApi* | [**validateDocument**](Apis/ValidateApi.md#validatedocument) | **POST** /v1/validate | Validate an identity document based on images |


<a name="documentation-for-models"></a>
## Documentation for Models

 - [CapturedDocument](./Models/CapturedDocument.md)
 - [Info](./Models/Info.md)
 - [ValidationErrorResponse](./Models/ValidationErrorResponse.md)
 - [ValidationResponse](./Models/ValidationResponse.md)
 - [ValidityProperty](./Models/ValidityProperty.md)


<a name="documentation-for-authorization"></a>
## Documentation for Authorization

All endpoints do not require authorization.
