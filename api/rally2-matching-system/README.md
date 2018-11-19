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
<h1>The Maryland Test Facility Matching System Interface</h1>
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
<li><a href="#_comparison_resource">Comparison</a></li>
<li><a href="#_extraction_resource">Extraction</a></li>
<li><a href="#_information_resource">Information</a></li>
</ul>
</li>
<li><a href="#_definitions">Definitions</a>
<ul class="sectlevel2">
<li><a href="#_comparelistrequest">CompareListRequest</a></li>
<li><a href="#_comparison">Comparison</a></li>
<li><a href="#_image">Image</a></li>
<li><a href="#_info">Info</a></li>
<li><a href="#_template">Template</a></li>
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
<p>DRAFT Applicaton Programming Interface for receiving matching system requests from the MdTF Backend.</p>
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
<p><em>Contact</em> : John Howard<br>
<em>Contact Email</em> : <a href="mailto:john@mdtf.org">john@mdtf.org</a></p>
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
<p>Comparison : Endpoints related to functionality that compares one or more biometric templates</p>
</li>
<li>
<p>Extraction : Endpoints related to functionality that extracts a biometric template from a biometric sample.</p>
</li>
<li>
<p>Information : Endpoints related to functionality that provides information on biometric algorithms</p>
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
<h3 id="_comparison_resource">Comparison</h3>
<div class="paragraph">
<p>Endpoints related to functionality that compares one or more biometric templates</p>
</div>
<div class="sect3">
<h4 id="_compare_list">Compares a template to a list of templates.</h4>
<div class="literalblock">
<div class="content">
<pre>POST /v1/compare-list/</pre>
</div>
</div>
<div class="sect4">
<h5 id="_description">Description</h5>
<div class="paragraph">
<p>This endpoint takes a template and a list of templates. It compares the single template to every other template in the list.
The result is a list of Comparison objects that holds a similarity score for each comparison.</p>
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
<td class="tableblock halign-left valign-middle"><p class="tableblock"><strong>compare_list_request</strong><br>
<em>required</em></p></td>
<td class="tableblock halign-left valign-middle"><p class="tableblock">A single template object and a list of templates that it will be compared to.</p></td>
<td class="tableblock halign-left valign-middle"><p class="tableblock"><a href="#_comparelistrequest">CompareListRequest</a></p></td>
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
<td class="tableblock halign-left valign-middle"><p class="tableblock">Successful response</p></td>
<td class="tableblock halign-left valign-middle"><p class="tableblock">&lt; <a href="#_comparison">Comparison</a> &gt; array</p></td>
</tr>
<tr>
<td class="tableblock halign-left valign-middle"><p class="tableblock"><strong>406</strong></p></td>
<td class="tableblock halign-left valign-middle"><p class="tableblock">Status not acceptable</p></td>
<td class="tableblock halign-left valign-middle"><p class="tableblock">No Content</p></td>
</tr>
<tr>
<td class="tableblock halign-left valign-middle"><p class="tableblock"><strong>500</strong></p></td>
<td class="tableblock halign-left valign-middle"><p class="tableblock">Internal server error</p></td>
<td class="tableblock halign-left valign-middle"><p class="tableblock">No Content</p></td>
</tr>
</tbody>
</table>
</div>
</div>
</div>
<div class="sect2">
<h3 id="_extraction_resource">Extraction</h3>
<div class="paragraph">
<p>Endpoints related to functionality that extracts a biometric template from a biometric sample.</p>
</div>
<div class="sect3">
<h4 id="_create_template">Generates a template from a biometric image.</h4>
<div class="literalblock">
<div class="content">
<pre>POST /v1/create-template/</pre>
</div>
</div>
<div class="sect4">
<h5 id="_description_2">Description</h5>
<div class="paragraph">
<p>This endpoint takes a base64 encoded PNG and performs a <code>feature extraction</code> operation resulting in a single template being produced</p>
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
<td class="tableblock halign-left valign-middle"><p class="tableblock"><strong>image</strong><br>
<em>required</em></p></td>
<td class="tableblock halign-left valign-middle"><p class="tableblock">Image that will be used for extraction. This object contains a base64 encoded PNG and a BiometricImageType.</p></td>
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
<td class="tableblock halign-left valign-middle"><p class="tableblock">Successful response</p></td>
<td class="tableblock halign-left valign-middle"><p class="tableblock"><a href="#_template">Template</a></p></td>
</tr>
<tr>
<td class="tableblock halign-left valign-middle"><p class="tableblock"><strong>406</strong></p></td>
<td class="tableblock halign-left valign-middle"><p class="tableblock">Status not acceptable</p></td>
<td class="tableblock halign-left valign-middle"><p class="tableblock">No Content</p></td>
</tr>
<tr>
<td class="tableblock halign-left valign-middle"><p class="tableblock"><strong>500</strong></p></td>
<td class="tableblock halign-left valign-middle"><p class="tableblock">Internal server error</p></td>
<td class="tableblock halign-left valign-middle"><p class="tableblock">No Content</p></td>
</tr>
</tbody>
</table>
</div>
</div>
</div>
<div class="sect2">
<h3 id="_information_resource">Information</h3>
<div class="paragraph">
<p>Endpoints related to functionality that provides information on biometric algorithms</p>
</div>
<div class="sect3">
<h4 id="_info">Returns basic information of the algorithm.</h4>
<div class="literalblock">
<div class="content">
<pre>GET /v1/info/</pre>
</div>
</div>
<div class="sect4">
<h5 id="_description_3">Description</h5>
<div class="paragraph">
<p>Returns information about the algorithm. This can include Name, SdkManufacturer, SdkModel, SdkVersion, Version, and ValidTypes.</p>
</div>
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
<td class="tableblock halign-left valign-middle"><p class="tableblock">Successful response</p></td>
<td class="tableblock halign-left valign-middle"><p class="tableblock"><a href="#_info">Info</a></p></td>
</tr>
<tr>
<td class="tableblock halign-left valign-middle"><p class="tableblock"><strong>406</strong></p></td>
<td class="tableblock halign-left valign-middle"><p class="tableblock">Status not acceptable</p></td>
<td class="tableblock halign-left valign-middle"><p class="tableblock">No Content</p></td>
</tr>
<tr>
<td class="tableblock halign-left valign-middle"><p class="tableblock"><strong>500</strong></p></td>
<td class="tableblock halign-left valign-middle"><p class="tableblock">Internal server error</p></td>
<td class="tableblock halign-left valign-middle"><p class="tableblock">No Content</p></td>
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
<h3 id="_comparelistrequest">CompareListRequest</h3>
<div class="paragraph">
<p>A single template object and a list of templates that it will be compared to.</p>
</div>
<table class="tableblock frame-all grid-all spread">
<colgroup>
<col style="width: 42.8571%;">
<col style="width: 57.1429%;">
</colgroup>
<thead>
<tr>
<th class="tableblock halign-left valign-middle">Name</th>
<th class="tableblock halign-left valign-middle">Schema</th>
</tr>
</thead>
<tbody>
<tr>
<td class="tableblock halign-left valign-middle"><p class="tableblock"><strong>SingleTemplate</strong><br>
<em>optional</em></p></td>
<td class="tableblock halign-left valign-middle"><p class="tableblock"><a href="#_template">Template</a></p></td>
</tr>
<tr>
<td class="tableblock halign-left valign-middle"><p class="tableblock"><strong>TemplateList</strong><br>
<em>optional</em></p></td>
<td class="tableblock halign-left valign-middle"><p class="tableblock">&lt; <a href="#_template">Template</a> &gt; array</p></td>
</tr>
</tbody>
</table>
</div>
<div class="sect2">
<h3 id="_comparison">Comparison</h3>
<div class="paragraph">
<p>Comparison contains a similarity score for a single templare comparison operation.</p>
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
<em>optional</em></p></td>
<td class="tableblock halign-left valign-middle"><p class="tableblock">Name of the algorithm used to create this comparison</p></td>
<td class="tableblock halign-left valign-middle"><p class="tableblock">string</p></td>
</tr>
<tr>
<td class="tableblock halign-left valign-middle"><p class="tableblock"><strong>NormalizedScore</strong><br>
<em>optional</em></p></td>
<td class="tableblock halign-left valign-middle"><p class="tableblock">Similarity score between 0 and 1, with 1 being the highest score the algorithm can produce</p></td>
<td class="tableblock halign-left valign-middle"><p class="tableblock">number (float)</p></td>
</tr>
<tr>
<td class="tableblock halign-left valign-middle"><p class="tableblock"><strong>Score</strong><br>
<em>optional</em></p></td>
<td class="tableblock halign-left valign-middle"><p class="tableblock">Un-normalized similarity score, raw score as produced by the algorithm</p></td>
<td class="tableblock halign-left valign-middle"><p class="tableblock">number (float)</p></td>
</tr>
</tbody>
</table>
</div>
<div class="sect2">
<h3 id="_image">Image</h3>
<div class="paragraph">
<p>Object that contains image data.</p>
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
<em>optional</em></p></td>
<td class="tableblock halign-left valign-middle"><p class="tableblock">The captured image data in PNG format, encoded as a base64 string. The data string shall not exceed 1MB.<br>
<strong>Example</strong> : <code>"iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAIAAACQd1PeAAAAEElEQVR4nGJiYGAABAAA//8ADAADcZGLFwAAAABJRU5ErkJggg=="</code></p></td>
<td class="tableblock halign-left valign-middle"><p class="tableblock">string</p></td>
</tr>
</tbody>
</table>
</div>
<div class="sect2">
<h3 id="_info">Info</h3>
<div class="paragraph">
<p>Basic information of the algorithm.</p>
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
<td class="tableblock halign-left valign-middle"><p class="tableblock"><strong>Name</strong><br>
<em>optional</em></p></td>
<td class="tableblock halign-left valign-middle"><p class="tableblock">Name of algorithm</p></td>
<td class="tableblock halign-left valign-middle"><p class="tableblock">string</p></td>
</tr>
<tr>
<td class="tableblock halign-left valign-middle"><p class="tableblock"><strong>SdkManufacturer</strong><br>
<em>optional</em></p></td>
<td class="tableblock halign-left valign-middle"><p class="tableblock">Name of the manufacturer of the SDK</p></td>
<td class="tableblock halign-left valign-middle"><p class="tableblock">string</p></td>
</tr>
<tr>
<td class="tableblock halign-left valign-middle"><p class="tableblock"><strong>SdkModel</strong><br>
<em>optional</em></p></td>
<td class="tableblock halign-left valign-middle"><p class="tableblock">Identifier of the model SDK</p></td>
<td class="tableblock halign-left valign-middle"><p class="tableblock">string</p></td>
</tr>
<tr>
<td class="tableblock halign-left valign-middle"><p class="tableblock"><strong>ValidTypes</strong><br>
<em>optional</em></p></td>
<td class="tableblock halign-left valign-middle"><p class="tableblock">Array of integer indexes of valid biometric types for this algorithm</p></td>
<td class="tableblock halign-left valign-middle"><p class="tableblock">&lt; integer &gt; array</p></td>
</tr>
<tr>
<td class="tableblock halign-left valign-middle"><p class="tableblock"><strong>Version</strong><br>
<em>optional</em></p></td>
<td class="tableblock halign-left valign-middle"><p class="tableblock">Algorithm version identifier</p></td>
<td class="tableblock halign-left valign-middle"><p class="tableblock">string</p></td>
</tr>
</tbody>
</table>
</div>
<div class="sect2">
<h3 id="_template">Template</h3>
<div class="paragraph">
<p>Object that contains a byte array of template data.</p>
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
<td class="tableblock halign-left valign-middle"><p class="tableblock"><strong>Template</strong><br>
<em>optional</em></p></td>
<td class="tableblock halign-left valign-middle"><p class="tableblock">A byte array that contains template data.</p></td>
<td class="tableblock halign-left valign-middle"><p class="tableblock">&lt; string (byte) &gt; array</p></td>
</tr>
</tbody>
</table>
</div>
</div>
</div>
</div>
<div id="footer">
<div id="footer-text">
Last updated 2018-11-19 16:14:39 -05:00
</div>
</div>
</body>
</html>