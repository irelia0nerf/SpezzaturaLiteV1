package services

import (
	"context"
	"spezzaturalitev1/internal/artifacts"
	"spezzaturalitev1/internal/core"
	"spezzaturalitev1/internal/risk"
	"spezzaturalitev1/internal/rules"
	"time"
)

type Orchestrator struct {
	engine        *core.TrustEngine
	riskEngine    *risk.ReactiveEngine
	proofGenerator *artifacts.Generator
}

func NewOrchestrator() *Orchestrator {
	return &Orchestrator{
		engine:        core.NewTrustEngine(),
		riskEngine:    risk.NewReactiveEngine(),
		proofGenerator: artifacts.NewGenerator(),
	}
}

func (o *Orchestrator) ProcessVerification(ctx context.Context, capTableData rules.CapTableInput) (float64, string, error) {
	// Zero-Persistence cleanup
	defer func() {
		capTableData.DeclaredEquity = nil
	}()

	// 1. Coleta de Sinais Soberanos (Mock do BigQuery)
	sovereignSignals := rules.SovereignData{
		DetectedTechnicalContributors: 2,
		CorporateRegistryFounders:     2,
	}

	// 2. Validação Histórica (Integridade)
	integrityScore, err := rules.EvaluateCapTableConsistency(ctx, capTableData, sovereignSignals)
	if err != nil {
		return 0, "", err
	}

	// 3. Avaliação de Risco Reativo P(x)
	// Simulando que a startup atualizou dados há 2 dias e sem notícias ruins (0.0)
	lastUpdate := time.Now().Add(-48 * time.Hour) 
	currentRisk := o.riskEngine.CalculatePx(lastUpdate, 0.0)

	// 4. Fusão Matemática (Engine + Poison Pill)
	finalTrustScore := o.engine.CalculateTrustScore(integrityScore, currentRisk)

	// 5. Geração de Artefato Criptográfico
	inputFingerprint := artifacts.GenerateInputFingerprint(capTableData.DeclaredFounders, capTableData.DeclaredEquity)
	artifactID := o.proofGenerator.CreateProof(inputFingerprint, finalTrustScore)

	return finalTrustScore, artifactID, nil
}
