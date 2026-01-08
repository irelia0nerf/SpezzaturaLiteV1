package service

import (
	"errors"

	"github.com/foundlab/spezzatura-lite/internal/artifacts"
	"github.com/foundlab/spezzatura-lite/internal/core"
	"github.com/foundlab/spezzatura-lite/internal/rules"
)

// EvaluationResult is the composed, non-advisory output of the orchestrator.
type EvaluationResult struct {
	CoreResult   core.Result           `json:"core_result"`
	Findings     []rules.Finding       `json:"findings"`
	Proof        artifacts.ProofOfCheck `json:"proof_of_check"`
}

// EngineDescriptor identifies the executing engine for audit.
var EngineDescriptor = artifacts.EngineDescriptor{
	Name:    "spezzatura-lite-orchestrator",
	Version: "v1.0.0",
}

// Evaluate orchestrates core evaluation, rules, scoring, and proof emission.
func Evaluate(input core.Input, ruleInput rules.Input) (EvaluationResult, error) {
	// 1) Core deterministic evaluation (canonical hash + base score)
	coreRes, err := core.Evaluate(input)
	if err != nil {
		return EvaluationResult{}, err
	}

	// 2) Rules evaluation (versioned rulepack)
	rp := rules.DefaultRulePack()
	findings, err := rp.Evaluate(ruleInput)
	if err != nil {
		return EvaluationResult{}, err
	}

	// 3) Aggregate signals and invariants
	signals := []string{}
	invariants := map[string]bool{
		"non_empty_match": coreRes.FailureMode == nil,
	}

	for _, f := range findings {
		signals = append(signals, f.Signals...)
		if !f.Passed {
			invariants["rules_all_passed"] = false
		}
	}
	if _, ok := invariants["rules_all_passed"]; !ok {
		invariants["rules_all_passed"] = true
	}

	// 4) Fail-closed if core already failed
	if coreRes.FailureMode != nil {
		return EvaluationResult{}, errors.New("core evaluation failed; proof not emitted")
	}

	// 5) Emit proof-of-check artifact
	proof, err := artifacts.NewProof(
		coreRes.CorrelationID,
		coreRes.InputHash,
		coreRes.TrustScore,
		signals,
		invariants,
		EngineDescriptor,
	)
	if err != nil {
		return EvaluationResult{}, err
	}

	return EvaluationResult{
		CoreResult: coreRes,
		Findings:   findings,
		Proof:      proof,
	}, nil
}
