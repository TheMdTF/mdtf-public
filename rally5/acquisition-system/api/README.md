# Documentation for The Maryland Test Facility Acquisition System Interface

<a name="documentation-for-api-endpoints"></a>
## Documentation for API Endpoints

All URIs are relative to *http://api.mdtf.org*

| Class | Method | HTTP request | Description |
|------------ | ------------- | ------------- | -------------|
| *ImageSubmissionApi* | [**v1FaceCapturesPost**](Apis/ImageSubmissionApi.md#v1facecapturespost) | **POST** /v1/face-captures | Associate a face image capture with the ongoing transaction |
*ImageSubmissionApi* | [**v1FingerCapturesPost**](Apis/ImageSubmissionApi.md#v1fingercapturespost) | **POST** /v1/finger-captures | Associate up to 10 finger captures with the ongoing transaction |
*ImageSubmissionApi* | [**v1IrisCapturesPost**](Apis/ImageSubmissionApi.md#v1iriscapturespost) | **POST** /v1/iris-captures | Associate an iris image capture (or pair of image captures) with the ongoing transaction |


<a name="documentation-for-models"></a>
## Documentation for Models

 - [Error](./Models/Error.md)
 - [FaceCapture](./Models/FaceCapture.md)
 - [FingerCapture](./Models/FingerCapture.md)
 - [FingerImage](./Models/FingerImage.md)
 - [FingerType](./Models/FingerType.md)
 - [IrisCapture](./Models/IrisCapture.md)


<a name="documentation-for-authorization"></a>
## Documentation for Authorization

All endpoints do not require authorization.
