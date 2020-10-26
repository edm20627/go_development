FROM golang:1.15.1

WORKDIR /go/src/app

ENV GO111MODULE=on

RUN apt-get update -qq && apt-get install -y vim git
RUN GO111MODULE=off go get github.com/pilu/fresh
RUN make devel-deps

EXPOSE 8080

# ENTRYPOINT [ "fresh" ]
