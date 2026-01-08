package rules

import (
	"errors"
	"sort"
)

// Version identifies the rulepack version. Increment ONLY with governance.
const Version = "v1.0.0"

// FailureMode defines deterministic rule failure reasons.
type FailureMode string

const (
	FailSchemaViolation  FailureMode = "SCHEMA_VIOLATION"
	FailInsufficientData FailureMode = "INSUFFICIENT_DATA"
	FailMismatch         FailureMode = "MISMATCH"
)

// Finding represents a deterministic rule evaluation output.
type Finding struct {
	RuleID   string      `json:"rule_id"`
	Passed   bool        `json:"passed"`
	Mode     *FailureMode `json:"failure_mode,omitempty"`
	Reason   string      `json:"reason"`
	Signals  []string    `json:"signals"`
}

// Input is the minimal normalized data required for rules.
// IMPORTANT: This must be non-sensitive, normalized, and stable.
type Input struct {
	Claims     map[string]string            `json:"claims"`     // key -> value (normalized)
	References map[string][]ReferenceRecord  `json:"references"` // key -> list of records
}

// ReferenceRecord is a non-sensitive reference data point.
type ReferenceRecord struct {
	Value  string `json:"value"`
	Source string `json:"source"`
}

// Rule defines a deterministic consistency rule.
type Rule interface {
	ID() string
	Apply(in Input) (Finding, error)
}

// RulePack is a versioned set of rules.
type RulePack struct {
	Version string
	Rules   []Rule
}

// DefaultRulePack returns the canonical rule pack.
func DefaultRulePack() RulePack {
	return RulePack{
		Version: Version,
		Rules: []Rule{
			CapTableFoundersCountRule{},
			CorporateIDStabilityRule{},
		},
	}
}

// Evaluate applies all rules deterministically and returns sorted findings.
func (rp RulePack) Evaluate(in Input) ([]Finding, error) {
	if rp.Version == "" {
		return nil, errors.New("rulepack version required")
	}
	findings := make([]Finding, 0, len(rp.Rules))

	// Stable order: sort by RuleID
	rules := make([]Rule, len(rp.Rules))
	copy(rules, rp.Rules)
	sort.Slice(rules, func(i, j int) bool { return rules[i].ID() < rules[j].ID() })

	for _, r := range rules {
		f, err := r.Apply(in)
		if err != nil {
			// Fail closed: surface deterministic failure
			return nil, err
		}
		f.Signals = stableStrings(f.Signals)
		findings = append(findings, f)
	}

	// Stable sort findings by RuleID
	sort.Slice(findings, func(i, j int) bool { return findings[i].RuleID < findings[j].RuleID })
	return findings, nil
}

func stableStrings(in []string) []string {
	out := make([]string, len(in))
	copy(out, in)
	sort.Strings(out)
	return out
}
