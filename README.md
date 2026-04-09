# Amartha Reconciliation Service

## Overview
The **Amartha Reconciliation Service (POC)** is a backend system built in Go (Golang) designed to reconcile financial transactions between internal systems and external bank records. It streamlines the process of matching transaction datasets, identifying discrepancies (exceptions), and resolving them accurately.

## Key Features
- **CSV Data Ingestion:** Upload transaction records (CSV) from both internal systems and partner banks.
- **Automated Reconciliation:** Process and match records between systems based on predefined logical rules.
- **Exception Handling:** Automatically flag mismatched or missing transactions as exceptions.
- **Resolution Flow:** Provides APIs to review, update, and resolve flagged exceptions.
- **Built on Fiber:** High-performance RESTful APIs built with the Fiber framework.

## Tech Stack
- **Language:** Go (Golang)
- **Web Framework:** [Fiber v2](https://gofiber.io/)
- **CLI Framework:** [Cobra](https://cobra.dev/) (for serving HTTP and background commands)

## Project Structure
- `cmd/` - CLI commands to run the server.
- `internal/` - Core business logic, domain services, and repositories.
- `presenter/` - API presentation layer (handlers and routing).
- `storage/` - Directory for uploaded CSV files and persistent assets.

## Getting Started

### Prerequisites
- Go 1.20+ (Check your `go.mod` for the exact minimum version)

### Running the Application

You can easily start the application's HTTP server using the provided `Makefile`:

```bash
make run
```

Alternatively, you can run the Cobra command directly:

```bash
go run main.go serveHttp
```

The server will start processing on `http://localhost:8080`.

## Documentation

### 1. Reconciliation Process

This explains the logic, payload, and usage of the Amartha Reconciliation Service process endpoint.

#### Algorithm Overview
The core reconciliation logic is designed to match System Transactions against Bank Statements based on date, amount, and a specific tolerance window.

- **Grouping and Sorting (Bucketing):** All internal System Transactions are grouped by Date and Amount, then sorted by time (FIFO).
- **Lookback Calculation:** Finds matches within the Bank Statement Date (T) and the previous day (T-1). If the bank date is a Monday, the lookback expands across the weekend to Friday. 
- **Exact vs. Tolerance Matching:** Tries an exact match on amount first. If unavailable, allows a Tolerance Limit (currently `5000` discrepancy).
- **Exception Handling:** Unmatched bank records become `BANK` exceptions; unclaimed system records become `SYSTEM` exceptions.

#### Process API Endpoint
**Endpoint:** `POST /api/v1/reconciliation/process`

```json
// Example Request Payload
{
    "start_date": "YYYY-MM-DD",
    "end_date": "YYYY-MM-DD"
}
```

---

### 2. Upload Reconciliation File

This flow is used to ingest internal system vs bank reconciliation records as CSV files.

#### Sequence
1. Client POSTs file to the `/upload/{type}` endpoint.
2. The endpoint verifies key `file` in the form-data.
3. It recursively saves the transaction records into the `storage/recon-files/` directory securely prefixed with a unix timestamp timestamp to prevent file collision.

#### Upload API Endpoints
*   **System Upload:** `POST /api/v1/reconciliation/upload/system`
*   **Bank Upload:** `POST /api/v1/reconciliation/upload/bank`

```bash
# Example Request
curl -X POST http://localhost:8080/api/v1/reconciliation/upload/system \
  -H "Content-Type: multipart/form-data" \
  -F "file=@/path/to/system_transactions.csv"
```

## Future Improvements (TODOs)

Since this project is currently a Proof of Concept (POC), the following enhancements are recommended before transitioning to a production-ready state:

- [ ] **Database Integration:** Move from file-based/in-memory storage to a robust relational database (e.g., PostgreSQL or MySQL) for persistent and reliable record keeping.
- [ ] **Asynchronous Processing:** Implement a message broker (e.g., Kafka, RabbitMQ, or Redis queue) to handle large CSV uploads and reconciliation matching tasks asynchronously in the background.
- [ ] **Authentication & Authorization:** Secure the API endpoints using robust methods like JWT, OAuth2, or API keys.
- [ ] **Comprehensive Testing:** Add extensive unit and integration test coverage for core business logic, repositories, and API handlers.
- [ ] **Observability:** Integrate structured logging (e.g., Zap or Logrus) and telemetry (e.g., OpenTelemetry, Prometheus, Grafana) for monitoring system health and performance.
- [ ] **Pagination & Filtering:** Add robust pagination, filtering, and sorting capabilities to the GET APIs (especially for listing exceptions and transactions).
- [ ] **CI/CD Pipeline:** Set up automated CI/CD pipelines (e.g., GitHub Actions, GitLab CI) for linting, testing, and building the application.