package artifacts

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"time"
)

// Version identifies the artifact schema version.
// Increment ONLY with explicit governance.
const Version = "v1.0.0"

// ProofOfCheck is a non-sensitive, immutable verification artifact.
// It is safe to persist and share.
type ProofOfCheck struct {
	SchemaVersion string            `json:"schema_version"`
	CorrelationID string            `json:"correlation_id"`
	InputHash     string            `json:"input_hash"`
	TrustScore    int               `json:"trust_score"`
	Signals       []string          `json:"signals"`
	Invariants    map[string]bool   `json:"invariants"`
	GeneratedAt   time.Time         `json:"generated_at_utc"`
	Engine        EngineDescriptor  `json:"engine"`
	Hash          string            `json:"artifact_hash"`
}

// EngineDescriptor fingerprints the execution engine.
type EngineDescriptor struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

// NewProof builds and seals a ProofOfCheck deterministically.
func NewProof(
	correlationID string,
	inputHash string,
	trustScore int,
	signals []string,
	invariants map[string]bool,
	engine EngineDescriptor,
) (ProofOfCheck, error) {

	if correlationID == "" || inputHash == "" {
		return ProofOfCheck{}, errors.New("correlation_id and input_hash required")
	}

	p := ProofOfCheck{
		SchemaVersion: Version,
		CorrelationID: correlationID,
		InputHash:     inputHash,
		TrustScore:    trustScore,
		Signals:       stableStrings(signals),
		Invariants:    stableMap(invariants),
		GeneratedAt:   time.Now().UTC(),
		Engine:        engine,
	}

	hash, err := HashArtifact(p)
	if err != nil {
		return ProofOfCheck{}, err
	}
	p.Hash = hash
	return p, nil
}

// HashArtifact computes the SHA-256 hash over the canonical JSON form.
func HashArtifact(p ProofOfCheck) (string, error) {
	canon, err := CanonicalJSON(p)
	if err != nil {
		return "", err
	}
	sum := sha256.Sum256([]byte(canon))
	return hex.EncodeToString(sum[:]), nil
}

// CanonicalJSON returns a stable JSON representation excluding the Hash field.
func CanonicalJSON(p ProofOfCheck) (string, error) {
	type canon struct {
		SchemaVersion string           `json:"schema_version"`
		CorrelationID string           `json:"correlation_id"`
		InputHash     string           `json:"input_hash"`
		TrustScore    int              `json:"trust_score"`
		Signals       []string         `json:"signals"`
		Invariants    map[string]bool  `json:"invariants"`
		GeneratedAt   int64            `json:"generated_at_utc"`
		Engine        EngineDescriptor `json:"engine"`
	}
	c := canon{
		SchemaVersion: p.SchemaVersion,
		CorrelationID: p.CorrelationID,
		InputHash:     p.InputHash,
		TrustScore:    p.TrustScore,
		Signals:       stableStrings(p.Signals),
		Invariants:    stableMap(p.Invariants),
		GeneratedAt:   p.GeneratedAt.UTC().Unix(),
		Engine:        p.Engine,
	}
	b, err := json.Marshal(c)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func stableStrings(in []string) []string {
	out := make([]string, len(in))
	copy(out, in)
	sortStrings(out)
	return out
}

func sortStrings(a []string) {
	for i := 0; i < len(a)-1; i++ {
		for j := i + 1; j < len(a); j++ {
			if a[j] < a[i] {
				a[i], a[j] = a[j], a[i]
			}
		}
	}
}

func stableMap(in map[string]bool) map[string]bool {
	out := make(map[string]bool, len(in))
	keys := make([]string, 0, len(in))
	for k := range in {
		keys = append(keys, k)
	}
	sortStrings(keys)
	for _, k := range keys {
		out[k] = in[k]
	}
	return out
}
