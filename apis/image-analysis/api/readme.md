---
title: The Maryland Test Facility Image Analysis Interface v1.0.0
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

<h1 id="the-maryland-test-facility-image-analysis-interface">The Maryland Test Facility Image Analysis Interface v1.0.0</h1>

> Scroll down for code samples, example requests and responses. Select a language for code samples from the tabs above or the mobile navigation menu.

Application programming interface for receiving image analysis requests.

Base URLs:

* <a href="http://172.17.0.2:8080">http://172.17.0.2:8080</a>

Email: <a href="mailto:info@mdtf.org">MdTF</a> Web: <a href="https://mdtf.org">MdTF</a> 
License: <a href="https://raw.githubusercontent.com/TheMdTF/mdtf-public/master/LICENSE.md">IDSL API License</a>

<h1 id="the-maryland-test-facility-image-analysis-interface-analysis-operations">Analysis Operations</h1>

## Returns an object containing the results of analysis for the provided biometric image.

<a id="opIdanalyze_image"></a>

> Code samples

```shell
# You can also use wget
curl -X POST http://172.17.0.2:8080/v1/analyze-image \
  -H 'Content-Type: application/json' \
  -H 'Accept: application/json'

```

```javascript
const inputBody = '{
  "ImageData": "iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAIAAACQd1PeAAAAEElEQVR4nGJiYGAABAAA//8ADAADcZGLFwAAAABJRU5ErkJggg=="
}';
const headers = {
  'Content-Type':'application/json',
  'Accept':'application/json'
};

fetch('http://172.17.0.2:8080/v1/analyze-image',
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
    req, err := http.NewRequest("POST", "http://172.17.0.2:8080/v1/analyze-image", data)
    req.Header = headers

    client := &http.Client{}
    resp, err := client.Do(req)
    // ...
}

```

`POST /v1/analyze-image`

This endpoint accepts a base64 encoded PNG image and returns an object containing the output of a generic image analysis routine on that image.

> Body parameter

```json
{
  "ImageData": "iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAIAAACQd1PeAAAAEElEQVR4nGJiYGAABAAA//8ADAADcZGLFwAAAABJRU5ErkJggg=="
}
```

<h3 id="returns-an-object-containing-the-results-of-analysis-for-the-provided-biometric-image.-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|body|body|[Image](#schemaimage)|true|The image that is being submitted for analysis.|

#### Detailed descriptions

**body**: The image that is being submitted for analysis.

> Example responses

> 200 Response

```json
{
  "AnalysisScore": "5.0002",
  "NormalizedScore": "0.50002",
  "ErrorLog": ""
}
```

