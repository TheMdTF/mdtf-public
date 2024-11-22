<!-- Generator: Widdershins v4.0.1 -->

<h1 id="the-maryland-test-facility-acquisition-system-interface">The Maryland Test Facility Acquisition System Interface v2.1.0</h1>

> Scroll down for code samples, example requests and responses. Select a language for code samples from the tabs above or the mobile navigation menu.

Application Programming Interface for sending acquisition records to the MdTF Backend.

Base URLs:

* <a href="http://acquisition-demo.mdtf.org">http://acquisition-demo.mdtf.org</a>

Email: <a href="mailto:rally@mdtf.org">The MdTF</a> Web: <a href="https://mdtf.org">The MdTF</a> 
License: <a href="https://raw.githubusercontent.com/TheMdTF/mdtf-public/master/LICENSE.md">IDSL API License</a>

<h1 id="the-maryland-test-facility-acquisition-system-interface-image-submission">Image Submission</h1>

Endpoints for submitting images to an ongoing transaction.

## Submit a captured image.

<a id="opIdCapture"></a>

> Code samples

```shell
# You can also use wget
curl -X POST http://acquisition-demo.mdtf.org/v1/capture \
  -H 'Content-Type: application/json' \
  -H 'Accept: application/json'

```

```javascript
const inputBody = '{
  "StationID": "Station_A",
  "Data": "iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAIAAACQd1PeAAAAEElEQVR4nGJiYGAABAAA//8ADAADcZGLFwAAAABJRU5ErkJggg=="
}';
const headers = {
  'Content-Type':'application/json',
  'Accept':'application/json'
};

fetch('http://acquisition-demo.mdtf.org/v1/capture',
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
    req, err := http.NewRequest("POST", "http://acquisition-demo.mdtf.org/v1/capture", data)
    req.Header = headers

    client := &http.Client{}
    resp, err := client.Do(req)
    // ...
}

```

`POST /v1/capture`

Acquisition system providers may submit image data through this endpoint.

> Body parameter

```json
{
  "StationID": "Station_A",
  "Data": "iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAIAAACQd1PeAAAAEElEQVR4nGJiYGAABAAA//8ADAADcZGLFwAAAABJRU5ErkJggg=="
}
```

<h3 id="submit-a-captured-image.-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|body|body|[Capture](#schemacapture)|true|Face image collected as part of a transaction.|

> Example responses

> 400 Response

```json
{
  "Code": 0,
  "Message": "string",
  "ErrorString": "string"
}
```

<h3 id="submit-a-captured-image.-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|Successfully associated the image capture with the ongoing transaction.|None|
|400|[Bad Request](https://tools.ietf.org/html/rfc7231#section-6.5.1)|The request is malformed.|[Error](#schemaerror)|
|404|[Not Found](https://tools.ietf.org/html/rfc7231#section-6.5.4)|No ongoing transaction was found for the specified StationID.|[Error](#schemaerror)|
|500|[Internal Server Error](https://tools.ietf.org/html/rfc7231#section-6.6.1)|This capture failed because of a server side issue.|[Error](#schemaerror)|

<aside class="success">
This operation does not require authentication
</aside>

# Schemas

<h2 id="tocS_Capture">Capture</h2>
<!-- backwards compatibility -->
<a id="schemacapture"></a>
<a id="schema_Capture"></a>
<a id="tocScapture"></a>
<a id="tocscapture"></a>

```json
{
  "StationID": "Station_A",
  "Data": "iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAIAAACQd1PeAAAAEElEQVR4nGJiYGAABAAA//8ADAADcZGLFwAAAABJRU5ErkJggg=="
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|StationID|string|true|none|The unique identifier for the station that the image was captured at.|
|Data|string(byte)|true|none|The raw bytes of the image data. The supported formats are jpeg and png. This data should be marshaled in this object as a base64 encoded string.|

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
|Code|integer|true|none|The error code.|
|Message|string|true|none|A description of the returned code.|
|ErrorString|string|true|none|A detailed description of the error that occurred.|

