package core

import (
	"testing"
	"time"
)

func TestEvaluate_DeterministicSameInputDifferentOrder(t *testing.T) {
	asOf := time.Date(2026, 1, 8, 12, 0, 0, 0, time.UTC)

	in1 := Input{
		CorrelationID: "dec_001",
		AsOfUTC:       asOf,
		Claims: []Claim{
			{Namespace: "cap_table", Key: "founders", Value: "3"},
			{Namespace: "corp", Key: "cnpj", Value: "40.822.202/0001-33"},
		},
		References: []Reference{
			{Namespace: "corp", Key: "cnpj", Value: "40.822.202/0001-33", Source: "registry"},
			{Namespace: "cap_table", Key: "founders", Value: "2", Source: "filing"},
			{Namespace: "cap_table", Key: "founders", Value: "3", Source: "filing"},
		},
	}

	// Same data, different order
	in2 := Input{
		CorrelationID: "dec_001",
		AsOfUTC:       asOf,
		Claims: []Claim{
			{Namespace: "corp", Key: "cnpj", Value: "40.822.202/0001-33"},
			{Namespace: "cap_table", Key: "founders", Value: "3"},
		},
		References: []Reference{
			{Namespace: "cap_table", Key: "founders", Value: "3", Source: "filing"},
			{Namespace: "cap_table", Key: "founders", Value: "2", Source: "filing"},
			{Namespace: "corp", Key: "cnpj", Value: "40.822.202/0001-33", Source: "registry"},
		},
	}

	r1, err1 := Evaluate(in1)
	r2, err2 := Evaluate(in2)

	if err1 != nil || err2 != nil {
		t.Fatalf("expected no errors, got err1=%v err2=%v", err1, err2)
	}
	if r1.InputHash != r2.InputHash {
		t.Fatalf("expected same InputHash, got %s vs %s", r1.InputHash, r2.InputHash)
	}
	if r1.TrustScore != r2.TrustScore {
		t.Fatalf("expected same TrustScore, got %d vs %d", r1.TrustScore, r2.TrustScore)
	}
	if len(r1.Inconsistencies) != len(r2.Inconsistencies) {
		t.Fatalf("expected same inconsistency count, got %d vs %d", len(r1.Inconsistencies), len(r2.Inconsistencies))
	}
}

func TestEvaluate_FailClosed_MissingCorrelationID(t *testing.T) {
	asOf := time.Date(2026, 1, 8, 12, 0, 0, 0, time.UTC)
	_, err := Evaluate(Input{
		CorrelationID: "",
		AsOfUTC:       asOf,
		Claims:        []Claim{{Namespace: "x", Key: "y", Value: "z"}},
		References:    []Reference{{Namespace: "x", Key: "y", Value: "z", Source: "s"}},
	})
	if err == nil {
		t.Fatalf("expected error for missing correlation_id")
	}
}

func TestEvaluate_FailClosed_ZeroMatchesInvariant(t *testing.T) {
	asOf := time.Date(2026, 1, 8, 12, 0, 0, 0, time.UTC)
	_, err := Evaluate(Input{
		CorrelationID: "dec_002",
		AsOfUTC:       asOf,
		Claims:        []Claim{{Namespace: "corp", Key: "cnpj", Value: "AAA"}},
		References:    []Reference{{Namespace: "corp", Key: "cnpj", Value: "BBB", Source: "registry"}},
	})
	if err == nil {
		t.Fatalf("expected invariant breach error")
	}
}
