# The Maryland Test Facility Passive Presentation Attack Detection System Interface

> Version 2.0.1

This document describes an application programming interface for a PAD system in which the PAD subsystem is logically distinct from the Data Capture Subsystem.  This allows for offline processing of PAD data.  This supports both image and video biometric sample inputs.

- License: [IDSL API License](https://raw.githubusercontent.com/TheMdTF/mdtf-public/master/LICENSE.md)
- Contact: The MdTF <rivr@mdtf.org> https://mdtf.org
- Server: https://api.mdtf.org/


## Path Table

| Method | Path | Description |
| --- | --- | --- |
| POST | [/v1/analyze-data-for-pad](#postv1analyze-data-for-pad) | Analyze biometric capture data for a presentation attack. |
| GET | [/v1/info](#getv1info) | Returns basic information for the algorithm. |

## Reference Table

| Name | Path | Description |
| --- | --- | --- |
| BiometricSample | [#/components/schemas/BiometricSample](#componentsschemasbiometricsample) | The biometric sample, encoded as a base64 string.  This can be an image, encoded as a PNG or JPEG or a short (<10s) video, encoded as a MOV or MP4. |
| BiometricSampleRequest | [#/components/schemas/BiometricSampleRequest](#componentsschemasbiometricsamplerequest) |  |
| PADAnalysis | [#/components/schemas/PADAnalysis](#componentsschemaspadanalysis) | Data transfer object for presentation attack information. |
| PADProperty | [#/components/schemas/PADProperty](#componentsschemaspadproperty) | Key value pair object. |
| Info | [#/components/schemas/Info](#componentsschemasinfo) | Basic information describing the PAD algorithm. |
| Error | [#/components/schemas/Error](#componentsschemaserror) |  |
| biometric_sample | [#/components/requestBodies/biometric_sample](#componentsrequestbodiesbiometric_sample) | A biometric sample to be analyzed for evidence of a presentation attack. |
| 400 | [#/components/responses/400](#componentsresponses400) | The request is malformed. |
| 500 | [#/components/responses/500](#componentsresponses500) | This request failed because of a server side issue. |
| pad_analysis_response | [#/components/responses/pad_analysis_response](#componentsresponsespad_analysis_response) | The result of an analysis of a biometric sample for a presentation attack. |
| info_response | [#/components/responses/info_response](#componentsresponsesinfo_response) | info response object |

## Path Details

***

### [POST]/v1/analyze-data-for-pad

- Summary  
  Analyze biometric capture data for a presentation attack.

- Description  
  Passive PAD systems receive data through this endpoint.  This request cannot exceed 100 Megabytes in size.

- Tags  
  Data Analysis

#### RequestBody

*Required*

- application/json

```typescript
{
  // The biometric sample, encoded as a base64 string.  This can be an image, encoded as a PNG or JPEG or a short (<10s) video, encoded as a MOV or MP4.
  BiometricSample?: string
}
```

#### Responses

- 200 The result of an analysis of a biometric sample for a presentation attack.

`application/json`

```typescript
// Data transfer object for presentation attack information.
{
  // Whether a presentation attack was determined to be detected (True) or not detected (False).  The outcome should be calibrated to produce a BPCER of 0.01.
  PADOutcome: boolean
  // A score corresponding to the level of confidence that a presentation attack was detected ranging between 0 and 1.
  PADScore: number //format: double
  // Key value pair object.
  PADProperties: {
    // Name of the property.
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
  This endpoint returns some basic information about the algorithm.

- Tags  
  Algorithm Information

#### Responses

- 200 info response object

`application/json`

```typescript
// Basic information describing the PAD algorithm.
{
  // Name of algorithm.
  AlgorithmName?: string
  // Algorithm version identifier.
  AlgorithmVersion?: string
  // A string enum describing the type of biometric modality the algorithm is meant to process.
  AlgorithmModality?: enum[Face]
  // Name of the company which produced the algorithm.
  CompanyName?: string
  // The email address of an engineer or other technical resource to contact in the event of an error running your service.
  // 
  TechnicalContactEmail?: string
  // The recommended allocation of CPUs for the deployed docker container.
  RecommendedCPUs?: number //format: float //maximum: 8
  // The recommended allocation of memory (MB) for the deployed docker container.
  RecommendedMem?: integer //maximum: 8192
  // A map of specific Bona-fide Classification Error Rates (BPCERs) to vendor-provided threshold values.  PADScore values optionally returned from calls to
  // /v1/analyze-data-for-pad indicate a presentation attack is occuring is if they are greater than the provided threshold value at the respective BPCER.
  // Note 1: Threshold values are floats stored as strings and cannot exceed a length of 10 characters.
  // Note 2: There are 4 required thresholds.
  // 
  Thresholds: {
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
// The biometric sample, encoded as a base64 string.  This can be an image, encoded as a PNG or JPEG or a short (<10s) video, encoded as a MOV or MP4.
string
```

### #/components/schemas/BiometricSampleRequest

```typescript
{
  // The biometric sample, encoded as a base64 string.  This can be an image, encoded as a PNG or JPEG or a short (<10s) video, encoded as a MOV or MP4.
  BiometricSample?: string
}
```

### #/components/schemas/PADAnalysis

```typescript
// Data transfer object for presentation attack information.
{
  // Whether a presentation attack was determined to be detected (True) or not detected (False).  The outcome should be calibrated to produce a BPCER of 0.01.
  PADOutcome: boolean
  // A score corresponding to the level of confidence that a presentation attack was detected ranging between 0 and 1.
  PADScore: number //format: double
  // Key value pair object.
  PADProperties: {
    // Name of the property.
    Property?: string
    // Value for the property.
    Value?: string
  }[]
}
```

### #/components/schemas/PADProperty

```typescript
// Key value pair object.
{
  // Name of the property.
  Property?: string
  // Value for the property.
  Value?: string
}
```

### #/components/schemas/Info

```typescript
// Basic information describing the PAD algorithm.
{
  // Name of algorithm.
  AlgorithmName?: string
  // Algorithm version identifier.
  AlgorithmVersion?: string
  // A string enum describing the type of biometric modality the algorithm is meant to process.
  AlgorithmModality?: enum[Face]
  // Name of the company which produced the algorithm.
  CompanyName?: string
  // The email address of an engineer or other technical resource to contact in the event of an error running your service.
  // 
  TechnicalContactEmail?: string
  // The recommended allocation of CPUs for the deployed docker container.
  RecommendedCPUs?: number //format: float //maximum: 8
  // The recommended allocation of memory (MB) for the deployed docker container.
  RecommendedMem?: integer //maximum: 8192
  // A map of specific Bona-fide Classification Error Rates (BPCERs) to vendor-provided threshold values.  PADScore values optionally returned from calls to
  // /v1/analyze-data-for-pad indicate a presentation attack is occuring is if they are greater than the provided threshold value at the respective BPCER.
  // Note 1: Threshold values are floats stored as strings and cannot exceed a length of 10 characters.
  // Note 2: There are 4 required thresholds.
  // 
  Thresholds: {
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
  // The biometric sample, encoded as a base64 string.  This can be an image, encoded as a PNG or JPEG or a short (<10s) video, encoded as a MOV or MP4.
  BiometricSample?: string
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

### #/components/responses/pad_analysis_response

- application/json

```typescript
// Data transfer object for presentation attack information.
{
  // Whether a presentation attack was determined to be detected (True) or not detected (False).  The outcome should be calibrated to produce a BPCER of 0.01.
  PADOutcome: boolean
  // A score corresponding to the level of confidence that a presentation attack was detected ranging between 0 and 1.
  PADScore: number //format: double
  // Key value pair object.
  PADProperties: {
    // Name of the property.
    Property?: string
    // Value for the property.
    Value?: string
  }[]
}
```

### #/components/responses/info_response

- application/json

```typescript
// Basic information describing the PAD algorithm.
{
  // Name of algorithm.
  AlgorithmName?: string
  // Algorithm version identifier.
  AlgorithmVersion?: string
  // A string enum describing the type of biometric modality the algorithm is meant to process.
  AlgorithmModality?: enum[Face]
  // Name of the company which produced the algorithm.
  CompanyName?: string
  // The email address of an engineer or other technical resource to contact in the event of an error running your service.
  // 
  TechnicalContactEmail?: string
  // The recommended allocation of CPUs for the deployed docker container.
  RecommendedCPUs?: number //format: float //maximum: 8
  // The recommended allocation of memory (MB) for the deployed docker container.
  RecommendedMem?: integer //maximum: 8192
  // A map of specific Bona-fide Classification Error Rates (BPCERs) to vendor-provided threshold values.  PADScore values optionally returned from calls to
  // /v1/analyze-data-for-pad indicate a presentation attack is occuring is if they are greater than the provided threshold value at the respective BPCER.
  // Note 1: Threshold values are floats stored as strings and cannot exceed a length of 10 characters.
  // Note 2: There are 4 required thresholds.
  // 
  Thresholds: {
    1:ten: string //maxLength: 10
    1:1e2: string //maxLength: 10
    1:1e3: string //maxLength: 10
    1:1e4: string //maxLength: 10
  }
}
```