# setting golang version
FROM golang:1.18.2-alpine as dev
# ログに出力する時間をJSTにするため、タイムゾーンを設定
ENV TZ /usr/share/zoneinfo/Asia/Tokyo
# 言語設定
ENV LANG C.UTF-8
ENV ROOT=/go/src/app
# Goはコンパイル時に、CGOを使ってC言語のライブラリを使うように設定する事ができる。
# しかし、scratchイメージにはこのライブラリは用意されていない。
# よって、今回はC言語のライブラリを使って欲しく無いので、CGO＿ENABLEDを0にしてこの機能をOFFにした。
ENV CGO_ENABLED 0
WORKDIR ${ROOT}
# アップデートとgitのインストール
RUN apk update && apk add git && apk add bash
# appineだとデフォルトでbashを利用できないため、bashを追加している
RUN apk add bash
# go.mod と go.sum をWORKDIRにCOPY
COPY go.mod go.sum ./
# パッケージをインストール
RUN go mod download
# コンテナのポートを設定
EXPOSE 8080
# migrate tool
RUN go install github.com/pressly/goose/v3/cmd/goose@latest
# hot realod config: Airをインストールし、コンテナ起動時に実行する
RUN go install github.com/cosmtrek/air@latest
CMD ["air", "go", "run", "server.go"]
