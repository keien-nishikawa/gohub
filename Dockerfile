# setting golang version
FROM golang:1.18.2-alpine as dev
# ログに出力する時間をJSTにするため、タイムゾーンを設定
ENV TZ /usr/share/zoneinfo/Asia/Tokyo
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
# hot realod config: Airをインストールし、コンテナ起動時に実行する
RUN go install github.com/cosmtrek/air@latest
CMD ["air"]
