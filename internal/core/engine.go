package core

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"sort"
	"strings"
	"time"
)

// Domain note:
// Spezzatura Lite produces CONSISTENCY SIGNALS, not truth assertions.
// This package intentionally contains pure deterministic logic only.

// FailureMode represents a fail-closed reason.
type FailureMode string

const (
	FailureUnknownInput     FailureMode = "UNKNOWN_INPUT"
	FailureSchemaViolation  FailureMode = "SCHEMA_VIOLATION"
	FailureInvariantBreach  FailureMode = "INVARIANT_BREACH"
	FailureInsufficientData FailureMode = "INSUFFICIENT_DATA"
)

// Claim represents a user-provided declaration (non-sensitive).
// Sensitive/raw payloads MUST be handled outside core and only passed as redacted/normalized values.
type Claim struct {
	// Namespace groups claims into domains (e.g., "cap_table", "corporate_history").
	Namespace string `json:"namespace"`
	// Key identifies the claim field.
	Key string `json:"key"`
	// Value is a normalized representation (string-encoded) to avoid floating drift.
	Value string `json:"value"`
}

// Reference represents an authoritative/sovereign reference signal (non-sensitive).
type Reference struct {
	Namespace string `json:"namespace"`
	Key       string `json:"key"`
	Value     string `json:"value"`
	// Source is a human-readable hint (never a secret), e.g. "registry_x".
	Source string `json:"source"`
}

// Input is the deterministic core input.
// Caller must ensure Zero-Persistence boundaries and redaction.
type Input struct {
	CorrelationID string      `json:"correlation_id"`
	AsOfUTC       time.Time   `json:"as_of_utc"`
	Claims        []Claim     `json:"claims"`
	References    []Reference `json:"references"`
}

// Inconsistency is an explicit signal that something does not match.
type Inconsistency struct {
	Namespace string `json:"namespace"`
	Key       string `json:"key"`
	Claim     string `json:"claim"`
	Ref       string `json:"ref"`
	Source    string `json:"source"`
	Reason    string `json:"reason"`
}

// Result is the deterministic output of evaluation.
type Result struct {
	CorrelationID string          `json:"correlation_id"`
	InputHash     string          `json:"input_hash"` // hash of canonicalized input
	TrustScore    int             `json:"trust_score"` // 0..100 consistency signal
	Inconsistencies []Inconsistency `json:"inconsistencies"`
	FailureMode   *FailureMode    `json:"failure_mode,omitempty"`
}

// Evaluate performs deterministic consistency evaluation.
// It fails closed when inputs are missing or malformed.
func Evaluate(in Input) (Result, error) {
	// Basic schema checks (fail closed)
	if strings.TrimSpace(in.CorrelationID) == "" {
		fm := FailureSchemaViolation
		return Result{FailureMode: &fm}, errors.New("correlation_id required")
	}
	if in.AsOfUTC.IsZero() {
		fm := FailureSchemaViolation
		return Result{CorrelationID: in.CorrelationID, FailureMode: &fm}, errors.New("as_of_utc required")
	}
	if len(in.Claims) == 0 {
		fm := FailureInsufficientData
		return Result{CorrelationID: in.CorrelationID, FailureMode: &fm}, errors.New("claims required")
	}
	if len(in.References) == 0 {
		fm := FailureInsufficientData
		return Result{CorrelationID: in.CorrelationID, FailureMode: &fm}, errors.New("references required")
	}

	// Canonicalize input and compute hash for audit artifact
	canon, err := CanonicalizeInput(in)
	if err != nil {
		fm := FailureUnknownInput
		return Result{CorrelationID: in.CorrelationID, FailureMode: &fm}, err
	}
	h := sha256.Sum256([]byte(canon))
	inputHash := hex.EncodeToString(h[:])

	// Build reference lookup (namespace|key -> list of refs)
	refMap := make(map[string][]Reference, len(in.References))
	for _, r := range in.References {
		if strings.TrimSpace(r.Namespace) == "" || strings.TrimSpace(r.Key) == "" {
			fm := FailureSchemaViolation
			return Result{CorrelationID: in.CorrelationID, InputHash: inputHash, FailureMode: &fm}, errors.New("reference namespace/key required")
		}
		k := r.Namespace + "|" + r.Key
		refMap[k] = append(refMap[k], r)
	}

	// Evaluate each claim against references deterministically
	var inconsistencies []Inconsistency
	matched := 0
	total := 0

	// Sort claims to ensure stable processing order regardless of input order
	claims := make([]Claim, len(in.Claims))
	copy(claims, in.Claims)
	sort.Slice(claims, func(i, j int) bool {
		if claims[i].Namespace != claims[j].Namespace {
			return claims[i].Namespace < claims[j].Namespace
		}
		if claims[i].Key != claims[j].Key {
			return claims[i].Key < claims[j].Key
		}
		return claims[i].Value < claims[j].Value
	})

	for _, c := range claims {
		if strings.TrimSpace(c.Namespace) == "" || strings.TrimSpace(c.Key) == "" {
			fm := FailureSchemaViolation
			return Result{CorrelationID: in.CorrelationID, InputHash: inputHash, FailureMode: &fm}, errors.New("claim namespace/key required")
		}
		total++

		k := c.Namespace + "|" + c.Key
		refs := refMap[k]
		if len(refs) == 0 {
			// no reference for this claim: inconsistency signal (missing ref)
			inconsistencies = append(inconsistencies, Inconsistency{
				Namespace: c.Namespace,
				Key:       c.Key,
				Claim:     c.Value,
				Ref:       "",
				Source:    "",
				Reason:    "NO_REFERENCE_AVAILABLE",
			})
			continue
		}

		// Deterministic match rule: any exact normalized value match counts as consistent.
		// More advanced rules belong in internal/rules package.
		found := false
		var bestRef Reference
		for _, r := range refs {
			if r.Value == c.Value {
				found = true
				bestRef = r
				break
			}
		}
		if found {
			matched++
		} else {
			// pick deterministic reference to report (sorted by Source then Value)
			sort.Slice(refs, func(i, j int) bool {
				if refs[i].Source != refs[j].Source {
					return refs[i].Source < refs[j].Source
				}
				return refs[i].Value < refs[j].Value
			})
			bestRef = refs[0]
			inconsistencies = append(inconsistencies, Inconsistency{
				Namespace: c.Namespace,
				Key:       c.Key,
				Claim:     c.Value,
				Ref:       bestRef.Value,
				Source:    bestRef.Source,
				Reason:    "VALUE_MISMATCH",
			})
		}
	}

	// Fail-closed invariant: at least one match must exist, otherwise reject as insufficient integrity.
	// This prevents "all-unknown" inputs from getting a deceptive score.
	if matched == 0 {
		fm := FailureInvariantBreach
		res := Result{
			CorrelationID:  in.CorrelationID,
			InputHash:      inputHash,
			TrustScore:     0,
			Inconsistencies: sortInconsistencies(inconsistencies),
			FailureMode:    &fm,
		}
		return res, errors.New("invariant breach: zero matches")
	}

	// Simple deterministic trustscore: matched/total scaled to 0..100.
	// Do not change this formula without explicit versioning.
	score := int((float64(matched) / float64(total)) * 100.0)

	res := Result{
		CorrelationID:  in.CorrelationID,
		InputHash:      inputHash,
		TrustScore:     clamp(score, 0, 100),
		Inconsistencies: sortInconsistencies(inconsistencies),
	}
	return res, nil
}

