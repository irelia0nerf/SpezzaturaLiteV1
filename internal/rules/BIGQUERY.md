# BigQuery Reference Integration

## Scope

This document describes how Spezzatura Lite MAY persist **derived, non-sensitive**
artifacts into BigQuery as a **reference implementation**.

This is NOT required for correctness and implies no partnership.

---

## What Goes to BigQuery

Only **Proof-of-Check artifacts** (see `internal/artifacts`) are eligible.

Allowed fields:
- schema_version
- correlation_id
- input_hash
- trust_score
- signals
- invariants
- generated_at_utc
- engine
- artifact_hash

Disallowed:
- raw claims
- raw references
- personal or confidential data

---

## Table Schema (Example)

Dataset: `spezzatura_lite`
Table: `proof_of_check`

- schema_version STRING
- correlation_id STRING
- input_hash STRING
- trust_score INT64
- signals ARRAY<STRING>
- invariants JSON
- generated_at_utc TIMESTAMP
- engine JSON
- artifact_hash STRING

Partitioning:
- by DATE(generated_at_utc)

Append-only semantics recommended.

---

## Guarantees

- Zero-Persistence preserved
- Artifacts are immutable
- Hash allows tamper detection
- Storage is optional and replaceable

---

## Non-Goals

- Analytics over raw inputs
- Reconstruction of sensitive data
- Decision storage
