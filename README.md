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