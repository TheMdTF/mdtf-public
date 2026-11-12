# The Maryland Test Facility Active Presentation Attack Detection System Interface

> Version 2.0.1

This document describes an application programming interface for a PAD system in which the PAD Subsystem is logically located within or tightly coupled with the Data Capture Subsystem.  This allows for what ISO/IEC 30107-1:2023 refers to as "through data capture subsystem" PAD methods (Table 2). It also allows for what the standard refers to as "challenge-response" actions (Subsection 5.2.1).  We collectively refer to these kinds of systems as "active" PAD or A-PAD systems.

- License: [IDSL API License](https://raw.githubusercontent.com/TheMdTF/mdtf-public/master/LICENSE.md)
- Contact: The MdTF <rivr@mdtf.org> https://mdtf.org
- Server: https://pad-demo.mdtf.org/


## Path Table

| Method | Path | Description |
| --- | --- | --- |
| POST | [/v1/capture-data-with-pad](#postv1capture-data-with-pad) | Create a biometric data capture with associated PAD information. |

## Reference Table

| Name | Path | Description |
| --- | --- | --- |
| PADDataCapture | [#/components/schemas/PADDataCapture](#componentsschemaspaddatacapture) | Data transfer object for biometric data capture and presentation attack information. |
| PADAnalysis | [#/components/schemas/PADAnalysis](#componentsschemaspadanalysis) | Data transfer object for presentation attack information. |
| PADProperty | [#/components/schemas/PADProperty](#componentsschemaspadproperty) | Key value pair object. |
| PADErrorResponse | [#/components/schemas/PADErrorResponse](#componentsschemaspaderrorresponse) | Presentation attack error response object. |
| Error | [#/components/schemas/Error](#componentsschemaserror) |  |
| pad_data_capture | [#/components/requestBodies/pad_data_capture](#componentsrequestbodiespad_data_capture) | PAD data collected as part of a biometric data capture process. |
| 400 | [#/components/responses/400](#componentsresponses400) | The request is malformed. |
| 404 | [#/components/responses/404](#componentsresponses404) | No ongoing transaction was found for the specified StationID. |
| 410 | [#/components/responses/410](#componentsresponses410) | The test volunteer has exited the station and the transaction has ended. |
| 429 | [#/components/responses/429](#componentsresponses429) | Too many requests for the ongoing transaction. |
| 500 | [#/components/responses/500](#componentsresponses500) | This capture failed because of a server side issue. |

## Path Details

***

### [POST]/v1/capture-data-with-pad

- Summary  
  Create a biometric data capture with associated PAD information.

- Description  
  Active PAD systems may submit data through this endpoint. This request cannot exceed 100 Megabytes in size.

- Tags  
  Data Submission

#### RequestBody

*Required*

PAD data collected as part of a biometric data capture process.

- application/json

```typescript
// Data transfer object for biometric data capture and presentation attack information.
{
  // ID of the station who is submitting this image. IDs will be provided to vendors upon installation for the test event.
  StationID: string
  MobilePlatform: enum[iOS, Android]
  // The captured biometric sample, encoded as a base64 string.  This can be an image, encoded as a PNG or JPEG or a short (<10s) video, encoded as a MOV or a MP4.
  BiometricSample: string
  // Data transfer object for presentation attack information.
  PADAnalysis: {
    // Whether a presentation attack was determined to be detected (True) or not detected (False).
    PADOutcome: boolean
    // A score corresponding to the level of confidence that a presentation attack was detected ranging between 0 and 1.
    PADScore: number //format: double
    // Key value pairs describing presentation attack properties and their relationship to the presentation attack outcome/score. There are no strictly defined properties. The inclusion of descriptive properties is encouraged to provide more context. Do not include image bytes in the properties. (optional)
    // Key value pair object.
    PADProperties: {
      // Name of the property.
      Property?: string
      // Value for the property.
      Value?: string
    }[]
  }
}
```

#### Responses

- 200 Successfully submitted the biometric data capture and PAD information.

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

- 404 No ongoing transaction was found for the specified StationID.

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

- 410 The test volunteer has exited the station and the transaction has ended.

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

- 429 Too many requests for the ongoing transaction.

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

- 500 This capture failed because of a server side issue.

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

### #/components/schemas/PADDataCapture

```typescript
// Data transfer object for biometric data capture and presentation attack information.
{
  // ID of the station who is submitting this image. IDs will be provided to vendors upon installation for the test event.
  StationID: string
  MobilePlatform: enum[iOS, Android]
  // The captured biometric sample, encoded as a base64 string.  This can be an image, encoded as a PNG or JPEG or a short (<10s) video, encoded as a MOV or a MP4.
  BiometricSample: string
  // Data transfer object for presentation attack information.
  PADAnalysis: {
    // Whether a presentation attack was determined to be detected (True) or not detected (False).
    PADOutcome: boolean
    // A score corresponding to the level of confidence that a presentation attack was detected ranging between 0 and 1.
    PADScore: number //format: double
    // Key value pairs describing presentation attack properties and their relationship to the presentation attack outcome/score. There are no strictly defined properties. The inclusion of descriptive properties is encouraged to provide more context. Do not include image bytes in the properties. (optional)
    // Key value pair object.
    PADProperties: {
      // Name of the property.
      Property?: string
      // Value for the property.
      Value?: string
    }[]
  }
}
```

### #/components/schemas/PADAnalysis

```typescript
// Data transfer object for presentation attack information.
{
  // Whether a presentation attack was determined to be detected (True) or not detected (False).
  PADOutcome: boolean
  // A score corresponding to the level of confidence that a presentation attack was detected ranging between 0 and 1.
  PADScore: number //format: double
  // Key value pairs describing presentation attack properties and their relationship to the presentation attack outcome/score. There are no strictly defined properties. The inclusion of descriptive properties is encouraged to provide more context. Do not include image bytes in the properties. (optional)
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

### #/components/schemas/PADErrorResponse

```typescript
// Presentation attack error response object.
{
  // Descriptive error string.
  error: string
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

### #/components/requestBodies/pad_data_capture

- application/json

```typescript
// Data transfer object for biometric data capture and presentation attack information.
{
  // ID of the station who is submitting this image. IDs will be provided to vendors upon installation for the test event.
  StationID: string
  MobilePlatform: enum[iOS, Android]
  // The captured biometric sample, encoded as a base64 string.  This can be an image, encoded as a PNG or JPEG or a short (<10s) video, encoded as a MOV or a MP4.
  BiometricSample: string
  // Data transfer object for presentation attack information.
  PADAnalysis: {
    // Whether a presentation attack was determined to be detected (True) or not detected (False).
    PADOutcome: boolean
    // A score corresponding to the level of confidence that a presentation attack was detected ranging between 0 and 1.
    PADScore: number //format: double
    // Key value pairs describing presentation attack properties and their relationship to the presentation attack outcome/score. There are no strictly defined properties. The inclusion of descriptive properties is encouraged to provide more context. Do not include image bytes in the properties. (optional)
    // Key value pair object.
    PADProperties: {
      // Name of the property.
      Property?: string
      // Value for the property.
      Value?: string
    }[]
  }
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

### #/components/responses/404

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

### #/components/responses/410

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

### #/components/responses/429

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