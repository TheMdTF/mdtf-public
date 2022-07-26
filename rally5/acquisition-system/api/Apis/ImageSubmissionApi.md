# ImageSubmissionApi

All URIs are relative to *http://api.mdtf.org*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**v1FaceCapturesPost**](ImageSubmissionApi.md#v1FaceCapturesPost) | **POST** /v1/face-captures | Associate a face image capture with the ongoing transaction. |


<a name="v1FaceCapturesPost"></a>
# **v1FaceCapturesPost**
> v1FaceCapturesPost(FaceCapture)

Associate a face image capture with the ongoing transaction.

    Acquisition System Providers may submit face image data through this endpoint. Submitted image captures will be stored and associated with the ongoing transaction. Multiple images may be submitted during a transaction. 

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **FaceCapture** | [**FaceCapture**](../Models/FaceCapture.md)| Face image collected as part of a transaction. | |

### Return type

null (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json
