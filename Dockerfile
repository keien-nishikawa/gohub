# setting golang version
FROM golang:1.18.2-alpine as dev
# udpate apk and install git
RUN apk update && apk add git
# lang config
ENV LANG C.UTF-8
# create app directory
RUN mkdir /go/src/gohub_api
# config working directory
WORKDIR /go/src/gohub_api
# move working directory and add to container files
ADD . /go/src/gohub_api
