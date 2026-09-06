# ---------- build stage ----------
FROM golang:1.26-alpine AS builder

WORKDIR /app

# install git (บาง lib ต้องใช้)
RUN apk add --no-cache git

# copy go mod ของ user-service
COPY user-service/go.mod user-service/go.sum ./user-service/

# copy contracts module
COPY contracts/go.mod ./contracts/

# download dependencies
WORKDIR /app/user-service
RUN go mod download

# copy source code
WORKDIR /app
COPY user-service/ ./user-service/
COPY contracts/ ./contracts/

# build binary
WORKDIR /app/user-service
RUN go build -o user-service ./cmd/server


# ---------- runtime stage ----------
FROM alpine:latest

WORKDIR /app

# copy binary from builder
COPY --from=builder /app/user-service/user-service .

# copy config folder
COPY --from=builder /app/user-service/configs ./configs

EXPOSE 8080 50051

CMD ["./user-service"]