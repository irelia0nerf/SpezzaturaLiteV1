package core

import (
	"math"
)

// TrustEngine encapsula a lógica matemática da Spezzatura.
type TrustEngine struct {
	// Configurações de sensibilidade do modelo
	RiskSensitivity float64 // Padrão: 10.0 (Curva agressiva)
	RiskThreshold   float64 // Padrão: 0.8 (Ponto de corte para "One Strike")
}

// NewTrustEngine inicializa o motor com parâmetros padrão de segurança.
func NewTrustEngine() *TrustEngine {
	return &TrustEngine{
		RiskSensitivity: 12.0, // Alta sensibilidade para disparar o bloqueio rápido
		RiskThreshold:   0.75, // Se 75% de probabilidade de fraude, ativa disjuntor
	}
}

// CalculateTrustScore combina o componente Histórico (T^2) com o Reativo P(x).
// Fórmula: Score = T2_Component * SigmoidGate(Px_Risk)
func (e *TrustEngine) CalculateTrustScore(historicalScore float64, reactiveRisk float64) float64 {
	gateMultiplier := e.sigmoidGate(reactiveRisk)
	
	// Aplica o disjuntor ao histórico
	finalScore := historicalScore * gateMultiplier

	// Clamp para garantir intervalo 0-10
	return math.Min(math.Max(finalScore, 0.0), 10.0)
}

// sigmoidGate atua como o fusível. Retorna um multiplicador entre 0.0 e 1.0.
// Se o risco for baixo, retorna ~1.0 (passa tudo).
// Se o risco for alto, retorna ~0.0 (bloqueia tudo).
func (e *TrustEngine) sigmoidGate(risk float64) float64 {
	// Função Sigmóide Invertida deslocada: 1 / (1 + e^(k * (x - x0)))
	// k = sensibilidade, x0 = limiar
	
	exponent := e.RiskSensitivity * (risk - e.RiskThreshold)
	denominator := 1 + math.Exp(exponent)
	
	return 1.0 / denominator
}
