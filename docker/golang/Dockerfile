FROM golang:1.15.1

WORKDIR /go/src/app

ENV GO111MODULE=on

RUN apt-get update -qq && apt-get install -y vim git
RUN GO111MODULE=off go get\
      github.com/pilu/fresh\
      github.com/motemen/gore/cmd/gore\
      github.com/mdempsky/gocode\
      github.com/k0kubun/pp\
      golang.org/x/tools/cmd/goimports\
      golang.org/x/lint/golint\
      github.com/motemen/gobump/cmd/gobump\
      github.com/Songmu/make2help/cmd/make2help

EXPOSE 8080

ENTRYPOINT [ "fresh" ]
