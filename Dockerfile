FROM golang:1.7.3-alpine

MAINTAINER Dongri Jin <dongrify@gmail.com>

RUN apk add --update git

ADD . /go/src/mitty.co/mitty-server
WORKDIR /go/src/mitty.co/mitty-server
RUN go install mitty.co/mitty-server

ENTRYPOINT /go/bin/mitty-server

EXPOSE 8000
