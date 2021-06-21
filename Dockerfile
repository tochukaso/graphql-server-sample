FROM golang:1.16.5 AS base

ENV GOPATH=/go
ENV GO111MODULE=on

WORKDIR /go/src/app

ADD . /go/src/app
