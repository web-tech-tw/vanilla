FROM golang:alpine

WORKDIR /app
ADD . /app

ENV GIN_MODE release

RUN cd /app/cmd/vhs && go build
RUN go clean -cache

ENTRYPOINT /app/cmd/vhs/vhs

EXPOSE 8080
