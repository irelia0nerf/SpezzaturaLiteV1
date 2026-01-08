package rules

import (
	"errors"
	"strings"
)

// CapTableFoundersCountRule checks that declared founders count matches at least one authoritative reference.
type CapTableFoundersCountRule struct{}

func (CapTableFoundersCountRule) ID() string { return "CAP_TABLE_FOUNDERS_COUNT_v1" }

func (CapTableFoundersCountRule) Apply(in Input) (Finding, error) {
	const key = "cap_table.founders_count"

	val, ok := in.Claims[key]
	if !ok || strings.TrimSpace(val) == "" {
		fm := FailInsufficientData
		return Finding{
			RuleID:  "CAP_TABLE_FOUNDERS_COUNT_v1",
			Passed:  false,
			Mode:    &fm,
			Reason:  "MISSING_CLAIM",
			Signals: []string{"CAP_TABLE_MISSING"},
		}, nil
	}

	refs, ok := in.References[key]
	if !ok || len(refs) == 0 {
		fm := FailInsufficientData
		return Finding{
			RuleID:  "CAP_TABLE_FOUNDERS_COUNT_v1",
			Passed:  false,
			Mode:    &fm,
			Reason:  "NO_REFERENCE_AVAILABLE",
			Signals: []string{"CAP_TABLE_NO_REFERENCE"},
		}, nil
	}

	// Deterministic rule: any exact normalized match passes.
	for _, r := range refs {
		if r.Value == val {
			return Finding{
				RuleID:  "CAP_TABLE_FOUNDERS_COUNT_v1",
				Passed:  true,
				Reason:  "MATCH",
				Signals: []string{"CAP_TABLE_CONSISTENT"},
			}, nil
		}
	}

	// Mismatch: deterministic failure, but not an error (still a valid signal).
	fm := FailMismatch
	return Finding{
		RuleID:  "CAP_TABLE_FOUNDERS_COUNT_v1",
		Passed:  false,
		Mode:    &fm,
		Reason:  "VALUE_MISMATCH",
		Signals: []string{"CAP_TABLE_INCONSISTENT"},
	}, nil
}

// compile-time check (avoid unused errors import in some toolchains)
var _ = errors.New
