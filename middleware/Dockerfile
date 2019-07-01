FROM golang:1.11.4-alpine3.8
WORKDIR /go/src/github.com/nkjmsss/class_3s_project_enshu/middleware
ENV GO111MODULE=on
RUN apk upgrade && \
    apk update && \
    apk add --no-cache \
      gcc \
      libc-dev \
      git \
      mplayer
COPY go.mod .
COPY go.sum .
RUN go mod download && \
    go get -u github.com/oxequa/realize
EXPOSE 1323 3000 8890 11111
