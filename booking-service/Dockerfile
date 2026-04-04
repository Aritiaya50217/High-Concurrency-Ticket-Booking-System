# ---------- build stage ----------
FROM golang:1.25-alpine AS builder

WORKDIR /app

# install git (บาง lib ต้องใช้)
RUN apk add --no-cache git

# copy go mod
COPY go.mod go.sum ./
RUN go mod download

# copy source code
COPY . .

# build binary
RUN go build -o booking-service ./cmd/server


# ---------- runtime stage ----------
FROM alpine:latest

WORKDIR /app

# copy binary from builder
COPY --from=builder /app/booking-service .

# copy config folder
COPY configs ./configs

EXPOSE 8081

CMD ["./booking-service"]