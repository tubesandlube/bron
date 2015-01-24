FROM golang:1.4-cross

MAINTAINER George Lewis <schvin@schvin.net>
MAINTAINER Charlie Lewis <defermat@defermat.net>

ENV REFRESHED_AT 2015-01-23

RUN apt-get update -y && apt-get install -y \
    cmake \
    git \
    libgit2-dev \
    libssh2-1-dev \
    libssl-dev \
    pkg-config

ADD . /go/src/github.com/gophergala/bron
RUN go get github.com/gophergala/bron

WORKDIR /go/src/github.com/gophergala/bron

ENTRYPOINT ["/go/bin/bron"]
CMD [""]
