FROM golang:1.11-alpine AS builder

# Install tools required to build the project
# We will need to run `docker build --no-cache .` to update those dependencies
RUN apk add --no-cache git build-base
RUN go get github.com/golang/dep/cmd/dep

#add the source
COPY . /go/src/project
WORKDIR /go/src/project/

# Install library dependencies
RUN dep ensure -vendor-only

#move the models onto the GOPATH
RUN mkdir -p $GOPATH/src/github.com/mdtf-public/rally2-matching-system/go-example
RUN mv models $GOPATH/src/github.com/mdtf-public/rally2-matching-system/go-example

#build the c implementation
RUN make -C /go/src/project/c

#build the go wrapper
RUN GOOS=linux GOARCH=amd64 go build -o ./rally2-matching-system .

#####################
#build a clean image
FROM alpine

COPY --from=builder /go/src/project/rally2-matching-system .

ENTRYPOINT ["/rally2-matching-system"]

EXPOSE 8080