package rules

import "strings"

// CorporateIDStabilityRule checks that corporate ID (e.g., CNPJ) matches references.
type CorporateIDStabilityRule struct{}

func (CorporateIDStabilityRule) ID() string { return "CORP_ID_STABILITY_v1" }

func (CorporateIDStabilityRule) Apply(in Input) (Finding, error) {
	const key = "corp.cnpj"

	val, ok := in.Claims[key]
	if !ok || strings.TrimSpace(val) == "" {
		fm := FailInsufficientData
		return Finding{
			RuleID:  "CORP_ID_STABILITY_v1",
			Passed:  false,
			Mode:    &fm,
			Reason:  "MISSING_CLAIM",
			Signals: []string{"CORP_ID_MISSING"},
		}, nil
	}

	refs, ok := in.References[key]
	if !ok || len(refs) == 0 {
		fm := FailInsufficientData
		return Finding{
			RuleID:  "CORP_ID_STABILITY_v1",
			Passed:  false,
			Mode:    &fm,
			Reason:  "NO_REFERENCE_AVAILABLE",
			Signals: []string{"CORP_ID_NO_REFERENCE"},
		}, nil
	}

	for _, r := range refs {
		if r.Value == val {
			return Finding{
				RuleID:  "CORP_ID_STABILITY_v1",
				Passed:  true,
				Reason:  "MATCH",
				Signals: []string{"CORP_ID_CONSISTENT"},
			}, nil
		}
	}

	fm := FailMismatch
	return Finding{
		RuleID:  "CORP_ID_STABILITY_v1",
		Passed:  false,
		Mode:    &fm,
		Reason:  "VALUE_MISMATCH",
		Signals: []string{"CORP_ID_INCONSISTENT"},
	}, nil
}
