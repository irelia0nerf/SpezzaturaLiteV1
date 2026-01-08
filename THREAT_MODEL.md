# Threat Model (STRIDE-lite)

## Scope
Spezzatura Lite external API and artifact storage.

## Key Threats & Mitigations
- Spoofing: correlation_id required, no auth implied
- Tampering: artifact_hash + append-only storage
- Repudiation: immutable proof-of-check
- Information Disclosure: zero-persistence, no raw data stored
- Denial of Service: stateless API, rate limiting external
- Elevation of Privilege: no state, no roles
