FROM golang:latest
RUN apt-get update && apt-get install -y curl git
ENV TERM=xterm-256color
WORKDIR /
COPY docker-entrypoint.sh ./
RUN chmod +x docker-entrypoint.sh
WORKDIR /work
ENTRYPOINT ["/docker-entrypoint.sh"]
