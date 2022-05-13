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
RUN apk update && apk add git
# go.mod と go.sum をWORKDIRにCOPY
COPY go.mod go.sum ./
# パッケージをインストール
RUN go mod download
# コンテナのポートを設定
EXPOSE 8080

# hot realod config: Airをインストールし、コンテナ起動時に実行する
RUN go install github.com/cosmtrek/air@latest
CMD ["air", "go", "run", "main.go"]


# 使用予定がないので一旦、コメントアウト
# FROM golang:1.18.2-alpine as builder

# ENV ROOT=/go/src/app
# WORKDIR ${ROOT}

# RUN apk update && apk add git
# COPY go.mod go.sum ./
# RUN go mod download

# COPY . ${ROOT}
# RUN CGO_ENABLED=0 GOOS=linux go build -o $ROOT/binary


# FROM scratch as prod

# ENV ROOT=/go/src/app
# WORKDIR ${ROOT}
# COPY --from=builder ${ROOT}/binary ${ROOT}

# EXPOSE 8080
# CMD ["/go/src/app/binary"]
