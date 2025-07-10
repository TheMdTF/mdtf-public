<!-- Generator: Widdershins v4.0.1 -->

<h1 id="the-maryland-test-facility-matching-system-interface">The Maryland Test Facility Matching System Interface v1.4.0</h1>

> Scroll down for code samples, example requests and responses. Select a language for code samples from the tabs above or the mobile navigation menu.

Application Programming Interface for receiving matching system requests from the MdTF Backend.
In version 1.4.0, `/v1/compare-list` was renamed to `/v1/verify`, and `/v1/info` expanded to advertise SupportedOperations, SupportedModalities, and SupportedImageTypes.

Base URLs:

* <a href="http://localhost:8080/">http://localhost:8080/</a>

Email: <a href="mailto:rally@mdtf.org">The MdTF</a> Web: <a href="https://mdtf.org">The MdTF</a> 
License: <a href="https://raw.githubusercontent.com/TheMdTF/mdtf-public/master/LICENSE.md">IDSL API License</a>

<h1 id="the-maryland-test-facility-matching-system-interface-biometric-operations">Biometric Operations</h1>

## Generate templates from a provided image.

<a id="opIdCreateTemplate"></a>

> Code samples

```shell
# You can also use wget
curl -X POST http://localhost:8080/v1/create-template \
  -H 'Content-Type: application/json' \
  -H 'Accept: application/json'

```

```javascript
const inputBody = '{
  "Modality": "Finger",
  "ImageType": "FlatLeftThumb",
  "ImageId": "image-abc123",
  "ImageData": "iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAIAAACQd1PeAAAAEElEQVR4nGJiYGAABAAA//8ADAADcZGLFwAAAABJRU5ErkJggg=="
}';
const headers = {
  'Content-Type':'application/json',
  'Accept':'application/json'
};

fetch('http://localhost:8080/v1/create-template',
{
  method: 'POST',
  body: inputBody,
  headers: headers
})
.then(function(res) {
    return res.json();
}).then(function(body) {
    console.log(body);
});

```

```go
package main

import (
       "bytes"
       "net/http"
)

func main() {

    headers := map[string][]string{
        "Content-Type": []string{"application/json"},
        "Accept": []string{"application/json"},
    }

    data := bytes.NewBuffer([]byte{jsonReq})
    req, err := http.NewRequest("POST", "http://localhost:8080/v1/create-template", data)
    req.Header = headers

    client := &http.Client{}
    resp, err := client.Do(req)
    // ...
}

```

`POST /v1/create-template`

> Body parameter

```json
{
  "Modality": "Finger",
  "ImageType": "FlatLeftThumb",
  "ImageId": "image-abc123",
  "ImageData": "iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAIAAACQd1PeAAAAEElEQVR4nGJiYGAABAAA//8ADAADcZGLFwAAAABJRU5ErkJggg=="
}
```

<h3 id="generate-templates-from-a-provided-image.-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|body|body|[Image](#schemaimage)|true|Biometric image to be submitted for template extraction.|

> Example responses

> 200 Response

```json
[
  {
    "Modality": "Finger",
    "ImageType": "FlatLeftThumb",
    "TemplateId": "target-001",
    "Template": "dGhpcyBzZW50ZW5jZSBpcyBhbiBleGFtcGxlIHRlbXBsYXRlLi4K",
    "Error": "string"
  }
]
```

> 400 Response

```json
{
  "message": "Invalid request: missing required field TemplateId"
}
```

> 413 Response

```json
{
  "message": "Request rejected: maximum templates exceeded (limit is 100)."
}
```

> 422 Response

```json
{
  "message": "Unprocessable entity: modality mismatch between probe and target."
}
```

> 500 Response

```json
{
  "message": "Internal server error: unexpected processing failure."
}
```

