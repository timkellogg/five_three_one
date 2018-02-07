FROM golang:1.8.3-alpine

MAINTAINER Tim Kellogg <tim.kellogg@gmail.com>

ENV GOBIN /go/bin
ENV GOPATH /go
ENV APPPATH $GOPATH/src/github.com/timkellogg/five_three_one
ENV PORT 3000

ADD . $APPPATH

WORKDIR $APPPATH

EXPOSE $PORT

RUN go build
RUN ./five_three_one
