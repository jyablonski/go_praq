FROM golang:1.23 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

RUN go build -o myprogram .

FROM alpine:latest

COPY --from=builder /app/myprogram /app/myprogram

ENTRYPOINT ["/app/myprogram"]
EXPOSE 8080