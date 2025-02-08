FROM golang:1.23 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY src/ ./

RUN go build -o todo src/main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/todo .

CMD ["./todo"] 

EXPOSE 8080