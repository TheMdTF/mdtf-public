# Sample Biometric Technology Rally II Matching System

## Overview
This folder contains a simple application, written in [The Go Programming Language](https://golang.org/).  This application implements the [Rally Matching System API](https://github.com/TheMdTF/mdtf-public/blob/master/api/rally2-matching-system/README.md).  Creating an application that implements this API is required for matching system providers to participate in the [2019 Biometric Technology Rally](https://mdtf.org/Rally2019).  Finished applications should be packaged and saved using [Docker](https://www.docker.com/).

It is not required that matching system providers implement the matching system API in Go, as the application in this folder does.  Matching system providers are free to use whatever technology stack they choose, as long as the final docker container conforms to the API.  Scripts that send API conformant requests to a user provided endpoint are provided in the `<mdtf-public>/rally2-matching-system/tests` folder.

## Running and Testing via Docker
*PREREQUISITES*:
 * A clone of the git repository
 * An installation of [docker](https://docs.docker.com/install/#supported-platforms)

To run the sample application via docker:
 * cd `./docker`
 * package the image using `./make-docker-image.sh`
 * run the image using `./run-docker-image.sh`

After the steps above, the server should be accessible at `localhost:8080`. The test scripts in `../tests` should send requests to this address and output the results to the command line.  The script `./docker/save-docker-image.sh` will output the docker image in the format we expect for delivery to the MdTF.  Please utilize the scripts' `[COMPANY_NAME]` command line parameter when preparing for delivery.

## Manually Building and Running With C/Go
*PREREQUISITES*:
 * A linux [Golang development environment](https://golang.org/doc/install)
 * gcc, cmake

To run the sample application using Go:
 * `go get -d github.com/TheMdTF/mdtf-public/rally2-matching-system/go-example`
 * `go get github.com/golang/dep/cmd/dep`
 * from `$GOPATH/src/github.com/TheMdTF/mdtf-public/rally2-matching-system/go-example`:
   * `dep ensure -vendor-only`
   * `make run`
