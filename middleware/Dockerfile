FROM golang:1.11.4-alpine3.8
WORKDIR /go/src/github.com/nkjmsss/class_3s_project_enshu/middleware
ENV GO111MODULE=on
COPY go.mod .
COPY go.sum .
RUN apk upgrade && \
    apk update && \
    apk add --no-cache \
      gcc \
      libc-dev \
      git \
    && \
    go mod download
