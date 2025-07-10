<!-- Generator: Widdershins v4.0.1 -->

<h1 id="mdtf-media-acquisition-api">MdTF Media Acquisition API v3.0.0</h1>

> Scroll down for code samples, example requests and responses. Select a language for code samples from the tabs above or the mobile navigation menu.

API for submitting media (images, video, audio) with support for multiple modalities.
Each media type and modality combination has its own endpoint.
Clients must use a single POST request with multipart/form-data to submit both the media file and its metadata.

Base URLs:

* <a href="https://api.mdtf.org">https://api.mdtf.org</a>

<h1 id="mdtf-media-acquisition-api-default">Default</h1>

## Upload face image

> Code samples

```shell
# You can also use wget
curl -X POST https://api.mdtf.org/image/face \
  -H 'Content-Type: multipart/form-data'

```

```javascript
const inputBody = '{
  "file": "string",
  "metadata": {
    "StationID": "string",
    "Timestamp": "2019-08-24T14:15:22Z",
    "Metadata": {
      "lighting": "normal"
    },
    "Format": "jpeg",
    "ImageType": "Unknown",
    "Modality": "Palm"
  }
}';
const headers = {
  'Content-Type':'multipart/form-data'
};

fetch('https://api.mdtf.org/image/face',
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
        "Content-Type": []string{"multipart/form-data"},
    }

    data := bytes.NewBuffer([]byte{jsonReq})
    req, err := http.NewRequest("POST", "https://api.mdtf.org/image/face", data)
    req.Header = headers

    client := &http.Client{}
    resp, err := client.Do(req)
    // ...
}

```

`POST /image/face`

> Body parameter

```yaml
file: string
metadata:
  StationID: string
  Timestamp: 2019-08-24T14:15:22Z
  Metadata:
    lighting: normal
  Format: jpeg
  ImageType: Unknown
  Modality: Palm

```

