FROM golang:1.15.3-alpine

RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go build -o main .

CMD ["/app/main"]