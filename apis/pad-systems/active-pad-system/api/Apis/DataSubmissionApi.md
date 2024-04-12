# DataSubmissionApi

All URIs are relative to *https://api.mdtf.org*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**v1CaptureDataWithPadPost**](DataSubmissionApi.md#v1CaptureDataWithPadPost) | **POST** /v1/capture-data-with-pad | Create a biometric data capture with associated PAD information. |


<a name="v1CaptureDataWithPadPost"></a>
# **v1CaptureDataWithPadPost**
> v1CaptureDataWithPadPost(PADDataCapture)

Create a biometric data capture with associated PAD information.

    Active PAD systems may submit data through this endpoint.  This request cannot exceed 100 Megabytes in size. 

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **PADDataCapture** | [**PADDataCapture**](../Models/PADDataCapture.md)| PAD data collected as part of a biometric data capture process. | |

### Return type

null (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

