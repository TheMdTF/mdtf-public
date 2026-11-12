# The Maryland Test Facility Biometric Deepfake Detection System Interface

> Version 1.0.0

This document describes an application programming interface for a Biometric Deepfake Detection (BDD) system.  BDD systems determine if a biometric sample contains AI-generated or manipulated biometric characteristics.  This supports both image and video biometric sample inputs.

- License: [IDSL API License](https://raw.githubusercontent.com/TheMdTF/mdtf-public/master/LICENSE.md)
- Contact: The MdTF <rivr@mdtf.org> https://mdtf.org


## Path Table

| Method | Path | Description |
| --- | --- | --- |
| POST | [/v1/analyze-data-for-bdd](#postv1analyze-data-for-bdd) | Analyze biometric capture data for a biometric deepfake. |
| GET | [/v1/info](#getv1info) | Returns basic information for the algorithm. |

## Reference Table

| Name | Path | Description |
| --- | --- | --- |
| BiometricSample | [#/components/schemas/BiometricSample](#componentsschemasbiometricsample) | The biometric sample, encoded as a base64 string.  This can be an individual image, encoded as a PNG or JPEG, or a short (~10s) video, encoded as a MOV or MP4. |
| BiometricSampleRequest | [#/components/schemas/BiometricSampleRequest](#componentsschemasbiometricsamplerequest) |  |
| BDDAnalysis | [#/components/schemas/BDDAnalysis](#componentsschemasbddanalysis) | Data transfer object for biometric deepfake detection results. |
| BDDProperty | [#/components/schemas/BDDProperty](#componentsschemasbddproperty) | Key value pair object describing a detected biometric deepfake property. |
| Info | [#/components/schemas/Info](#componentsschemasinfo) | Basic information describing the BDD algorithm. |
| Error | [#/components/schemas/Error](#componentsschemaserror) |  |
| biometric_sample | [#/components/requestBodies/biometric_sample](#componentsrequestbodiesbiometric_sample) | A biometric sample to be analyzed for evidence of a biometric deepfake. |
| 400 | [#/components/responses/400](#componentsresponses400) | The request is malformed. |
| 500 | [#/components/responses/500](#componentsresponses500) | This request failed because of a server side issue. |
| bdd_analysis_response | [#/components/responses/bdd_analysis_response](#componentsresponsesbdd_analysis_response) | The result of an analysis of a biometric sample for a biometric deepfake.  Failure to respond with BDD analysis should set an appropriate status code indicating a failure and return an explanatory error response. |
| info_response | [#/components/responses/info_response](#componentsresponsesinfo_response) | info response object |

## Path Details

***

### [POST]/v1/analyze-data-for-bdd

- Summary  
  Analyze biometric capture data for a biometric deepfake.

- Operation id  
  AnalyzeData

- Description  
  BDD systems receive data through this endpoint.  This endpoint accepts individual, base64-encoded PNG or JPEG image bytes or a short (~10s) video, encoded as a MOV or MP4.

- Tags  
  Data Analysis

#### RequestBody

*Required*

A biometric sample to be analyzed for evidence of a biometric deepfake.

- application/json

```typescript
{
  // The biometric sample, encoded as a base64 string.  This can be an individual image, encoded as a PNG or JPEG, or a short (~10s) video, encoded as a MOV or MP4.
  BiometricSample?: string //format: byte
}
```

#### Responses

- 200 The result of an analysis of a biometric sample for a biometric deepfake.  Failure to respond with BDD analysis should set an appropriate status code indicating a failure and return an explanatory error response.

`application/json`

```typescript
// Data transfer object for biometric deepfake detection results.
{
  // Whether a biometric deepfake was determined to be detected (True) or not detected (False).
  BDDOutcome: boolean
  // A score corresponding to the level of confidence that a biometric deepfake was detected, ranging between 0 and 1.
  BDDScore: number //format: double
  // Key value pairs describing biometric deepfake properties and their relationship to the deepfake detection outcome/score.  Common properties include ExcessiveSmoothing, AudioVideoMismatch, EyeGazeAnomaly, FacialAsymmetry, SkinTextureAnomaly, BlinkingPatternAnomaly, LipSyncMismatch, and BackgroundInconsistency.  The inclusion of descriptive properties is encouraged to provide more context. Do not include image bytes in the properties. (optional)
  // Key value pair object describing a detected biometric deepfake property.
  BDDProperties: {
    // Name of the deepfake property.  Common properties include: ExcessiveSmoothing, AudioVideoMismatch, EyeGazeAnomaly, FacialAsymmetry, SkinTextureAnomaly, BlinkingPatternAnomaly, LipSyncMismatch, BackgroundInconsistency.
    // 
    Property?: string
    // Value for the property.
    Value?: string
  }[]
}
```

- 400 The request is malformed.

`application/json`

```typescript
{
  // The error code.
  Code?: integer
  // A description of the returned code.
  Message?: string
  // A detailed description of the error that occurred.
  ErrorString?: string
}
```

- 500 This request failed because of a server side issue.

`application/json`

```typescript
{
  // The error code.
  Code?: integer
  // A description of the returned code.
  Message?: string
  // A detailed description of the error that occurred.
  ErrorString?: string
}
```

***

### [GET]/v1/info

- Summary  
  Returns basic information for the algorithm.

- Operation id  
  GetInfo

- Description  
  This endpoint returns some basic information about the BDD algorithm, including provider-specific details such as name, version, modality, thresholds, contact information, recommended detection thresholds, CPUs, and memory.

- Tags  
  Algorithm Information

#### Responses

- 200 info response object

`application/json`

```typescript
// Basic information describing the BDD algorithm.
{
  // Name of algorithm.
  AlgorithmName: string
  // Algorithm version identifier.
  AlgorithmVersion: string
  // A string enum describing the type of biometric modality the algorithm is meant to process.
  AlgorithmModality: enum[Face]
  // Name of the company which produced the algorithm.
  CompanyName: string
  // The email address of an engineer or other technical resource to contact in the event of an error running your service.
  // 
  TechnicalContactEmail: string
  // The recommended allocation of CPUs for the deployed docker container.
  RecommendedCPUs: number //format: float //maximum: 8
  // The recommended allocation of memory (MB) for the deployed docker container.
  RecommendedMem: integer //maximum: 8192
  // The name of the test event.
  Event: enum[RIVR_BDD]
  // Whether the algorithm requires a GPU to run.
  RequiresGPU: boolean
  // Whether the algorithm supports GPU acceleration.
  SupportsGPUAcceleration: boolean
  // A map of specific False Reject Rates (FRRs) to vendor-provided recommended detection threshold values.  BDDScore values returned from calls to
  // /v1/analyze-data-for-bdd indicate a biometric deepfake is detected if they are greater than the provided threshold value at the respective BPCER.
  // Note 1: Threshold values are floats stored as strings and cannot exceed a length of 10 characters.
  // Note 2: There are 4 required thresholds.
  // 
  RecommendedDetectionThresholds: {
    1:ten: string //maxLength: 10
    1:1e2: string //maxLength: 10
    1:1e3: string //maxLength: 10
    1:1e4: string //maxLength: 10
  }
}
```

- 500 This request failed because of a server side issue.

`application/json`

```typescript
{
  // The error code.
  Code?: integer
  // A description of the returned code.
  Message?: string
  // A detailed description of the error that occurred.
  ErrorString?: string
}
```

## References

### #/components/schemas/BiometricSample

```typescript
// The biometric sample, encoded as a base64 string.  This can be an individual image, encoded as a PNG or JPEG, or a short (~10s) video, encoded as a MOV or MP4.
string //format: byte
```

### #/components/schemas/BiometricSampleRequest

```typescript
{
  // The biometric sample, encoded as a base64 string.  This can be an individual image, encoded as a PNG or JPEG, or a short (~10s) video, encoded as a MOV or MP4.
  BiometricSample?: string //format: byte
}
```

### #/components/schemas/BDDAnalysis

```typescript
// Data transfer object for biometric deepfake detection results.
{
  // Whether a biometric deepfake was determined to be detected (True) or not detected (False).
  BDDOutcome: boolean
  // A score corresponding to the level of confidence that a biometric deepfake was detected, ranging between 0 and 1.
  BDDScore: number //format: double
  // Key value pairs describing biometric deepfake properties and their relationship to the deepfake detection outcome/score.  Common properties include ExcessiveSmoothing, AudioVideoMismatch, EyeGazeAnomaly, FacialAsymmetry, SkinTextureAnomaly, BlinkingPatternAnomaly, LipSyncMismatch, and BackgroundInconsistency.  The inclusion of descriptive properties is encouraged to provide more context. Do not include image bytes in the properties. (optional)
  // Key value pair object describing a detected biometric deepfake property.
  BDDProperties: {
    // Name of the deepfake property.  Common properties include: ExcessiveSmoothing, AudioVideoMismatch, EyeGazeAnomaly, FacialAsymmetry, SkinTextureAnomaly, BlinkingPatternAnomaly, LipSyncMismatch, BackgroundInconsistency.
    // 
    Property?: string
    // Value for the property.
    Value?: string
  }[]
}
```

### #/components/schemas/BDDProperty

```typescript
// Key value pair object describing a detected biometric deepfake property.
{
  // Name of the deepfake property.  Common properties include: ExcessiveSmoothing, AudioVideoMismatch, EyeGazeAnomaly, FacialAsymmetry, SkinTextureAnomaly, BlinkingPatternAnomaly, LipSyncMismatch, BackgroundInconsistency.
  // 
  Property?: string
  // Value for the property.
  Value?: string
}
```

### #/components/schemas/Info

```typescript
// Basic information describing the BDD algorithm.
{
  // Name of algorithm.
  AlgorithmName: string
  // Algorithm version identifier.
  AlgorithmVersion: string
  // A string enum describing the type of biometric modality the algorithm is meant to process.
  AlgorithmModality: enum[Face]
  // Name of the company which produced the algorithm.
  CompanyName: string
  // The email address of an engineer or other technical resource to contact in the event of an error running your service.
  // 
  TechnicalContactEmail: string
  // The recommended allocation of CPUs for the deployed docker container.
  RecommendedCPUs: number //format: float //maximum: 8
  // The recommended allocation of memory (MB) for the deployed docker container.
  RecommendedMem: integer //maximum: 8192
  // The name of the test event.
  Event: enum[RIVR_BDD]
  // Whether the algorithm requires a GPU to run.
  RequiresGPU: boolean
  // Whether the algorithm supports GPU acceleration.
  SupportsGPUAcceleration: boolean
  // A map of specific False Reject Rates (FRRs) to vendor-provided recommended detection threshold values.  BDDScore values returned from calls to
  // /v1/analyze-data-for-bdd indicate a biometric deepfake is detected if they are greater than the provided threshold value at the respective BPCER.
  // Note 1: Threshold values are floats stored as strings and cannot exceed a length of 10 characters.
  // Note 2: There are 4 required thresholds.
  // 
  RecommendedDetectionThresholds: {
    1:ten: string //maxLength: 10
    1:1e2: string //maxLength: 10
    1:1e3: string //maxLength: 10
    1:1e4: string //maxLength: 10
  }
}
```

### #/components/schemas/Error

```typescript
{
  // The error code.
  Code?: integer
  // A description of the returned code.
  Message?: string
  // A detailed description of the error that occurred.
  ErrorString?: string
}
```

### #/components/requestBodies/biometric_sample

- application/json

```typescript
{
  // The biometric sample, encoded as a base64 string.  This can be an individual image, encoded as a PNG or JPEG, or a short (~10s) video, encoded as a MOV or MP4.
  BiometricSample?: string //format: byte
}
```

### #/components/responses/400

- application/json

```typescript
{
  // The error code.
  Code?: integer
  // A description of the returned code.
  Message?: string
  // A detailed description of the error that occurred.
  ErrorString?: string
}
```

### #/components/responses/500

- application/json

```typescript
{
  // The error code.
  Code?: integer
  // A description of the returned code.
  Message?: string
  // A detailed description of the error that occurred.
  ErrorString?: string
}
```

### #/components/responses/bdd_analysis_response

- application/json

```typescript
// Data transfer object for biometric deepfake detection results.
{
  // Whether a biometric deepfake was determined to be detected (True) or not detected (False).
  BDDOutcome: boolean
  // A score corresponding to the level of confidence that a biometric deepfake was detected, ranging between 0 and 1.
  BDDScore: number //format: double
  // Key value pairs describing biometric deepfake properties and their relationship to the deepfake detection outcome/score.  Common properties include ExcessiveSmoothing, AudioVideoMismatch, EyeGazeAnomaly, FacialAsymmetry, SkinTextureAnomaly, BlinkingPatternAnomaly, LipSyncMismatch, and BackgroundInconsistency.  The inclusion of descriptive properties is encouraged to provide more context. Do not include image bytes in the properties. (optional)
  // Key value pair object describing a detected biometric deepfake property.
  BDDProperties: {
    // Name of the deepfake property.  Common properties include: ExcessiveSmoothing, AudioVideoMismatch, EyeGazeAnomaly, FacialAsymmetry, SkinTextureAnomaly, BlinkingPatternAnomaly, LipSyncMismatch, BackgroundInconsistency.
    // 
    Property?: string
    // Value for the property.
    Value?: string
  }[]
}
```

### #/components/responses/info_response

- application/json

```typescript
// Basic information describing the BDD algorithm.
{
  // Name of algorithm.
  AlgorithmName: string
  // Algorithm version identifier.
  AlgorithmVersion: string
  // A string enum describing the type of biometric modality the algorithm is meant to process.
  AlgorithmModality: enum[Face]
  // Name of the company which produced the algorithm.
  CompanyName: string
  // The email address of an engineer or other technical resource to contact in the event of an error running your service.
  // 
  TechnicalContactEmail: string
  // The recommended allocation of CPUs for the deployed docker container.
  RecommendedCPUs: number //format: float //maximum: 8
  // The recommended allocation of memory (MB) for the deployed docker container.
  RecommendedMem: integer //maximum: 8192
  // The name of the test event.
  Event: enum[RIVR_BDD]
  // Whether the algorithm requires a GPU to run.
  RequiresGPU: boolean
  // Whether the algorithm supports GPU acceleration.
  SupportsGPUAcceleration: boolean
  // A map of specific False Reject Rates (FRRs) to vendor-provided recommended detection threshold values.  BDDScore values returned from calls to
  // /v1/analyze-data-for-bdd indicate a biometric deepfake is detected if they are greater than the provided threshold value at the respective BPCER.
  // Note 1: Threshold values are floats stored as strings and cannot exceed a length of 10 characters.
  // Note 2: There are 4 required thresholds.
  // 
  RecommendedDetectionThresholds: {
    1:ten: string //maxLength: 10
    1:1e2: string //maxLength: 10
    1:1e3: string //maxLength: 10
    1:1e4: string //maxLength: 10
  }
}
```