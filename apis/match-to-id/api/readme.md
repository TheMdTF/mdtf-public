# The Maryland Test Facility Match-to-ID Interface

> Version 0.0.1

This document specifies the API requirements for MdTF testing of algorithms that match facial biometric samples to identity document images (``match-to-id'').  Match-to-ID testing at the MdTF is supported by the Department of Homeland Security, Science and Technology Directorate (DHS S&T) as part of the Remote Identity Validation Technology Demonstration (RIVTD).  For more information please visit [https://mdtf.org](https://mdtf.org) and [https://www.dhs.gov/science-and-technology/BI-TC](https://www.dhs.gov/science-and-technology/BI-TC)

- License: [IDSL API License](https://raw.githubusercontent.com/TheMdTF/mdtf-public/master/LICENSE.md)
- Contact: The MdTF <info@mdtf.org> https://mdtf.org


## Path Table

| Method | Path | Description |
| --- | --- | --- |
| POST | [/v1/create-template](#postv1create-template) | Generate a template from the provided facial biometric sample or the identity document image. |
| POST | [/v1/compare-list](#postv1compare-list) | Compare a single template to a list of target templates. |
| GET | [/v1/info](#getv1info) | Returns basic information for the algorithm. |

## Reference Table

| Name | Path | Description |
| --- | --- | --- |
| Image | [#/components/schemas/Image](#componentsschemasimage) | Data transfer object for an image. |
| Template | [#/components/schemas/Template](#componentsschemastemplate) | Data transfer object for a template. |
| Comparison | [#/components/schemas/Comparison](#componentsschemascomparison) | A similarity score for a template comparison operation. |
| CompareListRequest | [#/components/schemas/CompareListRequest](#componentsschemascomparelistrequest) | A single probe template object and a list of target templates that it will be compared to. |
| Info | [#/components/schemas/Info](#componentsschemasinfo) | Basic information describing the algorithm. |
| RequestError | [#/components/schemas/RequestError](#componentsschemasrequesterror) | Relevant and concise diagnostic information in the event of a client side error (e.g. malformed requests, or invalid<br>image encoding). |
| ServerError | [#/components/schemas/ServerError](#componentsschemasservererror) | Relevant and concise diagnostic information in the event of a server side error<br>(e.g. invalid license, or failure to generate a template). |
| image | [#/components/requestBodies/image](#componentsrequestbodiesimage) | The facial biometric sample or identity document that is being submitted.  Note there is no request body information<br>to distinguish between the submission of these two types (facial sample vs. ID document).  Algorithm providers must <br>make this distinction using internal logic. |
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
  Generate a template from the provided facial biometric sample or the identity document image.

- Operation id  
  GenerateTemplate

- Description  
  This endpoint accepts a base64 encoded PNG or JPG and attempts to perform a feature extraction operation producing a  
  single template.

- Tags  
  Biometric Operations

#### RequestBody

*Required*

The facial biometric sample or identity document that is being submitted.  Note there is no request body information
to distinguish between the submission of these two types (facial sample vs. ID document).  Algorithm providers must 
make this distinction using internal logic.

- application/json

```typescript
// Data transfer object for an image.
// 
{
  // The captured image data in PNG or JPG format, encoded as a base64 string.
  ImageData: string //format: byte
}
```

#### Responses

- 200 Successful Response

`application/json`

```typescript
// Data transfer object for a template.
// 
{
  // The template data, encoded as a base64 string.
  Template: string //format: byte
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
  Compare a single template to a list of target templates.

- Operation id  
  CompareTemplateList

- Description  
  This endpoint accepts a template and a list of templates. It compares the single template to every target template in the  
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
  SingleTemplate: {
    // The template data, encoded as a base64 string.
    Template: string //format: byte
  }
  // Data transfer object for a template.
  // 
  TargetTemplateList: {
    // The template data, encoded as a base64 string.
    Template: string //format: byte
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
// A similarity score for a template comparison operation.
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
  AlgorithmModality: enum[Face]
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
  Test: enum[MDTF_RIVTD_TRACK2]
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
  // The captured image data in PNG or JPG format, encoded as a base64 string.
  ImageData: string //format: byte
}
```

### #/components/schemas/Template

```typescript
// Data transfer object for a template.
// 
{
  // The template data, encoded as a base64 string.
  Template: string //format: byte
}
```

### #/components/schemas/Comparison

```typescript
// A similarity score for a template comparison operation.
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
  SingleTemplate: {
    // The template data, encoded as a base64 string.
    Template: string //format: byte
  }
  // Data transfer object for a template.
  // 
  TargetTemplateList: {
    // The template data, encoded as a base64 string.
    Template: string //format: byte
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
  AlgorithmModality: enum[Face]
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
  Test: enum[MDTF_RIVTD_TRACK2]
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
  // The captured image data in PNG or JPG format, encoded as a base64 string.
  ImageData: string //format: byte
}
```

### #/components/requestBodies/compare_list_request

- application/json

```typescript
// A single probe template object and a list of target templates that it will be compared to.
{
  // Data transfer object for a template.
  // 
  SingleTemplate: {
    // The template data, encoded as a base64 string.
    Template: string //format: byte
  }
  // Data transfer object for a template.
  // 
  TargetTemplateList: {
    // The template data, encoded as a base64 string.
    Template: string //format: byte
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
  AlgorithmModality: enum[Face]
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
  Test: enum[MDTF_RIVTD_TRACK2]
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
  // The template data, encoded as a base64 string.
  Template: string //format: byte
}
```

### #/components/responses/compareListResponse

- application/json

```typescript
// An array of comparison results. This list MUST contain the same number of elements AND be in
// the same order as the provided list of templates.
// 
// A similarity score for a template comparison operation.
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