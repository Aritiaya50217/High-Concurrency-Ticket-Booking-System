# High Concurrency Ticket Booking System

A high-concurrency ticket booking platform built with **Golang, Microservices, Event-Driven Architecture, Kafka, gRPC, Clean Architecture, and Domain-Driven Design (DDD)**.

The system is designed to handle concurrent seat reservations safely while maintaining consistency using **Inbox/Outbox Patterns, database locking, optimistic concurrency control, and idempotent event processing**.

---

## Architecture Overview

The system consists of three microservices:

```text
                         ┌─────────────────┐
                         │   User Service  │
                         │  REST + gRPC    │
                         │   :8080/:50051  │
                         └────────┬────────┘
                                  │
                               gRPC
                                  │
┌──────────────┐           ┌──────▼────────┐
│    Client    │── REST ──▶│ Booking       │
└──────────────┘           │ Service :8081 │
                           └───────┬────────┘
                                   │
                              Outbox Event
                                   │
                                   ▼
                              ┌─────────┐
                              │  Kafka  │
                              └────┬────┘
                                   │
                                   ▼
                           ┌───────▼────────┐
                           │ Event Service  │
                           │    :8082       │
                           └────────────────┘
```

### User Service

Responsible for user-related operations and provides synchronous communication through REST and gRPC.

Responsibilities:

* User management
* JWT-based authentication
* Provide user information through gRPC
* PostgreSQL persistence
* gRPC server

Acts as the **source of truth for user information**.

---

### Booking Service

Responsible for booking lifecycle management.

Responsibilities:

* Create booking requests
* Manage booking expiration
* Publish domain events through Outbox Pattern
* Maintain booking state (`PENDING`, `CONFIRMED`, `CANCELLED`)
* Communicate with User Service through gRPC
* Run booking expiration worker

Acts as the **source of truth for booking information**.

---

### Event Service

Responsible for event and seat management.

Responsibilities:

* Consume booking events from Kafka
* Manage event and seat information
* Reserve seats
* Release expired seats
* Prevent duplicate event processing using Inbox Pattern
* Maintain seat consistency under high concurrency

Acts as the **source of truth for seat availability**.

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

#### `interface`

* HTTP handlers
* Routers
* DTOs
* Middleware

#### `application`

* Use cases
* Business workflows

#### `domain`

* Aggregates
* Entities
* Value Objects
* Domain Events
* Repository interfaces

#### `infrastructure`

* PostgreSQL
* Kafka
* gRPC
* Repository implementations
* External services
* Configuration

---

## Domain-Driven Design

The project applies DDD concepts to model the booking domain.

### Aggregates

* Event

### Entities

* Seat
* Booking

### Value Objects

* SeatStatus
* BookingStatus

### Domain Events

* `BookingCreated`
* `BookingCancelled`
* `SeatReserved`
* `SeatReleased`

---

## High Concurrency Strategy

The system uses multiple consistency mechanisms to prevent race conditions and conflicting seat reservations.

### Row-Level Locking

Prevents concurrent transactions from modifying the same seat simultaneously.

```sql
SELECT ... FOR UPDATE;
```

This is used when processing critical seat reservation operations.

---

### Optimistic Concurrency Control

Seat updates use version checking to detect stale updates.

```sql
UPDATE seats
SET ...
WHERE id = ?
  AND version = ?;
```

If the version has changed, the update does not affect the expected row.

---

### Unique Constraints

Database constraints provide an additional consistency boundary and prevent duplicate seat creation.

```text
UNIQUE(event_id, seat_number)
```

---

## Event-Driven Workflow

### Booking Flow

```text
User
  │
  ▼
Booking Service
  │
  ├── Create Booking
  │      Status = PENDING
  │
  └── Create Outbox Event
             │
             ▼
           Kafka
             │
             ▼
       Event Service
             │
             ▼
        Reserve Seat
             │
             ▼
       Seat = RESERVED
```

The booking and seat management responsibilities are separated between services and connected through asynchronous events.

---

### Booking Expiration Flow

```text
Booking Expiration Worker
          │
          ▼
   Find expired bookings
          │
          ▼
Booking Status = CANCELLED
          │
          ▼
Create booking.cancelled event
          │
          ▼
    Outbox Worker
          │
          ▼
        Kafka
          │
          ▼
 Event Service Consumer
          │
          ▼
      Release Seat
          │
          ▼
   Seat = AVAILABLE
```

Expired bookings are automatically cancelled and the corresponding seats are released through an event-driven workflow.

---

## gRPC Communication

The Booking Service uses **gRPC for synchronous service-to-service communication** with the User Service.

```text
Booking Service
      │
      │ gRPC
      ▼
User Service
      │
      ▼
GetUser(...)
```

Protocol definitions are maintained under:

```text
contracts/
```

The project demonstrates:

* Protocol Buffers
* gRPC client/server communication
* Service contracts
* Synchronous inter-service communication
* Generated protobuf code

---

## Inbox Pattern

The Inbox Pattern is used to prevent duplicate event processing.

Kafka provides at-least-once delivery semantics, meaning a consumer may receive the same event more than once.

### Process

