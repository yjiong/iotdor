FROM golang:1.18-alpine AS development

ENV PROJECT_PATH=/mydevelop/iotdor
ENV PATH=$PATH:$PROJECT_PATH/build
ENV CGO_ENABLED=0
ENV GO_EXTRA_BUILD_ARGS="-a -installsuffix cgo"

RUN apk add --no-cache ca-certificates make git bash protobuf alpine-sdk 

RUN mkdir -p $PROJECT_PATH
COPY . $PROJECT_PATH
WORKDIR $PROJECT_PATH

RUN mkdir -p /etc/iotdor/

RUN make build

FROM alpine:latest AS production
MAINTAINER yjiong@msn.com

RUN echo "http://mirrors.aliyun.com/alpine/latest-stable/main/" > /etc/apk/repositories && \
    echo "http://mirrors.aliyun.com/alpine/latest-stable/community/" >> /etc/apk/repositories

RUN apk update && \
    apk add --no-cache bash tree tzdata && \
    cp -rf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime

ENV LANG=C.UTF-8

WORKDIR /root/
RUN apk --no-cache add ca-certificates
COPY --from=development /mydevelop/iotdor/build/iotdor .
ENTRYPOINT ["./iotdor"]

