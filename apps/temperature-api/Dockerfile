FROM golang:1.21-alpine
WORKDIR /app
COPY . .
RUN go build -o temperature-api main.go
EXPOSE 8081
CMD ["/app/temperature-api"]
