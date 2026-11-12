# The Maryland Test Facility Matching System Interface

> Version 1.1.1

Application Programming Interface for receiving matching system requests from the MdTF Backend.

- License: [IDSL API License](https://raw.githubusercontent.com/TheMdTF/mdtf-public/master/LICENSE.md)
- Contact: The MdTF <rally@mdtf.org> https://mdtf.org
- Server: http://localhost:8080/


## Path Table

| Method | Path | Description |
| --- | --- | --- |
| POST | [/v1/create-template](#postv1create-template) | Generate a template from the provided biometric image. |
| POST | [/v1/compare-list](#postv1compare-list) | Compare a single probe template to a list of target templates. |
| GET | [/v1/info](#getv1info) | Returns basic information for the algorithm. |

## Reference Table

| Name | Path | Description |
| --- | --- | --- |
| Image | [#/components/schemas/Image](#componentsschemasimage) | Data transfer object for an image. |
| Template | [#/components/schemas/Template](#componentsschemastemplate) | Data transfer object for a template. |
| Comparison | [#/components/schemas/Comparison](#componentsschemascomparison) | A similarity score for a single (1:1) template comparison operation. |
| CompareListRequest | [#/components/schemas/CompareListRequest](#componentsschemascomparelistrequest) | A single probe template object and a list of target templates that it will be compared to. |
| Info | [#/components/schemas/Info](#componentsschemasinfo) | Basic information describing the algorithm. |
| RequestError | [#/components/schemas/RequestError](#componentsschemasrequesterror) | Relevant and concise diagnostic information in the event of a client side error (e.g. malformed requests, or invalid<br>image encoding). |
| ServerError | [#/components/schemas/ServerError](#componentsschemasservererror) | Relevant and concise diagnostic information in the event of a server side error<br>(e.g. invalid license, or failure to generate a template). |
| image | [#/components/requestBodies/image](#componentsrequestbodiesimage) | The biometric image that is being submitted for feature extraction. |
| compare_list_request | [#/components/requestBodies/compare_list_request](#componentsrequestbodiescompare_list_request) | A single template object and a list of templates that it will be compared to. |
| infoResponse | [#/components/responses/infoResponse](#componentsresponsesinforesponse) | info response object |
| templateResponse | [#/components/responses/templateResponse](#componentsresponsestemplateresponse) | Successful Response |
| compareListResponse | [#/components/responses/compareListResponse](#componentsresponsescomparelistresponse) | Successful Response |
| BadRequest | [#/components/responses/BadRequest](#componentsresponsesbadrequest) | Bad Request |
| InternalServerError | [#/components/responses/InternalServerError](#componentsresponsesinternalservererror) | Internal Server Error |

## Path Details

***

### [POST]/v1/create-template

- Summary  
  Generate a template from the provided biometric image.

- Operation id  
  GenerateTemplate

- Description  
  This endpoint accepts a base64 encoded PNG and attempts to perform a feature extraction operation producing a  
  single template.

- Tags  
  Biometric Operations

#### RequestBody

*Required*

The biometric image that is being submitted for feature extraction.

- application/json

```typescript
// Data transfer object for an image.
// 
{
  // The captured image data in PNG format, encoded as a base64 string.
  ImageData: string
}
```

#### Responses

- 200 Successful Response

`application/json`

```typescript
// Data transfer object for a template.
// 
{
  // The template data, encoded as a base64 string.  The data string shall not exceed 1 MB.
  Template: string
}
```

- 400 Bad Request

`application/json`

```typescript
// Relevant and concise diagnostic information in the event of a client side error (e.g. malformed requests, or invalid
// image encoding).
// 
string
```

- 500 Internal Server Error

`application/json`

```typescript
// Relevant and concise diagnostic information in the event of a server side error
// (e.g. invalid license, or failure to generate a template).
// 
string
```

***

### [POST]/v1/compare-list

- Summary  
  Compare a single probe template to a list of target templates.

- Operation id  
  CompareTemplateList

- Description  
  This endpoint accepts a template and a list of templates. It compares the probe template to every target template in the  
  provided list. The result is a list of Comparison objects that holds a similarity score for each comparison.   
  The returned list of comparisons MUST contain the same number of elements AND be in the same order as the provided list  
  of templates.

- Tags  
  Biometric Operations

#### RequestBody

*Required*

A single template object and a list of templates that it will be compared to.

- application/json

```typescript
// A single probe template object and a list of target templates that it will be compared to.
{
  // Data transfer object for a template.
  // 
  ProbeTemplate: {
    // The template data, encoded as a base64 string.  The data string shall not exceed 1 MB.
    Template: string
  }
  // Data transfer object for a template.
  // 
  TargetTemplateList: {
    // The template data, encoded as a base64 string.  The data string shall not exceed 1 MB.
    Template: string
  }[]
}
```

#### Responses

- 200 Successful Response

`application/json`

```typescript
// An array of comparison results. This list MUST contain the same number of elements AND be in
// the same order as the provided list of templates.
// 
// A similarity score for a single (1:1) template comparison operation.
// 
{
  // A similarity score, as produced by the algorithm.
  Score: number //format: float
}[]
```

- 400 Bad Request

`application/json`

```typescript
// Relevant and concise diagnostic information in the event of a client side error (e.g. malformed requests, or invalid
// image encoding).
// 
string
```

- 500 Internal Server Error

`application/json`

```typescript
// Relevant and concise diagnostic information in the event of a server side error
// (e.g. invalid license, or failure to generate a template).
// 
string
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
// Basic information describing the algorithm.
{
  // Name of algorithm.
  AlgorithmName: string
  // Algorithm version identifier.
  AlgorithmVersion: string
  // A string enum describing the type of biometric images the algorithm is meant to process.
  AlgorithmModality: enum[Face, Finger, Iris, Palm]
  // Name of the Company which produced the algorithm.
  CompanyName: string
  // The email address of an engineer or other technical resource to contact in the event of
  // an error running your service.
  // 
  TechnicalContactEmail: string
  // The recommended allocation of CPUs for the deployed docker container.
  RecommendedCPUs: number //format: float //maximum: 8
  // The recommended allocation of memory (MB) for the deployed docker container.
  RecommendedMem: integer //maximum: 8192
  // A string enum describing which collection event the algorithm is being submitted for.
  Test: enum[MDTF_2022_RALLY]
  // A map of preset False Match Rates (FMR) to vendor-provided threshold values. Score values returned from calls to
  // v1/compare-list indicate a matching determination by the algorithm if they are greater than the provided threshold
  // value at the respective FMR.
  // Note that threshold values are floats stored as strings and cannot exceed a length of 10 characters.
  // There are 5 required thresholds.
  // 
  Thresholds: {
    1:500: string //maxLength: 10
    1:1e3: string //maxLength: 10
    1:1e4: string //maxLength: 10
    1:1e5: string //maxLength: 10
    1:1e6: string //maxLength: 10
  }
}
```

- 500 Internal Server Error

`application/json`

```typescript
// Relevant and concise diagnostic information in the event of a server side error
// (e.g. invalid license, or failure to generate a template).
// 
string
```

## References

### #/components/schemas/Image

```typescript
// Data transfer object for an image.
// 
{
  // The captured image data in PNG format, encoded as a base64 string.
  ImageData: string
}
```

### #/components/schemas/Template

```typescript
// Data transfer object for a template.
// 
{
  // The template data, encoded as a base64 string.  The data string shall not exceed 1 MB.
  Template: string
}
```

### #/components/schemas/Comparison

```typescript
// A similarity score for a single (1:1) template comparison operation.
// 
{
  // A similarity score, as produced by the algorithm.
  Score: number //format: float
}
```

### #/components/schemas/CompareListRequest

```typescript
// A single probe template object and a list of target templates that it will be compared to.
{
  // Data transfer object for a template.
  // 
  ProbeTemplate: {
    // The template data, encoded as a base64 string.  The data string shall not exceed 1 MB.
    Template: string
  }
  // Data transfer object for a template.
  // 
  TargetTemplateList: {
    // The template data, encoded as a base64 string.  The data string shall not exceed 1 MB.
    Template: string
  }[]
}
```

### #/components/schemas/Info

```typescript
// Basic information describing the algorithm.
{
  // Name of algorithm.
  AlgorithmName: string
  // Algorithm version identifier.
  AlgorithmVersion: string
  // A string enum describing the type of biometric images the algorithm is meant to process.
  AlgorithmModality: enum[Face, Finger, Iris, Palm]
  // Name of the Company which produced the algorithm.
  CompanyName: string
  // The email address of an engineer or other technical resource to contact in the event of
  // an error running your service.
  // 
  TechnicalContactEmail: string
  // The recommended allocation of CPUs for the deployed docker container.
  RecommendedCPUs: number //format: float //maximum: 8
  // The recommended allocation of memory (MB) for the deployed docker container.
  RecommendedMem: integer //maximum: 8192
  // A string enum describing which collection event the algorithm is being submitted for.
  Test: enum[MDTF_2022_RALLY]
  // A map of preset False Match Rates (FMR) to vendor-provided threshold values. Score values returned from calls to
  // v1/compare-list indicate a matching determination by the algorithm if they are greater than the provided threshold
  // value at the respective FMR.
  // Note that threshold values are floats stored as strings and cannot exceed a length of 10 characters.
  // There are 5 required thresholds.
  // 
  Thresholds: {
    1:500: string //maxLength: 10
    1:1e3: string //maxLength: 10
    1:1e4: string //maxLength: 10
    1:1e5: string //maxLength: 10
    1:1e6: string //maxLength: 10
  }
}
```

### #/components/schemas/RequestError

```typescript
// Relevant and concise diagnostic information in the event of a client side error (e.g. malformed requests, or invalid
// image encoding).
// 
string
```

### #/components/schemas/ServerError

```typescript
// Relevant and concise diagnostic information in the event of a server side error
// (e.g. invalid license, or failure to generate a template).
// 
string
```

### #/components/requestBodies/image

- application/json

```typescript
// Data transfer object for an image.
// 
{
  // The captured image data in PNG format, encoded as a base64 string.
  ImageData: string
}
```

### #/components/requestBodies/compare_list_request

- application/json

```typescript
// A single probe template object and a list of target templates that it will be compared to.
{
  // Data transfer object for a template.
  // 
  ProbeTemplate: {
    // The template data, encoded as a base64 string.  The data string shall not exceed 1 MB.
    Template: string
  }
  // Data transfer object for a template.
  // 
  TargetTemplateList: {
    // The template data, encoded as a base64 string.  The data string shall not exceed 1 MB.
    Template: string
  }[]
}
```

### #/components/responses/infoResponse

- application/json

```typescript
// Basic information describing the algorithm.
{
  // Name of algorithm.
  AlgorithmName: string
  // Algorithm version identifier.
  AlgorithmVersion: string
  // A string enum describing the type of biometric images the algorithm is meant to process.
  AlgorithmModality: enum[Face, Finger, Iris, Palm]
  // Name of the Company which produced the algorithm.
  CompanyName: string
  // The email address of an engineer or other technical resource to contact in the event of
  // an error running your service.
  // 
  TechnicalContactEmail: string
  // The recommended allocation of CPUs for the deployed docker container.
  RecommendedCPUs: number //format: float //maximum: 8
  // The recommended allocation of memory (MB) for the deployed docker container.
  RecommendedMem: integer //maximum: 8192
  // A string enum describing which collection event the algorithm is being submitted for.
  Test: enum[MDTF_2022_RALLY]
  // A map of preset False Match Rates (FMR) to vendor-provided threshold values. Score values returned from calls to
  // v1/compare-list indicate a matching determination by the algorithm if they are greater than the provided threshold
  // value at the respective FMR.
  // Note that threshold values are floats stored as strings and cannot exceed a length of 10 characters.
  // There are 5 required thresholds.
  // 
  Thresholds: {
    1:500: string //maxLength: 10
    1:1e3: string //maxLength: 10
    1:1e4: string //maxLength: 10
    1:1e5: string //maxLength: 10
    1:1e6: string //maxLength: 10
  }
}
```

### #/components/responses/templateResponse

- application/json

```typescript
// Data transfer object for a template.
// 
{
  // The template data, encoded as a base64 string.  The data string shall not exceed 1 MB.
  Template: string
}
```

### #/components/responses/compareListResponse

- application/json

```typescript
// An array of comparison results. This list MUST contain the same number of elements AND be in
// the same order as the provided list of templates.
// 
// A similarity score for a single (1:1) template comparison operation.
// 
{
  // A similarity score, as produced by the algorithm.
  Score: number //format: float
}[]
```

### #/components/responses/BadRequest

- application/json

```typescript
// Relevant and concise diagnostic information in the event of a client side error (e.g. malformed requests, or invalid
// image encoding).
// 
string
```

### #/components/responses/InternalServerError

- application/json

```typescript
// Relevant and concise diagnostic information in the event of a server side error
// (e.g. invalid license, or failure to generate a template).
// 
string
```