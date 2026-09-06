# ---------- build stage ----------
FROM golang:1.26-alpine AS builder

WORKDIR /app

RUN apk add --no-cache git

# copy go.mod
COPY booking-service/go.mod booking-service/go.sum ./booking-service/

# copy contracts module
COPY contracts/go.mod ./contracts/

# download dependencies
WORKDIR /app/booking-service
RUN go mod download

# copy source
WORKDIR /app
COPY booking-service/ ./booking-service/
COPY contracts/ ./contracts/

# build
WORKDIR /app/booking-service
RUN go build -o booking-service ./cmd/server


# ---------- runtime stage ----------
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/booking-service/booking-service .
COPY --from=builder /app/booking-service/configs ./configs

EXPOSE 8081

CMD ["./booking-service"]