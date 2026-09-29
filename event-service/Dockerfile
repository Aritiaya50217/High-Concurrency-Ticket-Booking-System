# ---------- build stage ----------
FROM golang:1.26-alpine AS builder

WORKDIR /app

RUN apk add --no-cache git

# copy go.mod
COPY event-service/go.mod event-service/go.sum ./event-service/

# copy contracts module
COPY contracts/go.mod ./contracts/

# download dependencies
WORKDIR /app/event-service
RUN go mod download

# copy source
WORKDIR /app
COPY event-service/ ./event-service/
COPY contracts/ ./contracts/

# build
WORKDIR /app/event-service
RUN go build -o event-service ./cmd/server


# ---------- runtime stage ----------
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/event-service/event-service .
COPY --from=builder /app/event-service/configs ./configs

EXPOSE 8082 50051

CMD ["./event-service"]