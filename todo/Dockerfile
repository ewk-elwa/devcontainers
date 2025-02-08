FROM golang:1.23 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY src/ ./src/

WORKDIR /app/src

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -installsuffix cgo -ldflags '-extldflags "-static"' -o todo main.go
# RUN go build -o todo main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/src/todo .

CMD ["./todo"] 

EXPOSE 8080
