FROM golang:1.4-cross

MAINTAINER George Lewis <schvin@schvin.net>
ENV REFRESHED_AT 2015-01-23

RUN apt-get update -y && apt-get install -y git cmake pkg-config libssl-dev libssh2-1-dev libgit2-dev

ADD . /go/src/github.com/gophergala/bron

RUN go install github.com/gophergala/bron
RUN go install github.com/gophergala/bron/api
RUN go get -d github.com/libgit2/git2go
