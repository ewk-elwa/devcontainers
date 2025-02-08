FROM golang:1.20 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY src/ ./src/

WORKDIR /app/src

RUN go build -o recipebook main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/src/recipebook .

CMD ["./recipebook"] 

EXPOSE 8080