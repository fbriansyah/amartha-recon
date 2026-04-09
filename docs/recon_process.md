# Reconciliation Process Documentation

This document explains the logic, payload, and usage of the Amartha Reconciliation Service process endpoint.

## Algorithm Overview

The core reconciliation logic is designed to match System Transactions against Bank Statements based on date, amount, and a specific tolerance window.

### 1. Grouping and Sorting (Bucketing)
All internal System Transactions are grouped into "buckets" first by `Date` and then by `Amount`. Within each bucket, the transactions are sorted by `Transaction Time` from oldest to newest. This ensures that a First-In-First-Out (FIFO) approach is used when matching bank statements to system records.

### 2. Lookback Calculation
For every Bank Statement, the system calculates acceptable "lookback dates" against the System Transactions.
- Normally, it checks the Bank Statement Date (T) and the previous day (T-1).
- If the Bank Statement falls on a **Monday**, it expands the lookback to include Sunday, Saturday, and Friday to account for the weekend gap. 
- Lookback dates are checked in order from oldest to newest.

### 3. Exact Matching
The service first tries to find an exact match for the Bank Statement `Amount` within the System Transaction buckets on the calculated lookback dates. If it finds one, it pairs the records and removes the matched System Transaction from the bucket.

### 4. Tolerance Matching
If an exact match is not found on a lookback date, the service scans the remaining unmet System Transactions for that day. It looks for a transaction with an amount difference that falls within an acceptable **Tolerance Limit** (currently configured at `5000`). If found, it marks it as a matched transaction with a discrepancy.

### 5. Exception Handling
After iterating over all Bank Statements:
- **Bank Exceptions**: Any Bank Statements that could not find a matching System Transaction (exact or within tolerance) are flagged as an exception with Source `"BANK"`.
- **System Exceptions**: Any remaining System Transactions that were not claimed by a Bank Statement are flagged as an exception with Source `"SYSTEM"`.

## API Endpoint Usage

The reconciliation process is triggered manually using a REST API endpoint. 

**Endpoint:** `POST /api/v1/reconciliation/process`

### Payload

The endpoint expects a JSON payload specifying the date boundary to process files for.

```json
{
    "start_date": "YYYY-MM-DD",
    "end_date": "YYYY-MM-DD"
}
```

### Curl Example

```bash
curl -X POST http://localhost:8080/api/v1/reconciliation/process \
  -H "Content-Type: application/json" \
  -d '{
    "start_date": "2023-01-01",
    "end_date": "2023-01-31"
  }'
```

### Success Response

The response includes the totals processed, matches made, and a breakdown of the generated exceptions.

```json
{
    "details": {
        "missing_in_bank": [
            // List of SYSTEM exceptions
        ],
        "missing_in_system": {
            "BANK": [
                // List of BANK exceptions
            ]
        }
    },
    "total_discrepancy": "0",
    "total_matched": 4,
    "total_processed": 10,
    "total_unmatched": 6
}
```