<h3 id="returns-an-object-containing-the-results-of-analysis-for-the-provided-biometric-image.-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|Successful Response|[AnalysisResult](#schemaanalysisresult)|
|400|[Bad Request](https://tools.ietf.org/html/rfc7231#section-6.5.1)|Bad Request|[RequestError](#schemarequesterror)|
|500|[Internal Server Error](https://tools.ietf.org/html/rfc7231#section-6.6.1)|Internal Server Error|[ServerError](#schemaservererror)|

<aside class="success">
This operation does not require authentication
</aside>

<h1 id="the-maryland-test-facility-image-analysis-interface-algorithm-information">Algorithm Information</h1>

## Returns basic information for the algorithm.

<a id="opIdinfo"></a>

> Code samples

```shell
# You can also use wget
curl -X GET http://172.17.0.2:8080/v1/info \
  -H 'Accept: application/json'

```

```javascript

const headers = {
  'Accept':'application/json'
};

fetch('http://172.17.0.2:8080/v1/info',
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
    req, err := http.NewRequest("GET", "http://172.17.0.2:8080/v1/info", data)
    req.Header = headers

    client := &http.Client{}
    resp, err := client.Do(req)
    // ...
}

```

`GET /v1/info`

This endpoint returns some basic information about the image analysis algorithm.

> Example responses

> 200 Response

```json
{
  "AlgorithmName": "OpenFace",
  "AlgorithmType": "Face",
  "AlgorithmVersion": "1.0.1",
  "CompanyName": "MdTF",
  "ImageDataset": "NIST_MEDS",
  "RecommendedCPUs": 4,
  "RecommendedMem": 2048,
  "TechnicalContactEmail": "info@mdtf.org"
}
```

<h3 id="returns-basic-information-for-the-algorithm.-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|Successful Response|[Info](#schemainfo)|
|500|[Internal Server Error](https://tools.ietf.org/html/rfc7231#section-6.6.1)|Internal Server Error|[ServerError](#schemaservererror)|

<aside class="success">
This operation does not require authentication
</aside>

# Schemas

<h2 id="tocS_Image">Image</h2>
<!-- backwards compatibility -->
<a id="schemaimage"></a>
<a id="schema_Image"></a>
<a id="tocSimage"></a>
<a id="tocsimage"></a>

```json
{
  "ImageData": "iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAIAAACQd1PeAAAAEElEQVR4nGJiYGAABAAA//8ADAADcZGLFwAAAABJRU5ErkJggg=="
}

```

Data transfer object for an image.

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|ImageData|string|true|none|The captured image data in PNG format, encoded as a base64 string. The data string shall not exceed 10MB.|

<h2 id="tocS_Info">Info</h2>
<!-- backwards compatibility -->
<a id="schemainfo"></a>
<a id="schema_Info"></a>
<a id="tocSinfo"></a>
<a id="tocsinfo"></a>

```json
{
  "AlgorithmName": "OpenFace",
  "AlgorithmType": "Face",
  "AlgorithmVersion": "1.0.1",
  "CompanyName": "MdTF",
  "ImageDataset": "NIST_MEDS",
  "RecommendedCPUs": 4,
  "RecommendedMem": 2048,
  "TechnicalContactEmail": "info@mdtf.org"
}

```

Basic information describing the algorithm.

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|AlgorithmName|string|true|none|Name of algorithm|
|AlgorithmType|string|true|none|A string enum describing the type of biometric images the algorithm is meant to process|
|AlgorithmVersion|string|true|none|Algorithm version identifier|
|CompanyName|string|true|none|Name of the company which produces the algorithm|
|ImageDataset|string|true|none|A string enum to select an MdTF dataset of biometric images for this submission to analyze|
|RecommendedCPUs|number(integer)|true|none|The recommended allocation of CPUs for the deployed docker container.|
|RecommendedMem|number(integer)|true|none|The recommended allocation of memory (MB) for the deployed docker container.|
|TechnicalContactEmail|string|true|none|The email address of an engineer or other technical resource to contact in the event of an error running your service. This field may be left blank if desired.|

#### Enumerated Values

|Property|Value|
|---|---|
|AlgorithmType|Face|
|AlgorithmType|Finger|
|AlgorithmType|Iris|
|ImageDataset|NIST_MEDS|

<h2 id="tocS_AnalysisResult">AnalysisResult</h2>
<!-- backwards compatibility -->
<a id="schemaanalysisresult"></a>
<a id="schema_AnalysisResult"></a>
<a id="tocSanalysisresult"></a>
<a id="tocsanalysisresult"></a>

```json
{
  "AnalysisScore": "5.0002",
  "NormalizedScore": "0.50002",
  "ErrorLog": ""
}

```

Object containing the results of a generic image analysis algorithm on a single image. Must include at least 1 and not more than 8 string properties in a JSON object with a depth of 1. The full contents shall not exceed 512 Bytes.

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|**additionalProperties**|string|false|none|The output of a generic image analysis algorithm on the image provided. Contents MAY NOT include portions of the raw image data. Note that the example fields are just that - examples.  In this example, 3 of 8 output property fields have been utilized. Field name and output values should be defined based on useful output for evaluating the performance of your image analysis algorithm.|

<h2 id="tocS_RequestError">RequestError</h2>
<!-- backwards compatibility -->
<a id="schemarequesterror"></a>
<a id="schema_RequestError"></a>
<a id="tocSrequesterror"></a>
<a id="tocsrequesterror"></a>

```json
"Unable to decode image data as a PNG."

```

Relevant and concise diagnostic information in the event of a client side error (e.g. malformed requests, or invalid image encoding).

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|*anonymous*|string|false|none|Relevant and concise diagnostic information in the event of a client side error (e.g. malformed requests, or invalid image encoding).|

<h2 id="tocS_ServerError">ServerError</h2>
<!-- backwards compatibility -->
<a id="schemaservererror"></a>
<a id="schema_ServerError"></a>
<a id="tocSservererror"></a>
<a id="tocsservererror"></a>

```json
"Error encoding response."

```

Relevant and concise diagnostic information in the event of a server side error (e.g. invalid license, or failure to generate a template).

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|*anonymous*|string|false|none|Relevant and concise diagnostic information in the event of a server side error (e.g. invalid license, or failure to generate a template).|

