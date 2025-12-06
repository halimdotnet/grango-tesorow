FROM golang:1.25-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o grango-tesorow cmd/rest/main.go

FROM alpine:latest
RUN apk add --no-cache curl ca-certificates
WORKDIR /app
COPY --from=builder /app/grango-tesorow .
COPY --from=builder /app/configs/ configs/
COPY --from=builder /app/.env .env
EXPOSE 8081
CMD ["./grango-tesorow"]