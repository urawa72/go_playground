FROM golang:latest
RUN mkdir /go/src/work
WORKDIR /go/src/work
ADD ./src /go/src/work
