FROM golang:latest

RUN mkdir /app

ADD . /app/

WORKDIR /app

RUN  go get -u -v all
RUN  go build -o main .

ENTRYPOINT ["/app/main"]
