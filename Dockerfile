FROM golang:latest

RUN mkdir /app

ADD . /app/

WORKDIR /app

ENTRYPOINT go run QueueBot
