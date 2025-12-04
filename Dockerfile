FROM golang:1.25-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o grango-tesorow cmd/rest/main.go

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/grango-tesorow .
COPY configs/ configs/
EXPOSE 8081
CMD ["./grango-tesorow"]