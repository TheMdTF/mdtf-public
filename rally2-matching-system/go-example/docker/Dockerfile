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
# Build C Library
########

#build
RUN make -C /go/src/project/c

########
# Build Go Wrapper
########

# Install tools
RUN go get github.com/golang/dep/cmd/dep

# Install library dependencies
RUN dep ensure -vendor-only

#move the models onto the GOPATH
RUN mkdir -p $GOPATH/src/github.com/TheMdTF/mdtf-public/rally2-matching-system/go-example
RUN mv models $GOPATH/src/github.com/TheMdTF/mdtf-public/rally2-matching-system/go-example

#build the go wrapper
RUN GOOS=linux GOARCH=amd64 go build -o ./rally2-matching-system .

########
# Package
########
FROM alpine

COPY --from=builder /go/src/project/rally2-matching-system .

ENTRYPOINT ["/rally2-matching-system"]

EXPOSE 8080
