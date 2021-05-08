FROM golang:1.14-alpine

RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh

WORKDIR /go/src/simple-rest
COPY . .

COPY go.mod go.sum ./

RUN go mod download
RUN go build main.go

EXPOSE 8080
CMD ["./main"]