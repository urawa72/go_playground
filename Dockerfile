FROM golang:latest
WORKDIR /go/src
RUN go get -u github.com/golang/dep/cmd/dep
RUN apt-get update && apt-get install -y curl git
ADD ./src /go/src
