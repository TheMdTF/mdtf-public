FROM golang:1.11-alpine AS builder

########
# Prep
########
#install apk tools
RUN apk add --no-cache git make gcc libc-dev

#add the source
COPY . /go/src/project
WORKDIR /go/src/project/

########
# Build Go Wrapper
########

# Install tools
RUN go get github.com/golang/dep/cmd/dep

# Install library dependencies
RUN dep ensure -vendor-only

# move the models and example algorithm onto the GOPATH
RUN mkdir -p $GOPATH/src/github.com/TheMdTF/mdtf-public/image-analysis/go-example
RUN mv models $GOPATH/src/github.com/TheMdTF/mdtf-public/image-analysis/go-example
RUN mv algorithm $GOPATH/src/github.com/TheMdTF/mdtf-public/image-analysis/go-example

#build the go wrapper
RUN GOOS=linux GOARCH=amd64 go build -o ./image-analysis .

########
# Package
########
FROM alpine

COPY --from=builder /go/src/project/image-analysis .

ENTRYPOINT ["/image-analysis"]

EXPOSE 8080
