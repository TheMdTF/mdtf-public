# Documentation for The Maryland Test Facility Passive Presentation Attack Detection System Interface

<a name="documentation-for-api-endpoints"></a>
## Documentation for API Endpoints

All URIs are relative to *https://pad-demo.mdtf.org*

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
