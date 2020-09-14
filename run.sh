#!/bin/sh

docker build -t my_golang .
docker run -it -v $(pwd)/src:/app --name my_golang my_golang:latest
