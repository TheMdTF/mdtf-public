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
<h1>The Maryland Test Facility Image Analysis Interface</h1>
<div id="toc" class="toc2">
<div id="toctitle">Table of Contents</div>
<ul class="sectlevel1">
<li><a href="#_overview">Overview</a>
<ul class="sectlevel2">
<li><a href="#_version_information">Version information</a></li>
<li><a href="#_contact_information">Contact information</a></li>
<li><a href="#_license_information">License information</a></li>
<li><a href="#_uri_scheme">URI scheme</a></li>
<li><a href="#_tags">Tags</a></li>
<li><a href="#_consumes">Consumes</a></li>
<li><a href="#_produces">Produces</a></li>
</ul>
</li>
<li><a href="#_paths">Resources</a>
<ul class="sectlevel2">
<li><a href="#_algorithm_information_resource">Algorithm Information</a></li>
<li><a href="#_analysis_operations_resource">Analysis Operations</a></li>
</ul>
</li>
<li><a href="#_definitions">Definitions</a>
<ul class="sectlevel2">
<li><a href="#_analysisresult">AnalysisResult</a></li>
<li><a href="#_image">Image</a></li>
<li><a href="#_info">Info</a></li>
<li><a href="#_requesterror">RequestError</a></li>
<li><a href="#_servererror">ServerError</a></li>
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
<p>Application programming interface for receiving image analysis requests.</p>
</div>
<div class="sect2">
<h3 id="_version_information">Version information</h3>
<div class="paragraph">
<p><em>Version</em> : 1.0.0</p>
</div>
</div>
<div class="sect2">
<h3 id="_contact_information">Contact information</h3>
<div class="paragraph">
<p><em>Contact</em> : MdTF<br>
<em>Contact Email</em> : <a href="mailto:info@mdtf.org">info@mdtf.org</a></p>
</div>
</div>
<div class="sect2">
<h3 id="_license_information">License information</h3>
<div class="paragraph">
<p><em>License</em> : Copyright (c) 2019, The Maryland Test Facility<br>
<em>Terms of service</em> : null</p>
</div>
</div>
<div class="sect2">
<h3 id="_uri_scheme">URI scheme</h3>
<div class="paragraph">
<p><em>Host</em> : 172.17.0.2:8080<br>
<em>Schemes</em> : HTTP</p>
</div>
</div>
<div class="sect2">
<h3 id="_tags">Tags</h3>
<div class="ulist">
<ul>
<li>
<p>Algorithm Information</p>
</li>
<li>
<p>Analysis Operations</p>
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
<h3 id="_algorithm_information_resource">Algorithm Information</h3>
<div class="sect3">
<h4 id="_info">Returns basic information for the algorithm.</h4>
<div class="literalblock">
<div class="content">
<pre>GET /v1/info</pre>
</div>
</div>
<div class="sect4">
<h5 id="_description">Description</h5>
<div class="paragraph">
<p>This endpoint returns some basic information about the image analysis algorithm.</p>
</div>
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
<td class="tableblock halign-left valign-middle"><p class="tableblock">Successful Response</p></td>
<td class="tableblock halign-left valign-middle"><p class="tableblock"><a href="#_info">Info</a></p></td>
</tr>
<tr>
<td class="tableblock halign-left valign-middle"><p class="tableblock"><strong>500</strong></p></td>
<td class="tableblock halign-left valign-middle"><p class="tableblock">Internal Server Error</p></td>
<td class="tableblock halign-left valign-middle"><p class="tableblock"><a href="#_servererror">ServerError</a></p></td>
</tr>
</tbody>
</table>
</div>
<div class="sect4">
<h5 id="_example_http_response">Example HTTP response</h5>
<div class="sect5">
<h6 id="_response_500">Response 500</h6>
<div class="listingblock">
<div class="content">
<pre class="highlight"><code class="language-json" data-lang="json">"\"Error encoding response.\""</code></pre>
</div>
</div>
</div>
</div>
</div>
</div>
<div class="sect2">
<h3 id="_analysis_operations_resource">Analysis Operations</h3>
<div class="sect3">
<h4 id="_analyze_image">Returns an object containing the results of analysis for the provided biometric image.</h4>
<div class="literalblock">
<div class="content">
<pre>POST /v1/analyze-image</pre>
</div>
</div>
<div class="sect4">
<h5 id="_description_2">Description</h5>
<div class="paragraph">
<p>This endpoint accepts a base64 encoded PNG image and returns an object containing the output of a generic image analysis routine on that image.</p>
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
<td class="tableblock halign-left valign-middle"><p class="tableblock"><strong>image</strong><br>
<em>required</em></p></td>
<td class="tableblock halign-left valign-middle"><p class="tableblock">The image that is being submitted for analysis.</p></td>
<td class="tableblock halign-left valign-middle"><p class="tableblock"><a href="#_image">Image</a></p></td>
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
<td class="tableblock halign-left valign-middle"><p class="tableblock">Successful Response</p></td>
<td class="tableblock halign-left valign-middle"><p class="tableblock"><a href="#_analysisresult">AnalysisResult</a></p></td>
</tr>
<tr>
<td class="tableblock halign-left valign-middle"><p class="tableblock"><strong>400</strong></p></td>
<td class="tableblock halign-left valign-middle"><p class="tableblock">Bad Request</p></td>
<td class="tableblock halign-left valign-middle"><p class="tableblock"><a href="#_requesterror">RequestError</a></p></td>
</tr>
<tr>
<td class="tableblock halign-left valign-middle"><p class="tableblock"><strong>500</strong></p></td>
<td class="tableblock halign-left valign-middle"><p class="tableblock">Internal Server Error</p></td>
<td class="tableblock halign-left valign-middle"><p class="tableblock"><a href="#_servererror">ServerError</a></p></td>
</tr>
</tbody>
</table>
</div>
<div class="sect4">
<h5 id="_example_http_response_2">Example HTTP response</h5>
<div class="sect5">
<h6 id="_response_200">Response 200</h6>
<div class="listingblock">
<div class="content">
<pre class="highlight"><code class="language-json" data-lang="json">{
  "AnalysisScore" : "5.0002",
  "NormalizedScore" : "0.50002",
  "ErrorLog" : ""
}</code></pre>
</div>
</div>
</div>
<div class="sect5">
<h6 id="_response_400">Response 400</h6>
<div class="listingblock">
<div class="content">
<pre class="highlight"><code class="language-json" data-lang="json">"\"Unable to decode image data as a PNG.\""</code></pre>
</div>
</div>
</div>
<div class="sect5">
<h6 id="_response_500_2">Response 500</h6>
<div class="listingblock">
<div class="content">
<pre class="highlight"><code class="language-json" data-lang="json">"\"Error encoding response.\""</code></pre>
</div>
</div>
</div>
</div>
</div>
</div>
</div>
</div>
<div class="sect1">
<h2 id="_definitions">Definitions</h2>
<div class="sectionbody">
<div class="sect2">
<h3 id="_analysisresult">AnalysisResult</h3>
<div class="paragraph">
<p>Object containing the results of a generic image analysis algorithm on a single image. Must include at least 1 and not more than 8 string properties in a JSON object with a depth of 1. The full contents shall not exceed 512 Bytes.</p>
</div>
<div class="paragraph">
<p><em>Type</em> : &lt; string, string &gt; map</p>
</div>
</div>
<div class="sect2">
<h3 id="_image">Image</h3>
<div class="paragraph">
<p>Data transfer object for an image.</p>
</div>
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
<td class="tableblock halign-left valign-middle"><p class="tableblock"><strong>ImageData</strong><br>
<em>required</em></p></td>
<td class="tableblock halign-left valign-middle"><p class="tableblock">The captured image data in PNG format, encoded as a base64 string. The data string shall not exceed 10MB.<br>
<strong>Example</strong> : <code>"iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAIAAACQd1PeAAAAEElEQVR4nGJiYGAABAAA//8ADAADcZGLFwAAAABJRU5ErkJggg=="</code></p></td>
<td class="tableblock halign-left valign-middle"><p class="tableblock">string</p></td>
</tr>
</tbody>
</table>
</div>
<div class="sect2">
<h3 id="_info">Info</h3>
<div class="paragraph">
<p>Basic information describing the algorithm.</p>
</div>
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
<td class="tableblock halign-left valign-middle"><p class="tableblock"><strong>AlgorithmName</strong><br>
<em>required</em></p></td>
<td class="tableblock halign-left valign-middle"><p class="tableblock">Name of algorithm<br>
<strong>Example</strong> : <code>"OpenFace"</code></p></td>
<td class="tableblock halign-left valign-middle"><p class="tableblock">string</p></td>
</tr>
<tr>
<td class="tableblock halign-left valign-middle"><p class="tableblock"><strong>AlgorithmType</strong><br>
<em>required</em></p></td>
<td class="tableblock halign-left valign-middle"><p class="tableblock">A string enum describing the type of biometric images the algorithm is meant to process</p></td>
<td class="tableblock halign-left valign-middle"><p class="tableblock">enum (Face, Finger, Iris)</p></td>
</tr>
<tr>
<td class="tableblock halign-left valign-middle"><p class="tableblock"><strong>AlgorithmVersion</strong><br>
<em>required</em></p></td>
<td class="tableblock halign-left valign-middle"><p class="tableblock">Algorithm version identifier<br>
<strong>Example</strong> : <code>"1.0.1"</code></p></td>
<td class="tableblock halign-left valign-middle"><p class="tableblock">string</p></td>
</tr>
<tr>
<td class="tableblock halign-left valign-middle"><p class="tableblock"><strong>CompanyName</strong><br>
<em>required</em></p></td>
<td class="tableblock halign-left valign-middle"><p class="tableblock">Name of the company which produces the algorithm<br>
<strong>Example</strong> : <code>"MdTF"</code></p></td>
<td class="tableblock halign-left valign-middle"><p class="tableblock">string</p></td>
</tr>
<tr>
<td class="tableblock halign-left valign-middle"><p class="tableblock"><strong>ImageDataset</strong><br>
<em>required</em></p></td>
<td class="tableblock halign-left valign-middle"><p class="tableblock">A string enum to select an MdTF dataset of biometric images for this submission to analyze</p></td>
<td class="tableblock halign-left valign-middle"><p class="tableblock">enum (NIST_MEDS, MDTF_RALLY2_ENROLLMENT)</p></td>
</tr>
<tr>
<td class="tableblock halign-left valign-middle"><p class="tableblock"><strong>RecommendedCPUs</strong><br>
<em>required</em></p></td>
<td class="tableblock halign-left valign-middle"><p class="tableblock">The recommended allocation of CPUs for the deployed docker container.<br>
<strong>Example</strong> : <code>4.0</code></p></td>
<td class="tableblock halign-left valign-middle"><p class="tableblock">number (integer)</p></td>
</tr>
<tr>
<td class="tableblock halign-left valign-middle"><p class="tableblock"><strong>RecommendedMem</strong><br>
<em>required</em></p></td>
<td class="tableblock halign-left valign-middle"><p class="tableblock">The recommended allocation of memory (MB) for the deployed docker container.<br>
<strong>Example</strong> : <code>2048.0</code></p></td>
<td class="tableblock halign-left valign-middle"><p class="tableblock">number (integer)</p></td>
</tr>
<tr>
<td class="tableblock halign-left valign-middle"><p class="tableblock"><strong>TechnicalContactEmail</strong><br>
<em>required</em></p></td>
<td class="tableblock halign-left valign-middle"><p class="tableblock">The email address of an engineer or other technical resource to contact in the event of an error running your service. This field may be left blank if desired.<br>
<strong>Example</strong> : <code>"<a href="mailto:info@mdtf.org">info@mdtf.org</a>"</code></p></td>
<td class="tableblock halign-left valign-middle"><p class="tableblock">string</p></td>
</tr>
</tbody>
</table>
</div>
<div class="sect2">
<h3 id="_requesterror">RequestError</h3>
<div class="paragraph">
<p>Relevant and concise diagnostic information in the event of a client side error (e.g. malformed requests, or invalid image encoding).</p>
</div>
<div class="paragraph">
<p><em>Type</em> : string</p>
</div>
</div>
<div class="sect2">
<h3 id="_servererror">ServerError</h3>
<div class="paragraph">
<p>Relevant and concise diagnostic information in the event of a server side error (e.g. invalid license, or failure to generate a template).</p>
</div>
<div class="paragraph">
<p><em>Type</em> : string</p>
</div>
</div>
</div>
</div>
</div>
<div id="footer">
<div id="footer-text">
Last updated 2020-02-18 09:59:03 -05:00
</div>
</div>
</body>
</html>