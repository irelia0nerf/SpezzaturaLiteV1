package services

import (
	"context"
	"spezzaturalitev1/internal/artifacts"
	"spezzaturalitev1/internal/core"
	"spezzaturalitev1/internal/risk"
	"spezzaturalitev1/internal/rules"
	"time"
)

// Orchestrator gerencia o fluxo de dados entre os componentes da Spezzatura.
// Ele segue estritamente o padrão Zero-Persistence.
type Orchestrator struct {
	engine         *core.TrustEngine
	riskEngine     *risk.ReactiveEngine
	proofGenerator *artifacts.Generator
}

// NewOrchestrator inicializa os serviços dependentes.
func NewOrchestrator() *Orchestrator {
	return &Orchestrator{
		engine:         core.NewTrustEngine(),
		riskEngine:     risk.NewReactiveEngine(),
		proofGenerator: artifacts.NewGenerator(),
	}
}

// ProcessVerification executa a lógica de Due Diligence Automatizada.
func (o *Orchestrator) ProcessVerification(ctx context.Context, capTableData rules.CapTableInput) (float64, string, error) {
	
	// --- ZERO-PERSISTENCE PROTOCOL ---
	// Assegura que referências a dados sensíveis sejam descartadas ao sair do escopo.
	defer func() {
		capTableData.DeclaredEquity = nil
	}()

	// 1. Coleta de Sinais Soberanos (Simulação de Cliente BigQuery)
	// Em produção, isso faria uma query autenticada no Data Warehouse.
	sovereignSignals := rules.SovereignData{
		DetectedTechnicalContributors: 2, // Ex: Encontramos 2 devs contribuindo código
		CorporateRegistryFounders:     2, // Ex: Encontramos 2 sócios no contrato social
	}

	// 2. Validação Histórica (Componente T^2)
	// Avalia a consistência entre o que foi dito (Pitch) e o que existe (Física).
	integrityScore, err := rules.EvaluateCapTableConsistency(ctx, capTableData, sovereignSignals)
	if err != nil {
		return 0, "", err
	}

	// 3. Avaliação de Risco Reativo P(x) (ATUALIZADO v2.0)
	// Calcula o decaimento temporal, volatilidade e eventos de risco (OSINT).
	
	// A. Simulação de "Última Atualização" (Time Decay)
	// Vamos assumir que a startup atualizou seus dados há 48 horas.
	lastUpdate := time.Now().Add(-48 * time.Hour)

	// B. Detecção de Eventos de Risco (Simulação de OSINT)
	// Em produção, isso viria de crawlers de notícias ou APIs legais.
	detectedRiskEvents := []risk.RiskEvent{
		{
			Type:      risk.RiskTypeReputation,
			Severity:  0.15, // Risco baixo (ex: reclamação isolada em rede social)
			Timestamp: time.Now().Add(-24 * time.Hour),
		},
	}

	// C. Cálculo de Volatilidade de Dados
	// Se a startup mudou o Cap Table 2 vezes no último mês (comportamento normal/baixo risco).
	changesInLastMonth := 2
	volatilityRisk := o.riskEngine.CalculateVolatility(changesInLastMonth)

	// Se a volatilidade for significativa, adicionamos como um evento de risco técnico
	if volatilityRisk > 0.1 {
		detectedRiskEvents = append(detectedRiskEvents, risk.RiskEvent{
			Type:      risk.RiskTypeVolatility,
			Severity:  volatilityRisk,
			Timestamp: time.Now(),
		})
	}

	// D. Cálculo do P(x) Final
	currentRisk := o.riskEngine.CalculatePx(lastUpdate, detectedRiskEvents)

	// 4. Fusão Matemática (Engine + Poison Pill)
	// O motor aplica a lógica Sigmóide: Score = T^2 * SigmoidGate(Px)
	finalTrustScore := o.engine.CalculateTrustScore(integrityScore, currentRisk)

	// 5. Geração de Artefato Criptográfico (Evidence)
	// Cria um hash cego dos inputs para permitir verificação futura sem guardar os dados.
	inputFingerprint := artifacts.GenerateInputFingerprint(capTableData.DeclaredFounders, capTableData.DeclaredEquity)
	artifactID := o.proofGenerator.CreateProof(inputFingerprint, finalTrustScore)

	return finalTrustScore, artifactID, nil
}
