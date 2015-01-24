FROM golang:1.4-cross

MAINTAINER George Lewis <schvin@schvin.net>
MAINTAINER Charlie Lewis <defermat@defermat.net>

ENV REFRESHED_AT 2015-01-23

RUN apt-get update -y && apt-get install -y git cmake pkg-config libssl-dev libssh2-1-dev libgit2-dev

ADD . /go/src/github.com/gophergala/bron

RUN go get github.com/gophergala/bron

ENTRYPOINT /go/bin/bron
