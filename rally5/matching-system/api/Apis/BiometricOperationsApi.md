# BiometricOperationsApi

All URIs are relative to *http://localhost:8080*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**compareTemplateList**](BiometricOperationsApi.md#compareTemplateList) | **POST** /v1/compare-list | Compare a single probe template to a list of target templates |
| [**generateTemplate**](BiometricOperationsApi.md#generateTemplate) | **POST** /v1/generate-template | Generate a template from the provided biometric image |


<a name="compareTemplateList"></a>
# **compareTemplateList**
> List compareTemplateList(CompareListRequest)

Compare a single probe template to a list of target templates

    This endpoint accepts a template and a list of templates. It compares the probe template to every target template in the provided list. The result is a list of Comparison objects that holds a similarity score for each comparison. &lt;br&gt;&lt;br&gt; The returned comparison list MUST contain the same number of elements AND be in the same order as the provided list of templates. 

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **CompareListRequest** | [**CompareListRequest**](../Models/CompareListRequest.md)| A single template object and a list of templates that it will be compared to. | |

### Return type

[**List**](../Models/Comparison.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="generateTemplate"></a>
# **generateTemplate**
> Template generateTemplate(Image)

Generate a template from the provided biometric image

    This endpoint accepts a base64 encoded PNG and attempts to perform a &#39;feature extraction&#39; operation producing a single template. 

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **Image** | [**Image**](../Models/Image.md)| The biometric image that is being submitted for feature extraction.  | |

### Return type

[**Template**](../Models/Template.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

