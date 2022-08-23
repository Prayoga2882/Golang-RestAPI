# first builder
FROM golang:1.18.1-alpine as builder

ENV APP /go/src/restapi
RUN apk --no-cache add build-base git mercurial gcc
WORKDIR ${APP}
COPY . ${APP}
RUN cd ${APP} \
    && go mod tidy \
    && CGO_ENABLED=0 go build 
CMD ["./Golang-RestAPI"]

# second runtime
# FROM alpine:3.15

# ENV APP_BINARY /go/src/restapi
# COPY --from=builder ${APP_BINARY}/restapi .
# CMD ["./Golang-RestAPI"]
