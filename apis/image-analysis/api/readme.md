# The Maryland Test Facility Image Analysis Interface

> Version 1.0.0

Application programming interface for receiving image analysis requests.

- License: [IDSL API License](https://raw.githubusercontent.com/TheMdTF/mdtf-public/master/LICENSE.md)
- Contact: MdTF <info@mdtf.org> https://mdtf.org
- Server: http://172.17.0.2:8080


## Path Table

| Method | Path | Description |
| --- | --- | --- |
| GET | [/v1/info](#getv1info) | Returns basic information for the algorithm. |
| POST | [/v1/analyze-image](#postv1analyze-image) | Returns an object containing the results of analysis for the provided biometric image. |

## Reference Table

| Name | Path | Description |
| --- | --- | --- |
| Image | [#/components/schemas/Image](#componentsschemasimage) | Data transfer object for an image. |
| Info | [#/components/schemas/Info](#componentsschemasinfo) | Basic information describing the algorithm. |
| AnalysisResult | [#/components/schemas/AnalysisResult](#componentsschemasanalysisresult) | Object containing the results of a generic image analysis algorithm on a single image. Must include at least 1 and not more than 8 string properties in a JSON object with a depth of 1. The full contents shall not exceed 512 Bytes. |
| RequestError | [#/components/schemas/RequestError](#componentsschemasrequesterror) | Relevant and concise diagnostic information in the event of a client side error (e.g. malformed requests, or invalid image encoding). |
| ServerError | [#/components/schemas/ServerError](#componentsschemasservererror) | Relevant and concise diagnostic information in the event of a server side error (e.g. invalid license, or failure to generate a template). |

## Path Details

***

### [GET]/v1/info

- Summary  
  Returns basic information for the algorithm.

- Operation id  
  info

- Description  
  This endpoint returns some basic information about the image analysis algorithm.

- Tags  
  Algorithm Information

#### Responses

- 200 Successful Response

`application/json`

```typescript
// Basic information describing the algorithm.
{
  // Name of algorithm
  AlgorithmName: string
  // A string enum describing the type of biometric images the algorithm is meant to process
  AlgorithmType: enum[Face, Finger, Iris]
  // Algorithm version identifier
  AlgorithmVersion: string
  // Name of the company which produces the algorithm
  CompanyName: string
  // A string enum to select an MdTF dataset of biometric images for this submission to analyze
  ImageDataset: enum[NIST_MEDS]
  // The recommended allocation of CPUs for the deployed docker container.
  RecommendedCPUs: number //format: integer
  // The recommended allocation of memory (MB) for the deployed docker container.
  RecommendedMem: number //format: integer
  // The email address of an engineer or other technical resource to contact in the event of an error running your service. This field may be left blank if desired.
  TechnicalContactEmail: string
}
```

- 500 Internal Server Error

`application/json`

```typescript
// Relevant and concise diagnostic information in the event of a server side error (e.g. invalid license, or failure to generate a template).
// 
string
```

***

### [POST]/v1/analyze-image

- Summary  
  Returns an object containing the results of analysis for the provided biometric image.

- Operation id  
  analyze_image

- Description  
  This endpoint accepts a base64 encoded PNG image and returns an object containing the output of a generic image analysis routine on that image.

- Tags  
  Analysis Operations

#### RequestBody

*Required*

The image that is being submitted for analysis.

- application/json

```typescript
// Data transfer object for an image.
// 
{
  // The captured image data in PNG format, encoded as a base64 string. The data string shall not exceed 10MB.
  ImageData: string
}
```

#### Responses

- 200 Successful Response

`application/json`

```typescript
// Object containing the results of a generic image analysis algorithm on a single image. Must include at least 1 and not more than 8 string properties in a JSON object with a depth of 1. The full contents shall not exceed 512 Bytes.
// 
{
  // The output of a generic image analysis algorithm on the image provided. Contents MAY NOT include portions of the raw image data. Note that the example fields are just that - examples.  In this example, 3 of 8 output property fields have been utilized. Field name and output values should be defined based on useful output for evaluating the performance of your image analysis algorithm.
  // 
  [key: string]: string
}
```

- 400 Bad Request

`application/json`

```typescript
// Relevant and concise diagnostic information in the event of a client side error (e.g. malformed requests, or invalid image encoding).
// 
string
```

- 500 Internal Server Error

`application/json`

```typescript
// Relevant and concise diagnostic information in the event of a server side error (e.g. invalid license, or failure to generate a template).
// 
string
```

## References

### #/components/schemas/Image

```typescript
// Data transfer object for an image.
// 
{
  // The captured image data in PNG format, encoded as a base64 string. The data string shall not exceed 10MB.
  ImageData: string
}
```

### #/components/schemas/Info

```typescript
// Basic information describing the algorithm.
{
  // Name of algorithm
  AlgorithmName: string
  // A string enum describing the type of biometric images the algorithm is meant to process
  AlgorithmType: enum[Face, Finger, Iris]
  // Algorithm version identifier
  AlgorithmVersion: string
  // Name of the company which produces the algorithm
  CompanyName: string
  // A string enum to select an MdTF dataset of biometric images for this submission to analyze
  ImageDataset: enum[NIST_MEDS]
  // The recommended allocation of CPUs for the deployed docker container.
  RecommendedCPUs: number //format: integer
  // The recommended allocation of memory (MB) for the deployed docker container.
  RecommendedMem: number //format: integer
  // The email address of an engineer or other technical resource to contact in the event of an error running your service. This field may be left blank if desired.
  TechnicalContactEmail: string
}
```

### #/components/schemas/AnalysisResult

```typescript
// Object containing the results of a generic image analysis algorithm on a single image. Must include at least 1 and not more than 8 string properties in a JSON object with a depth of 1. The full contents shall not exceed 512 Bytes.
// 
{
  // The output of a generic image analysis algorithm on the image provided. Contents MAY NOT include portions of the raw image data. Note that the example fields are just that - examples.  In this example, 3 of 8 output property fields have been utilized. Field name and output values should be defined based on useful output for evaluating the performance of your image analysis algorithm.
  // 
  [key: string]: string
}
```

### #/components/schemas/RequestError

```typescript
// Relevant and concise diagnostic information in the event of a client side error (e.g. malformed requests, or invalid image encoding).
// 
string
```

### #/components/schemas/ServerError

```typescript
// Relevant and concise diagnostic information in the event of a server side error (e.g. invalid license, or failure to generate a template).
// 
string
```