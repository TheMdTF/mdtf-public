# The Maryland Test Facility Identity Document Validation Interface

> Version 0.0.1

This documents the API requirements for MdTF testing of identity document validation algorithms.  Identity Document validation testing at the MdTF is supported by the Department of Homeland Security Science and Technology Directorate (DHS S&T).  For more information please visit [https://mdtf.org](https://mdtf.org) and [https://www.dhs.gov/science-and-technology/BI-TC](https://www.dhs.gov/science-and-technology/BI-TC).

- License: [IDSL API License](https://raw.githubusercontent.com/TheMdTF/mdtf-public/master/LICENSE.md)
- Contact: The MdTF <info@mdtf.org> https://mdtf.org


## Path Table

| Method | Path | Description |
| --- | --- | --- |
| GET | [/v1/info](#getv1info) | Return information about the document validation algorithm. |
| POST | [/v1/validate](#postv1validate) | Validate an identity document based on images. |

## Reference Table

| Name | Path | Description |
| --- | --- | --- |
| Info | [#/components/schemas/Info](#componentsschemasinfo) | Basic information describing the document validation algorithm. |
| CapturedDocument | [#/components/schemas/CapturedDocument](#componentsschemascaptureddocument) | Object representation of a document. Consists of an image of the front and back of a document.  Back document may be an empty string and this case should not in and of itself be conisdered a validation failure. |
| ValidityProperty | [#/components/schemas/ValidityProperty](#componentsschemasvalidityproperty) | Key value pair object. |
| ValidationResponse | [#/components/schemas/ValidationResponse](#componentsschemasvalidationresponse) | Validation response object. |
| ValidationErrorResponse | [#/components/schemas/ValidationErrorResponse](#componentsschemasvalidationerrorresponse) | Validation error response object. |

## Path Details

***

### [GET]/v1/info

- Summary  
  Return information about the document validation algorithm.

- Operation id  
  Info

- Description  
  This endpoint returns information about the validation algorithm.

- Tags  
  Algorithm Information

#### Responses

- 200 Successful operation

`application/json`

```typescript
// Basic information describing the document validation algorithm.
{
  // Name of the algorithm.
  AlgorithmName: string
  // Algorithm version identifier.
  AlgorithmVersion: string
  // Name of the company that produced the algorithm.
  CompanyName: string
  // The name of the test event. Please use RIVR_IDV.
  Event: enum[RIVR_IDV]
  // The recommended allocation of CPUs for the deployed docker container.
  RecommendedCPUs: integer //format: int64
  // The recommended allocation of memory (MB) for the deployed docker container.
  RecommendedMem: integer //format: int64
  // The email address of an engineer or other technical resource to contact in the event of an error running your service.
  TechnicalContactEmail: string
}
```

- 500 Server error

`application/json`

```typescript
// Validation error response object.
{
  // Descriptive error string.
  error: string
}
```

***

### [POST]/v1/validate

- Summary  
  Validate an identity document based on images.

- Operation id  
  ValidateDocument

- Description  
  Receive base64 encoded PNG or JPEG images of the front and back of a document. Validation requires that the document is judged to be authentic. See [NIST SP 800-63A](https://pages.nist.gov/800-63-3-Implementation-Resources/63A/ial2remote/) Section A.10.2: Identity Validation for recommended best practices. Image orientation is not guaranteed. Subsystem must distinguish and process accordingly. The expected request duration is less than 30 seconds. Data will be passed to your system as a request including DocumentFront and DocumentBack base64 encoded images. For passports, DocumentFront will be an image of the data page and DocumentBack will be an empty string.

- Tags  
  Validate

#### RequestBody

*Required*

- application/json

```typescript
// Object representation of a document. Consists of an image of the front and back of a document.  Back document may be an empty string and this case should not in and of itself be conisdered a validation failure.
{
  // Image of front of document; base64 encoded bytes.
  DocumentFront: string //format: byte
  // Image of back of document; base64 encoded bytes.  May be an empty string if back of document is not applicable or available.
  DocumentBack: string //format: byte
}
```

#### Responses

- 200 Successful validation operation

`application/json`

```typescript
// Validation response object.
{
  // Whether the input document is determined to be valid (True) or invalid (False).
  ValidityOutcome: boolean
  // A score corresponding to the level of confidence that the document is valid ranging between 0 and 1 (optional).
  ValidityScore?: number //format: double
  // Key value pairs describing document properties and their relationship to the validity decision. There are no strictly defined  properties. The inclusion of descriptive properties is encouraged to provide more context. Do not return any personally identifiable information or image bytes from the documents in the properties. (optional)
  // Key value pair object.
  ValidityProperties: {
    // Name of the property.
    Property?: string
    // Value for the property.
    Value?: string
  }[]
}
```

- 400 Client error

`application/json`

```typescript
// Validation error response object.
{
  // Descriptive error string.
  error: string
}
```

- 500 Server error

`application/json`

```typescript
// Validation error response object.
{
  // Descriptive error string.
  error: string
}
```

## References

### #/components/schemas/Info

```typescript
// Basic information describing the document validation algorithm.
{
  // Name of the algorithm.
  AlgorithmName: string
  // Algorithm version identifier.
  AlgorithmVersion: string
  // Name of the company that produced the algorithm.
  CompanyName: string
  // The name of the test event. Please use RIVR_IDV.
  Event: enum[RIVR_IDV]
  // The recommended allocation of CPUs for the deployed docker container.
  RecommendedCPUs: integer //format: int64
  // The recommended allocation of memory (MB) for the deployed docker container.
  RecommendedMem: integer //format: int64
  // The email address of an engineer or other technical resource to contact in the event of an error running your service.
  TechnicalContactEmail: string
}
```

### #/components/schemas/CapturedDocument

```typescript
// Object representation of a document. Consists of an image of the front and back of a document.  Back document may be an empty string and this case should not in and of itself be conisdered a validation failure.
{
  // Image of front of document; base64 encoded bytes.
  DocumentFront: string //format: byte
  // Image of back of document; base64 encoded bytes.  May be an empty string if back of document is not applicable or available.
  DocumentBack: string //format: byte
}
```

### #/components/schemas/ValidityProperty

```typescript
// Key value pair object.
{
  // Name of the property.
  Property?: string
  // Value for the property.
  Value?: string
}
```

### #/components/schemas/ValidationResponse

```typescript
// Validation response object.
{
  // Whether the input document is determined to be valid (True) or invalid (False).
  ValidityOutcome: boolean
  // A score corresponding to the level of confidence that the document is valid ranging between 0 and 1 (optional).
  ValidityScore?: number //format: double
  // Key value pairs describing document properties and their relationship to the validity decision. There are no strictly defined  properties. The inclusion of descriptive properties is encouraged to provide more context. Do not return any personally identifiable information or image bytes from the documents in the properties. (optional)
  // Key value pair object.
  ValidityProperties: {
    // Name of the property.
    Property?: string
    // Value for the property.
    Value?: string
  }[]
}
```

### #/components/schemas/ValidationErrorResponse

```typescript
// Validation error response object.
{
  // Descriptive error string.
  error: string
}
```