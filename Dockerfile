FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

# Download and cache dependencies (this step can be cached)
RUN go mod tidy

COPY . .

RUN go build -o main .

# Runner small image
FROM alpine:latest

WORKDIR /root/app

COPY --from=builder /app/main .

EXPOSE 3000

CMD ["./main"]