func clamp(v, lo, hi int) int {
	if v < lo {
		return lo
	}
	if v > hi {
		return hi
	}
	return v
}

func sortInconsistencies(in []Inconsistency) []Inconsistency {
	out := make([]Inconsistency, len(in))
	copy(out, in)
	sort.Slice(out, func(i, j int) bool {
		if out[i].Namespace != out[j].Namespace {
			return out[i].Namespace < out[j].Namespace
		}
		if out[i].Key != out[j].Key {
			return out[i].Key < out[j].Key
		}
		if out[i].Reason != out[j].Reason {
			return out[i].Reason < out[j].Reason
		}
		if out[i].Source != out[j].Source {
			return out[i].Source < out[j].Source
		}
		return out[i].Claim < out[j].Claim
	})
	return out
}

// CanonicalizeInput returns a stable JSON string for hashing/audit.
// It sorts Claims and References deterministically and normalizes whitespace.
func CanonicalizeInput(in Input) (string, error) {
	type canonRef struct {
		Namespace string `json:"namespace"`
		Key       string `json:"key"`
		Value     string `json:"value"`
		Source    string `json:"source"`
	}
	type canonIn struct {
		CorrelationID string     `json:"correlation_id"`
		AsOfUnix      int64      `json:"as_of_unix"`
		Claims        []Claim    `json:"claims"`
		References    []canonRef `json:"references"`
	}

	claims := make([]Claim, len(in.Claims))
	copy(claims, in.Claims)
	sort.Slice(claims, func(i, j int) bool {
		if claims[i].Namespace != claims[j].Namespace {
			return claims[i].Namespace < claims[j].Namespace
		}
		if claims[i].Key != claims[j].Key {
			return claims[i].Key < claims[j].Key
		}
		return claims[i].Value < claims[j].Value
	})

	refs := make([]canonRef, 0, len(in.References))
	for _, r := range in.References {
		refs = append(refs, canonRef{Namespace: r.Namespace, Key: r.Key, Value: r.Value, Source: r.Source})
	}
	sort.Slice(refs, func(i, j int) bool {
		if refs[i].Namespace != refs[j].Namespace {
			return refs[i].Namespace < refs[j].Namespace
		}
		if refs[i].Key != refs[j].Key {
			return refs[i].Key < refs[j].Key
		}
		if refs[i].Source != refs[j].Source {
			return refs[i].Source < refs[j].Source
		}
		return refs[i].Value < refs[j].Value
	})

	c := canonIn{
		CorrelationID: strings.TrimSpace(in.CorrelationID),
		AsOfUnix:      in.AsOfUTC.UTC().Unix(),
		Claims:        claims,
		References:    refs,
	}
	b, err := json.Marshal(c)
	if err != nil {
		return "", err
	}
	// Normalize to compact JSON (Marshal already does) and avoid trailing spaces.
	return string(b), nil
}
