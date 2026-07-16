# ---------- build stage ----------
FROM golang:1.26-alpine AS builder

WORKDIR /app

# install git (บาง lib ต้องใช้)
RUN apk add --no-cache git

# copy go mod
COPY go.mod go.sum ./
RUN go mod download

# copy source code
COPY . .

# build binary
RUN CGO_ENABLED=0 go build -o event-service ./cmd/server


# ---------- runtime stage ----------
FROM alpine:latest

WORKDIR /app

# copy binary from builder
COPY --from=builder /app/event-service .

# copy config folder
COPY configs ./configs

EXPOSE 8082

CMD ["./event-service"]