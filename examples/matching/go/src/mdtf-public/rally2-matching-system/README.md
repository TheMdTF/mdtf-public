# Sample Biometric Technology Rally II Matching System

##Overview
This folder contains a simple application, written in [The Go Programming Language](https://golang.org/).  This application implements the [Rally Matching System API](https://github.com/TheMdTF/mdtf-public/blob/master/api/rally2-matching-system/README.md).  Creating an application that implements this API is required for matching system providers to participate in the [2019 Biometric Technology Rally](https://mdtf.org/Rally2019).  Finished applications should be packaged and saved using [Docker](https://www.docker.com/).  Scripts that package and save this application are in the `rally2-matching-system/docker` folder.

It is not required that matching system providers implement the matching system API in Go, as the application in this folder does.  Matching system providers are free to use whatever technology stack they choose, as long as the final docker container conforms to the API.  Scripts that send API conformant requests to a user provided endpoint are provided in the `/api/rally2-matching-system/test` folder.

##Running and Testing via Docker
The sample application can be packaged and run using the scripts in `rally2-matching-system/docker` and tested using the scripts in `/api/rally2-matching-system/test`.

The script `docker/save-docker-image.sh` is an example which outputs the docker image in the format we will expect for delivery to the MdTF.

##Manually Building and Running
To run the sample application using Go:
 * Set your `GOPATH` to `mdtf-public/go`
 * Install the [Go Disintegration Imaging](https://godoc.org/github.com/disintegration/imaging) package via `go get https://github.com/disintegration/imaging`
 * Build the underlying C implementation of the sample algorithm service in `/c` with `make` 

The application can then be run via `go run main.go` or compiled via `go build -o ./rally2-mathcing-system`.