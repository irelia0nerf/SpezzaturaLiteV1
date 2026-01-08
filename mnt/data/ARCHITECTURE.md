# Spezzatura Lite — Architecture

## Purpose

This document describes the **reference architecture** of Spezzatura Lite as a
**Consistency-as-a-Service (CaaS)** system.

It is intended for:
- engineers integrating the system
- platform reviewers
- security and architecture assessments

This document is **descriptive**, not a partnership or deployment claim.

---

## Architectural Objective

Spezzatura Lite is designed to:

- Reduce information asymmetry in startup investment data
- Convert subjective declarations into **deterministic consistency signals**
- Operate with **Zero-Persistence** for sensitive inputs
- Emit **verifiable, non-sensitive proof-of-check artifacts**

The system is explicitly **non-decisional** and **non-advisory**.

---

## High-Level Flow

```
Client / Platform
   │
   ▼
Input Normalization & Redaction
   │
   ▼
Deterministic Core (Pure Functions)
   │
   ▼
Consistency Signals + TrustScore
   │
   ▼
Proof-of-Check Artifacts
   │
   ▼
Cloud Analytics / Audit Storage (Derived Only)
```

---

## Core Components

### 1. Input Boundary (Zero-Persistence Gate)

Responsibilities:
- Receive claims and references
- Normalize values (string-encoded, canonical formats)
- Redact or hash sensitive fields
- Enforce explicit data lifecycle boundaries

Guarantees:
- No raw sensitive payloads cross into persistence layers
- No raw payloads logged or cached

---

### 2. Deterministic Core (`internal/core`)

Responsibilities:
- Schema validation (fail-closed)
- Invariant enforcement
- Deterministic consistency evaluation
- Canonical input hashing

Characteristics:
- Pure functions
- Order-independent execution
- Stable serialization
- Explicit failure modes

This layer contains **no I/O**, **no storage**, and **no external dependencies**.

---

### 3. Rules Engine (`internal/rules`)

Responsibilities:
- Domain-specific consistency rules
- Explicit versioning of rule sets
- Deterministic evaluation only

Examples:
- Cap table internal coherence
- Funding round temporal consistency
- Corporate identifier stability

Rules MUST:
- Be reproducible
- Produce explicit reasons on failure
- Never infer or predict outcomes

---

### 4. Scoring (`internal/scoring`)

Responsibilities:
- Convert rule outcomes into a bounded **TrustScore**
- Preserve determinism across versions
- Fail closed when signal quality is insufficient

Constraints:
- Score represents **consistency level**, not quality or performance
- Formula changes require explicit versioning

---

### 5. Proof-of-Check Artifacts (`internal/artifacts`)

Responsibilities:
- Emit non-sensitive verification artifacts
- Provide third-party verifiability without exposing inputs

Artifact contents:
- correlation_id / decision_id
- canonical input hash
- rules evaluated
- score and signals
- timestamps and version identifiers

Artifacts are:
- Immutable once emitted
- Safe to persist
- Safe to share

---

## Cloud-Native Reference Layer

Spezzatura Lite is designed to integrate with cloud-native analytics platforms
for **derived, non-sensitive artifacts only**.

### BigQuery (Reference Implementation)

Used for:
- Append-only storage of proof-of-check artifacts
- Reproducible analytical queries
- Deterministic aggregation

Constraints:
- No raw claims or references stored
- No personal or confidential data
- Append-only semantics preferred

This layer is **replaceable** and **not required** for core correctness.

---

## Zero-Persistence Model (Explicit)

| Data Type | Stored? | Notes |
|---------|--------|------|
| Raw claims | ❌ | Ephemeral only |
| Raw references | ❌ | Ephemeral only |
| Normalized values | ❌ | Memory-bound |
| Canonical hashes | ✅ | Non-sensitive |
| Proof-of-check artifacts | ✅ | Derived only |

---

## Determinism Guarantees

- Stable input canonicalization
- Sorted collections before evaluation
- Explicit float handling (or integer basis points)
- No randomness without seeding
- Idempotent evaluations

---

## Failure Philosophy

Spezzatura Lite fails **closed**.

Examples:
- Missing references → inconsistency signal
- Zero matching evidence → invariant breach
- Schema ambiguity → hard failure

Silence is treated as risk.

---

## Non-Goals

- Truth registries
- Ownership assertion
- Investment recommendations
- Predictive analytics
- Compliance opinions

---

## Extensibility

Designed extension points:
- Rules engine (versioned)
- Artifact format (versioned)
- Storage adapters (cloud-agnostic)
- API layer (optional)

All extensions must preserve:
- determinism
- zero-persistence
- non-advisory semantics

---

## Status

Reference architecture — active development.
