---
title: The Maryland Test Facility Acquisition System Interface v2.0.0
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

<h1 id="the-maryland-test-facility-acquisition-system-interface">The Maryland Test Facility Acquisition System Interface v2.0.0</h1>

> Scroll down for code samples, example requests and responses. Select a language for code samples from the tabs above or the mobile navigation menu.

Application Programming Interface for sending acquisition records to the MdTF Backend.

Base URLs:

* <a href="https://api.mdtf.org/">https://api.mdtf.org/</a>

    * **host** -  Default: ids.lab

    * **port** -  Default: 8080

        * 8080

Email: <a href="mailto:rally@mdtf.org">The MdTF</a> Web: <a href="https://mdtf.org">The MdTF</a> 
 License: Copyright (c) 2022, The Maryland Test Facility

<h1 id="the-maryland-test-facility-acquisition-system-interface-image-submission">Image Submission</h1>

Endpoints for submitting images to an ongoing transaction.

## Associate a face image capture with the ongoing transaction.

> Code samples

```shell
# You can also use wget
curl -X POST https://api.mdtf.org/v1/face-captures \
  -H 'Content-Type: application/json' \
  -H 'Accept: application/json'

```

```javascript
const inputBody = '{
  "StationID": "Station_A",
  "FaceImageData": "iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAIAAACQd1PeAAAAEElEQVR4nGJiYGAABAAA//8ADAADcZGLFwAAAABJRU5ErkJggg=="
}';
const headers = {
  'Content-Type':'application/json',
  'Accept':'application/json'
};

fetch('https://api.mdtf.org/v1/face-captures',
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
    req, err := http.NewRequest("POST", "https://api.mdtf.org/v1/face-captures", data)
    req.Header = headers

    client := &http.Client{}
    resp, err := client.Do(req)
    // ...
}

```

`POST /v1/face-captures`

Acquisition system providers may submit face image data through this endpoint. Submitted image captures will be
stored and associated with the ongoing transaction. Multiple images may be submitted during a transaction.

> Body parameter

```json
{
  "StationID": "Station_A",
  "FaceImageData": "iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAIAAACQd1PeAAAAEElEQVR4nGJiYGAABAAA//8ADAADcZGLFwAAAABJRU5ErkJggg=="
}
```

<h3 id="associate-a-face-image-capture-with-the-ongoing-transaction.-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|body|body|[FaceCapture](#schemafacecapture)|true|Face image collected as part of a transaction.|

> Example responses

> 400 Response

```json
{
  "Code": 0,
  "Message": "string",
  "ErrorString": "string"
}
```

<h3 id="associate-a-face-image-capture-with-the-ongoing-transaction.-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|Successfully associated the image capture with the ongoing transaction.|None|
|400|[Bad Request](https://tools.ietf.org/html/rfc7231#section-6.5.1)|The request is malformed.|[Error](#schemaerror)|
|404|[Not Found](https://tools.ietf.org/html/rfc7231#section-6.5.4)|No ongoing transaction was found for the specified StationID.|[Error](#schemaerror)|
|410|[Gone](https://tools.ietf.org/html/rfc7231#section-6.5.9)|The test volunteer has exited the station and the transaction has ended.|[Error](#schemaerror)|
|429|[Too Many Requests](https://tools.ietf.org/html/rfc6585#section-4)|Too many requests for the ongoing transaction.|[Error](#schemaerror)|
|500|[Internal Server Error](https://tools.ietf.org/html/rfc7231#section-6.6.1)|This capture failed because of a server side issue.|[Error](#schemaerror)|

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
|Code|integer|false|none|The error code.|
|Message|string|false|none|A description of the returned code.|
|ErrorString|string|false|none|A detailed description of the error that occurred.|

