<!-- Generator: Widdershins v4.0.1 -->

<h1 id="the-maryland-test-facility-passive-presentation-attack-detection-system-interface">The Maryland Test Facility Passive Presentation Attack Detection System Interface v2.0.1</h1>

> Scroll down for code samples, example requests and responses. Select a language for code samples from the tabs above or the mobile navigation menu.

This document describes an application programming interface for a PAD system in which the PAD subsystem is logically distinct from the Data Capture Subsystem.  This allows for offline processing of PAD data.  This supports both image and video biometric sample inputs.

Base URLs:

* <a href="https://api.mdtf.org/">https://api.mdtf.org/</a>

    * **host** -  Default: mdtf.org

    * **port** -  Default: 8080

        * 8080

Email: <a href="mailto:rivtd@mdtf.org">The MdTF</a> Web: <a href="https://mdtf.org">The MdTF</a> 
License: <a href="https://raw.githubusercontent.com/TheMdTF/mdtf-public/master/LICENSE.md">IDSL API License</a>

<h1 id="the-maryland-test-facility-passive-presentation-attack-detection-system-interface-data-analysis">Data Analysis</h1>

## Analyze biometric capture data for a presentation attack.

> Code samples

```shell
# You can also use wget
curl -X POST https://api.mdtf.org/v1/analyze-data-for-pad \
  -H 'Content-Type: application/json' \
  -H 'Accept: application/json'

```

```javascript
const inputBody = '{
  "BiometricSample": "iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAIAAACQd1PeAAAAEElEQVR4nGJiYGAABAAA//8ADAADcZGLFwAAAABJRU5ErkJggg=="
}';
const headers = {
  'Content-Type':'application/json',
  'Accept':'application/json'
};

fetch('https://api.mdtf.org/v1/analyze-data-for-pad',
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
    req, err := http.NewRequest("POST", "https://api.mdtf.org/v1/analyze-data-for-pad", data)
    req.Header = headers

    client := &http.Client{}
    resp, err := client.Do(req)
    // ...
}

```

`POST /v1/analyze-data-for-pad`

Passive PAD systems receive data through this endpoint.  This request cannot exceed 100 Megabytes in size.

> Body parameter

```json
{
  "BiometricSample": "iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAIAAACQd1PeAAAAEElEQVR4nGJiYGAABAAA//8ADAADcZGLFwAAAABJRU5ErkJggg=="
}
```

<h3 id="analyze-biometric-capture-data-for-a-presentation-attack.-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|body|body|[BiometricSampleRequest](#schemabiometricsamplerequest)|true|none|

> Example responses

> 200 Response

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

