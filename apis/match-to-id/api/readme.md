<!-- Generator: Widdershins v4.0.1 -->

<h1 id="the-maryland-test-facility-match-to-id-interface">The Maryland Test Facility Match-to-ID Interface v0.0.1</h1>

> Scroll down for code samples, example requests and responses. Select a language for code samples from the tabs above or the mobile navigation menu.

This document specifies the API requirements for MdTF testing of algorithms that match facial biometric samples to identity document images (``match-to-id'').  Match-to-ID testing at the MdTF is supported by the Department of Homeland Security, Science and Technology Directorate (DHS S&T) as part of the Remote Identity Validation Technology Demonstration (RIVTD).  For more information please visit [https://mdtf.org](https://mdtf.org) and [https://www.dhs.gov/science-and-technology/BI-TC](https://www.dhs.gov/science-and-technology/BI-TC)

Email: <a href="mailto:info@mdtf.org">The MdTF</a> Web: <a href="https://mdtf.org">The MdTF</a> 
License: <a href="https://raw.githubusercontent.com/TheMdTF/mdtf-public/master/LICENSE.md">IDSL API License</a>

<h1 id="the-maryland-test-facility-match-to-id-interface-biometric-operations">Biometric Operations</h1>

## Generate a template from the provided facial biometric sample or the identity document image.

<a id="opIdGenerateTemplate"></a>

> Code samples

```shell
# You can also use wget
curl -X POST /v1/create-template \
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

fetch('/v1/create-template',
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
    req, err := http.NewRequest("POST", "/v1/create-template", data)
    req.Header = headers

    client := &http.Client{}
    resp, err := client.Do(req)
    // ...
}

```

`POST /v1/create-template`

This endpoint accepts a base64 encoded PNG or JPG and attempts to perform a feature extraction operation producing a
single template.

> Body parameter

```json
{
  "ImageData": "iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAIAAACQd1PeAAAAEElEQVR4nGJiYGAABAAA//8ADAADcZGLFwAAAABJRU5ErkJggg=="
}
```

<h3 id="generate-a-template-from-the-provided-facial-biometric-sample-or-the-identity-document-image.-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|body|body|[Image](#schemaimage)|true|The facial biometric sample or identity document that is being submitted.  Note there is no request body information|

#### Detailed descriptions

**body**: The facial biometric sample or identity document that is being submitted.  Note there is no request body information
to distinguish between the submission of these two types (facial sample vs. ID document).  Algorithm providers must 
make this distinction using internal logic.

> Example responses

> 200 Response

```json
{
  "Template": "dGhpcyBzZW50ZW5jZSBpcyBhbiBleGFtcGxlIHRlbXBsYXRlLi4K"
}
```

<h3 id="generate-a-template-from-the-provided-facial-biometric-sample-or-the-identity-document-image.-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|Successful Response|[Template](#schematemplate)|
|400|[Bad Request](https://tools.ietf.org/html/rfc7231#section-6.5.1)|Bad Request|[RequestError](#schemarequesterror)|
|500|[Internal Server Error](https://tools.ietf.org/html/rfc7231#section-6.6.1)|Internal Server Error|[ServerError](#schemaservererror)|

<aside class="success">
This operation does not require authentication
</aside>

## Compare a single template to a list of target templates.

<a id="opIdCompareTemplateList"></a>

> Code samples

```shell
# You can also use wget
curl -X POST /v1/compare-list \
  -H 'Content-Type: application/json' \
  -H 'Accept: application/json'

```

```javascript
const inputBody = '{
  "SingleTemplate": {
    "Template": "dGhpcyBzZW50ZW5jZSBpcyBhbiBleGFtcGxlIHRlbXBsYXRlLi4K"
  },
  "TargetTemplateList": [
    {
      "Template": "dGhpcyBzZW50ZW5jZSBpcyBhbiBleGFtcGxlIHRlbXBsYXRlLi4K"
    }
  ]
}';
const headers = {
  'Content-Type':'application/json',
  'Accept':'application/json'
};

fetch('/v1/compare-list',
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
    req, err := http.NewRequest("POST", "/v1/compare-list", data)
    req.Header = headers

    client := &http.Client{}
    resp, err := client.Do(req)
    // ...
}

```

`POST /v1/compare-list`

This endpoint accepts a template and a list of templates. It compares the single template to every target template in the
provided list. The result is a list of Comparison objects that holds a similarity score for each comparison. 
The returned list of comparisons MUST contain the same number of elements AND be in the same order as the provided list
of templates.

> Body parameter

```json
{
  "SingleTemplate": {
    "Template": "dGhpcyBzZW50ZW5jZSBpcyBhbiBleGFtcGxlIHRlbXBsYXRlLi4K"
  },
  "TargetTemplateList": [
    {
      "Template": "dGhpcyBzZW50ZW5jZSBpcyBhbiBleGFtcGxlIHRlbXBsYXRlLi4K"
    }
  ]
}
```

<h3 id="compare-a-single-template-to-a-list-of-target-templates.-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|body|body|[CompareListRequest](#schemacomparelistrequest)|true|A single template object and a list of templates that it will be compared to.|

> Example responses

> 200 Response

```json
[
  {
    "Score": 8734
  }
]
```

<h3 id="compare-a-single-template-to-a-list-of-target-templates.-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|Successful Response|Inline|
|400|[Bad Request](https://tools.ietf.org/html/rfc7231#section-6.5.1)|Bad Request|[RequestError](#schemarequesterror)|
|500|[Internal Server Error](https://tools.ietf.org/html/rfc7231#section-6.6.1)|Internal Server Error|[ServerError](#schemaservererror)|

<h3 id="compare-a-single-template-to-a-list-of-target-templates.-responseschema">Response Schema</h3>

Status Code **200**

*An array of comparison results. This list MUST contain the same number of elements AND be in
the same order as the provided list of templates.
*

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|*anonymous*|[[Comparison](#schemacomparison)]|false|none|An array of comparison results. This list MUST contain the same number of elements AND be in<br>the same order as the provided list of templates.|
|» Score|number(float)|true|none|A similarity score, as produced by the algorithm.|

<aside class="success">
This operation does not require authentication
</aside>

<h1 id="the-maryland-test-facility-match-to-id-interface-algorithm-information">Algorithm Information</h1>

## Returns basic information for the algorithm.

<a id="opIdGetInfo"></a>

> Code samples

```shell
# You can also use wget
curl -X GET /v1/info \
  -H 'Accept: application/json'

```

```javascript

const headers = {
  'Accept':'application/json'
};

fetch('/v1/info',
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
    req, err := http.NewRequest("GET", "/v1/info", data)
    req.Header = headers

    client := &http.Client{}
    resp, err := client.Do(req)
    // ...
}

```

`GET /v1/info`

This endpoint returns some basic information about the algorithm.

> Example responses

> 200 Response

```json
{
  "AlgorithmName": "AlwaysTrue",
  "AlgorithmVersion": "1.0.1",
  "AlgorithmModality": "Face",
  "CompanyName": "MdTF",
  "TechnicalContactEmail": "john@mdtf.org",
  "RecommendedCPUs": 0.5,
  "RecommendedMem": 512,
  "Test": "MDTF_RIVTD_TRACK2",
  "Thresholds": {
    "1:500": "7543",
    "1:1e3": "8730",
    "1:1e4": "9321",
    "1:1e5": "9863",
    "1:1e6": "9972"
  }
}
```

<h3 id="returns-basic-information-for-the-algorithm.-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|info response object|[Info](#schemainfo)|
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
|ImageData|string(byte)|true|none|The captured image data in PNG or JPG format, encoded as a base64 string.|

<h2 id="tocS_Template">Template</h2>
<!-- backwards compatibility -->
<a id="schematemplate"></a>
<a id="schema_Template"></a>
<a id="tocStemplate"></a>
<a id="tocstemplate"></a>

```json
{
  "Template": "dGhpcyBzZW50ZW5jZSBpcyBhbiBleGFtcGxlIHRlbXBsYXRlLi4K"
}

```

Data transfer object for a template.

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|Template|string(byte)|true|none|The template data, encoded as a base64 string.|

<h2 id="tocS_Comparison">Comparison</h2>
<!-- backwards compatibility -->
<a id="schemacomparison"></a>
<a id="schema_Comparison"></a>
<a id="tocScomparison"></a>
<a id="tocscomparison"></a>

```json
{
  "Score": 8734
}

```

A similarity score for a template comparison operation.

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|Score|number(float)|true|none|A similarity score, as produced by the algorithm.|

<h2 id="tocS_CompareListRequest">CompareListRequest</h2>
<!-- backwards compatibility -->
<a id="schemacomparelistrequest"></a>
<a id="schema_CompareListRequest"></a>
<a id="tocScomparelistrequest"></a>
<a id="tocscomparelistrequest"></a>

```json
{
  "SingleTemplate": {
    "Template": "dGhpcyBzZW50ZW5jZSBpcyBhbiBleGFtcGxlIHRlbXBsYXRlLi4K"
  },
  "TargetTemplateList": [
    {
      "Template": "dGhpcyBzZW50ZW5jZSBpcyBhbiBleGFtcGxlIHRlbXBsYXRlLi4K"
    }
  ]
}

```

A single probe template object and a list of target templates that it will be compared to.

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|SingleTemplate|[Template](#schematemplate)|true|none|Data transfer object for a template.|
|TargetTemplateList|[[Template](#schematemplate)]|true|none|[Data transfer object for a template.<br>]|

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
  "AlgorithmModality": "Face",
  "CompanyName": "MdTF",
  "TechnicalContactEmail": "john@mdtf.org",
  "RecommendedCPUs": 0.5,
  "RecommendedMem": 512,
  "Test": "MDTF_RIVTD_TRACK2",
  "Thresholds": {
    "1:500": "7543",
    "1:1e3": "8730",
    "1:1e4": "9321",
    "1:1e5": "9863",
    "1:1e6": "9972"
  }
}

```

Basic information describing the algorithm.

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|AlgorithmName|string|true|none|Name of algorithm.|
|AlgorithmVersion|string|true|none|Algorithm version identifier.|
|AlgorithmModality|string|true|none|A string enum describing the type of biometric images the algorithm is meant to process.|
|CompanyName|string|true|none|Name of the Company which produced the algorithm.|
|TechnicalContactEmail|string|true|none|The email address of an engineer or other technical resource to contact in the event of<br>an error running your service.|
|RecommendedCPUs|number(float)|true|none|The recommended allocation of CPUs for the deployed docker container.|
|RecommendedMem|integer|true|none|The recommended allocation of memory (MB) for the deployed docker container.|
|Test|string|true|none|A string enum describing which collection event the algorithm is being submitted for.|
|Thresholds|object|true|none|A map of preset False Match Rates (FMR) to vendor-provided threshold values. Score values returned from calls to<br>v1/compare-list indicate a matching determination by the algorithm if they are greater than the provided threshold<br>value at the respective FMR.<br>Note that threshold values are floats stored as strings and cannot exceed a length of 10 characters.<br>There are 5 required thresholds.|
|» 1:500|string|true|none|none|
|» 1:1e3|string|true|none|none|
|» 1:1e4|string|true|none|none|
|» 1:1e5|string|true|none|none|
|» 1:1e6|string|true|none|none|

#### Enumerated Values

|Property|Value|
|---|---|
|AlgorithmModality|Face|
|Test|MDTF_RIVTD_TRACK2|

<h2 id="tocS_RequestError">RequestError</h2>
<!-- backwards compatibility -->
<a id="schemarequesterror"></a>
<a id="schema_RequestError"></a>
<a id="tocSrequesterror"></a>
<a id="tocsrequesterror"></a>

```json
"Unable to decode image data as a PNG or JPG."

```

Relevant and concise diagnostic information in the event of a client side error (e.g. malformed requests, or invalid
image encoding).

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|*anonymous*|string|false|none|Relevant and concise diagnostic information in the event of a client side error (e.g. malformed requests, or invalid<br>image encoding).|

<h2 id="tocS_ServerError">ServerError</h2>
<!-- backwards compatibility -->
<a id="schemaservererror"></a>
<a id="schema_ServerError"></a>
<a id="tocSservererror"></a>
<a id="tocsservererror"></a>

```json
"The internal license has expired."

```

Relevant and concise diagnostic information in the event of a server side error
(e.g. invalid license, or failure to generate a template).

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|*anonymous*|string|false|none|Relevant and concise diagnostic information in the event of a server side error<br>(e.g. invalid license, or failure to generate a template).|

