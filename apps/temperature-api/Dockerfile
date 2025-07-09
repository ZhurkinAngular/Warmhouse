FROM golang:1.22-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o temperature-api .

FROM alpine:latest

RUN apk --no-cache add ca-certificates wget
WORKDIR /root/

COPY --from=builder /app/temperature-api .

EXPOSE 8081

CMD ["./temperature-api"]