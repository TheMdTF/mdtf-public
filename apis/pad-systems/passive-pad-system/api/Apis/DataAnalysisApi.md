# DataAnalysisApi

All URIs are relative to *https://api.mdtf.org*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**v1AnalyzeDataForPadPost**](DataAnalysisApi.md#v1AnalyzeDataForPadPost) | **POST** /v1/analyze-data-for-pad | Analyze biometric capture data for a presentation attack. |


<a name="v1AnalyzeDataForPadPost"></a>
# **v1AnalyzeDataForPadPost**
> PADAnalysis v1AnalyzeDataForPadPost(body)

Analyze biometric capture data for a presentation attack.

    Passive PAD systems receive data through this endpoint 

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **body** | **String**| A biometric sample to be analyzed for evidence of a presentation attack. | |

### Return type

[**PADAnalysis**](../Models/PADAnalysis.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

