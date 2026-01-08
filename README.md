# Spezzatura Lite

**Consistency as a Service (CaaS)**  
Deterministic consistency validation for startup investment data.

---

## Overview

**Spezzatura Lite** is a lightweight infrastructure that provides **consistency validation as a service** for the startup investment ecosystem.

It reduces information asymmetry by converting **subjective claims** (e.g. cap table structure, corporate history, declared traction) into **deterministic consistency signals**, derived from reproducible cross-validation against authoritative data sources.

Spezzatura Lite does **not** make decisions, recommendations, or predictions.  
It produces **signals**, not opinions.

---

## Core Principles

### 1) Trust by Physics
All trust signals are derived from:
- Deterministic logic
- Reproducible processes
- Explicit validation rules

Narratives, intuition, or subjective judgment are explicitly excluded.

### 2) Consistency, Not Truth
Spezzatura Lite does **not** assert factual truth.

It answers a narrower, safer question:

> “Is what is being declared internally consistent with authoritative references and with itself?”

### 3) Non-Advisory by Design
The system **never**:
- Recommends investments
- Rates startups
- Predicts outcomes
- Acts as legal, financial, or compliance advisor

### 4) Zero-Persistence Security Model
Sensitive data is:
- Processed ephemerally (in memory)
- Never stored at rest
- Never logged in raw form

Only **non-sensitive derived artifacts** may be emitted.

---

## Ecosystem Alignment

Spezzatura Lite is architected to complement **startup data platforms** by providing **verifiable integrity signals** in areas where authoritative data is fragmented or unavailable — most notably **cap table consistency and historical coherence**.

This project does not replace data platforms.  
It provides **deterministic consistency signals** that can enrich ranking, intake, and screening workflows.

---

## Cap Table Consistency Signals

One of the primary motivations behind Spezzatura Lite is the structural lack of authoritative, verifiable cap table information in the ecosystem.

Spezzatura Lite does NOT attempt to assert ownership truth.

Instead, it evaluates:
- internal consistency of declared cap tables
- historical coherence across funding events
- alignment between disclosures and authoritative references (when available)

The output is a **consistency signal**, not a statement of factual ownership.

---

## Cloud-Native Reference Implementation

Spezzatura Lite is designed for execution on cloud-native analytics stacks, leveraging scalable data warehouses for **non-sensitive derived artifacts** and audit metadata.

**BigQuery** is used as a reference implementation for:
- deterministic aggregation
- reproducible analytics
- append-only verification records (non-sensitive)

---

## What Spezzatura Lite Produces

- **Deterministic TrustScore**  
  A reproducible score representing consistency level — not quality or performance.

- **Inconsistency Signals**  
  Explicit flags when declarations conflict with references or internal structure.

- **Proof-of-Check Artifacts**  
  Immutable, non-sensitive verification artifacts suitable for third-party reliance.

---

## Architecture (High Level)

```
Input Claims
   │
   ▼
Schema & Invariant Validation
   │
   ▼
Deterministic Consistency Engine
   │
   ▼
TrustScore + Inconsistency Signals
   │
   ▼
Proof-of-Check Artifacts (Non-Sensitive)
```

### Architectural Guarantees
- Deterministic outputs
- Stable serialization
- Fail-closed behavior
- Explicit data lifecycle

---

## Go Stack

- Go 1.22+
- go modules (`go.mod`)
- gofmt mandatory
- golangci-lint (recommended)
- built-in `testing` for tests

---

## Repository Structure (Planned)

```
spezzatura-lite/
├── cmd/                    # entrypoints (api/worker)
├── internal/
│   ├── core/               # pure deterministic logic
│   ├── rules/              # invariants and consistency rules
│   ├── scoring/            # trustscore computation
│   ├── artifacts/          # proof-of-check generation
│   ├── adapters/           # external interfaces (no persistence)
│   └── audit/              # structured audit events (non-sensitive)
└── test/                   # integration tests (optional)
```

---

## Determinism & Testing

- Pure functions preferred
- No hidden global state
- Stable serialization (sorted keys / deterministic slices)
- Boundary conditions and failure modes tested
- Ambiguity fails closed

---

## Security & Privacy

- No raw sensitive data stored
- No sensitive data in logs
- Explicit lifecycle boundaries
- Derived artifacts contain no personal or confidential information

---

## What This Project Is Not

- ❌ An investment recommendation engine  
- ❌ A compliance platform  
- ❌ A fraud accusation system  
- ❌ A source of cap table truth  
- ❌ A shareholder registry  

---

## Status

**Active development**  
Initial public version focused on core consistency validation primitives.

---

## License

To be defined.
