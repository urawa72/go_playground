FROM golang:latest
WORKDIR /go/src
RUN go get -u github.com/golang/dep/cmd/dep
RUN apt-get update && apt-get install -y curl git
ENV TERM=xterm-256color
ADD ./src /go/src
WORKDIR /
COPY docker-entrypoint.sh ./
RUN chmod +x docker-entrypoint.sh
ENTRYPOINT ["./docker-entrypoint.sh"]
