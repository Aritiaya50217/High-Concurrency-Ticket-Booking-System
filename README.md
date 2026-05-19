# High Concurrency Ticket Booking System (Event-Driven Microservices)

ระบบจองตั๋วแบบ High Concurrency โดยใช้ **Microservices + Kafka + Clean Architecture + DDD**  
ออกแบบเพื่อรองรับ concurrent booking สูง พร้อม consistency แบบ event-driven + idempotent consumer

---

# Architecture Overview

ระบบแบ่งเป็น 2 services หลัก

## 1. Booking Service (Source of Truth)
- รับ request จาก user
- สร้าง booking
- publish event ไป Kafka

## 2. Event Service (Seat Management)
- consume event จาก Kafka (`booking.created`)
- reserve seat
- update database
- ป้องกัน duplicate event (Inbox Pattern)

---

# Architecture Pattern

## Clean Architecture

---

### Layers:
- interface: handler / router
- usecase: business logic
- domain: entity + aggregate + rules
- infrastructure: DB / Kafka / external system

---

## Domain Driven Design (DDD)

### Core Domain:
- Aggregate: Event
- Entity: Seat
- Value Object: SeatStatus
- Domain Event: BookingCreated, SeatReserved

---

# Key Features

## High Concurrency Safe
- PostgreSQL row lock (`SELECT ... FOR UPDATE`)
- Prevent race condition on seat booking

## Event Driven Architecture
- Kafka-based async communication
- Decoupled services

# Event Flow

User Request
   ↓
Booking Service
   ↓
Kafka (booking.created)
   ↓
Event Service Consumer
   ↓
HandleBookingCreated()
   ↓
Reserve Seat (DB Lock)
   ↓
Update Database
   ↓
Inbox (processed_events)

# Idempotency (Inbox Pattern)

ระบบใช้ตาราง processed_events เพื่อป้องกัน Kafka duplicate event

### Process Logic:
1. Check event_id in processed_events
2. If exists → skip
3. If not exists → process event
4. Insert into processed_events

# Kafka Topics

| Topic | Description |
|------|-------------|
| booking.created | Event จาก booking-service |
| seat.reserved | Event หลังจองสำเร็จ |

# Tech Stack

- Go (Golang)
- Gin
- Kafka (segmentio/kafka-go)
- PostgreSQL
- GORM
- Docker
- Clean Architecture
- DDD

# Project Structure

internal/
 ├── application
 ├── domain
 │    ├── aggregate
 │    ├── entity
 │    ├── event
 │    ├── repository
 │    └── valueobject
 ├── infrastructure
 │    ├── kafka
 │    ├── database
 │    ├── repository
 ├── interface
 └── worker

# How It Works

1. Booking service creates booking
2. Event published to Kafka
3. Event service consumes message
4. Event is validated & checked (Inbox Pattern)
5. Seat is locked using database row-level lock
6. Seat status updated
7. Event marked as processed