# The Maryland Test Facility Public Code and Data Repository

Public code and data from the [Maryland Test Facility](https://mdtf.org).

### apis
This directory has all the APIs developed for events at the MdTF

#### acquisition-system
* api

  Contains OAS3 Spec and a generated [README.md](apis/acquisition-system/api/README.md) for an acquisition system API. Acquisition system providers will interact with this API to submit acquired images.

#### document-validation
* api

    Contains the OAS3 Spec and generated [README.md](apis/document-validation/api/README.md) for an identity document validation system. Document validation providers will implement this API to validate documents.
#### match-to-id
* api

    Contains the OAS3 Spec and generated [README.md](apis/match-to-id/api/README.md) for systems that match face images to ID documents.  Matching system providers will implement this API to verify identities based on documents.
  
#### image-analysis
* api

    This folder contains the API definition for generic image analysis systems submitted to the MdTF for evaluation.

* go-example

    This folder contains a working go example implementation of the image-analysis API definition. See example README for more details.

#### matching-system
   * api

        Contains OAS3 Spec and a generated [README.md](apis/matching-system/api/README.md) for an matching system API. Matching system providers will produce a matching system container image implementing the provided API to generate and compare biometric templates.
   * go-example

        Has an example of a matching system, developed in Go.
   * tests

       Contains examples of some tests used to validate the functionality of a matching system.

### datasets

Contains datasets for use with provider systems, or for providers to test their systems

### past-events

Contains older APIs from past collection events, rallies, etc. An API may still be in use, but after a test event is completed, a copy is saved here.

#### rally2
* rally2-acquisition-system
  * API
  
      This folder contains the API definition for acquisition systems participating in the 2019 Biometric Technology Rally. The API is open for comment as of November 11, 2018. Comments can be made via GitHub Issues.
* rally2-matching-system

  * API

      This folder contains the API definition for matching systems participating in the 2019 Biometric Technology Rally. The API is open for comment as of November 11, 2018. Comments can be made via GitHub Issues.

  * example

      This folder contains a working example of a matching system implemented in C with endpoints exposed via golang. See example README for more details.
#### rally5
Has the acquisition system and matching system for the 2022 Biometric Rally.
* acquisition-system
    * api

        Contains OAS3 Spec and a generated [README.md](apis/acquisition-system/api/README.md) for an acquisition system API. Acquisition system providers will interact with this API to submit acquired images.
* matching-system
  * api

    Contains OAS3 Spec and a generated [README.md](apis/matching-system/api/README.md) for an matching system API. Matching system providers will produce a matching system container image implementing the provided API to generate and compare biometric templates.
  * go-example

    Has an example of a matching system, developed in Go.
  * tests

    Contains examples of some tests used to validate the functionality of a matching system.

### Tools

README Documentation from OpenAPI Specifications has been generated with [widdershins](https://github.com/Mermade/widdershins). Installation instructions can be found on their GitHub page.

   
