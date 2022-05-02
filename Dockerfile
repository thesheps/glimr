FROM golang:alpine

COPY ./app /app
WORKDIR /app

RUN apk add make
RUN make compile