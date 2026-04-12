FROM golang:1.25 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app

FROM alpine:latest

WORKDIR /app

RUN mkdir -p /data

COPY --from=builder /app/app .
COPY --from=builder /app/web ./web

ENV TODO_DBFILE=/app/scheduler.db

EXPOSE 7540

CMD ["./app"]