# Documentation for The Maryland Test Facility Passive Presentation Attack Detection (P-PAD) System Interface

<a name="documentation-for-api-endpoints"></a>

## Call for Contributions

This API is currently available in DRAFT form.  We are seeking comments from parties interested in PAD testing, specifically on the following items:

* Any additional endpoints or PAD capabilities that would be needed to evaluate modern PAD systems
* Appropriate video formats (.mov, .mp4, etc.)
* Communication or classification of PAD techniques

## Documentation for API Endpoints

All URIs are relative to *https://api.mdtf.org*

| Class | Method | HTTP request | Description |
|------------ | ------------- | ------------- | -------------|
| *AlgorithmInformationApi* | [**getInfo**](Apis/AlgorithmInformationApi.md#getinfo) | **GET** /v1/info | Returns basic information for the algorithm. |
| *DataAnalysisApi* | [**v1AnalyzeDataForPadPost**](Apis/DataAnalysisApi.md#v1analyzedataforpadpost) | **POST** /v1/analyze-data-for-pad | Analyze biometric capture data for a presentation attack. |


<a name="documentation-for-models"></a>
## Documentation for Models

 - [Error](./Models/Error.md)
 - [Info](./Models/Info.md)
 - [Info_Thresholds](./Models/Info_Thresholds.md)
 - [PADAnalysis](./Models/PADAnalysis.md)
 - [PADProperty](./Models/PADProperty.md)


<a name="documentation-for-authorization"></a>
## Documentation for Authorization

All endpoints do not require authorization.
