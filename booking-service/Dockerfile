# ---------- build stage ----------
FROM golang:1.26-alpine AS builder

WORKDIR /app

RUN apk add --no-cache git

# booking-service dependencies
COPY booking-service/go.mod booking-service/go.sum ./booking-service/
COPY contracts/go.mod ./contracts/go.mod

RUN cd booking-service && go mod download

# copy source
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