FROM golang:alpine

RUN apk update && apk add --no-cache bash
RUN apk add --no-cache git

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

RUN go install github.com/air-verse/air@latest

COPY . .

EXPOSE 9090

CMD ["air", "-c", ".air.toml"]