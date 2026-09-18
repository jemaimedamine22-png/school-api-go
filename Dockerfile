# ==========================================
# Stage 1: Build Stage 
# ==========================================
FROM golang:1.26.6-alpine AS builder


RUN apk add --no-cache git ca-certificates tzdata


WORKDIR /app


COPY go.mod go.sum ./
RUN go mod download


COPY . .

COPY app.env .


RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o main main.go

# ==========================================
# Stage 2: Final Security Stage 
# ==========================================
FROM alpine:3.21


RUN apk add --no-cache ca-certificates tzdata


RUN addgroup -g 10001 appgroup && \
    adduser -D -u 10001 -G appgroup appuser

WORKDIR /app


COPY --from=builder /app/main .


RUN chown -R appuser:appgroup /app


USER appuser


EXPOSE 8080


CMD ["./main"]