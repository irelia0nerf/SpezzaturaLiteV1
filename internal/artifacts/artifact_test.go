package artifacts

import "testing"

func TestProofDeterministicHash(t *testing.T) {
	engine := EngineDescriptor{Name: "spezzatura-core", Version: "v1"}

	p1, err := NewProof(
		"dec_001",
		"hash_abc",
		80,
		[]string{"VALUE_MISMATCH", "NO_REFERENCE"},
		map[string]bool{"non_empty_match": true},
		engine,
	)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	p2, err := NewProof(
		"dec_001",
		"hash_abc",
		80,
		[]string{"NO_REFERENCE", "VALUE_MISMATCH"},
		map[string]bool{"non_empty_match": true},
		engine,
	)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if p1.Hash != p2.Hash {
		t.Fatalf("expected deterministic hash, got %s vs %s", p1.Hash, p2.Hash)
	}
}
