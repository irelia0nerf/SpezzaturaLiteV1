package rules

import "testing"

func TestRulePack_DeterministicOrder(t *testing.T) {
	rp := DefaultRulePack()

	in := Input{
		Claims: map[string]string{
			"corp.cnpj":                 "40.822.202/0001-33",
			"cap_table.founders_count":  "3",
		},
		References: map[string][]ReferenceRecord{
			"corp.cnpj": {
				{Value: "40.822.202/0001-33", Source: "registry"},
			},
			"cap_table.founders_count": {
				{Value: "2", Source: "filing"},
				{Value: "3", Source: "filing"},
			},
		},
	}

	f1, err := rp.Evaluate(in)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	// Re-run with same input (maps are randomized in Go, so we assert stable ordering & results).
	f2, err := rp.Evaluate(in)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(f1) != len(f2) {
		t.Fatalf("expected same length, got %d vs %d", len(f1), len(f2))
	}
	for i := range f1 {
		if f1[i].RuleID != f2[i].RuleID || f1[i].Passed != f2[i].Passed || f1[i].Reason != f2[i].Reason {
			t.Fatalf("expected deterministic findings, got %#v vs %#v", f1[i], f2[i])
		}
	}
}

func TestCapTableFoundersCount_Mismatch(t *testing.T) {
	r := CapTableFoundersCountRule{}

	in := Input{
		Claims: map[string]string{
			"cap_table.founders_count": "3",
		},
		References: map[string][]ReferenceRecord{
			"cap_table.founders_count": {
				{Value: "2", Source: "filing"},
			},
		},
	}

	f, err := r.Apply(in)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if f.Passed {
		t.Fatalf("expected mismatch (not passed)")
	}
	if f.Mode == nil || *f.Mode != FailMismatch {
		t.Fatalf("expected FailMismatch, got %#v", f.Mode)
	}
}

func TestCorporateID_MissingClaim(t *testing.T) {
	r := CorporateIDStabilityRule{}

	in := Input{
		Claims:     map[string]string{},
		References: map[string][]ReferenceRecord{"corp.cnpj": {{Value: "x", Source: "registry"}}},
	}

	f, err := r.Apply(in)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if f.Passed {
		t.Fatalf("expected fail due to missing claim")
	}
	if f.Mode == nil || *f.Mode != FailInsufficientData {
		t.Fatalf("expected FailInsufficientData, got %#v", f.Mode)
	}
}
