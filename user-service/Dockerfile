# ---------- build stage ----------
FROM golang:1.26-alpine AS builder

WORKDIR /app

RUN apk add --no-cache git

# copy go.mod
COPY user-service/go.mod user-service/go.sum ./user-service/

# copy contracts module
COPY contracts/go.mod ./contracts/

# download dependencies
WORKDIR /app/user-service
RUN go mod download

# copy source
WORKDIR /app
COPY user-service/ ./user-service/
COPY contracts/ ./contracts/

# build
WORKDIR /app/user-service
RUN go build -o user-service ./cmd/server


# ---------- runtime stage ----------
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/user-service/user-service .
COPY --from=builder /app/user-service/configs ./configs

EXPOSE 8080 50051

CMD ["./user-service"]