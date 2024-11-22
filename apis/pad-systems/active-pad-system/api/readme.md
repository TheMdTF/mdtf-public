---
title: The Maryland Test Facility Active Presentation Attack Detection System
  Interface v2.0.1
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

<h1 id="the-maryland-test-facility-active-presentation-attack-detection-system-interface">The Maryland Test Facility Active Presentation Attack Detection System Interface v2.0.1</h1>

> Scroll down for code samples, example requests and responses. Select a language for code samples from the tabs above or the mobile navigation menu.

This document describes an application programming interface for a PAD system in which the PAD Subsystem is logically located within or tightly coupled with the Data Capture Subsystem.  This allows for what ISO/IEC 30107-1:2023 refers to as "through data capture subsystem" PAD methods (Table 2). It also allows for what the standard refers to as "challenge-response" actions (Subsection 5.2.1).  We collectively refer to these kinds of systems as "active" PAD or A-PAD systems.

Base URLs:

* <a href="https://pad-demo.mdtf.org/">https://pad-demo.mdtf.org/</a>

    * **host** -  Default: mdtf.org

    * **port** -  Default: 8080

        * 8080

Email: <a href="mailto:rivtd@mdtf.org">The MdTF</a> Web: <a href="https://mdtf.org">The MdTF</a> 
License: <a href="https://raw.githubusercontent.com/TheMdTF/mdtf-public/master/LICENSE.md">IDSL API License</a>

<h1 id="the-maryland-test-facility-active-presentation-attack-detection-system-interface-data-submission">Data Submission</h1>

## Create a biometric data capture with associated PAD information.

> Code samples

```shell
# You can also use wget
curl -X POST https://pad-demo.mdtf.org/v1/capture-data-with-pad \
  -H 'Content-Type: application/json' \
  -H 'Accept: application/json'

```

```javascript
const inputBody = '{
  "StationID": "Station_A",
  "MobilePlatform": "iOS",
  "BiometricSample": "iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAIAAACQd1PeAAAAEElEQVR4nGJiYGAABAAA//8ADAADcZGLFwAAAABJRU5ErkJggg==",
  "PADAnalysis": {
    "PADOutcome": true,
    "PADScore": 0.8,
    "PADProperties": [
      {
        "Property": "EyesMoving",
        "Value": "true"
      },
      {
        "Property": "MouthMoving",
        "Value": "true"
      },
      {
        "Property": "PupilsResponsive",
        "Value": "true"
      },
      {
        "Property": "NonconformantIlluminationDetected",
        "Value": "true"
      },
      {
        "Property": "MoirePatternDetected",
        "Value": "true"
      },
      {
        "Property": "ObsscurationDetected",
        "Value": "true"
      }
    ]
  }
}';
const headers = {
  'Content-Type':'application/json',
  'Accept':'application/json'
};

fetch('https://pad-demo.mdtf.org/v1/capture-data-with-pad',
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
    req, err := http.NewRequest("POST", "https://pad-demo.mdtf.org/v1/capture-data-with-pad", data)
    req.Header = headers

    client := &http.Client{}
    resp, err := client.Do(req)
    // ...
}

```

`POST /v1/capture-data-with-pad`

Active PAD systems may submit data through this endpoint.  This request cannot exceed 100 Megabytes in size.

> Body parameter

```json
{
  "StationID": "Station_A",
  "MobilePlatform": "iOS",
  "BiometricSample": "iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAIAAACQd1PeAAAAEElEQVR4nGJiYGAABAAA//8ADAADcZGLFwAAAABJRU5ErkJggg==",
  "PADAnalysis": {
    "PADOutcome": true,
    "PADScore": 0.8,
    "PADProperties": [
      {
        "Property": "EyesMoving",
        "Value": "true"
      },
      {
        "Property": "MouthMoving",
        "Value": "true"
      },
      {
        "Property": "PupilsResponsive",
        "Value": "true"
      },
      {
        "Property": "NonconformantIlluminationDetected",
        "Value": "true"
      },
      {
        "Property": "MoirePatternDetected",
        "Value": "true"
      },
      {
        "Property": "ObsscurationDetected",
        "Value": "true"
      }
    ]
  }
}
```

<h3 id="create-a-biometric-data-capture-with-associated-pad-information.-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|body|body|[PADDataCapture](#schemapaddatacapture)|true|PAD data collected as part of a biometric data capture process.|

> Example responses

> 400 Response

```json
{
  "Code": 0,
  "Message": "string",
  "ErrorString": "string"
}
```

