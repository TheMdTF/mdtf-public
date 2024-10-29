---
title: The Maryland Test Facility Acquisition System Interface v1.2.0
language_tabs:
  - shell: Shell
  - javascript: JavaScript
  - go: Go
language_clients:
  - shell: ""
  - javascript: ""
  - go: ""
toc_footers: []
includes: []
search: true
highlight_theme: darkula
headingLevel: 2

---

<!-- Generator: Widdershins v4.0.1 -->

<h1 id="the-maryland-test-facility-acquisition-system-interface">The Maryland Test Facility Acquisition System Interface v1.2.0</h1>

> Scroll down for code samples, example requests and responses. Select a language for code samples from the tabs above or the mobile navigation menu.

Application Programming Interface for sending acquisition records to the MdTF Backend.

Base URLs:

* <a href="http://api.mdtf.org">http://api.mdtf.org</a>

Email: <a href="mailto:rally@mdtf.org">The MdTF</a> Web: <a href="https://mdtf.org">The MdTF</a> 
 License: Copyright (c) 2021, The Maryland Test Facility

<h1 id="the-maryland-test-facility-acquisition-system-interface-image-submission">Image Submission</h1>

Endpoints for submitting images to an ongoing transaction.

## Associate a face image capture with the ongoing transaction

> Code samples

```shell
# You can also use wget
curl -X POST http://api.mdtf.org/v1/face-captures \
  -H 'Accept: application/json'

```

```javascript

const headers = {
  'Accept':'application/json'
};

fetch('http://api.mdtf.org/v1/face-captures',
{
  method: 'POST',

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
    req, err := http.NewRequest("POST", "http://api.mdtf.org/v1/face-captures", data)
    req.Header = headers

    client := &http.Client{}
    resp, err := client.Do(req)
    // ...
}

```

`POST /v1/face-captures`

Acquisition System Providers may submit face image data through this endpoint. Submitted image captures will be
stored and associated with the ongoing transaction. Multiple images may be submitted during a transaction.

> Example responses

> 400 Response

```json
{
  "Code": 0,
  "Message": "string",
  "ErrorString": "string"
}
```

