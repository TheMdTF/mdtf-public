# The Maryland Test Facility Public Code and Data Repository

Public API specifications and code examples from the [Maryland Test Facility](https://mdtf.org).

## Definitions

The following categories of systems are supported:
- **Acquisition** systems are systems that arrive on-site, collect or process data, and must communicate with an API during the span of a collection event. These systems often have an opportunity to test against a demo version of the API before they arrive on-site. You can think of these systems as client systems to the APIs.
- **Analysis** systems must implement an API and be submitted to process an analysis workflow. These systems will run in a controlled environment and be orchestrated as part of an analysis workflow. They must pass a validation process in order to be accepted for collaboration for testing. You can think of these systems as server implementations of the APIs.

## APIs

This directory catalogs all supported APIs for acquisition and analysis systems.

- **acquisition-system** [\[Acquisition\]](#Definitions)  
  This is the default API for acquisition systems that arrive at our facilities. Acquisition systems must be able to submit acquired data to this API when arriving at the facility for respective test events.
- **document-validation** [\[Analysis\]](#Definitions)  
  This API supports document validation systems. These analysis systems will be presented with documents and tasked with determining their validity, providing any additional data used to make that decision.
- **match-to-id** [\[Analysis\]](#Definitions)  
  This API supports match to ID systems. These analysis systems will be presented with face images and document images and tasked with comparison of the two artifacts to determine identity.
- **image-analysis** [\[Analysis\]](#Definitions)  
  This API is intended to support generic analysis workflows with image inputs.
- **matching-system** [\[Analysis\]](#Definitions)  
  This is the default API for matching systems that are submitted for analysis. It supports images producing biometric templates and the comparison of those templates.
- **pad-systems**
  - **active-pad-system** [\[Acquisition\]](#Definitions)  
    This API supports acquisition for active PAD systems that arrive at our facilities. Systems must submit acquired data as well as presentation attack detection results and analysis parameters.
  - **passive-pad-system** [\[Analysis\]](#Definitions)  
    This API supports passive PAD systems that are submitted for analysis. These systems must produce presentation attack detection results and analysis parameters when sent biometric data.

## Datasets

Contains example datasets for providers to test their systems.

## Tools

README Documentation from OpenAPI Specifications has been generated with [Widdershins](https://github.com/Mermade/widdershins). Installation instructions can be found on their GitHub page.


#### [LICENSE](./LICENSE.md)
