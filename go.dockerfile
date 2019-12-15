FROM golang:1.13-alpine

ARG GITHUB_USER_NAME
ENV GITHUB_USER_NAME $GITHUB_USER_NAME
ENV CGO_ENABLED 0

WORKDIR /go/src/github.com/${GITHUB_USER_NAME}

RUN apk --no-cache update \
    && apk add --no-cache git \
    # Hot reload
    && go get github.com/oxequa/realize