<h3 id="associate-a-face-image-capture-with-the-ongoing-transaction-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|Successfully associated the image capture with the ongoing transaction.|None|
|400|[Bad Request](https://tools.ietf.org/html/rfc7231#section-6.5.1)|The request is malformed.|[Error](#schemaerror)|
|404|[Not Found](https://tools.ietf.org/html/rfc7231#section-6.5.4)|No ongoing transaction was found for the specified StationID.|[Error](#schemaerror)|
|410|[Gone](https://tools.ietf.org/html/rfc7231#section-6.5.9)|The test volunteer has exited the station and the transaction has ended.|[Error](#schemaerror)|
|429|[Too Many Requests](https://tools.ietf.org/html/rfc6585#section-4)|Too many requests for the ongoing transaction.|[Error](#schemaerror)|
|500|[Internal Server Error](https://tools.ietf.org/html/rfc7231#section-6.6.1)|This capture failed because of a server side issue|[Error](#schemaerror)|
|default|Default|An error occurred.|[Error](#schemaerror)|

<aside class="success">
This operation does not require authentication
</aside>

## Associate an iris image capture (or pair of image captures) with the ongoing transaction

> Code samples

```shell
# You can also use wget
curl -X POST http://api.mdtf.org/v1/iris-captures \
  -H 'Accept: application/json'

```

```javascript

const headers = {
  'Accept':'application/json'
};

fetch('http://api.mdtf.org/v1/iris-captures',
{
  method: 'POST',

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
    req, err := http.NewRequest("POST", "http://api.mdtf.org/v1/iris-captures", data)
    req.Header = headers

    client := &http.Client{}
    resp, err := client.Do(req)
    // ...
}

```

`POST /v1/iris-captures`

Acquisition System Providers may submit iris image data through this endpoint. Submitted image captures will be
stored and associated with the ongoing transaction. Multiple images may be submitted during a transaction.

> Example responses

> 400 Response

```json
{
  "Code": 0,
  "Message": "string",
  "ErrorString": "string"
}
```

<h3 id="associate-an-iris-image-capture-(or-pair-of-image-captures)-with-the-ongoing-transaction-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|Successfully associated the image capture(s) with the ongoing transaction.|None|
|400|[Bad Request](https://tools.ietf.org/html/rfc7231#section-6.5.1)|The request is malformed.|[Error](#schemaerror)|
|404|[Not Found](https://tools.ietf.org/html/rfc7231#section-6.5.4)|No ongoing transaction was found for the specified StationID.|[Error](#schemaerror)|
|410|[Gone](https://tools.ietf.org/html/rfc7231#section-6.5.9)|The test volunteer has exited the station and the transaction has ended.|[Error](#schemaerror)|
|429|[Too Many Requests](https://tools.ietf.org/html/rfc6585#section-4)|Too many requests for the ongoing transaction.|[Error](#schemaerror)|
|500|[Internal Server Error](https://tools.ietf.org/html/rfc7231#section-6.6.1)|This capture failed because of a server side issue|[Error](#schemaerror)|
|default|Default|An error occurred.|[Error](#schemaerror)|

<aside class="success">
This operation does not require authentication
</aside>

## Associate up to 10 finger captures with the ongoing transaction

> Code samples

```shell
# You can also use wget
curl -X POST http://api.mdtf.org/v1/finger-captures \
  -H 'Accept: application/json'

```

```javascript

const headers = {
  'Accept':'application/json'
};

fetch('http://api.mdtf.org/v1/finger-captures',
{
  method: 'POST',

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
    req, err := http.NewRequest("POST", "http://api.mdtf.org/v1/finger-captures", data)
    req.Header = headers

    client := &http.Client{}
    resp, err := client.Do(req)
    // ...
}

```

`POST /v1/finger-captures`

Acquisition System Providers may submit finger image data through this endpoint. Submitted image captures will be stored and associated with the ongoing transaction. Multiple images may be submitted during a transaction.

> Example responses

> 400 Response

```json
{
  "Code": 0,
  "Message": "string",
  "ErrorString": "string"
}
```

<h3 id="associate-up-to-10-finger-captures-with-the-ongoing-transaction-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|Successfully associated the image capture(s) with the ongoing transaction.|None|
|400|[Bad Request](https://tools.ietf.org/html/rfc7231#section-6.5.1)|The request is malformed.|[Error](#schemaerror)|
|404|[Not Found](https://tools.ietf.org/html/rfc7231#section-6.5.4)|No ongoing transaction was found for the specified StationID.|[Error](#schemaerror)|
|410|[Gone](https://tools.ietf.org/html/rfc7231#section-6.5.9)|The test volunteer has exited the station and the transaction has ended.|[Error](#schemaerror)|
|429|[Too Many Requests](https://tools.ietf.org/html/rfc6585#section-4)|Too many requests for the ongoing transaction.|[Error](#schemaerror)|
|500|[Internal Server Error](https://tools.ietf.org/html/rfc7231#section-6.6.1)|This capture failed because of a server side issue|[Error](#schemaerror)|
|default|Default|An error occurred.|[Error](#schemaerror)|

<aside class="success">
This operation does not require authentication
</aside>

# Schemas

<h2 id="tocS_FaceCapture">FaceCapture</h2>
<!-- backwards compatibility -->
<a id="schemafacecapture"></a>
<a id="schema_FaceCapture"></a>
<a id="tocSfacecapture"></a>
<a id="tocsfacecapture"></a>

```json
{
  "StationID": "Station_A",
  "FaceImageData": "iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAIAAACQd1PeAAAAEElEQVR4nGJiYGAABAAA//8ADAADcZGLFwAAAABJRU5ErkJggg=="
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|StationID|string|true|none|The unique identifier for the station that the image was captured at.|
|FaceImageData|string|true|none|The captured face image data in PNG format, encoded as a base64 string. The data string shall not exceed 1MB.|

<h2 id="tocS_IrisCapture">IrisCapture</h2>
<!-- backwards compatibility -->
<a id="schemairiscapture"></a>
<a id="schema_IrisCapture"></a>
<a id="tocSiriscapture"></a>
<a id="tocsiriscapture"></a>

```json
{
  "StationID": "Station_A",
  "LeftIrisImageData": "iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAIAAACQd1PeAAAAEElEQVR4nGJiYGAABAAA//8ADAADcZGLFwAAAABJRU5ErkJggg==",
  "RightIrisImageData": "iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAIAAACQd1PeAAAAEElEQVR4nGJiYGAABAAA//8ADAADcZGLFwAAAABJRU5ErkJggg=="
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|StationID|string|true|none|The unique identifier for the station that the image was captured at.|
|LeftIrisImageData|string|false|none|The captured left iris image data, encoded as a base64 string. The data string shall not exceed 512kB. At least one iris image data field MUST be filled.|
|RightIrisImageData|string|false|none|The captured right iris image data, encoded as a base64 string. The data string shall not exceed 512kB. At least one iris image data field MUST be filled.|

<h2 id="tocS_FingerCapture">FingerCapture</h2>
<!-- backwards compatibility -->
<a id="schemafingercapture"></a>
<a id="schema_FingerCapture"></a>
<a id="tocSfingercapture"></a>
<a id="tocsfingercapture"></a>

```json
{
  "StationID": "Station_A",
  "FingerImages": [
    {
      "FingerType": "RightIndexFinger",
      "FingerImageData": "iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAIAAACQd1PeAAAAEElEQVR4nGJiYGAABAAA//8ADAADcZGLFwAAAABJRU5ErkJggg=="
    },
    {
      "FingerType": "RightMiddleFinger",
      "FingerImageData": "iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAIAAACQd1PeAAAAEElEQVR4nGJiYGAABAAA//8ADAADcZGLFwAAAABJRU5ErkJggg=="
    }
  ]
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|StationID|string|true|none|The unique identifier for the station that the image was captured at.|
|FingerImages|[[FingerImage](#schemafingerimage)]|false|none|The captured finger(s). A minimum of one (1) FingerImage must be submitted with a request. No more than ten (10) FingerImages may be submitted with a request. Each submitted FingerImage must have a distinct type (i.e. there cannot be two "RightIndexFingers" in a request).|

<h2 id="tocS_FingerImage">FingerImage</h2>
<!-- backwards compatibility -->
<a id="schemafingerimage"></a>
<a id="schema_FingerImage"></a>
<a id="tocSfingerimage"></a>
<a id="tocsfingerimage"></a>

```json
{
  "FingerType": "RightThumb",
  "FingerImageData": "iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAIAAACQd1PeAAAAEElEQVR4nGJiYGAABAAA//8ADAADcZGLFwAAAABJRU5ErkJggg=="
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|FingerType|[FingerType](#schemafingertype)|true|none|The allowed string names for finger images.|
|FingerImageData|string|false|none|The captured finger image data, encoded as a base64 string. The data string shall not exceed 100kB.|

<h2 id="tocS_Error">Error</h2>
<!-- backwards compatibility -->
<a id="schemaerror"></a>
<a id="schema_Error"></a>
<a id="tocSerror"></a>
<a id="tocserror"></a>

```json
{
  "Code": 0,
  "Message": "string",
  "ErrorString": "string"
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|Code|integer(int32)|false|none|The error code.|
|Message|string|false|none|A description of the returned code.|
|ErrorString|string|false|none|A detailed description of the error that occurred.|

<h2 id="tocS_FingerType">FingerType</h2>
<!-- backwards compatibility -->
<a id="schemafingertype"></a>
<a id="schema_FingerType"></a>
<a id="tocSfingertype"></a>
<a id="tocsfingertype"></a>

```json
"RightThumb"

```

The allowed string names for finger images.

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|*anonymous*|string|false|none|The allowed string names for finger images.|

#### Enumerated Values

|Property|Value|
|---|---|
|*anonymous*|RightThumb|
|*anonymous*|RightIndexFinger|
|*anonymous*|RightMiddleFinger|
|*anonymous*|RightRingFinger|
|*anonymous*|RightLittleFinger|
|*anonymous*|LeftThumb|
|*anonymous*|LeftIndexFinger|
|*anonymous*|LeftMiddleFinger|
|*anonymous*|LeftRingFinger|
|*anonymous*|LeftLittleFinger|

