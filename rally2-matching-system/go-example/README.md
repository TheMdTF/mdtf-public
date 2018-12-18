# Sample Biometric Technology Rally II Matching System

## Overview
This folder contains a simple application, written in [The Go Programming Language](https://golang.org/).  This application implements the [Rally Matching System API](https://github.com/TheMdTF/mdtf-public/blob/master/api/rally2-matching-system/README.md).  Creating an application that implements this API is required for matching system providers to participate in the [2019 Biometric Technology Rally](https://mdtf.org/Rally2019).  Finished applications should be packaged and saved using [Docker](https://www.docker.com/).

It is not required that matching system providers implement the matching system API in Go, as the application in this folder does.  Matching system providers are free to use whatever technology stack they choose, as long as the final docker container conforms to the API.  Scripts that send API conformant requests to a user provided endpoint are provided in the `/api/rally2-matching-system/test` folder.

## Running and Testing via Docker
The sample application can be packaged and run using the scripts in `./docker` and tested using the scripts in `<mdtf-public>/api/rally2-matching-system/test`.

The script `./docker/save-docker-image.sh` is an example which outputs the docker image in the format we will expect for delivery to the MdTF.

## Manually Building and Running
To run the sample application using Go:
 * `go get github.com/mdtf-public/rally2-matching-system/example`	
 * `go get github.com/golang/dep/cmd/dep`
 * from the package directory on your GOPATH
   * dep ensure -vendor-only
   * `make run`