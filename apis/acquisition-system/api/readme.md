# The Maryland Test Facility Acquisition System Interface

> Version 2.1.0

Application Programming Interface for sending acquisition records to the MdTF Backend.

- License: [IDSL API License](https://raw.githubusercontent.com/TheMdTF/mdtf-public/master/LICENSE.md)
- Contact: The MdTF <rally@mdtf.org> https://mdtf.org
- Server: http://acquisition-demo.mdtf.org


## Path Table

| Method | Path | Description |
| --- | --- | --- |
| POST | [/v1/capture](#postv1capture) | Submit a captured image. |

## Reference Table

| Name | Path | Description |
| --- | --- | --- |
| Capture | [#/components/schemas/Capture](#componentsschemascapture) |  |
| Error | [#/components/schemas/Error](#componentsschemaserror) |  |
| capture | [#/components/requestBodies/capture](#componentsrequestbodiescapture) | Face image collected as part of a transaction. |
| 400 | [#/components/responses/400](#componentsresponses400) | The request is malformed. |
| 404 | [#/components/responses/404](#componentsresponses404) | No ongoing transaction was found for the specified StationID. |
| 500 | [#/components/responses/500](#componentsresponses500) | This capture failed because of a server side issue. |

## Path Details

***

### [POST]/v1/capture

- Summary  
  Submit a captured image.

- Operation id  
  Capture

- Description  
  Acquisition system providers may submit image data through this endpoint.

- Tags  
  Image Submission

#### RequestBody

*Required*

Face image collected as part of a transaction.

- application/json

```typescript
{
  // The unique identifier for the station that the image was captured at.
  StationID: string
  // The raw bytes of the image data. The supported formats are jpeg and png. This data should be marshaled in this object as a base64 encoded string.
  Data: string //format: byte
}
```

#### Responses

- 200 Successfully associated the image capture with the ongoing transaction.

- 400 The request is malformed.

`application/json`

```typescript
{
  // The error code.
  Code: integer
  // A description of the returned code.
  Message: string
  // A detailed description of the error that occurred.
  ErrorString: string
}
```

- 404 No ongoing transaction was found for the specified StationID.

`application/json`

```typescript
{
  // The error code.
  Code: integer
  // A description of the returned code.
  Message: string
  // A detailed description of the error that occurred.
  ErrorString: string
}
```

- 500 This capture failed because of a server side issue.

`application/json`

```typescript
{
  // The error code.
  Code: integer
  // A description of the returned code.
  Message: string
  // A detailed description of the error that occurred.
  ErrorString: string
}
```

## References

### #/components/schemas/Capture

```typescript
{
  // The unique identifier for the station that the image was captured at.
  StationID: string
  // The raw bytes of the image data. The supported formats are jpeg and png. This data should be marshaled in this object as a base64 encoded string.
  Data: string //format: byte
}
```

### #/components/schemas/Error

```typescript
{
  // The error code.
  Code: integer
  // A description of the returned code.
  Message: string
  // A detailed description of the error that occurred.
  ErrorString: string
}
```

### #/components/requestBodies/capture

- application/json

```typescript
{
  // The unique identifier for the station that the image was captured at.
  StationID: string
  // The raw bytes of the image data. The supported formats are jpeg and png. This data should be marshaled in this object as a base64 encoded string.
  Data: string //format: byte
}
```

### #/components/responses/400

- application/json

```typescript
{
  // The error code.
  Code: integer
  // A description of the returned code.
  Message: string
  // A detailed description of the error that occurred.
  ErrorString: string
}
```

### #/components/responses/404

- application/json

```typescript
{
  // The error code.
  Code: integer
  // A description of the returned code.
  Message: string
  // A detailed description of the error that occurred.
  ErrorString: string
}
```

### #/components/responses/500

- application/json

```typescript
{
  // The error code.
  Code: integer
  // A description of the returned code.
  Message: string
  // A detailed description of the error that occurred.
  ErrorString: string
}
```