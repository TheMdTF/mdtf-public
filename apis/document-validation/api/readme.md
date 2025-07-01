<!-- Generator: Widdershins v4.0.1 -->

<h1 id="the-maryland-test-facility-document-validation-interface">The Maryland Test Facility Document Validation Interface
 v0.0.1</h1>

> Scroll down for code samples, example requests and responses. Select a language for code samples from the tabs above or the mobile navigation menu.

This documents the API requirements for MdTF testing of document validation
algorithms.  Document validation testing at the MdTF is supported by the
Department of Homeland Security Science and Technology Directorate (DHS
S&T).  For more information please visit
[https://mdtf.org](https://mdtf.org) and
[https://www.dhs.gov/science-and-technology/BI-TC](https://www.dhs.gov/science-and-technology/BI-TC).

Email: <a href="mailto:info@mdtf.org">The MdTF</a> Web: <a href="https://mdtf.org">The MdTF</a> 
License: <a href="https://raw.githubusercontent.com/TheMdTF/mdtf-public/master/LICENSE.md">IDSL API License</a>

<h1 id="the-maryland-test-facility-document-validation-interface-validate">Validate</h1>

Return validation decision based on images of an identity document

## Validate an identity document based on images.

<a id="opIdValidateDocument"></a>

> Code samples

```shell
# You can also use wget
curl -X POST /v1/validate \
  -H 'Content-Type: application/json' \
  -H 'Accept: application/json'

```

```javascript
const inputBody = '{
  "DocumentFront": "iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAQAAAC1HAwCAAAAC0lEQVQYV2NgYAAAAAMAAWgmWQ0AAAAASUVORK5CYII=\n",
  "DocumentBack": "iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAQAAAC1HAwCAAAAC0lEQVQYV2NgYAAAAAMAAWgmWQ0AAAAASUVORK5CYII=\n"
}';
const headers = {
  'Content-Type':'application/json',
  'Accept':'application/json'
};

fetch('/v1/validate',
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
    req, err := http.NewRequest("POST", "/v1/validate", data)
    req.Header = headers

    client := &http.Client{}
    resp, err := client.Do(req)
    // ...
}

```

`POST /v1/validate`

Receive base64 encoded PNG or JPEG images of the front and back of a
document.  Validation requires that the document is judged to be
authentic.  See [NIST SP
800-63A](https://pages.nist.gov/800-63-3-Implementation-Resources/63A/ial2remote/)
Section A.10.2: Identity Validation for recommended best practices.

> Body parameter

```json
{
  "DocumentFront": "iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAQAAAC1HAwCAAAAC0lEQVQYV2NgYAAAAAMAAWgmWQ0AAAAASUVORK5CYII=\n",
  "DocumentBack": "iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAQAAAC1HAwCAAAAC0lEQVQYV2NgYAAAAAMAAWgmWQ0AAAAASUVORK5CYII=\n"
}
```

<h3 id="validate-an-identity-document-based-on-images.-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|body|body|[CapturedDocument](#schemacaptureddocument)|true|none|

> Example responses

> 200 Response

```json
{
  "ValidityOutcome": true,
  "ValidityScore": 0.8,
  "ValidityProperties": [
    {
      "Property": "InformationFormatCorrect",
      "Value": true
    },
    {
      "Property": "InformationCorrect",
      "Value": true
    },
    {
      "Property": "InformationComplete",
      "Value": true
    },
    {
      "Property": "CheckSumsCorrect",
      "Value": true
    },
    {
      "Property": "MRZCorrect",
      "Value": true
    },
    {
      "Property": "PhotoTampered",
      "Value": true
    },
    {
      "Property": "LivenessCheck",
      "Value": "Passed"
    },
    {
      "Property": "InformationCompleteness",
      "Value": 0.9
    },
    {
      "Property": "SecurityFeaturesTampered",
      "Value": "None"
    },
    {
      "Property": "DataFieldsTampered",
      "Value": "Age changed"
    }
  ]
}
```

<h3 id="validate-an-identity-document-based-on-images.-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|Successful validation operation|[ValidationResponse](#schemavalidationresponse)|
|400|[Bad Request](https://tools.ietf.org/html/rfc7231#section-6.5.1)|Client error|[ValidationErrorResponse](#schemavalidationerrorresponse)|
|500|[Internal Server Error](https://tools.ietf.org/html/rfc7231#section-6.6.1)|Server error|[ValidationErrorResponse](#schemavalidationerrorresponse)|

<aside class="success">
This operation does not require authentication
</aside>

<h1 id="the-maryland-test-facility-document-validation-interface-algorithm-information">Algorithm Information</h1>

## Return information about the document validation algorithm.

<a id="opIdInfo"></a>

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

This endpoint returns information about the validation algorithm.

> Example responses

> 200 Response

```json
{
  "AlgorithmName": "Docvalidate",
  "AlgorithmVersion": "1.0.0",
  "CompanyName": "IDSL",
  "Event": "2023 Document Validation Demonstration",
  "RecommendedCPUs": 4,
  "RecommendedMem": 256,
  "TechnicalContactEmail": "info@mdtf.org"
}
```

<h3 id="return-information-about-the-document-validation-algorithm.-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|Successful operation|[Info](#schemainfo)|
|500|[Internal Server Error](https://tools.ietf.org/html/rfc7231#section-6.6.1)|Server error|[ValidationErrorResponse](#schemavalidationerrorresponse)|

<aside class="success">
This operation does not require authentication
</aside>

# Schemas

<h2 id="tocS_Info">Info</h2>
<!-- backwards compatibility -->
<a id="schemainfo"></a>
<a id="schema_Info"></a>
<a id="tocSinfo"></a>
<a id="tocsinfo"></a>

```json
{
  "AlgorithmName": "Docvalidate",
  "AlgorithmVersion": "1.0.0",
  "CompanyName": "IDSL",
  "Event": "2023 Document Validation Demonstration",
  "RecommendedCPUs": 4,
  "RecommendedMem": 256,
  "TechnicalContactEmail": "info@mdtf.org"
}

```

Basic information describing the document validation algorithm.

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|AlgorithmName|string|true|none|Name of the algorithm.|
|AlgorithmVersion|string|true|none|Algorithm version identifier.|
|CompanyName|string|true|none|Name of the company that produced the algorithm.|
|Event|string|true|none|The name of the test event.|
|RecommendedCPUs|integer(int64)|true|none|The recommended allocation of CPUs for the deployed docker<br>container.|
|RecommendedMem|integer(int64)|true|none|The recommended allocation of memory (MB) for the deployed docker<br>container.|
|TechnicalContactEmail|string|true|none|The email address of an engineer or other technical resource to<br>contact in the event of an error running your service.|

#### Enumerated Values

|Property|Value|
|---|---|
|Event|2023 Document Validation Demonstration|

<h2 id="tocS_CapturedDocument">CapturedDocument</h2>
<!-- backwards compatibility -->
<a id="schemacaptureddocument"></a>
<a id="schema_CapturedDocument"></a>
<a id="tocScaptureddocument"></a>
<a id="tocscaptureddocument"></a>

```json
{
  "DocumentFront": "iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAQAAAC1HAwCAAAAC0lEQVQYV2NgYAAAAAMAAWgmWQ0AAAAASUVORK5CYII=\n",
  "DocumentBack": "iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAQAAAC1HAwCAAAAC0lEQVQYV2NgYAAAAAMAAWgmWQ0AAAAASUVORK5CYII=\n"
}

```

Object representation of a document. Consists of an image of the front
and back of a document.

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|DocumentFront|string(byte)|true|none|Image of front of document; base64 encoded bytes.|
|DocumentBack|string(byte)|true|none|Image of back of document; base64 encoded bytes.|

<h2 id="tocS_ValidityProperty">ValidityProperty</h2>
<!-- backwards compatibility -->
<a id="schemavalidityproperty"></a>
<a id="schema_ValidityProperty"></a>
<a id="tocSvalidityproperty"></a>
<a id="tocsvalidityproperty"></a>

```json
{
  "Property": "InformationFormatCorrect",
  "Value": true
}

```

Key value pair object.

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|Property|string|false|none|Name of the property.|
|Value|string|false|none|Value for the property.|

<h2 id="tocS_ValidationResponse">ValidationResponse</h2>
<!-- backwards compatibility -->
<a id="schemavalidationresponse"></a>
<a id="schema_ValidationResponse"></a>
<a id="tocSvalidationresponse"></a>
<a id="tocsvalidationresponse"></a>

```json
{
  "ValidityOutcome": true,
  "ValidityScore": 0.8,
  "ValidityProperties": [
    {
      "Property": "InformationFormatCorrect",
      "Value": true
    },
    {
      "Property": "InformationCorrect",
      "Value": true
    },
    {
      "Property": "InformationComplete",
      "Value": true
    },
    {
      "Property": "CheckSumsCorrect",
      "Value": true
    },
    {
      "Property": "MRZCorrect",
      "Value": true
    },
    {
      "Property": "PhotoTampered",
      "Value": true
    },
    {
      "Property": "LivenessCheck",
      "Value": "Passed"
    },
    {
      "Property": "InformationCompleteness",
      "Value": 0.9
    },
    {
      "Property": "SecurityFeaturesTampered",
      "Value": "None"
    },
    {
      "Property": "DataFieldsTampered",
      "Value": "Age changed"
    }
  ]
}

```

Validation response object.

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|ValidityOutcome|boolean|true|none|Whether the input document is determined to be valid (True) or<br>invalid (False).|
|ValidityScore|number(double)|false|none|A score corresponding to the level of confidence that the document<br>is valid ranging between 0 and 1 (optional).|
|ValidityProperties|[[ValidityProperty](#schemavalidityproperty)]|false|none|Key value pairs describing document properties and their<br>relationship to the validity decision.  There are no strictly<br>defined  properties. The inclusion of descriptive properties is<br>encouraged to provide more context. (optional)|

<h2 id="tocS_ValidationErrorResponse">ValidationErrorResponse</h2>
<!-- backwards compatibility -->
<a id="schemavalidationerrorresponse"></a>
<a id="schema_ValidationErrorResponse"></a>
<a id="tocSvalidationerrorresponse"></a>
<a id="tocsvalidationerrorresponse"></a>

```json
{
  "error": "string"
}

```

Validation error response object.

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|error|string|true|none|Descriptive error string.|