<h3 id="analyze-biometric-capture-data-for-a-presentation-attack.-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|The result of an analysis of a biometric sample for a presentation attack.|[PADAnalysis](#schemapadanalysis)|
|400|[Bad Request](https://tools.ietf.org/html/rfc7231#section-6.5.1)|The request is malformed.|[Error](#schemaerror)|
|500|[Internal Server Error](https://tools.ietf.org/html/rfc7231#section-6.6.1)|This request failed because of a server side issue.|[Error](#schemaerror)|

<aside class="success">
This operation does not require authentication
</aside>

<h1 id="the-maryland-test-facility-passive-presentation-attack-detection-system-interface-algorithm-information">Algorithm Information</h1>

## Returns basic information for the algorithm.

<a id="opIdGetInfo"></a>

> Code samples

```shell
# You can also use wget
curl -X GET https://api.mdtf.org/v1/info \
  -H 'Accept: application/json'

```

```javascript

const headers = {
  'Accept':'application/json'
};

fetch('https://api.mdtf.org/v1/info',
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
    req, err := http.NewRequest("GET", "https://api.mdtf.org/v1/info", data)
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
  "AlgorithmName": "EyeDetector",
  "AlgorithmVersion": "1.0.1",
  "AlgorithmModality": "Face",
  "CompanyName": "MdTF",
  "TechnicalContactEmail": "john@mdtf.org",
  "RecommendedCPUs": 0.5,
  "RecommendedMem": 512,
  "Thresholds": {
    "1:ten": "0.75",
    "1:1e2": "0.87",
    "1:1e3": "0.93",
    "1:1e4": "0.98"
  }
}
```

<h3 id="returns-basic-information-for-the-algorithm.-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|info response object|[Info](#schemainfo)|
|500|[Internal Server Error](https://tools.ietf.org/html/rfc7231#section-6.6.1)|This request failed because of a server side issue.|[Error](#schemaerror)|

<aside class="success">
This operation does not require authentication
</aside>

# Schemas

<h2 id="tocS_BiometricSample">BiometricSample</h2>
<!-- backwards compatibility -->
<a id="schemabiometricsample"></a>
<a id="schema_BiometricSample"></a>
<a id="tocSbiometricsample"></a>
<a id="tocsbiometricsample"></a>

```json
"iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAIAAACQd1PeAAAAEElEQVR4nGJiYGAABAAA//8ADAADcZGLFwAAAABJRU5ErkJggg=="

```

The biometric sample, encoded as a base64 string.  This can be an image, encoded as a PNG or JPEG or a short (<10s) video, encoded as a MOV or MP4.

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|*anonymous*|string|false|none|The biometric sample, encoded as a base64 string.  This can be an image, encoded as a PNG or JPEG or a short (<10s) video, encoded as a MOV or MP4.|

<h2 id="tocS_BiometricSampleRequest">BiometricSampleRequest</h2>
<!-- backwards compatibility -->
<a id="schemabiometricsamplerequest"></a>
<a id="schema_BiometricSampleRequest"></a>
<a id="tocSbiometricsamplerequest"></a>
<a id="tocsbiometricsamplerequest"></a>

```json
{
  "BiometricSample": "iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAIAAACQd1PeAAAAEElEQVR4nGJiYGAABAAA//8ADAADcZGLFwAAAABJRU5ErkJggg=="
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|BiometricSample|[BiometricSample](#schemabiometricsample)|false|none|The biometric sample, encoded as a base64 string.  This can be an image, encoded as a PNG or JPEG or a short (<10s) video, encoded as a MOV or MP4.|

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
|PADOutcome|boolean|true|none|Whether a presentation attack was determined to be detected (True) or not detected (False).  The outcome should be calibrated to produce a BPCER of 0.01.|
|PADScore|number(double)|true|none|A score corresponding to the level of confidence that a presentation attack was detected ranging between 0 and 1.|
|PADProperties|[[PADProperty](#schemapadproperty)]|false|none|Key value pairs describing presentation attack properties and their relationship to the presentation attack outcome/score.  There are no strictly defined properties.  The inclusion of descriptive properties is encouraged to provide more context.  (optional)|

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

<h2 id="tocS_Info">Info</h2>
<!-- backwards compatibility -->
<a id="schemainfo"></a>
<a id="schema_Info"></a>
<a id="tocSinfo"></a>
<a id="tocsinfo"></a>

```json
{
  "AlgorithmName": "EyeDetector",
  "AlgorithmVersion": "1.0.1",
  "AlgorithmModality": "Face",
  "CompanyName": "MdTF",
  "TechnicalContactEmail": "john@mdtf.org",
  "RecommendedCPUs": 0.5,
  "RecommendedMem": 512,
  "Thresholds": {
    "1:ten": "0.75",
    "1:1e2": "0.87",
    "1:1e3": "0.93",
    "1:1e4": "0.98"
  }
}

```

Basic information describing the PAD algorithm.

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|AlgorithmName|string|false|none|Name of algorithm.|
|AlgorithmVersion|string|false|none|Algorithm version identifier.|
|AlgorithmModality|string|false|none|A string enum describing the type of biometric modality the algorithm is meant to process.|
|CompanyName|string|false|none|Name of the company which produced the algorithm.|
|TechnicalContactEmail|string|false|none|The email address of an engineer or other technical resource to contact in the event of an error running your service.|
|RecommendedCPUs|number(float)|false|none|The recommended allocation of CPUs for the deployed docker container.|
|RecommendedMem|integer|false|none|The recommended allocation of memory (MB) for the deployed docker container.|
|Thresholds|object|false|none|A map of specific Bona-fide Classification Error Rates (BPCERs) to vendor-provided threshold values.  PADScore values optionally returned from calls to<br>/v1/analyze-data-for-pad indicate a presentation attack is occuring is if they are greater than the provided threshold value at the respective BPCER.<br>Note 1: Threshold values are floats stored as strings and cannot exceed a length of 10 characters.<br>Note 2: There are 4 required thresholds.|
|» 1:ten|string|true|none|none|
|» 1:1e2|string|true|none|none|
|» 1:1e3|string|true|none|none|
|» 1:1e4|string|true|none|none|

#### Enumerated Values

|Property|Value|
|---|---|
|AlgorithmModality|Face|

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

