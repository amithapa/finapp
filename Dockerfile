FROM golang:1.15.3 as builder
ENV DATA_DIRECTORY /go/src/github.com/amithapa/finapp
WORKDIR $DATA_DIRECTORY
ARG APP_VERSION
ARG CGO_ENABLED=0

COPY . .
RUN go build -ldflags="-X github.com/amithapa/finapp/internal/config.Version=$APP_VERSION" github.com/amithapa/finapp/cmd/server

FROM alpine:3.12
ENV DATA_DIRECTORY /go/src/github.com/amithapa/finapp
RUN apk add --update --no-cache \
    ca-certificates
COPY --from=builder ${DATA_DIRECTORY}/server /finapp

ENTRYPOINT ["/finapp"]