FROM golang:1.8-alpine

RUN apk --update upgrade && apk add ca-certificates && apk add git && rm -rf /var/cache/apk/*
WORKDIR /go/src/github.com/rattrap/docker-demo
COPY . .

RUN go-wrapper download
RUN go-wrapper install

CMD ["go-wrapper", "run"]