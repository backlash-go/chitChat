FROM golang:1.14 AS build-env

ENV WORKSPACE=/workspace
ENV GOPROXY=https://goproxy.cn

RUN mkdir $WORKSPACE

WORKDIR $WORKSPACE
COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o ./serve

FROM alpine
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories
RUN apk add --no-cache libc6-compat tzdata curl \
&& echo "Asia/Shanghai" > /etc/timezone \
&& ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
RUN mkdir /app

COPY --from=build-env /workspace/* /app/

RUN ls /app/

WORKDIR /app

ENTRYPOINT ./serve -conf config.yaml
