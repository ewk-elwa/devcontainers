FROM golang:1.23 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY src/ ./src/

WORKDIR /app/src

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -installsuffix cgo -ldflags '-extldflags "-static"' -o recipebook main.go
# RUN go build -o recipebook main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/src/recipebook .

CMD ["./recipebook"] 

EXPOSE 8080