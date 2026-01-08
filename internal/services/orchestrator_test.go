package service

import (
	"testing"
	"time"

	"github.com/foundlab/spezzatura-lite/internal/core"
	"github.com/foundlab/spezzatura-lite/internal/rules"
)

func TestEvaluate_EndToEnd(t *testing.T) {
	asOf := time.Date(2026, 1, 8, 12, 0, 0, 0, time.UTC)

	coreIn := core.Input{
		CorrelationID: "dec_100",
		AsOfUTC:       asOf,
		Claims: []core.Claim{
			{Namespace: "cap_table", Key: "founders_count", Value: "3"},
			{Namespace: "corp", Key: "cnpj", Value: "40.822.202/0001-33"},
		},
		References: []core.Reference{
			{Namespace: "cap_table", Key: "founders_count", Value: "3", Source: "filing"},
			{Namespace: "corp", Key: "cnpj", Value: "40.822.202/0001-33", Source: "registry"},
		},
	}

	ruleIn := rules.Input{
		Claims: map[string]string{
			"cap_table.founders_count": "3",
			"corp.cnpj":                "40.822.202/0001-33",
		},
		References: map[string][]rules.ReferenceRecord{
			"cap_table.founders_count": {{Value: "3", Source: "filing"}},
			"corp.cnpj":                {{Value: "40.822.202/0001-33", Source: "registry"}},
		},
	}

	res, err := Evaluate(coreIn, ruleIn)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if res.Proof.Hash == "" {
		t.Fatalf("expected proof hash")
	}
	if res.CoreResult.TrustScore <= 0 {
		t.Fatalf("expected positive trust score")
	}
}
