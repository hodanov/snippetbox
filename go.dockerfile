FROM golang:1.13-alpine

ARG GITHUB_USER_NAME
ENV GITHUB_USER_NAME $GITHUB_USER_NAME

WORKDIR /go/src/github.com/${GITHUB_USER_NAME}

RUN apk --no-cache update \
    && apk add --no-cache git \
    # Hot reload
    && go get github.com/oxequa/realize \
    # MySQL driver
    && go get github.com/go-sql-driver/mysql