```text
Receive Kafka Event
        │
        ▼
Check processed_events
        │
   ┌────┴────┐
   │         │
Exists     Not Found
   │         │
   ▼         ▼
 Skip      Process
             │
             ▼
     Store processed event
```

This provides idempotent event processing.

---

## Outbox Pattern

The Outbox Pattern is used to reliably publish domain events to Kafka.

### Process

```text
Database Transaction
        │
        ├── Save Business Data
        │
        └── Save Outbox Event
                  │
                  ▼
            Outbox Worker
                  │
                  ▼
                Kafka
                  │
                  ▼
          Mark Event as Sent
```

The business data and outbox event are stored within the same database transaction, reducing the risk of database/Kafka inconsistency.

---

## Kafka Topics

| Topic               | Description                       |
| ------------------- | --------------------------------- |
| `booking.created`   | Booking successfully created      |
| `booking.cancelled` | Booking expired or cancelled      |
| `seat.reserved`     | Seat reserved successfully        |
| `seat.released`     | Seat released and available again |

---

## Observability

The project includes a monitoring stack using **Prometheus and Grafana**.

```text
Services
    │
    ▼
Prometheus
    │
    ▼
Grafana
```

Prometheus is used for metrics collection and Grafana is used for visualization and monitoring.

---

## Technology Stack

| Category           | Technology         |
| ------------------ | ------------------ |
| Language           | Golang             |
| API                | REST / HTTP        |
| RPC                | gRPC               |
| Serialization      | Protocol Buffers   |
| Web Framework      | Gin                |
| Database           | PostgreSQL         |
| ORM                | GORM               |
| Messaging          | Apache Kafka       |
| Authentication     | JWT                |
| Architecture       | Microservices      |
| Architecture Style | Clean Architecture |
| Domain Modeling    | DDD                |
| Containerization   | Docker             |
| Orchestration      | Docker Compose     |
| Metrics            | Prometheus         |
| Visualization      | Grafana            |
| CI/CD              | GitHub Actions     |

---

## Project Structure

```text
High-Concurrency-Ticket-Booking-System/
│
├── booking-service/
│   ├── cmd/
│   ├── internal/
│   │   ├── application/
│   │   ├── domain/
│   │   ├── infrastructure/
│   │   └── interface/
│   ├── configs/
│   ├── Dockerfile
│   └── ...
│
├── event-service/
│   ├── cmd/
│   ├── internal/
│   │   ├── application/
│   │   ├── domain/
│   │   ├── infrastructure/
│   │   └── interface/
│   ├── configs/
│   ├── Dockerfile
│   └── ...
│
├── user-service/
│   ├── cmd/
│   ├── internal/
│   │   ├── application/
│   │   ├── domain/
│   │   ├── infrastructure/
│   │   └── interface/
│   ├── configs/
│   ├── Dockerfile
│   └── ...
│
├── contracts/
│   └── proto/
│
├── init-db/
│
├── monitoring/
│   └── grafana/
│
├── .github/
│   └── workflows/
│
├── docker-compose.yml
├── prometheus.yml
├── Makefile
└── README.md
```

---

## Running the Project

### Prerequisites

* Docker
* Docker Compose
* Go 1.26+

### Clone Repository

```bash
git clone https://github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System.git

cd High-Concurrency-Ticket-Booking-System
```

### Start Services

```bash
docker compose up --build
```

---

## Service Endpoints

| Service           | Endpoint                |
| ----------------- | ----------------------- |
| User Service      | `http://localhost:8080` |
| User Service gRPC | `localhost:50051`       |
| Booking Service   | `http://localhost:8081` |
| Event Service     | `http://localhost:8082` |
| PostgreSQL        | `localhost:5433`        |
| Kafka             | `localhost:9092`        |
| Prometheus        | `http://localhost:9090` |
| Grafana           | `http://localhost:3000` |

---

## Key Features

* High-concurrency seat booking
* Microservices architecture
* Event-driven communication
* REST APIs
* gRPC service-to-service communication
* Protocol Buffers
* Inbox Pattern
* Outbox Pattern
* Booking expiration handling
* Background workers
* Optimistic concurrency control
* Row-level locking
* Idempotent event processing
* Eventual consistency
* JWT authentication
* PostgreSQL persistence
* Docker Compose
* Prometheus metrics
* Grafana monitoring
* GitHub Actions CI/CD

---

## Design Goals

The main goal of this project is to explore how a distributed booking system can maintain consistency under concurrent requests while keeping services loosely coupled.

The project focuses on:

1. **Consistency** — prevent conflicting seat reservations.
2. **Reliability** — reliably publish and process domain events.
3. **Concurrency** — handle multiple requests competing for the same resources.
4. **Decoupling** — separate synchronous and asynchronous communication.
5. **Observability** — expose metrics for system monitoring.
6. **Maintainability** — keep business logic independent through Clean Architecture.

---

## Future Improvements

Potential areas for further development:

* Payment Service
* Saga Pattern
* Retry / backoff strategy
* Dead Letter Queue (DLQ)
* Distributed tracing
* Kubernetes deployment
* More comprehensive integration and load testing
* Advanced Kafka reliability configuration
