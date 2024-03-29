# BiometricOperationsApi

All URIs are relative to *http://localhost*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**compareTemplateList**](BiometricOperationsApi.md#compareTemplateList) | **POST** /v1/compare-list | Compare a single template to a list of target templates. |
| [**generateTemplate**](BiometricOperationsApi.md#generateTemplate) | **POST** /v1/create-template | Generate a template from the provided facial biometric sample or the identity document image. |


<a name="compareTemplateList"></a>
# **compareTemplateList**
> List compareTemplateList(CompareListRequest)

Compare a single template to a list of target templates.

    This endpoint accepts a template and a list of templates. It compares the single template to every target template in the provided list. The result is a list of Comparison objects that holds a similarity score for each comparison.  The returned list of comparisons MUST contain the same number of elements AND be in the same order as the provided list of templates. 

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

Generate a template from the provided facial biometric sample or the identity document image.

    This endpoint accepts a base64 encoded PNG or JPG and attempts to perform a feature extraction operation producing a single template. 

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **Image** | [**Image**](../Models/Image.md)| The facial biometric sample or identity document that is being submitted.  Note there is no request body information to distinguish between the submission of these two types (facial sample vs. ID document).  Algorithm providers must  make this distinction using internal logic.  | |

### Return type

[**Template**](../Models/Template.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