<h3 id="generate-templates-from-a-provided-image.-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|Template creation success response.|Inline|
|400|[Bad Request](https://tools.ietf.org/html/rfc7231#section-6.5.1)|Bad Request|[Error](#schemaerror)|
|413|[Payload Too Large](https://tools.ietf.org/html/rfc7231#section-6.5.11)|Payload Too Large|[Error](#schemaerror)|
|422|[Unprocessable Entity](https://tools.ietf.org/html/rfc2518#section-10.3)|Unprocessable Entity|[Error](#schemaerror)|
|500|[Internal Server Error](https://tools.ietf.org/html/rfc7231#section-6.6.1)|Internal Server Error|[Error](#schemaerror)|

<h3 id="generate-templates-from-a-provided-image.-responseschema">Response Schema</h3>

Status Code **200**

*A list of templates generated from an input image.
*

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|*anonymous*|[allOf]|false|none|A list of templates generated from an input image.|

*allOf*

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|» *anonymous*|[BiometricData](#schemabiometricdata)|false|none|Base object for biometric data, including modality and image type.|
|»» Modality|[Modality](#schemamodality)|true|none|Type of biometric modality processed.|
|»» ImageType|[ImageType](#schemaimagetype)|true|none|Type of biometric image.|

*and*

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|» *anonymous*|object|false|none|none|
|»» TemplateId|string|true|none|Unique identifier for the template, assigned by the client.|
|»» Template|string(byte)|true|none|The template data encoded as a base64 string.<br>The data string shall not exceed 1 MB.|
|»» Error|string|false|none|An error produced from creating the biometric data|

#### Enumerated Values

|Property|Value|
|---|---|
|Modality|Face|
|Modality|Finger|
|Modality|Iris|
|Modality|Palm|
|Modality|Unknown|
|ImageType|EjiOrTip|
|ImageType|Face|
|ImageType|FlatLeftAndRightThumbs|
|ImageType|FlatLeftFourFingers|
|ImageType|FlatLeftThumb|
|ImageType|FlatRightFourFingers|
|ImageType|FlatRightThumb|
|ImageType|Iris|
|ImageType|LeftExtraDigit|
|ImageType|LeftIndexAndMiddle|
|ImageType|LeftIndexFinger|
|ImageType|LeftIris|
|ImageType|LeftLittleFinger|
|ImageType|LeftMiddleFinger|
|ImageType|LeftPalm|
|ImageType|LeftRingAndLittle|
|ImageType|LeftRingFinger|
|ImageType|LeftThumb|
|ImageType|RightExtraDigit|
|ImageType|RightIndexAndMiddle|
|ImageType|RightIndexFinger|
|ImageType|RightIris|
|ImageType|RightLittleFinger|
|ImageType|RightMiddleFinger|
|ImageType|RightPalm|
|ImageType|RightRingAndLittle|
|ImageType|RightRingFinger|
|ImageType|RightThumb|
|ImageType|RightTwo|
|ImageType|RolledLeftExtraDigit|
|ImageType|RolledLeftIndexFinger|
|ImageType|RolledLeftLittleFinger|
|ImageType|RolledLeftMiddleFinger|
|ImageType|RolledLeftRingFinger|
|ImageType|RolledLeftThumb|
|ImageType|RolledRightExtraDigit|
|ImageType|RolledRightIndexFinger|
|ImageType|RolledRightLittleFinger|
|ImageType|RolledRightMiddleFinger|
|ImageType|RolledRightRingFinger|
|ImageType|RolledRightThumb|
|ImageType|Unknown|
|ImageType|UnknownFinger|
|ImageType|UnknownFrictionRidge|

<aside class="success">
This operation does not require authentication
</aside>

## Verify a probe template against a list of target templates.

<a id="opIdVerifyTemplates"></a>

> Code samples

```shell
# You can also use wget
curl -X POST http://localhost:8080/v1/verify \
  -H 'Content-Type: application/json' \
  -H 'Accept: application/json'

```

```javascript
const inputBody = '{
  "ProbeTemplate": {
    "Modality": "Finger",
    "ImageType": "FlatLeftThumb",
    "TemplateId": "target-001",
    "Template": "dGhpcyBzZW50ZW5jZSBpcyBhbiBleGFtcGxlIHRlbXBsYXRlLi4K",
    "Error": "string"
  },
  "TargetTemplateList": [
    {
      "Modality": "Finger",
      "ImageType": "FlatLeftThumb",
      "TemplateId": "target-001",
      "Template": "dGhpcyBzZW50ZW5jZSBpcyBhbiBleGFtcGxlIHRlbXBsYXRlLi4K",
      "Error": "string"
    }
  ]
}';
const headers = {
  'Content-Type':'application/json',
  'Accept':'application/json'
};

fetch('http://localhost:8080/v1/verify',
{
  method: 'POST',
  body: inputBody,
  headers: headers
})
.then(function(res) {
    return res.json();
}).then(function(body) {
    console.log(body);
});

```

```go
package main

import (
       "bytes"
       "net/http"
)

func main() {

    headers := map[string][]string{
        "Content-Type": []string{"application/json"},
        "Accept": []string{"application/json"},
    }

    data := bytes.NewBuffer([]byte{jsonReq})
    req, err := http.NewRequest("POST", "http://localhost:8080/v1/verify", data)
    req.Header = headers

    client := &http.Client{}
    resp, err := client.Do(req)
    // ...
}

```

`POST /v1/verify`

> Body parameter

```json
{
  "ProbeTemplate": {
    "Modality": "Finger",
    "ImageType": "FlatLeftThumb",
    "TemplateId": "target-001",
    "Template": "dGhpcyBzZW50ZW5jZSBpcyBhbiBleGFtcGxlIHRlbXBsYXRlLi4K",
    "Error": "string"
  },
  "TargetTemplateList": [
    {
      "Modality": "Finger",
      "ImageType": "FlatLeftThumb",
      "TemplateId": "target-001",
      "Template": "dGhpcyBzZW50ZW5jZSBpcyBhbiBleGFtcGxlIHRlbXBsYXRlLi4K",
      "Error": "string"
    }
  ]
}
```

<h3 id="verify-a-probe-template-against-a-list-of-target-templates.-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|body|body|[VerifyRequest](#schemaverifyrequest)|true|A probe template and list of target templates for verification.|

> Example responses

> 200 Response

```json
[
  {
    "TargetTemplateId": "target-001",
    "Score": 8734,
    "Error": "string"
  }
]
```

> 400 Response

```json
{
  "message": "Invalid request: missing required field TemplateId"
}
```

> 413 Response

```json
{
  "message": "Request rejected: maximum templates exceeded (limit is 100)."
}
```

> 422 Response

```json
{
  "message": "Unprocessable entity: modality mismatch between probe and target."
}
```

> 500 Response

```json
{
  "message": "Internal server error: unexpected processing failure."
}
```

<h3 id="verify-a-probe-template-against-a-list-of-target-templates.-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|Verification success response.|Inline|
|400|[Bad Request](https://tools.ietf.org/html/rfc7231#section-6.5.1)|Bad Request|[Error](#schemaerror)|
|413|[Payload Too Large](https://tools.ietf.org/html/rfc7231#section-6.5.11)|Payload Too Large|[Error](#schemaerror)|
|422|[Unprocessable Entity](https://tools.ietf.org/html/rfc2518#section-10.3)|Unprocessable Entity|[Error](#schemaerror)|
|500|[Internal Server Error](https://tools.ietf.org/html/rfc7231#section-6.6.1)|Internal Server Error|[Error](#schemaerror)|

<h3 id="verify-a-probe-template-against-a-list-of-target-templates.-responseschema">Response Schema</h3>

Status Code **200**

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|*anonymous*|[[Comparison](#schemacomparison)]|false|none|[A similarity score result tied to a specific target template.<br>]|
|» TargetTemplateId|string|true|none|Target TemplateId for which this comparison score applies.|
|» Score|number(float)|true|none|Similarity score produced by the algorithm.|
|» Error|string|false|none|An error produced from creating the biometric data|

<aside class="success">
This operation does not require authentication
</aside>

<h1 id="the-maryland-test-facility-matching-system-interface-algorithm-information">Algorithm Information</h1>

## Retrieve server and algorithm information.

<a id="opIdGetInfo"></a>

> Code samples

```shell
# You can also use wget
curl -X GET http://localhost:8080/v1/info \
  -H 'Accept: application/json'

```

```javascript

const headers = {
  'Accept':'application/json'
};

fetch('http://localhost:8080/v1/info',
{
  method: 'GET',

  headers: headers
})
.then(function(res) {
    return res.json();
}).then(function(body) {
    console.log(body);
});

```

```go
package main

import (
       "bytes"
       "net/http"
)

func main() {

    headers := map[string][]string{
        "Accept": []string{"application/json"},
    }

    data := bytes.NewBuffer([]byte{jsonReq})
    req, err := http.NewRequest("GET", "http://localhost:8080/v1/info", data)
    req.Header = headers

    client := &http.Client{}
    resp, err := client.Do(req)
    // ...
}

```

`GET /v1/info`

> Example responses

> 200 Response

```json
{
  "AlgorithmName": "AlwaysTrue",
  "AlgorithmVersion": "1.0.1",
  "CompanyName": "MdTF",
  "TechnicalContactEmail": "john@mdtf.org",
  "RecommendedCPUs": 0.5,
  "RecommendedMem": 512,
  "Thresholds": {
    "1:500": 7543,
    "1:1e3": 8730.5,
    "1:1e4": 9321.2,
    "1:1e5": 9863.7,
    "1:1e6": 9972.8
  },
  "SupportedOperations": [
    "Verify"
  ],
  "SupportedModalities": [
    "Finger"
  ],
  "SupportedImageTypes": [
    "FlatLeftThumb"
  ]
}
```

> 500 Response

```json
{
  "message": "Internal server error: unexpected processing failure."
}
```

<h3 id="retrieve-server-and-algorithm-information.-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|Algorithm info response.|[Info](#schemainfo)|
|500|[Internal Server Error](https://tools.ietf.org/html/rfc7231#section-6.6.1)|Internal Server Error|[Error](#schemaerror)|

<aside class="success">
This operation does not require authentication
</aside>

# Schemas

<h2 id="tocS_Operation">Operation</h2>
<!-- backwards compatibility -->
<a id="schemaoperation"></a>
<a id="schema_Operation"></a>
<a id="tocSoperation"></a>
<a id="tocsoperation"></a>

```json
"Verify"

```

An operation supported by the server.

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|*anonymous*|string|false|none|An operation supported by the server.|

#### Enumerated Values

|Property|Value|
|---|---|
|*anonymous*|CreateTemplate|
|*anonymous*|Verify|

<h2 id="tocS_Modality">Modality</h2>
<!-- backwards compatibility -->
<a id="schemamodality"></a>
<a id="schema_Modality"></a>
<a id="tocSmodality"></a>
<a id="tocsmodality"></a>

```json
"Finger"

```

Type of biometric modality processed.

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|*anonymous*|string|false|none|Type of biometric modality processed.|

#### Enumerated Values

|Property|Value|
|---|---|
|*anonymous*|Face|
|*anonymous*|Finger|
|*anonymous*|Iris|
|*anonymous*|Palm|
|*anonymous*|Unknown|

<h2 id="tocS_ImageType">ImageType</h2>
<!-- backwards compatibility -->
<a id="schemaimagetype"></a>
<a id="schema_ImageType"></a>
<a id="tocSimagetype"></a>
<a id="tocsimagetype"></a>

```json
"FlatLeftThumb"

```

Type of biometric image.

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|*anonymous*|string|false|none|Type of biometric image.|

#### Enumerated Values

|Property|Value|
|---|---|
|*anonymous*|EjiOrTip|
|*anonymous*|Face|
|*anonymous*|FlatLeftAndRightThumbs|
|*anonymous*|FlatLeftFourFingers|
|*anonymous*|FlatLeftThumb|
|*anonymous*|FlatRightFourFingers|
|*anonymous*|FlatRightThumb|
|*anonymous*|Iris|
|*anonymous*|LeftExtraDigit|
|*anonymous*|LeftIndexAndMiddle|
|*anonymous*|LeftIndexFinger|
|*anonymous*|LeftIris|
|*anonymous*|LeftLittleFinger|
|*anonymous*|LeftMiddleFinger|
|*anonymous*|LeftPalm|
|*anonymous*|LeftRingAndLittle|
|*anonymous*|LeftRingFinger|
|*anonymous*|LeftThumb|
|*anonymous*|RightExtraDigit|
|*anonymous*|RightIndexAndMiddle|
|*anonymous*|RightIndexFinger|
|*anonymous*|RightIris|
|*anonymous*|RightLittleFinger|
|*anonymous*|RightMiddleFinger|
|*anonymous*|RightPalm|
|*anonymous*|RightRingAndLittle|
|*anonymous*|RightRingFinger|
|*anonymous*|RightThumb|
|*anonymous*|RightTwo|
|*anonymous*|RolledLeftExtraDigit|
|*anonymous*|RolledLeftIndexFinger|
|*anonymous*|RolledLeftLittleFinger|
|*anonymous*|RolledLeftMiddleFinger|
|*anonymous*|RolledLeftRingFinger|
|*anonymous*|RolledLeftThumb|
|*anonymous*|RolledRightExtraDigit|
|*anonymous*|RolledRightIndexFinger|
|*anonymous*|RolledRightLittleFinger|
|*anonymous*|RolledRightMiddleFinger|
|*anonymous*|RolledRightRingFinger|
|*anonymous*|RolledRightThumb|
|*anonymous*|Unknown|
|*anonymous*|UnknownFinger|
|*anonymous*|UnknownFrictionRidge|

<h2 id="tocS_BiometricData">BiometricData</h2>
<!-- backwards compatibility -->
<a id="schemabiometricdata"></a>
<a id="schema_BiometricData"></a>
<a id="tocSbiometricdata"></a>
<a id="tocsbiometricdata"></a>

```json
{
  "Modality": "Finger",
  "ImageType": "FlatLeftThumb"
}

```

Base object for biometric data, including modality and image type.

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|Modality|[Modality](#schemamodality)|true|none|Type of biometric modality processed.|
|ImageType|[ImageType](#schemaimagetype)|true|none|Type of biometric image.|

<h2 id="tocS_Image">Image</h2>
<!-- backwards compatibility -->
<a id="schemaimage"></a>
<a id="schema_Image"></a>
<a id="tocSimage"></a>
<a id="tocsimage"></a>

```json
{
  "Modality": "Finger",
  "ImageType": "FlatLeftThumb",
  "ImageId": "image-abc123",
  "ImageData": "iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAIAAACQd1PeAAAAEElEQVR4nGJiYGAABAAA//8ADAADcZGLFwAAAABJRU5ErkJggg=="
}

```

### Properties

allOf

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|*anonymous*|[BiometricData](#schemabiometricdata)|false|none|Base object for biometric data, including modality and image type.|

and

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|*anonymous*|object|false|none|none|
|» ImageId|string|false|none|Optional client-provided ID for the image, used to correlate multiple templates from the same image.|
|» ImageData|string(byte)|true|none|The image data, encoded in base64 PNG format.|

<h2 id="tocS_Template">Template</h2>
<!-- backwards compatibility -->
<a id="schematemplate"></a>
<a id="schema_Template"></a>
<a id="tocStemplate"></a>
<a id="tocstemplate"></a>

```json
{
  "Modality": "Finger",
  "ImageType": "FlatLeftThumb",
  "TemplateId": "target-001",
  "Template": "dGhpcyBzZW50ZW5jZSBpcyBhbiBleGFtcGxlIHRlbXBsYXRlLi4K",
  "Error": "string"
}

```

### Properties

allOf

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|*anonymous*|[BiometricData](#schemabiometricdata)|false|none|Base object for biometric data, including modality and image type.|

and

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|*anonymous*|object|false|none|none|
|» TemplateId|string|true|none|Unique identifier for the template, assigned by the client.|
|» Template|string(byte)|true|none|The template data encoded as a base64 string.<br>The data string shall not exceed 1 MB.|
|» Error|string|false|none|An error produced from creating the biometric data|

<h2 id="tocS_Comparison">Comparison</h2>
<!-- backwards compatibility -->
<a id="schemacomparison"></a>
<a id="schema_Comparison"></a>
<a id="tocScomparison"></a>
<a id="tocscomparison"></a>

```json
{
  "TargetTemplateId": "target-001",
  "Score": 8734,
  "Error": "string"
}

```

A similarity score result tied to a specific target template.

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|TargetTemplateId|string|true|none|Target TemplateId for which this comparison score applies.|
|Score|number(float)|true|none|Similarity score produced by the algorithm.|
|Error|string|false|none|An error produced from creating the biometric data|

<h2 id="tocS_VerifyRequest">VerifyRequest</h2>
<!-- backwards compatibility -->
<a id="schemaverifyrequest"></a>
<a id="schema_VerifyRequest"></a>
<a id="tocSverifyrequest"></a>
<a id="tocsverifyrequest"></a>

```json
{
  "ProbeTemplate": {
    "Modality": "Finger",
    "ImageType": "FlatLeftThumb",
    "TemplateId": "target-001",
    "Template": "dGhpcyBzZW50ZW5jZSBpcyBhbiBleGFtcGxlIHRlbXBsYXRlLi4K",
    "Error": "string"
  },
  "TargetTemplateList": [
    {
      "Modality": "Finger",
      "ImageType": "FlatLeftThumb",
      "TemplateId": "target-001",
      "Template": "dGhpcyBzZW50ZW5jZSBpcyBhbiBleGFtcGxlIHRlbXBsYXRlLi4K",
      "Error": "string"
    }
  ]
}

```

A probe template and a list of target templates for verification.

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|ProbeTemplate|[Template](#schematemplate)|true|none|none|
|TargetTemplateList|[[Template](#schematemplate)]|true|none|none|

<h2 id="tocS_Info">Info</h2>
<!-- backwards compatibility -->
<a id="schemainfo"></a>
<a id="schema_Info"></a>
<a id="tocSinfo"></a>
<a id="tocsinfo"></a>

```json
{
  "AlgorithmName": "AlwaysTrue",
  "AlgorithmVersion": "1.0.1",
  "CompanyName": "MdTF",
  "TechnicalContactEmail": "john@mdtf.org",
  "RecommendedCPUs": 0.5,
  "RecommendedMem": 512,
  "Thresholds": {
    "1:500": 7543,
    "1:1e3": 8730.5,
    "1:1e4": 9321.2,
    "1:1e5": 9863.7,
    "1:1e6": 9972.8
  },
  "SupportedOperations": [
    "Verify"
  ],
  "SupportedModalities": [
    "Finger"
  ],
  "SupportedImageTypes": [
    "FlatLeftThumb"
  ]
}

```

Information describing the algorithm and supported server capabilities.

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|AlgorithmName|string|true|none|none|
|AlgorithmVersion|string|true|none|none|
|CompanyName|string|true|none|none|
|TechnicalContactEmail|string|true|none|none|
|RecommendedCPUs|number(float)|true|none|none|
|RecommendedMem|integer|true|none|none|
|Thresholds|object|true|none|Threshold values mapped to false match rates.|
|» 1:500|string|true|none|none|
|» 1:1e3|string|true|none|none|
|» 1:1e4|string|true|none|none|
|» 1:1e5|string|true|none|none|
|» 1:1e6|string|true|none|none|
|SupportedOperations|[[Operation](#schemaoperation)]|true|none|[An operation supported by the server.<br>]|
|SupportedModalities|[[Modality](#schemamodality)]|true|none|[Type of biometric modality processed.<br>]|
|SupportedImageTypes|[[ImageType](#schemaimagetype)]|true|none|[Type of biometric image.<br>]|

<h2 id="tocS_Error">Error</h2>
<!-- backwards compatibility -->
<a id="schemaerror"></a>
<a id="schema_Error"></a>
<a id="tocSerror"></a>
<a id="tocserror"></a>

```json
{
  "message": "string"
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|message|string|true|none|none|

