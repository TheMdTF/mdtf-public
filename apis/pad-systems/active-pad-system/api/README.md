# Documentation for The Maryland Test Facility Active Presentation Attack Detection (A-PAD) System Interface

<a name="documentation-for-api-endpoints"></a>

![Active PAD](PAD%20API%20Diagrams%20-%20Active.png)

## Call for Contributions

This API is currently available in DRAFT form.  We are seeking comments from parties interested in PAD testing, specifically on the following items:

* Any additional endpoints or PAD capabilities that would be needed to evaluate modern PAD systems
* Appropriate video formats (.mov, .mp4, etc.)
* Communication or classification of PAD techniques

## Documentation for API Endpoints

All URIs are relative to *https://api.mdtf.org*

| Class | Method | HTTP request | Description |
|------------ | ------------- | ------------- | -------------|
| *DataSubmissionApi* | [**v1CaptureDataWithPadPost**](Apis/DataSubmissionApi.md#v1capturedatawithpadpost) | **POST** /v1/capture-data-with-pad | Create a biometric data capture with associated PAD information. |


<a name="documentation-for-models"></a>
## Documentation for Models

 - [Error](./Models/Error.md)
 - [PADAnalysis](./Models/PADAnalysis.md)
 - [PADDataCapture](./Models/PADDataCapture.md)
 - [PADErrorResponse](./Models/PADErrorResponse.md)
 - [PADProperty](./Models/PADProperty.md)


<a name="documentation-for-authorization"></a>
## Documentation for Authorization

All endpoints do not require authorization.
