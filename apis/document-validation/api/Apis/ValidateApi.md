# ValidateApi

All URIs are relative to *http://localhost*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**validateDocument**](ValidateApi.md#validateDocument) | **POST** /v1/validate | Validate an identity document based on images |


<a name="validateDocument"></a>
# **validateDocument**
> ValidationResponse validateDocument(CapturedDocument)

Validate an identity document based on images

    Receive base64 encoded PNG images of the front and back of a document. Validation requires that the document is judged to be authentic. See [NIST SP 800-63A](https://pages.nist.gov/800-63-3-Implementation-Resources/63A/ial2remote/) Section A.10.2: Identity Validation for recommended best practices.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **CapturedDocument** | [**CapturedDocument**](../Models/CapturedDocument.md)|  | |

### Return type

[**ValidationResponse**](../Models/ValidationResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

