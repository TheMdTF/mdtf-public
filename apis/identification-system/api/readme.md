# The Maryland Test Facility Identification System Interface

> Version 1.0.0

Application Programming Interface for receiving identification system requests from the MdTF Backend. Identification systems maintain an enrolled gallery of biometric templates and perform 1:N searches against that gallery.

- License: [IDSL API License](https://raw.githubusercontent.com/TheMdTF/mdtf-public/master/LICENSE.md)
- Contact: The MdTF <rally@mdtf.org> https://mdtf.org
- Server: http://localhost:8080/


## Path Table

| Method | Path | Description |
| --- | --- | --- |
| POST | [/v1/enroll](#postv1enroll) | Enroll a subject into the gallery. |
| POST | [/v1/identify](#postv1identify) | Search the gallery for candidates matching the probe image. |
| GET | [/v1/info](#getv1info) | Returns basic information for the algorithm. |

## Reference Table

| Name | Path | Description |
| --- | --- | --- |
| Image | [#/components/schemas/Image](#componentsschemasimage) | Data transfer object for an image. |
| EnrollmentRequest | [#/components/schemas/EnrollmentRequest](#componentsschemasenrollmentrequest) | A request to enroll a single biometric image into the gallery. The SubjectId<br>uniquely identifies the subject across all enrollment and identification operations<br>within a single test session. Enrolling a SubjectId that already exists in the<br>gallery MUST overwrite the existing entry. |
| Candidate | [#/components/schemas/Candidate](#componentsschemascandidate) | A single candidate returned from a 1:N identification search. Each candidate<br>represents one enrolled subject and carries the similarity score produced<br>by the algorithm for that subject. |
| IdentificationRequest | [#/components/schemas/IdentificationRequest](#componentsschemasidentificationrequest) | A request to search the enrolled gallery for the subject that best matches<br>the provided probe image. |
| IdentificationResult | [#/components/schemas/IdentificationResult](#componentsschemasidentificationresult) | The result of a 1:N identification search against the enrolled gallery.<br>Candidates are ordered from most to least similar to the probe. |
| Info | [#/components/schemas/Info](#componentsschemasinfo) | Basic information describing the algorithm. |
| RequestError | [#/components/schemas/RequestError](#componentsschemasrequesterror) | Relevant and concise diagnostic information in the event of a client side error (e.g. malformed requests, or invalid<br>image encoding). |
| ServerError | [#/components/schemas/ServerError](#componentsschemasservererror) | Relevant and concise diagnostic information in the event of a server side error<br>(e.g. invalid license, or failure to generate a template). |
| enrollment_request | [#/components/requestBodies/enrollment_request](#componentsrequestbodiesenrollment_request) | A biometric image and a subject identifier to be enrolled into the gallery. |
| identification_request | [#/components/requestBodies/identification_request](#componentsrequestbodiesidentification_request) | A probe biometric image and search parameters for a 1:N gallery search. |
| infoResponse | [#/components/responses/infoResponse](#componentsresponsesinforesponse) | info response object |
| enrollResponse | [#/components/responses/enrollResponse](#componentsresponsesenrollresponse) | Successful enrollment. The subject has been added to or updated in the gallery. |
| identifyResponse | [#/components/responses/identifyResponse](#componentsresponsesidentifyresponse) | Successful identification search. |
| BadRequest | [#/components/responses/BadRequest](#componentsresponsesbadrequest) | Bad Request |
| InternalServerError | [#/components/responses/InternalServerError](#componentsresponsesinternalservererror) | Internal Server Error |

## Path Details

***

### [POST]/v1/enroll

- Summary  
  Enroll a subject into the gallery.

- Operation id  
  EnrollSubject

- Description  
  This endpoint accepts a base64 encoded PNG image and a SubjectId. The algorithm  
  must extract a biometric template from the image and store it in an internal  
  gallery, keyed by SubjectId.

  If the SubjectId already exists in the gallery, the existing entry MUST be  
  overwritten with the newly provided image.

  The gallery persists across calls to this endpoint and is searched by  
  subsequent calls to /v1/identify. The gallery is cleared only when the  
  service is restarted.

- Tags  
  Biometric Operations

#### RequestBody

*Required*

A biometric image and a subject identifier to be enrolled into the gallery.

- application/json

```typescript
// A request to enroll a single biometric image into the gallery. The SubjectId
// uniquely identifies the subject across all enrollment and identification operations
// within a single test session. Enrolling a SubjectId that already exists in the
// gallery MUST overwrite the existing entry.
// 
{
  // A unique identifier for the subject being enrolled.
  SubjectId: string
  // Data transfer object for an image.
  // 
  Image: {
    // The captured image data in PNG format, encoded as a base64 string.
    ImageData: string
  }
}
```

#### Responses

- 200 Successful enrollment. The subject has been added to or updated in the gallery.

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

### [POST]/v1/identify

- Summary  
  Search the gallery for candidates matching the probe image.

- Operation id  
  IdentifySubject

- Description  
  This endpoint accepts a base64 encoded PNG probe image and a MaxCandidates  
  value. The algorithm must extract a biometric template from the probe image  
  and compare it against all subjects currently enrolled in the gallery.

  The response is a list of Candidate objects sorted in descending order by  
  Score (most similar first). The list MUST contain at most MaxCandidates  
  entries.

  If the gallery is empty, the response MUST be a 200 with an empty  
  Candidates list.

- Tags  
  Biometric Operations

#### RequestBody

*Required*

A probe biometric image and search parameters for a 1:N gallery search.

- application/json

```typescript
// A request to search the enrolled gallery for the subject that best matches
// the provided probe image.
// 
{
  // Data transfer object for an image.
  // 
  Image: {
    // The captured image data in PNG format, encoded as a base64 string.
    ImageData: string
  }
  // The maximum number of candidates to return. The response list MUST contain
  // at most this many entries. If fewer enrolled subjects exist than MaxCandidates,
  // all enrolled subjects may be returned. Must be greater than or equal to 1.
  // 
  MaxCandidates: integer //minimum: 1
}
```

#### Responses

- 200 Successful identification search.

`application/json`

```typescript
// The result of a 1:N identification search against the enrolled gallery.
// Candidates are ordered from most to least similar to the probe.
// 
{
  // A single candidate returned from a 1:N identification search. Each candidate
  // represents one enrolled subject and carries the similarity score produced
  // by the algorithm for that subject.
  // 
  Candidates: {
    // The unique identifier of the enrolled subject.
    SubjectId: string
    // A similarity score produced by the algorithm for this candidate. Higher values
    // indicate greater similarity to the probe. Candidates MUST be returned in
    // descending order of Score (highest score first).
    // 
    Score: number //format: float
  }[]
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
  Test: enum[MDTF_2025_IDENTIFICATION]
  // A map of preset False Positive Identification Rates (FPIR) to vendor-provided threshold values.
  // Score values returned in Candidate objects from calls to /v1/identify indicate a positive
  // identification determination by the algorithm if they are greater than the provided threshold
  // value at the respective FPIR.
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

### #/components/schemas/EnrollmentRequest

```typescript
// A request to enroll a single biometric image into the gallery. The SubjectId
// uniquely identifies the subject across all enrollment and identification operations
// within a single test session. Enrolling a SubjectId that already exists in the
// gallery MUST overwrite the existing entry.
// 
{
  // A unique identifier for the subject being enrolled.
  SubjectId: string
  // Data transfer object for an image.
  // 
  Image: {
    // The captured image data in PNG format, encoded as a base64 string.
    ImageData: string
  }
}
```

### #/components/schemas/Candidate

```typescript
// A single candidate returned from a 1:N identification search. Each candidate
// represents one enrolled subject and carries the similarity score produced
// by the algorithm for that subject.
// 
{
  // The unique identifier of the enrolled subject.
  SubjectId: string
  // A similarity score produced by the algorithm for this candidate. Higher values
  // indicate greater similarity to the probe. Candidates MUST be returned in
  // descending order of Score (highest score first).
  // 
  Score: number //format: float
}
```

### #/components/schemas/IdentificationRequest

```typescript
// A request to search the enrolled gallery for the subject that best matches
// the provided probe image.
// 
{
  // Data transfer object for an image.
  // 
  Image: {
    // The captured image data in PNG format, encoded as a base64 string.
    ImageData: string
  }
  // The maximum number of candidates to return. The response list MUST contain
  // at most this many entries. If fewer enrolled subjects exist than MaxCandidates,
  // all enrolled subjects may be returned. Must be greater than or equal to 1.
  // 
  MaxCandidates: integer //minimum: 1
}
```

### #/components/schemas/IdentificationResult

```typescript
// The result of a 1:N identification search against the enrolled gallery.
// Candidates are ordered from most to least similar to the probe.
// 
{
  // A single candidate returned from a 1:N identification search. Each candidate
  // represents one enrolled subject and carries the similarity score produced
  // by the algorithm for that subject.
  // 
  Candidates: {
    // The unique identifier of the enrolled subject.
    SubjectId: string
    // A similarity score produced by the algorithm for this candidate. Higher values
    // indicate greater similarity to the probe. Candidates MUST be returned in
    // descending order of Score (highest score first).
    // 
    Score: number //format: float
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
  Test: enum[MDTF_2025_IDENTIFICATION]
  // A map of preset False Positive Identification Rates (FPIR) to vendor-provided threshold values.
  // Score values returned in Candidate objects from calls to /v1/identify indicate a positive
  // identification determination by the algorithm if they are greater than the provided threshold
  // value at the respective FPIR.
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

### #/components/requestBodies/enrollment_request

- application/json

```typescript
// A request to enroll a single biometric image into the gallery. The SubjectId
// uniquely identifies the subject across all enrollment and identification operations
// within a single test session. Enrolling a SubjectId that already exists in the
// gallery MUST overwrite the existing entry.
// 
{
  // A unique identifier for the subject being enrolled.
  SubjectId: string
  // Data transfer object for an image.
  // 
  Image: {
    // The captured image data in PNG format, encoded as a base64 string.
    ImageData: string
  }
}
```

### #/components/requestBodies/identification_request

- application/json

```typescript
// A request to search the enrolled gallery for the subject that best matches
// the provided probe image.
// 
{
  // Data transfer object for an image.
  // 
  Image: {
    // The captured image data in PNG format, encoded as a base64 string.
    ImageData: string
  }
  // The maximum number of candidates to return. The response list MUST contain
  // at most this many entries. If fewer enrolled subjects exist than MaxCandidates,
  // all enrolled subjects may be returned. Must be greater than or equal to 1.
  // 
  MaxCandidates: integer //minimum: 1
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
  Test: enum[MDTF_2025_IDENTIFICATION]
  // A map of preset False Positive Identification Rates (FPIR) to vendor-provided threshold values.
  // Score values returned in Candidate objects from calls to /v1/identify indicate a positive
  // identification determination by the algorithm if they are greater than the provided threshold
  // value at the respective FPIR.
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

### #/components/responses/enrollResponse

Successful enrollment. The subject has been added to or updated in the gallery.

### #/components/responses/identifyResponse

- application/json

```typescript
// The result of a 1:N identification search against the enrolled gallery.
// Candidates are ordered from most to least similar to the probe.
// 
{
  // A single candidate returned from a 1:N identification search. Each candidate
  // represents one enrolled subject and carries the similarity score produced
  // by the algorithm for that subject.
  // 
  Candidates: {
    // The unique identifier of the enrolled subject.
    SubjectId: string
    // A similarity score produced by the algorithm for this candidate. Higher values
    // indicate greater similarity to the probe. Candidates MUST be returned in
    // descending order of Score (highest score first).
    // 
    Score: number //format: float
  }[]
}
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