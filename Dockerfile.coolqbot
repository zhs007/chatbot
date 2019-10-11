FROM golang:1.12 as builder

MAINTAINER zerro "zerrozhao@gmail.com"

ENV GOPROXY https://goproxy.io

WORKDIR /src/chatbot

COPY ./go.* /src/chatbot/

RUN go mod download

COPY . /src/chatbot

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o coolqbot ./coolqnode \
    && mkdir /app \
    && mkdir /app/coolqbot \
    && mkdir /app/coolqbot/cfg \
    && mkdir /app/coolqbot/logs \
    && cp ./coolqbot /app/coolqbot/ \
    && cp ./cfg/coolq.yaml.sample /app/coolqbot/cfg/config.yaml

FROM alpine
RUN apk upgrade && apk add --no-cache ca-certificates
WORKDIR /app/coolqbot
COPY --from=builder /usr/local/go/lib/time/zoneinfo.zip /usr/local/go/lib/time/zoneinfo.zip
COPY --from=builder /app/coolqbot /app/coolqbot
CMD ["./coolqbot"]