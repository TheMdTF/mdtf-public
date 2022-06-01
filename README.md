# The Maryland Test Facility Public Code and Data Repository

Public code and data from the [Maryland Test Facility](https://mdtf.org).  

### datasets

This folder contains publicly releasable datasets from the Maryland Test Facility designed to further research across the biometrics and computer visison communities.

### image-analysis

* api

	This folder contains the API definition for generic image analysis systems submitted to the MdTF for evaluation.
	
* go-example

	This folder contains a working go example implementation of the image-analysis API definition. See example README for more details.

### rally5 

 * acquisition-system

    * api
    
      Contains OAS3 Spec and a generated [README.md](./rally5/acquisition-system/api/README.md) for Rally5 acquisition system api. Acquisition System vendors will interact with this API to submit acquired images.

 * matching-system

   * api 

     Contains OAS3 Spec and a generated [README.md](./rally5/matching-system/api/README.md) for Rally5 matching system api. Matching System vendors will produce a matching system container image implementing the provided API to generate and compare biometric templates.

### deprecated

Contains deprecated APIs from past Collection Events, Rallies, ETC.

* #### rally2
  * rally2-acquisition-system

    * api
    
        This folder contains the API definition for acquisition systems participating in the 2019 Biometric Technology Rally. The API is open for comment as of November 11, 2018. Comments can be made via GitHub Issues.
  * rally2-matching-system

    * api

        This folder contains the API definition for matching systems participating in the 2019 Biometric Technology Rally. The API is open for comment as of November 11, 2018. Comments can be made via GitHub Issues.

    * example

        This folder contains a working example of a matching system implemented in C with endpoints exposed via golang. See example README for more details.