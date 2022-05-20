# ImageSubmissionApi

All URIs are relative to *http://api.mdtf.org*

Method | HTTP request | Description
------------- | ------------- | -------------
[**v1FaceCapturesPost**](ImageSubmissionApi.md#v1FaceCapturesPost) | **POST** /v1/face-captures | Associate a face image capture with the ongoing transaction
[**v1FingerCapturesPost**](ImageSubmissionApi.md#v1FingerCapturesPost) | **POST** /v1/finger-captures | Associate up to 10 finger captures with the ongoing transaction
[**v1IrisCapturesPost**](ImageSubmissionApi.md#v1IrisCapturesPost) | **POST** /v1/iris-captures | Associate an iris image capture (or pair of image captures) with the ongoing transaction


<a name="v1FaceCapturesPost"></a>
# **v1FaceCapturesPost**
> v1FaceCapturesPost(FaceCapture)

Associate a face image capture with the ongoing transaction

    Acquisition System Providers may submit face image data through this endpoint. Submitted image captures will be stored and associated with the ongoing transaction. Multiple images may be submitted during a transaction. 

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **FaceCapture** | [**FaceCapture**](../Models/FaceCapture.md)| face image collected as part of a transaction |

### Return type

null (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="v1FingerCapturesPost"></a>
# **v1FingerCapturesPost**
> v1FingerCapturesPost(FingerCapture)

Associate up to 10 finger captures with the ongoing transaction

    Acquisition System Providers may submit finger image data through this endpoint. Submitted image captures will be stored and associated with the ongoing transaction. Multiple images may be submitted during a transaction. 

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **FingerCapture** | [**FingerCapture**](../Models/FingerCapture.md)| finger image(s) collected as part of a transaction |

### Return type

null (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="v1IrisCapturesPost"></a>
# **v1IrisCapturesPost**
> v1IrisCapturesPost(IrisCapture)

Associate an iris image capture (or pair of image captures) with the ongoing transaction

    Acquisition System Providers may submit iris image data through this endpoint. Submitted image captures will be stored and associated with the ongoing transaction. Multiple images may be submitted during a transaction. 

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **IrisCapture** | [**IrisCapture**](../Models/IrisCapture.md)| iris image(s) collected as part of a transaction |

### Return type

null (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

