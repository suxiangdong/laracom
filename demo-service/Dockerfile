FROM golang:1.14-alpine as builder

ENV GO111MODULE on
ENV GOPROXY https://goproxy.cn,direct

# 更新安装源
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories

RUN apk update && apk --no-cache add git

WORKDIR /app/laracom-demo-service

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o laracom-demo-service

FROM alpine:latest

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories

RUN apk update && apk --no-cache add bash supervisor ca-certificates

RUN mkdir /app

WORKDIR /app

COPY --from=builder /app/laracom-demo-service/laracom-demo-service .

#ADD supervisord.conf /etc/supervisord.conf

CMD ["./laracom-demo-service"]