<h3 id="create-a-biometric-data-capture-with-associated-pad-information.-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|Successfully submitted the biometric data capture and PAD information.|None|
|400|[Bad Request](https://tools.ietf.org/html/rfc7231#section-6.5.1)|The request is malformed.|[Error](#schemaerror)|
|404|[Not Found](https://tools.ietf.org/html/rfc7231#section-6.5.4)|No ongoing transaction was found for the specified StationID.|[Error](#schemaerror)|
|410|[Gone](https://tools.ietf.org/html/rfc7231#section-6.5.9)|The test volunteer has exited the station and the transaction has ended.|[Error](#schemaerror)|
|429|[Too Many Requests](https://tools.ietf.org/html/rfc6585#section-4)|Too many requests for the ongoing transaction.|[Error](#schemaerror)|
|500|[Internal Server Error](https://tools.ietf.org/html/rfc7231#section-6.6.1)|This capture failed because of a server side issue.|[Error](#schemaerror)|

<aside class="success">
This operation does not require authentication
</aside>

# Schemas

<h2 id="tocS_PADDataCapture">PADDataCapture</h2>
<!-- backwards compatibility -->
<a id="schemapaddatacapture"></a>
<a id="schema_PADDataCapture"></a>
<a id="tocSpaddatacapture"></a>
<a id="tocspaddatacapture"></a>

```json
{
  "StationID": "Station_A",
  "MobilePlatform": "iOS",
  "BiometricSample": "iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAIAAACQd1PeAAAAEElEQVR4nGJiYGAABAAA//8ADAADcZGLFwAAAABJRU5ErkJggg==",
  "PADAnalysis": {
    "PADOutcome": true,
    "PADScore": 0.8,
    "PADProperties": [
      {
        "Property": "EyesMoving",
        "Value": "true"
      },
      {
        "Property": "MouthMoving",
        "Value": "true"
      },
      {
        "Property": "PupilsResponsive",
        "Value": "true"
      },
      {
        "Property": "NonconformantIlluminationDetected",
        "Value": "true"
      },
      {
        "Property": "MoirePatternDetected",
        "Value": "true"
      },
      {
        "Property": "ObsscurationDetected",
        "Value": "true"
      }
    ]
  }
}

```

Data transfer object for biometric data capture and presentation attack information.

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|StationID|string|true|none|ID of the station who is submitting this image. IDs will be provided to vendors upon installation for the test event.|
|MobilePlatform|string|true|none|none|
|BiometricSample|string|true|none|The captured biometric sample, encoded as a base64 string.  This can be an image, encoded as a PNG or JPEG or a short (<10s) video, encoded as a MOV or a MP4.|
|PADAnalysis|[PADAnalysis](#schemapadanalysis)|true|none|Data transfer object for presentation attack information.|

#### Enumerated Values

|Property|Value|
|---|---|
|MobilePlatform|iOS|
|MobilePlatform|Android|

<h2 id="tocS_PADAnalysis">PADAnalysis</h2>
<!-- backwards compatibility -->
<a id="schemapadanalysis"></a>
<a id="schema_PADAnalysis"></a>
<a id="tocSpadanalysis"></a>
<a id="tocspadanalysis"></a>

```json
{
  "PADOutcome": true,
  "PADScore": 0.8,
  "PADProperties": [
    {
      "Property": "EyesMoving",
      "Value": "true"
    },
    {
      "Property": "MouthMoving",
      "Value": "true"
    },
    {
      "Property": "PupilsResponsive",
      "Value": "true"
    },
    {
      "Property": "NonconformantIlluminationDetected",
      "Value": "true"
    },
    {
      "Property": "MoirePatternDetected",
      "Value": "true"
    },
    {
      "Property": "ObsscurationDetected",
      "Value": "true"
    }
  ]
}

```

Data transfer object for presentation attack information.

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|PADOutcome|boolean|true|none|Whether a presentation attack was determined to be detected (True) or not detected (False).|
|PADScore|number(double)|true|none|A score corresponding to the level of confidence that a presentation attack was detected ranging between 0 and 1.|
|PADProperties|[[PADProperty](#schemapadproperty)]|false|none|Key value pairs describing presentation attack properties and their relationship to the presentation attack outcome/score. There are no strictly defined  properties. The inclusion of descriptive properties is encouraged to provide more context. (optional)|

<h2 id="tocS_PADProperty">PADProperty</h2>
<!-- backwards compatibility -->
<a id="schemapadproperty"></a>
<a id="schema_PADProperty"></a>
<a id="tocSpadproperty"></a>
<a id="tocspadproperty"></a>

```json
{
  "Property": "EyesMoving",
  "Value": true
}

```

Key value pair object.

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|Property|string|false|none|Name of the property.|
|Value|string|false|none|Value for the property.|

<h2 id="tocS_PADErrorResponse">PADErrorResponse</h2>
<!-- backwards compatibility -->
<a id="schemapaderrorresponse"></a>
<a id="schema_PADErrorResponse"></a>
<a id="tocSpaderrorresponse"></a>
<a id="tocspaderrorresponse"></a>

```json
{
  "error": "string"
}

```

Presentation attack error response object.

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|error|string|true|none|Descriptive error string.|

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

