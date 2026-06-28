# High Concurrency Ticket Booking System

A high-concurrency ticket booking platform built with **Microservices, Event-Driven Architecture, Kafka, Clean Architecture, and Domain-Driven Design (DDD)**.

The system is designed to handle concurrent seat reservations safely while maintaining consistency using **Inbox/Outbox Patterns**, **database locking**, and **idempotent event processing**.

---

## Architecture Overview

The system consists of two main microservices:

### Booking Service

Responsible for booking lifecycle management.

Responsibilities:

* Create booking requests
* Manage booking expiration
* Publish domain events through Outbox Pattern
* Maintain booking state (`PENDING`, `CONFIRMED`, `CANCELLED`)

Acts as the **source of truth** for booking information.

---

### Event Service

Responsible for seat management.

Responsibilities:

* Consume booking events from Kafka
* Reserve seats
* Release expired seats
* Prevent duplicate event processing using Inbox Pattern
* Maintain seat consistency under high concurrency

Acts as the **source of truth** for seat availability.

---

## Architecture Patterns

### Clean Architecture

The application is organized into four layers:

```text
interface
↓
application
↓
domain
↓
infrastructure
```

### Layers

* **interface**

  * HTTP handlers
  * Routers
  * DTOs
  * Middleware

* **application**

  * Use cases
  * Business workflows

* **domain**

  * Aggregates
  * Entities
  * Value Objects
  * Domain Events
  * Repository interfaces

* **infrastructure**

  * PostgreSQL
  * Kafka
  * Repository implementations
  * External services

---

## Domain Driven Design

### Aggregate

* Event

### Entities

* Seat
* Booking

### Value Objects

* SeatStatus
* BookingStatus

### Domain Events

* BookingCreated
* BookingCancelled
* SeatReserved
* SeatReleased

---

## High Concurrency Strategy

### Row-Level Locking

Prevents race conditions when multiple users attempt to reserve the same seat.

```sql
SELECT ... FOR UPDATE
```

---

### Optimistic Locking

Seat updates use version checking to prevent lost updates.

```text
WHERE id = ? AND version = ?
```

---

### Unique Constraints

Prevents duplicate seat creation.

```text
UNIQUE(event_id, seat_number)
```

---

## Event-Driven Workflow

### Booking Flow

```text
User
↓
Booking Service
↓
Create Booking (PENDING)
↓
Outbox Event Created
↓
Kafka (booking.created)
↓
Event Service Consumer
↓
Reserve Seat
↓
Seat Status = RESERVED
```

---

### Booking Expiration Flow

```text
Booking Expiration Worker
↓
Find expired bookings
↓
Booking Status = CANCELLED
↓
Create booking.cancelled event
↓
Outbox Worker
↓
Kafka
↓
Event Service Consumer
↓
Release Seat
↓
Seat Status = AVAILABLE
```

---

## Inbox Pattern

Used to prevent duplicate event processing caused by Kafka's at-least-once delivery guarantee.

### Process

1. Receive Kafka event
2. Check processed_events table
3. If event already exists → skip
4. Otherwise process event
5. Store processed event record

---

## Outbox Pattern

Ensures reliable event delivery between database transactions and Kafka.

### Process

1. Store event in outbox table within the same transaction
2. Outbox worker polls pending events
3. Publish events to Kafka
4. Mark event as sent

---

## Kafka Topics

| Topic             | Description                       |
| ----------------- | --------------------------------- |
| booking.created   | Booking successfully created      |
| booking.cancelled | Booking expired or cancelled      |
| seat.reserved     | Seat reserved successfully        |
| seat.released     | Seat released and available again |

---

## Technology Stack

* Golang
* Gin
* PostgreSQL
* GORM
* Apache Kafka
* Docker
* Clean Architecture
* Domain Driven Design (DDD)

---

## Project Structure

```text
internal/
├── application
│   └── usecase
├── domain
│   ├── aggregate
│   ├── entity
│   ├── event
│   ├── repository
│   └── valueobject
├── infrastructure
│   ├── config
│   ├── database
│   ├── kafka
│   ├── repository
│   └── security
├── interface
│   ├── dto
│   ├── handler
│   ├── middleware
│   └── router
└── worker
```

---

## Key Features

* High concurrency seat booking
* Event-driven microservices
* Inbox Pattern
* Outbox Pattern
* Booking expiration handling
* Optimistic locking
* Row-level locking
* Idempotent event processing
* Decoupled services
* Eventual consistency

---

## Future Improvements

* Payment Service integration
* Saga Pattern orchestration
* Dead Letter Queue (DLQ)
* Retry strategy
* Distributed tracing
* Prometheus and Grafana monitoring
* Kubernetes deployment
