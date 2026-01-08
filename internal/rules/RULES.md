# Rules Governance

## Purpose

This document defines how **consistency rules** in Spezzatura Lite are authored,
versioned, reviewed, and deprecated.

---

## Principles

- Deterministic
- Reproducible
- Non-advisory
- Fail-closed

Rules NEVER:
- infer intent
- predict outcomes
- assert truth

---

## Versioning

Each rule has:
- stable RuleID
- explicit version suffix (e.g. `_v1`)
- documented rationale

Rule packs are versioned independently (e.g. `v1.0.0`).

Breaking changes REQUIRE:
- new RuleID or version
- explicit governance approval

---

## Authoring Checklist

Before adding a rule:
- [ ] Input schema defined
- [ ] Normalization rules explicit
- [ ] Failure modes enumerated
- [ ] Signals documented
- [ ] Unit tests for match/mismatch/missing

---

## Deprecation

Deprecated rules:
- remain readable
- are not executed in default pack
- keep historical artifacts valid

---

## Auditability

Every rule evaluation yields:
- rule_id
- passed flag
- failure_mode (if any)
- reason
- signals

This enables third-party replay and review.
