<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<!--[if IE]><meta http-equiv="X-UA-Compatible" content="IE=edge"><![endif]-->
<meta name="viewport" content="width=device-width, initial-scale=1.0">
<meta name="generator" content="Asciidoctor 1.5.4">
<link rel="stylesheet" href="https://fonts.googleapis.com/css?family=Open+Sans:300,300italic,400,400italic,600,600italic%7CNoto+Serif:400,400italic,700,700italic%7CDroid+Sans+Mono:400,700">
<link rel="stylesheet" href="./asciidoctor.css">
</head>
<body class="article toc2 toc-left">
<div id="header">
<h1>The Maryland Test Facility Acquisition System Interface</h1>
<div id="toc" class="toc2">
<div id="toctitle">Table of Contents</div>
<ul class="sectlevel1">
<li><a href="#_overview">Overview</a>
<ul class="sectlevel2">
<li><a href="#_version_information">Version information</a></li>
<li><a href="#_contact_information">Contact information</a></li>
<li><a href="#_uri_scheme">URI scheme</a></li>
<li><a href="#_tags">Tags</a></li>
<li><a href="#_consumes">Consumes</a></li>
<li><a href="#_produces">Produces</a></li>
</ul>
</li>
<li><a href="#_paths">Resources</a>
<ul class="sectlevel2">
<li><a href="#_image_submission_resource">Image Submission</a></li>
</ul>
</li>
<li><a href="#_definitions">Definitions</a>
<ul class="sectlevel2">
<li><a href="#_error">Error</a></li>
<li><a href="#_facecapture">FaceCapture</a></li>
<li><a href="#_fingercapture">FingerCapture</a></li>
<li><a href="#_fingerimage">FingerImage</a></li>
<li><a href="#_fingertype">FingerType</a></li>
<li><a href="#_iriscapture">IrisCapture</a></li>
</ul>
</li>
</ul>
</div>
</div>
<div id="content">
<div class="sect1">
<h2 id="_overview">Overview</h2>
<div class="sectionbody">
<div class="paragraph">
<p>DRAFT Application Programming Interface for sending acquisition records to the MdTF Backend.</p>
</div>
<div class="sect2">
<h3 id="_version_information">Version information</h3>
<div class="paragraph">
<p><em>Version</em> : 1.1.0</p>
</div>
</div>
<div class="sect2">
<h3 id="_contact_information">Contact information</h3>
<div class="paragraph">
<p><em>Contact</em> : Andrew Blanchard<br>
<em>Contact Email</em> : <a href="mailto:andrew@mdtf.org">andrew@mdtf.org</a></p>
</div>
</div>
<div class="sect2">
<h3 id="_uri_scheme">URI scheme</h3>
<div class="paragraph">
<p><em>Host</em> : api.mdtf.org<br>
<em>Schemes</em> : HTTP</p>
</div>
</div>
<div class="sect2">
<h3 id="_tags">Tags</h3>
<div class="ulist">
<ul>
<li>
<p>Image Submission : Endpoints for submitting images to an ongoing transaction.</p>
</li>
</ul>
</div>
</div>
<div class="sect2">
<h3 id="_consumes">Consumes</h3>
<div class="ulist">
<ul>
<li>
<p><code>application/json</code></p>
</li>
</ul>
</div>
</div>
<div class="sect2">
<h3 id="_produces">Produces</h3>
<div class="ulist">
<ul>
<li>
<p><code>application/json</code></p>
</li>
</ul>
</div>
</div>
</div>
</div>
<div class="sect1">
<h2 id="_paths">Resources</h2>
<div class="sectionbody">
<div class="sect2">
<h3 id="_image_submission_resource">Image Submission</h3>
<div class="paragraph">
<p>Endpoints for submitting images to an ongoing transaction.</p>
</div>
<div class="sect3">
<h4 id="_v1_face-captures_post">Associate a face image capture with the ongoing transaction</h4>
<div class="literalblock">
<div class="content">
<pre>POST /v1/face-captures</pre>
</div>
</div>
<div class="sect4">
<h5 id="_description">Description</h5>
<div class="paragraph">
<p>Rally Participants will submit face image data through this endpoint. Submitted image captures will be
stored and associated with the ongoing transaction. Multiple images may be submitted during a transaction.</p>
</div>
</div>
<div class="sect4">
<h5 id="_parameters">Parameters</h5>
<table class="tableblock frame-all grid-all spread">
<colgroup>
<col style="width: 11.1111%;">
<col style="width: 16.6666%;">
<col style="width: 50%;">
<col style="width: 22.2223%;">
</colgroup>
<thead>
<tr>
<th class="tableblock halign-left valign-middle">Type</th>
<th class="tableblock halign-left valign-middle">Name</th>
<th class="tableblock halign-left valign-middle">Description</th>
<th class="tableblock halign-left valign-middle">Schema</th>
</tr>
</thead>
<tbody>
<tr>
<td class="tableblock halign-left valign-middle"><p class="tableblock"><strong>Body</strong></p></td>
<td class="tableblock halign-left valign-middle"><p class="tableblock"><strong>face_capture</strong><br>
<em>required</em></p></td>
<td class="tableblock halign-left valign-middle"><p class="tableblock">face image collected as part of a transaction</p></td>
<td class="tableblock halign-left valign-middle"><p class="tableblock"><a href="#_facecapture">FaceCapture</a></p></td>
</tr>
</tbody>
</table>
</div>
<div class="sect4">
<h5 id="_responses">Responses</h5>
<table class="tableblock frame-all grid-all spread">
<colgroup>
<col style="width: 10%;">
<col style="width: 70%;">
<col style="width: 20%;">
</colgroup>
<thead>
<tr>
<th class="tableblock halign-left valign-middle">HTTP Code</th>
<th class="tableblock halign-left valign-middle">Description</th>
<th class="tableblock halign-left valign-middle">Schema</th>
</tr>
</thead>
<tbody>
<tr>
<td class="tableblock halign-left valign-middle"><p class="tableblock"><strong>200</strong></p></td>
<td class="tableblock halign-left valign-middle"><p class="tableblock">Successfully associated the image capture with the ongoing transaction.</p></td>
<td class="tableblock halign-left valign-middle"><p class="tableblock">No Content</p></td>
</tr>
<tr>
<td class="tableblock halign-left valign-middle"><p class="tableblock"><strong>400</strong></p></td>
<td class="tableblock halign-left valign-middle"><p class="tableblock">The request is malformed.</p></td>
<td class="tableblock halign-left valign-middle"><p class="tableblock"><a href="#_error">Error</a></p></td>
</tr>
<tr>
<td class="tableblock halign-left valign-middle"><p class="tableblock"><strong>404</strong></p></td>
<td class="tableblock halign-left valign-middle"><p class="tableblock">No ongoing transaction was found for the specified StationID.</p></td>
<td class="tableblock halign-left valign-middle"><p class="tableblock"><a href="#_error">Error</a></p></td>
</tr>
<tr>
<td class="tableblock halign-left valign-middle"><p class="tableblock"><strong>429</strong></p></td>
<td class="tableblock halign-left valign-middle"><p class="tableblock">Too many requests for the ongoing transaction.</p></td>
<td class="tableblock halign-left valign-middle"><p class="tableblock"><a href="#_error">Error</a></p></td>
</tr>
<tr>
<td class="tableblock halign-left valign-middle"><p class="tableblock"><strong>500</strong></p></td>
<td class="tableblock halign-left valign-middle"><p class="tableblock">This capture failed because of a server side issue</p></td>
<td class="tableblock halign-left valign-middle"><p class="tableblock"><a href="#_error">Error</a></p></td>
</tr>
<tr>
<td class="tableblock halign-left valign-middle"><p class="tableblock"><strong>default</strong></p></td>
<td class="tableblock halign-left valign-middle"><p class="tableblock">An error occurred.</p></td>
<td class="tableblock halign-left valign-middle"><p class="tableblock"><a href="#_error">Error</a></p></td>
</tr>
</tbody>
</table>
</div>
</div>
<div class="sect3">
<h4 id="_v1_finger-captures_post">Associate up to 10 finger captures with the ongoing transaction</h4>
<div class="literalblock">
<div class="content">
<pre>POST /v1/finger-captures</pre>
</div>
</div>
<div class="sect4">
<h5 id="_description_2">Description</h5>
<div class="paragraph">
<p>Rally Participants may submit finger image data through this endpoint. Submitted image captures will be stored and associated with the ongoing transaction. Multiple images may be submitted during a transaction.</p>
</div>
</div>
<div class="sect4">
<h5 id="_parameters_2">Parameters</h5>
<table class="tableblock frame-all grid-all spread">
<colgroup>
<col style="width: 11.1111%;">
<col style="width: 16.6666%;">
<col style="width: 50%;">
<col style="width: 22.2223%;">
</colgroup>
<thead>
<tr>
<th class="tableblock halign-left valign-middle">Type</th>
<th class="tableblock halign-left valign-middle">Name</th>
<th class="tableblock halign-left valign-middle">Description</th>
<th class="tableblock halign-left valign-middle">Schema</th>
</tr>
</thead>
<tbody>
<tr>
<td class="tableblock halign-left valign-middle"><p class="tableblock"><strong>Body</strong></p></td>
<td class="tableblock halign-left valign-middle"><p class="tableblock"><strong>finger_capture</strong><br>
<em>required</em></p></td>
<td class="tableblock halign-left valign-middle"><p class="tableblock">finger image(s) collected as part of a transaction</p></td>
<td class="tableblock halign-left valign-middle"><p class="tableblock"><a href="#_fingercapture">FingerCapture</a></p></td>
</tr>
</tbody>
</table>
</div>
<div class="sect4">
<h5 id="_responses_2">Responses</h5>
<table class="tableblock frame-all grid-all spread">
<colgroup>
<col style="width: 10%;">
<col style="width: 70%;">
<col style="width: 20%;">
</colgroup>
<thead>
<tr>
<th class="tableblock halign-left valign-middle">HTTP Code</th>
<th class="tableblock halign-left valign-middle">Description</th>
<th class="tableblock halign-left valign-middle">Schema</th>
</tr>
</thead>
<tbody>
<tr>
<td class="tableblock halign-left valign-middle"><p class="tableblock"><strong>200</strong></p></td>
<td class="tableblock halign-left valign-middle"><p class="tableblock">Successfully associated the image capture(s) with the ongoing transaction.</p></td>
<td class="tableblock halign-left valign-middle"><p class="tableblock">No Content</p></td>
</tr>
<tr>
<td class="tableblock halign-left valign-middle"><p class="tableblock"><strong>400</strong></p></td>
<td class="tableblock halign-left valign-middle"><p class="tableblock">The request is malformed.</p></td>
<td class="tableblock halign-left valign-middle"><p class="tableblock"><a href="#_error">Error</a></p></td>
</tr>
<tr>
<td class="tableblock halign-left valign-middle"><p class="tableblock"><strong>404</strong></p></td>
<td class="tableblock halign-left valign-middle"><p class="tableblock">No ongoing transaction was found for the specified StationID.</p></td>
<td class="tableblock halign-left valign-middle"><p class="tableblock"><a href="#_error">Error</a></p></td>
</tr>
<tr>
<td class="tableblock halign-left valign-middle"><p class="tableblock"><strong>429</strong></p></td>
<td class="tableblock halign-left valign-middle"><p class="tableblock">Too many requests for the ongoing transaction.</p></td>
<td class="tableblock halign-left valign-middle"><p class="tableblock"><a href="#_error">Error</a></p></td>
</tr>
<tr>
<td class="tableblock halign-left valign-middle"><p class="tableblock"><strong>500</strong></p></td>
<td class="tableblock halign-left valign-middle"><p class="tableblock">This capture failed because of a server side issue</p></td>
<td class="tableblock halign-left valign-middle"><p class="tableblock"><a href="#_error">Error</a></p></td>
</tr>
<tr>
<td class="tableblock halign-left valign-middle"><p class="tableblock"><strong>default</strong></p></td>
<td class="tableblock halign-left valign-middle"><p class="tableblock">An error occurred.</p></td>
<td class="tableblock halign-left valign-middle"><p class="tableblock"><a href="#_error">Error</a></p></td>
</tr>
</tbody>
</table>
</div>
</div>
<div class="sect3">
<h4 id="_v1_iris-captures_post">Associate an iris image capture (or pair of image captures) with the ongoing transaction</h4>
<div class="literalblock">
<div class="content">
<pre>POST /v1/iris-captures</pre>
</div>
</div>
<div class="sect4">
<h5 id="_description_3">Description</h5>
<div class="paragraph">
<p>Rally Participants may submit iris image data through this endpoint. Submitted image captures will be
stored and associated with the ongoing transaction. Multiple images may be submitted during a transaction.</p>
</div>
</div>
<div class="sect4">
<h5 id="_parameters_3">Parameters</h5>
<table class="tableblock frame-all grid-all spread">
<colgroup>
<col style="width: 11.1111%;">
<col style="width: 16.6666%;">
<col style="width: 50%;">
<col style="width: 22.2223%;">
</colgroup>
<thead>
<tr>
<th class="tableblock halign-left valign-middle">Type</th>
<th class="tableblock halign-left valign-middle">Name</th>
<th class="tableblock halign-left valign-middle">Description</th>
<th class="tableblock halign-left valign-middle">Schema</th>
</tr>
</thead>
<tbody>
<tr>
<td class="tableblock halign-left valign-middle"><p class="tableblock"><strong>Body</strong></p></td>
<td class="tableblock halign-left valign-middle"><p class="tableblock"><strong>iris_capture</strong><br>
<em>required</em></p></td>
<td class="tableblock halign-left valign-middle"><p class="tableblock">iris image(s) collected as part of a transaction</p></td>
<td class="tableblock halign-left valign-middle"><p class="tableblock"><a href="#_iriscapture">IrisCapture</a></p></td>
</tr>
</tbody>
</table>
</div>
<div class="sect4">
<h5 id="_responses_3">Responses</h5>
<table class="tableblock frame-all grid-all spread">
<colgroup>
<col style="width: 10%;">
<col style="width: 70%;">
<col style="width: 20%;">
</colgroup>
<thead>
<tr>
<th class="tableblock halign-left valign-middle">HTTP Code</th>
<th class="tableblock halign-left valign-middle">Description</th>
<th class="tableblock halign-left valign-middle">Schema</th>
</tr>
</thead>
<tbody>
<tr>
<td class="tableblock halign-left valign-middle"><p class="tableblock"><strong>200</strong></p></td>
<td class="tableblock halign-left valign-middle"><p class="tableblock">Successfully associated the image capture(s) with the ongoing transaction.</p></td>
<td class="tableblock halign-left valign-middle"><p class="tableblock">No Content</p></td>
</tr>
<tr>
<td class="tableblock halign-left valign-middle"><p class="tableblock"><strong>400</strong></p></td>
<td class="tableblock halign-left valign-middle"><p class="tableblock">The request is malformed.</p></td>
<td class="tableblock halign-left valign-middle"><p class="tableblock"><a href="#_error">Error</a></p></td>
</tr>
<tr>
<td class="tableblock halign-left valign-middle"><p class="tableblock"><strong>404</strong></p></td>
<td class="tableblock halign-left valign-middle"><p class="tableblock">No ongoing transaction was found for the specified StationID.</p></td>
<td class="tableblock halign-left valign-middle"><p class="tableblock"><a href="#_error">Error</a></p></td>
</tr>
<tr>
<td class="tableblock halign-left valign-middle"><p class="tableblock"><strong>429</strong></p></td>
<td class="tableblock halign-left valign-middle"><p class="tableblock">Too many requests for the ongoing transaction.</p></td>
<td class="tableblock halign-left valign-middle"><p class="tableblock"><a href="#_error">Error</a></p></td>
</tr>
<tr>
<td class="tableblock halign-left valign-middle"><p class="tableblock"><strong>500</strong></p></td>
<td class="tableblock halign-left valign-middle"><p class="tableblock">This capture failed because of a server side issue</p></td>
<td class="tableblock halign-left valign-middle"><p class="tableblock"><a href="#_error">Error</a></p></td>
</tr>
<tr>
<td class="tableblock halign-left valign-middle"><p class="tableblock"><strong>default</strong></p></td>
<td class="tableblock halign-left valign-middle"><p class="tableblock">An error occurred.</p></td>
<td class="tableblock halign-left valign-middle"><p class="tableblock"><a href="#_error">Error</a></p></td>
</tr>
</tbody>
</table>
</div>
</div>
</div>
</div>
</div>
<div class="sect1">
<h2 id="_definitions">Definitions</h2>
<div class="sectionbody">
<div class="sect2">
<h3 id="_error">Error</h3>
<table class="tableblock frame-all grid-all spread">
<colgroup>
<col style="width: 16.6666%;">
<col style="width: 61.1111%;">
<col style="width: 22.2223%;">
</colgroup>
<thead>
<tr>
<th class="tableblock halign-left valign-middle">Name</th>
<th class="tableblock halign-left valign-middle">Description</th>
<th class="tableblock halign-left valign-middle">Schema</th>
</tr>
</thead>
<tbody>
<tr>
<td class="tableblock halign-left valign-middle"><p class="tableblock"><strong>Code</strong><br>
<em>optional</em></p></td>
<td class="tableblock halign-left valign-middle"><p class="tableblock">The error code.</p></td>
<td class="tableblock halign-left valign-middle"><p class="tableblock">integer (int32)</p></td>
</tr>
<tr>
<td class="tableblock halign-left valign-middle"><p class="tableblock"><strong>ErrorString</strong><br>
<em>optional</em></p></td>
<td class="tableblock halign-left valign-middle"><p class="tableblock">A detailed description of the error that occurred.</p></td>
<td class="tableblock halign-left valign-middle"><p class="tableblock">string</p></td>
</tr>
<tr>
<td class="tableblock halign-left valign-middle"><p class="tableblock"><strong>Message</strong><br>
<em>optional</em></p></td>
<td class="tableblock halign-left valign-middle"><p class="tableblock">A description of the returned code.</p></td>
<td class="tableblock halign-left valign-middle"><p class="tableblock">string</p></td>
</tr>
</tbody>
</table>
</div>
<div class="sect2">
<h3 id="_facecapture">FaceCapture</h3>
<table class="tableblock frame-all grid-all spread">
<colgroup>
<col style="width: 16.6666%;">
<col style="width: 61.1111%;">
<col style="width: 22.2223%;">
</colgroup>
<thead>
<tr>
<th class="tableblock halign-left valign-middle">Name</th>
<th class="tableblock halign-left valign-middle">Description</th>
<th class="tableblock halign-left valign-middle">Schema</th>
</tr>
</thead>
<tbody>
<tr>
<td class="tableblock halign-left valign-middle"><p class="tableblock"><strong>FaceImageData</strong><br>
<em>required</em></p></td>
<td class="tableblock halign-left valign-middle"><p class="tableblock">The captured face image data in PNG format, encoded as a base64 string. The data string shall not exceed 1MB.<br>
<strong>Example</strong> : <code>"iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAIAAACQd1PeAAAAEElEQVR4nGJiYGAABAAA//8ADAADcZGLFwAAAABJRU5ErkJggg=="</code></p></td>
<td class="tableblock halign-left valign-middle"><p class="tableblock">string</p></td>
</tr>
<tr>
<td class="tableblock halign-left valign-middle"><p class="tableblock"><strong>StationID</strong><br>
<em>required</em></p></td>
<td class="tableblock halign-left valign-middle"><p class="tableblock">The unique identifier for the station that the image was captured at.<br>
<strong>Example</strong> : <code>"MdTF_Station"</code></p></td>
<td class="tableblock halign-left valign-middle"><p class="tableblock">string</p></td>
</tr>
</tbody>
</table>
</div>
<div class="sect2">
<h3 id="_fingercapture">FingerCapture</h3>
<table class="tableblock frame-all grid-all spread">
<colgroup>
<col style="width: 16.6666%;">
<col style="width: 61.1111%;">
<col style="width: 22.2223%;">
</colgroup>
<thead>
<tr>
<th class="tableblock halign-left valign-middle">Name</th>
<th class="tableblock halign-left valign-middle">Description</th>
<th class="tableblock halign-left valign-middle">Schema</th>
</tr>
</thead>
<tbody>
<tr>
<td class="tableblock halign-left valign-middle"><p class="tableblock"><strong>FingerImages</strong><br>
<em>optional</em></p></td>
<td class="tableblock halign-left valign-middle"><p class="tableblock">The captured finger(s). At least one FingerImage must be submitted with the request.</p></td>
<td class="tableblock halign-left valign-middle"><p class="tableblock">&lt; <a href="#_fingerimage">FingerImage</a> &gt; array</p></td>
</tr>
<tr>
<td class="tableblock halign-left valign-middle"><p class="tableblock"><strong>StationID</strong><br>
<em>required</em></p></td>
<td class="tableblock halign-left valign-middle"><p class="tableblock">The unique identifier for the station that the image was captured at.<br>
<strong>Example</strong> : <code>"MdTF_Station"</code></p></td>
<td class="tableblock halign-left valign-middle"><p class="tableblock">string</p></td>
</tr>
</tbody>
</table>
</div>
<div class="sect2">
<h3 id="_fingerimage">FingerImage</h3>
<table class="tableblock frame-all grid-all spread">
<colgroup>
<col style="width: 16.6666%;">
<col style="width: 61.1111%;">
<col style="width: 22.2223%;">
</colgroup>
<thead>
<tr>
<th class="tableblock halign-left valign-middle">Name</th>
<th class="tableblock halign-left valign-middle">Description</th>
<th class="tableblock halign-left valign-middle">Schema</th>
</tr>
</thead>
<tbody>
<tr>
<td class="tableblock halign-left valign-middle"><p class="tableblock"><strong>FingerImageData</strong><br>
<em>optional</em></p></td>
<td class="tableblock halign-left valign-middle"><p class="tableblock">The captured finger image data, encoded as a base64 string. The data string shall not exceed 100kB.<br>
<strong>Example</strong> : <code>"iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAIAAACQd1PeAAAAEElEQVR4nGJiYGAABAAA//8ADAADcZGLFwAAAABJRU5ErkJggg=="</code></p></td>
<td class="tableblock halign-left valign-middle"><p class="tableblock">string</p></td>
</tr>
<tr>
<td class="tableblock halign-left valign-middle"><p class="tableblock"><strong>FingerType</strong><br>
<em>required</em></p></td>
<td class="tableblock halign-left valign-middle"></td>
<td class="tableblock halign-left valign-middle"><p class="tableblock"><a href="#_fingertype">FingerType</a></p></td>
</tr>
</tbody>
</table>
</div>
<div class="sect2">
<h3 id="_fingertype">FingerType</h3>
<div class="paragraph">
<p>The allowed string names for finger images.</p>
</div>
<div class="paragraph">
<p><em>Type</em> : enum (RightThumb, RightIndexFinger, RightMiddleFinger, RightRingFinger, RightLittleFinger, LeftThumb, LeftIndexFinger, LeftMiddleFinger, LeftRingFinger, LeftLittleFinger)</p>
</div>
</div>
<div class="sect2">
<h3 id="_iriscapture">IrisCapture</h3>
<table class="tableblock frame-all grid-all spread">
<colgroup>
<col style="width: 16.6666%;">
<col style="width: 61.1111%;">
<col style="width: 22.2223%;">
</colgroup>
<thead>
<tr>
<th class="tableblock halign-left valign-middle">Name</th>
<th class="tableblock halign-left valign-middle">Description</th>
<th class="tableblock halign-left valign-middle">Schema</th>
</tr>
</thead>
<tbody>
<tr>
<td class="tableblock halign-left valign-middle"><p class="tableblock"><strong>LeftIrisImageData</strong><br>
<em>optional</em></p></td>
<td class="tableblock halign-left valign-middle"><p class="tableblock">The captured left iris image data, encoded as a base64 string. The data string shall not exceed 512kB. At least one iris image data field MUST be filled.<br>
<strong>Example</strong> : <code>"iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAIAAACQd1PeAAAAEElEQVR4nGJiYGAABAAA//8ADAADcZGLFwAAAABJRU5ErkJggg=="</code></p></td>
<td class="tableblock halign-left valign-middle"><p class="tableblock">string</p></td>
</tr>
<tr>
<td class="tableblock halign-left valign-middle"><p class="tableblock"><strong>RightIrisImageData</strong><br>
<em>optional</em></p></td>
<td class="tableblock halign-left valign-middle"><p class="tableblock">The captured right iris image data, encoded as a base64 string. The data string shall not exceed 512kB. At least one iris image data field MUST be filled.<br>
<strong>Example</strong> : <code>"iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAIAAACQd1PeAAAAEElEQVR4nGJiYGAABAAA//8ADAADcZGLFwAAAABJRU5ErkJggg=="</code></p></td>
<td class="tableblock halign-left valign-middle"><p class="tableblock">string</p></td>
</tr>
<tr>
<td class="tableblock halign-left valign-middle"><p class="tableblock"><strong>StationID</strong><br>
<em>required</em></p></td>
<td class="tableblock halign-left valign-middle"><p class="tableblock">The unique identifier for the station that the image was captured at.<br>
<strong>Example</strong> : <code>"MdTF_Station"</code></p></td>
<td class="tableblock halign-left valign-middle"><p class="tableblock">string</p></td>
</tr>
</tbody>
</table>
</div>
</div>
</div>
</div>
<div id="footer">
<div id="footer-text">
Last updated 2018-11-15 16:28:55 -05:00
</div>
</div>
</body>
</html>