<h3 id="upload-face-image-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|body|body|object|true|none|
|» file|body|string(binary)|true|none|
|» metadata|body|any|true|none|
|»» *anonymous*|body|[UploadMetadataBase](#schemauploadmetadatabase)|false|none|
|»»» StationID|body|string|true|none|
|»»» Timestamp|body|string(date-time)|false|none|
|»»» Metadata|body|object|false|none|
|»» *anonymous*|body|object|false|none|
|»»» Format|body|[ImageFormat](#schemaimageformat)|true|none|
|»»» ImageType|body|[BiometricImageType](#schemabiometricimagetype)|false|The allowed string names for biometric image types.|
|»»» Modality|body|string|true|none|

#### Enumerated Values

|Parameter|Value|
|---|---|
|»»» Format|jpeg|
|»»» Format|png|
|»»» Format|wsq|
|»»» Format|bmp|
|»»» ImageType|Unknown|
|»»» ImageType|UnknownFinger|
|»»» ImageType|RightThumb|
|»»» ImageType|RightIndexFinger|
|»»» ImageType|RightMiddleFinger|
|»»» ImageType|RightRingFinger|
|»»» ImageType|RightLittleFinger|
|»»» ImageType|RightTwo|
|»»» ImageType|LeftThumb|
|»»» ImageType|LeftIndexFinger|
|»»» ImageType|LeftMiddleFinger|
|»»» ImageType|LeftRingFinger|
|»»» ImageType|LeftLittleFinger|
|»»» ImageType|FlatRightThumb|
|»»» ImageType|FlatLeftThumb|
|»»» ImageType|FlatRightFourFingers|
|»»» ImageType|FlatLeftFourFingers|
|»»» ImageType|FlatLeftAndRightThumbs|
|»»» ImageType|RightExtraDigit|
|»»» ImageType|LeftExtraDigit|
|»»» ImageType|UnknownFrictionRidge|
|»»» ImageType|EjiOrTip|
|»»» ImageType|Iris|
|»»» ImageType|RightIris|
|»»» ImageType|LeftIris|
|»»» ImageType|Face|
|»»» ImageType|RightIndexAndMiddle|
|»»» ImageType|RolledRightThumb|
|»»» ImageType|RolledRightIndexFinger|
|»»» ImageType|RolledRightMiddleFinger|
|»»» ImageType|RolledRightRingFinger|
|»»» ImageType|RolledRightLittleFinger|
|»»» ImageType|RolledLeftThumb|
|»»» ImageType|RolledLeftIndexFinger|
|»»» ImageType|RolledLeftMiddleFinger|
|»»» ImageType|RolledLeftRingFinger|
|»»» ImageType|RolledLeftLittleFinger|
|»»» ImageType|RolledLeftExtraDigit|
|»»» ImageType|RolledRightExtraDigit|
|»»» ImageType|RightRingAndLittle|
|»»» ImageType|LeftIndexAndMiddle|
|»»» ImageType|LeftRingAndLittle|
|»»» Modality|Palm|
|»»» Modality|Face|
|»»» Modality|Iris|
|»»» Modality|Finger|

<h3 id="upload-face-image-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|Upload successful|None|
|400|[Bad Request](https://tools.ietf.org/html/rfc7231#section-6.5.1)|Malformed request or metadata|None|

<aside class="success">
This operation does not require authentication
</aside>

## Upload iris image (left/right)

> Code samples

```shell
# You can also use wget
curl -X POST https://api.mdtf.org/image/iris \
  -H 'Content-Type: multipart/form-data'

```

```javascript
const inputBody = '{
  "file": "string",
  "metadata": {
    "StationID": "string",
    "Timestamp": "2019-08-24T14:15:22Z",
    "Metadata": {
      "lighting": "normal"
    },
    "Format": "jpeg",
    "ImageType": "Unknown",
    "Modality": "Palm"
  }
}';
const headers = {
  'Content-Type':'multipart/form-data'
};

fetch('https://api.mdtf.org/image/iris',
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
        "Content-Type": []string{"multipart/form-data"},
    }

    data := bytes.NewBuffer([]byte{jsonReq})
    req, err := http.NewRequest("POST", "https://api.mdtf.org/image/iris", data)
    req.Header = headers

    client := &http.Client{}
    resp, err := client.Do(req)
    // ...
}

```

`POST /image/iris`

> Body parameter

```yaml
file: string
metadata:
  StationID: string
  Timestamp: 2019-08-24T14:15:22Z
  Metadata:
    lighting: normal
  Format: jpeg
  ImageType: Unknown
  Modality: Palm

```

<h3 id="upload-iris-image-(left/right)-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|body|body|object|true|none|
|» file|body|string(binary)|true|none|
|» metadata|body|any|true|none|
|»» *anonymous*|body|[UploadMetadataBase](#schemauploadmetadatabase)|false|none|
|»»» StationID|body|string|true|none|
|»»» Timestamp|body|string(date-time)|false|none|
|»»» Metadata|body|object|false|none|
|»» *anonymous*|body|object|false|none|
|»»» Format|body|[ImageFormat](#schemaimageformat)|true|none|
|»»» ImageType|body|[BiometricImageType](#schemabiometricimagetype)|false|The allowed string names for biometric image types.|
|»»» Modality|body|string|true|none|

#### Enumerated Values

|Parameter|Value|
|---|---|
|»»» Format|jpeg|
|»»» Format|png|
|»»» Format|wsq|
|»»» Format|bmp|
|»»» ImageType|Unknown|
|»»» ImageType|UnknownFinger|
|»»» ImageType|RightThumb|
|»»» ImageType|RightIndexFinger|
|»»» ImageType|RightMiddleFinger|
|»»» ImageType|RightRingFinger|
|»»» ImageType|RightLittleFinger|
|»»» ImageType|RightTwo|
|»»» ImageType|LeftThumb|
|»»» ImageType|LeftIndexFinger|
|»»» ImageType|LeftMiddleFinger|
|»»» ImageType|LeftRingFinger|
|»»» ImageType|LeftLittleFinger|
|»»» ImageType|FlatRightThumb|
|»»» ImageType|FlatLeftThumb|
|»»» ImageType|FlatRightFourFingers|
|»»» ImageType|FlatLeftFourFingers|
|»»» ImageType|FlatLeftAndRightThumbs|
|»»» ImageType|RightExtraDigit|
|»»» ImageType|LeftExtraDigit|
|»»» ImageType|UnknownFrictionRidge|
|»»» ImageType|EjiOrTip|
|»»» ImageType|Iris|
|»»» ImageType|RightIris|
|»»» ImageType|LeftIris|
|»»» ImageType|Face|
|»»» ImageType|RightIndexAndMiddle|
|»»» ImageType|RolledRightThumb|
|»»» ImageType|RolledRightIndexFinger|
|»»» ImageType|RolledRightMiddleFinger|
|»»» ImageType|RolledRightRingFinger|
|»»» ImageType|RolledRightLittleFinger|
|»»» ImageType|RolledLeftThumb|
|»»» ImageType|RolledLeftIndexFinger|
|»»» ImageType|RolledLeftMiddleFinger|
|»»» ImageType|RolledLeftRingFinger|
|»»» ImageType|RolledLeftLittleFinger|
|»»» ImageType|RolledLeftExtraDigit|
|»»» ImageType|RolledRightExtraDigit|
|»»» ImageType|RightRingAndLittle|
|»»» ImageType|LeftIndexAndMiddle|
|»»» ImageType|LeftRingAndLittle|
|»»» Modality|Palm|
|»»» Modality|Face|
|»»» Modality|Iris|
|»»» Modality|Finger|

<h3 id="upload-iris-image-(left/right)-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|Upload successful|None|
|400|[Bad Request](https://tools.ietf.org/html/rfc7231#section-6.5.1)|Malformed request or metadata|None|

<aside class="success">
This operation does not require authentication
</aside>

## Upload fingerprint image

> Code samples

```shell
# You can also use wget
curl -X POST https://api.mdtf.org/image/finger \
  -H 'Content-Type: multipart/form-data'

```

```javascript
const inputBody = '{
  "file": "string",
  "metadata": {
    "StationID": "string",
    "Timestamp": "2019-08-24T14:15:22Z",
    "Metadata": {
      "lighting": "normal"
    },
    "Format": "jpeg",
    "ImageType": "Unknown",
    "Modality": "Palm"
  }
}';
const headers = {
  'Content-Type':'multipart/form-data'
};

fetch('https://api.mdtf.org/image/finger',
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
        "Content-Type": []string{"multipart/form-data"},
    }

    data := bytes.NewBuffer([]byte{jsonReq})
    req, err := http.NewRequest("POST", "https://api.mdtf.org/image/finger", data)
    req.Header = headers

    client := &http.Client{}
    resp, err := client.Do(req)
    // ...
}

```

`POST /image/finger`

> Body parameter

```yaml
file: string
metadata:
  StationID: string
  Timestamp: 2019-08-24T14:15:22Z
  Metadata:
    lighting: normal
  Format: jpeg
  ImageType: Unknown
  Modality: Palm

```

<h3 id="upload-fingerprint-image-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|body|body|object|true|none|
|» file|body|string(binary)|true|none|
|» metadata|body|any|true|none|
|»» *anonymous*|body|[UploadMetadataBase](#schemauploadmetadatabase)|false|none|
|»»» StationID|body|string|true|none|
|»»» Timestamp|body|string(date-time)|false|none|
|»»» Metadata|body|object|false|none|
|»» *anonymous*|body|object|false|none|
|»»» Format|body|[ImageFormat](#schemaimageformat)|true|none|
|»»» ImageType|body|[BiometricImageType](#schemabiometricimagetype)|false|The allowed string names for biometric image types.|
|»»» Modality|body|string|true|none|

#### Enumerated Values

|Parameter|Value|
|---|---|
|»»» Format|jpeg|
|»»» Format|png|
|»»» Format|wsq|
|»»» Format|bmp|
|»»» ImageType|Unknown|
|»»» ImageType|UnknownFinger|
|»»» ImageType|RightThumb|
|»»» ImageType|RightIndexFinger|
|»»» ImageType|RightMiddleFinger|
|»»» ImageType|RightRingFinger|
|»»» ImageType|RightLittleFinger|
|»»» ImageType|RightTwo|
|»»» ImageType|LeftThumb|
|»»» ImageType|LeftIndexFinger|
|»»» ImageType|LeftMiddleFinger|
|»»» ImageType|LeftRingFinger|
|»»» ImageType|LeftLittleFinger|
|»»» ImageType|FlatRightThumb|
|»»» ImageType|FlatLeftThumb|
|»»» ImageType|FlatRightFourFingers|
|»»» ImageType|FlatLeftFourFingers|
|»»» ImageType|FlatLeftAndRightThumbs|
|»»» ImageType|RightExtraDigit|
|»»» ImageType|LeftExtraDigit|
|»»» ImageType|UnknownFrictionRidge|
|»»» ImageType|EjiOrTip|
|»»» ImageType|Iris|
|»»» ImageType|RightIris|
|»»» ImageType|LeftIris|
|»»» ImageType|Face|
|»»» ImageType|RightIndexAndMiddle|
|»»» ImageType|RolledRightThumb|
|»»» ImageType|RolledRightIndexFinger|
|»»» ImageType|RolledRightMiddleFinger|
|»»» ImageType|RolledRightRingFinger|
|»»» ImageType|RolledRightLittleFinger|
|»»» ImageType|RolledLeftThumb|
|»»» ImageType|RolledLeftIndexFinger|
|»»» ImageType|RolledLeftMiddleFinger|
|»»» ImageType|RolledLeftRingFinger|
|»»» ImageType|RolledLeftLittleFinger|
|»»» ImageType|RolledLeftExtraDigit|
|»»» ImageType|RolledRightExtraDigit|
|»»» ImageType|RightRingAndLittle|
|»»» ImageType|LeftIndexAndMiddle|
|»»» ImageType|LeftRingAndLittle|
|»»» Modality|Palm|
|»»» Modality|Face|
|»»» Modality|Iris|
|»»» Modality|Finger|

<h3 id="upload-fingerprint-image-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|Upload successful|None|
|400|[Bad Request](https://tools.ietf.org/html/rfc7231#section-6.5.1)|Malformed request or metadata|None|

<aside class="success">
This operation does not require authentication
</aside>

## Upload palm image

> Code samples

```shell
# You can also use wget
curl -X POST https://api.mdtf.org/image/palm \
  -H 'Content-Type: multipart/form-data'

```

```javascript
const inputBody = '{
  "file": "string",
  "metadata": {
    "StationID": "string",
    "Timestamp": "2019-08-24T14:15:22Z",
    "Metadata": {
      "lighting": "normal"
    },
    "Format": "jpeg",
    "ImageType": "Unknown",
    "Modality": "Palm"
  }
}';
const headers = {
  'Content-Type':'multipart/form-data'
};

fetch('https://api.mdtf.org/image/palm',
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
        "Content-Type": []string{"multipart/form-data"},
    }

    data := bytes.NewBuffer([]byte{jsonReq})
    req, err := http.NewRequest("POST", "https://api.mdtf.org/image/palm", data)
    req.Header = headers

    client := &http.Client{}
    resp, err := client.Do(req)
    // ...
}

```

`POST /image/palm`

> Body parameter

```yaml
file: string
metadata:
  StationID: string
  Timestamp: 2019-08-24T14:15:22Z
  Metadata:
    lighting: normal
  Format: jpeg
  ImageType: Unknown
  Modality: Palm

```

<h3 id="upload-palm-image-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|body|body|object|true|none|
|» file|body|string(binary)|true|none|
|» metadata|body|any|true|none|
|»» *anonymous*|body|[UploadMetadataBase](#schemauploadmetadatabase)|false|none|
|»»» StationID|body|string|true|none|
|»»» Timestamp|body|string(date-time)|false|none|
|»»» Metadata|body|object|false|none|
|»» *anonymous*|body|object|false|none|
|»»» Format|body|[ImageFormat](#schemaimageformat)|true|none|
|»»» ImageType|body|[BiometricImageType](#schemabiometricimagetype)|false|The allowed string names for biometric image types.|
|»»» Modality|body|string|true|none|

#### Enumerated Values

|Parameter|Value|
|---|---|
|»»» Format|jpeg|
|»»» Format|png|
|»»» Format|wsq|
|»»» Format|bmp|
|»»» ImageType|Unknown|
|»»» ImageType|UnknownFinger|
|»»» ImageType|RightThumb|
|»»» ImageType|RightIndexFinger|
|»»» ImageType|RightMiddleFinger|
|»»» ImageType|RightRingFinger|
|»»» ImageType|RightLittleFinger|
|»»» ImageType|RightTwo|
|»»» ImageType|LeftThumb|
|»»» ImageType|LeftIndexFinger|
|»»» ImageType|LeftMiddleFinger|
|»»» ImageType|LeftRingFinger|
|»»» ImageType|LeftLittleFinger|
|»»» ImageType|FlatRightThumb|
|»»» ImageType|FlatLeftThumb|
|»»» ImageType|FlatRightFourFingers|
|»»» ImageType|FlatLeftFourFingers|
|»»» ImageType|FlatLeftAndRightThumbs|
|»»» ImageType|RightExtraDigit|
|»»» ImageType|LeftExtraDigit|
|»»» ImageType|UnknownFrictionRidge|
|»»» ImageType|EjiOrTip|
|»»» ImageType|Iris|
|»»» ImageType|RightIris|
|»»» ImageType|LeftIris|
|»»» ImageType|Face|
|»»» ImageType|RightIndexAndMiddle|
|»»» ImageType|RolledRightThumb|
|»»» ImageType|RolledRightIndexFinger|
|»»» ImageType|RolledRightMiddleFinger|
|»»» ImageType|RolledRightRingFinger|
|»»» ImageType|RolledRightLittleFinger|
|»»» ImageType|RolledLeftThumb|
|»»» ImageType|RolledLeftIndexFinger|
|»»» ImageType|RolledLeftMiddleFinger|
|»»» ImageType|RolledLeftRingFinger|
|»»» ImageType|RolledLeftLittleFinger|
|»»» ImageType|RolledLeftExtraDigit|
|»»» ImageType|RolledRightExtraDigit|
|»»» ImageType|RightRingAndLittle|
|»»» ImageType|LeftIndexAndMiddle|
|»»» ImageType|LeftRingAndLittle|
|»»» Modality|Palm|
|»»» Modality|Face|
|»»» Modality|Iris|
|»»» Modality|Finger|

<h3 id="upload-palm-image-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|Upload successful|None|
|400|[Bad Request](https://tools.ietf.org/html/rfc7231#section-6.5.1)|Malformed request or metadata|None|

<aside class="success">
This operation does not require authentication
</aside>

## Upload face video

> Code samples

```shell
# You can also use wget
curl -X POST https://api.mdtf.org/video/face \
  -H 'Content-Type: multipart/form-data'

```

```javascript
const inputBody = '{
  "file": "string",
  "metadata": {
    "StationID": "string",
    "Timestamp": "2019-08-24T14:15:22Z",
    "Metadata": {
      "lighting": "normal"
    },
    "Format": "mp4",
    "Modality": "Face"
  }
}';
const headers = {
  'Content-Type':'multipart/form-data'
};

fetch('https://api.mdtf.org/video/face',
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
        "Content-Type": []string{"multipart/form-data"},
    }

    data := bytes.NewBuffer([]byte{jsonReq})
    req, err := http.NewRequest("POST", "https://api.mdtf.org/video/face", data)
    req.Header = headers

    client := &http.Client{}
    resp, err := client.Do(req)
    // ...
}

```

`POST /video/face`

> Body parameter

```yaml
file: string
metadata:
  StationID: string
  Timestamp: 2019-08-24T14:15:22Z
  Metadata:
    lighting: normal
  Format: mp4
  Modality: Face

```

<h3 id="upload-face-video-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|body|body|object|true|none|
|» file|body|string(binary)|true|none|
|» metadata|body|any|true|none|
|»» *anonymous*|body|[UploadMetadataBase](#schemauploadmetadatabase)|false|none|
|»»» StationID|body|string|true|none|
|»»» Timestamp|body|string(date-time)|false|none|
|»»» Metadata|body|object|false|none|
|»» *anonymous*|body|object|false|none|
|»»» Format|body|[VideoFormat](#schemavideoformat)|true|none|
|»»» Modality|body|string|true|none|

#### Enumerated Values

|Parameter|Value|
|---|---|
|»»» Format|mp4|
|»»» Format|mov|
|»»» Modality|Face|

<h3 id="upload-face-video-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|Upload successful|None|
|400|[Bad Request](https://tools.ietf.org/html/rfc7231#section-6.5.1)|Malformed request or metadata|None|

<aside class="success">
This operation does not require authentication
</aside>

## Upload voice audio

> Code samples

```shell
# You can also use wget
curl -X POST https://api.mdtf.org/audio/voice \
  -H 'Content-Type: multipart/form-data'

```

```javascript
const inputBody = '{
  "file": "string",
  "metadata": {
    "StationID": "string",
    "Timestamp": "2019-08-24T14:15:22Z",
    "Metadata": {
      "lighting": "normal"
    },
    "Format": "wav",
    "Modality": "Voice"
  }
}';
const headers = {
  'Content-Type':'multipart/form-data'
};

fetch('https://api.mdtf.org/audio/voice',
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
        "Content-Type": []string{"multipart/form-data"},
    }

    data := bytes.NewBuffer([]byte{jsonReq})
    req, err := http.NewRequest("POST", "https://api.mdtf.org/audio/voice", data)
    req.Header = headers

    client := &http.Client{}
    resp, err := client.Do(req)
    // ...
}

```

`POST /audio/voice`

> Body parameter

```yaml
file: string
metadata:
  StationID: string
  Timestamp: 2019-08-24T14:15:22Z
  Metadata:
    lighting: normal
  Format: wav
  Modality: Voice

```

<h3 id="upload-voice-audio-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|body|body|object|true|none|
|» file|body|string(binary)|true|none|
|» metadata|body|any|true|none|
|»» *anonymous*|body|[UploadMetadataBase](#schemauploadmetadatabase)|false|none|
|»»» StationID|body|string|true|none|
|»»» Timestamp|body|string(date-time)|false|none|
|»»» Metadata|body|object|false|none|
|»» *anonymous*|body|object|false|none|
|»»» Format|body|[AudioFormat](#schemaaudioformat)|true|none|
|»»» Modality|body|string|true|none|

#### Enumerated Values

|Parameter|Value|
|---|---|
|»»» Format|wav|
|»»» Format|mp3|
|»»» Format|flac|
|»»» Modality|Voice|

<h3 id="upload-voice-audio-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|Upload successful|None|
|400|[Bad Request](https://tools.ietf.org/html/rfc7231#section-6.5.1)|Malformed request or metadata|None|

<aside class="success">
This operation does not require authentication
</aside>

# Schemas

<h2 id="tocS_MediaFormat">MediaFormat</h2>
<!-- backwards compatibility -->
<a id="schemamediaformat"></a>
<a id="schema_MediaFormat"></a>
<a id="tocSmediaformat"></a>
<a id="tocsmediaformat"></a>

```json
"jpeg"

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|*anonymous*|string|false|none|none|

#### Enumerated Values

|Property|Value|
|---|---|
|*anonymous*|jpeg|
|*anonymous*|png|
|*anonymous*|wsq|
|*anonymous*|bmp|
|*anonymous*|mp4|
|*anonymous*|mov|
|*anonymous*|wav|
|*anonymous*|mp3|
|*anonymous*|flac|

<h2 id="tocS_ImageFormat">ImageFormat</h2>
<!-- backwards compatibility -->
<a id="schemaimageformat"></a>
<a id="schema_ImageFormat"></a>
<a id="tocSimageformat"></a>
<a id="tocsimageformat"></a>

```json
"jpeg"

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|*anonymous*|string|false|none|none|

#### Enumerated Values

|Property|Value|
|---|---|
|*anonymous*|jpeg|
|*anonymous*|png|
|*anonymous*|wsq|
|*anonymous*|bmp|

<h2 id="tocS_VideoFormat">VideoFormat</h2>
<!-- backwards compatibility -->
<a id="schemavideoformat"></a>
<a id="schema_VideoFormat"></a>
<a id="tocSvideoformat"></a>
<a id="tocsvideoformat"></a>

```json
"mp4"

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|*anonymous*|string|false|none|none|

#### Enumerated Values

|Property|Value|
|---|---|
|*anonymous*|mp4|
|*anonymous*|mov|

<h2 id="tocS_AudioFormat">AudioFormat</h2>
<!-- backwards compatibility -->
<a id="schemaaudioformat"></a>
<a id="schema_AudioFormat"></a>
<a id="tocSaudioformat"></a>
<a id="tocsaudioformat"></a>

```json
"wav"

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|*anonymous*|string|false|none|none|

#### Enumerated Values

|Property|Value|
|---|---|
|*anonymous*|wav|
|*anonymous*|mp3|
|*anonymous*|flac|

<h2 id="tocS_BiometricImageType">BiometricImageType</h2>
<!-- backwards compatibility -->
<a id="schemabiometricimagetype"></a>
<a id="schema_BiometricImageType"></a>
<a id="tocSbiometricimagetype"></a>
<a id="tocsbiometricimagetype"></a>

```json
"Unknown"

```

The allowed string names for biometric image types.

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|*anonymous*|string|false|none|The allowed string names for biometric image types.|

#### Enumerated Values

|Property|Value|
|---|---|
|*anonymous*|Unknown|
|*anonymous*|UnknownFinger|
|*anonymous*|RightThumb|
|*anonymous*|RightIndexFinger|
|*anonymous*|RightMiddleFinger|
|*anonymous*|RightRingFinger|
|*anonymous*|RightLittleFinger|
|*anonymous*|RightTwo|
|*anonymous*|LeftThumb|
|*anonymous*|LeftIndexFinger|
|*anonymous*|LeftMiddleFinger|
|*anonymous*|LeftRingFinger|
|*anonymous*|LeftLittleFinger|
|*anonymous*|FlatRightThumb|
|*anonymous*|FlatLeftThumb|
|*anonymous*|FlatRightFourFingers|
|*anonymous*|FlatLeftFourFingers|
|*anonymous*|FlatLeftAndRightThumbs|
|*anonymous*|RightExtraDigit|
|*anonymous*|LeftExtraDigit|
|*anonymous*|UnknownFrictionRidge|
|*anonymous*|EjiOrTip|
|*anonymous*|Iris|
|*anonymous*|RightIris|
|*anonymous*|LeftIris|
|*anonymous*|Face|
|*anonymous*|RightIndexAndMiddle|
|*anonymous*|RolledRightThumb|
|*anonymous*|RolledRightIndexFinger|
|*anonymous*|RolledRightMiddleFinger|
|*anonymous*|RolledRightRingFinger|
|*anonymous*|RolledRightLittleFinger|
|*anonymous*|RolledLeftThumb|
|*anonymous*|RolledLeftIndexFinger|
|*anonymous*|RolledLeftMiddleFinger|
|*anonymous*|RolledLeftRingFinger|
|*anonymous*|RolledLeftLittleFinger|
|*anonymous*|RolledLeftExtraDigit|
|*anonymous*|RolledRightExtraDigit|
|*anonymous*|RightRingAndLittle|
|*anonymous*|LeftIndexAndMiddle|
|*anonymous*|LeftRingAndLittle|

<h2 id="tocS_UploadMetadataBase">UploadMetadataBase</h2>
<!-- backwards compatibility -->
<a id="schemauploadmetadatabase"></a>
<a id="schema_UploadMetadataBase"></a>
<a id="tocSuploadmetadatabase"></a>
<a id="tocsuploadmetadatabase"></a>

```json
{
  "StationID": "string",
  "Timestamp": "2019-08-24T14:15:22Z",
  "Metadata": {
    "lighting": "normal"
  }
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|StationID|string|true|none|none|
|Timestamp|string(date-time)|false|none|none|
|Metadata|object|false|none|none|

<h2 id="tocS_UploadImageMetadata">UploadImageMetadata</h2>
<!-- backwards compatibility -->
<a id="schemauploadimagemetadata"></a>
<a id="schema_UploadImageMetadata"></a>
<a id="tocSuploadimagemetadata"></a>
<a id="tocsuploadimagemetadata"></a>

```json
{
  "StationID": "string",
  "Timestamp": "2019-08-24T14:15:22Z",
  "Metadata": {
    "lighting": "normal"
  },
  "Format": "jpeg",
  "ImageType": "Unknown",
  "Modality": "Palm"
}

```

### Properties

allOf

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|*anonymous*|[UploadMetadataBase](#schemauploadmetadatabase)|false|none|none|

and

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|*anonymous*|object|false|none|none|
|» Format|[ImageFormat](#schemaimageformat)|true|none|none|
|» ImageType|[BiometricImageType](#schemabiometricimagetype)|false|none|The allowed string names for biometric image types.|
|» Modality|string|true|none|none|

#### Enumerated Values

|Property|Value|
|---|---|
|Modality|Palm|
|Modality|Face|
|Modality|Iris|
|Modality|Finger|

<h2 id="tocS_UploadVideoMetadata">UploadVideoMetadata</h2>
<!-- backwards compatibility -->
<a id="schemauploadvideometadata"></a>
<a id="schema_UploadVideoMetadata"></a>
<a id="tocSuploadvideometadata"></a>
<a id="tocsuploadvideometadata"></a>

```json
{
  "StationID": "string",
  "Timestamp": "2019-08-24T14:15:22Z",
  "Metadata": {
    "lighting": "normal"
  },
  "Format": "mp4",
  "Modality": "Face"
}

```

### Properties

allOf

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|*anonymous*|[UploadMetadataBase](#schemauploadmetadatabase)|false|none|none|

and

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|*anonymous*|object|false|none|none|
|» Format|[VideoFormat](#schemavideoformat)|true|none|none|
|» Modality|string|true|none|none|

#### Enumerated Values

|Property|Value|
|---|---|
|Modality|Face|

<h2 id="tocS_UploadAudioMetadata">UploadAudioMetadata</h2>
<!-- backwards compatibility -->
<a id="schemauploadaudiometadata"></a>
<a id="schema_UploadAudioMetadata"></a>
<a id="tocSuploadaudiometadata"></a>
<a id="tocsuploadaudiometadata"></a>

```json
{
  "StationID": "string",
  "Timestamp": "2019-08-24T14:15:22Z",
  "Metadata": {
    "lighting": "normal"
  },
  "Format": "wav",
  "Modality": "Voice"
}

```

### Properties

allOf

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|*anonymous*|[UploadMetadataBase](#schemauploadmetadatabase)|false|none|none|

and

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|*anonymous*|object|false|none|none|
|» Format|[AudioFormat](#schemaaudioformat)|true|none|none|
|» Modality|string|true|none|none|

#### Enumerated Values

|Property|Value|
|---|---|
|Modality|Voice|

