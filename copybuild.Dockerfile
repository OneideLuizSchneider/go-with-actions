FROM golang:1.15.3-alpine

RUN mkdir /app
ADD . /app
WORKDIR /app

COPY main main

CMD ["/app/main"]