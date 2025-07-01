<!-- Generator: Widdershins v4.0.1 -->

<h1 id="the-maryland-test-facility-image-analysis-interface">The Maryland Test Facility Image Analysis Interface v1.0.0</h1>

> Scroll down for code samples, example requests and responses. Select a language for code samples from the tabs above or the mobile navigation menu.

Application programming interface for receiving image analysis requests.

Email: <a href="mailto:info@mdtf.org">MdTF</a> Web: <a href="https://mdtf.org">MdTF</a> 
License: <a href="https://raw.githubusercontent.com/TheMdTF/mdtf-public/master/LICENSE.md">IDSL API License</a>

<h1 id="the-maryland-test-facility-image-analysis-interface-analysis-operations">Analysis Operations</h1>

## Returns an object containing the results of analysis for the provided
biometric image.

<a id="opIdanalyze_image"></a>

> Code samples

```shell
# You can also use wget
curl -X POST /v1/analyze-image

```

```javascript

fetch('/v1/analyze-image',
{
  method: 'POST'

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

    data := bytes.NewBuffer([]byte{jsonReq})
    req, err := http.NewRequest("POST", "/v1/analyze-image", data)
    req.Header = headers

    client := &http.Client{}
    resp, err := client.Do(req)
    // ...
}

```

`POST /v1/analyze-image`

This endpoint accepts a base64 encoded PNG image and returns an object
containing the output of a generic image analysis routine on that
image.

<h3 id="returns-an-object-containing-the-results-of-analysis-for-the-provided
biometric-image.
-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|image|body|[#/definitions/Image](#schema#/definitions/image)|true|The image that is being submitted for analysis.|

#### Detailed descriptions

**image**: The image that is being submitted for analysis.

<h3 id="returns-an-object-containing-the-results-of-analysis-for-the-provided
biometric-image.
-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|Successful Response|None|
|400|[Bad Request](https://tools.ietf.org/html/rfc7231#section-6.5.1)|Bad Request|None|
|500|[Internal Server Error](https://tools.ietf.org/html/rfc7231#section-6.6.1)|Internal Server Error|None|

<aside class="success">
This operation does not require authentication
</aside>

<h1 id="the-maryland-test-facility-image-analysis-interface-algorithm-information">Algorithm Information</h1>

## Returns basic information for the algorithm.

<a id="opIdinfo"></a>

> Code samples

```shell
# You can also use wget
curl -X GET /v1/info

```

```javascript

fetch('/v1/info',
{
  method: 'GET'

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

    data := bytes.NewBuffer([]byte{jsonReq})
    req, err := http.NewRequest("GET", "/v1/info", data)
    req.Header = headers

    client := &http.Client{}
    resp, err := client.Do(req)
    // ...
}

```

`GET /v1/info`

This endpoint returns some basic information about the image analysis
algorithm.

<h3 id="returns-basic-information-for-the-algorithm.-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|Successful Response|None|
|500|[Internal Server Error](https://tools.ietf.org/html/rfc7231#section-6.6.1)|Internal Server Error|None|

<aside class="success">
This operation does not require authentication
</aside>

