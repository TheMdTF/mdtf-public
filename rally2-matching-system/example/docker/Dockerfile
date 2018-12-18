FROM golang:1.11.2 as builder

#add the source
RUN mkdir /build
ADD . /build/
WORKDIR /build

#move the models onto the GOPATH
RUN mkdir -p $GOPATH/src/github.com/mdtf-public/rally2-matching-system/example
RUN mv models $GOPATH/src/github.com/mdtf-public/rally2-matching-system/example

#build the c implementation
RUN make -C /build/c

#install dependencies
RUN go get github.com/disintegration/imaging

#build the go wrapper
RUN GOOS=linux GOARCH=amd64 go build -o ./rally2-matching-system .

#####################
#build a clean image
FROM ubuntu:16.04

COPY --from=builder /build/rally2-matching-system .

ENTRYPOINT ["/rally2-matching-system"]

EXPOSE 8080