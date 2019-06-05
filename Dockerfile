FROM golang:1.11.4-alpine3.8
WORKDIR /go/src/drone_middleware
RUN apk upgrade && \
    apk update && \
    apk add --no-cache \
      gcc \
      libc-dev \
      git
