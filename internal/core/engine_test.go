package core

import (
	"testing"
	"math"
)

func TestSigmoidGate_PoisonPill(t *testing.T) {
	engine := NewTrustEngine()

	// Cenário 1: Startup Perfeita (Histórico Alto, Risco Baixo)
	// T^2 = 9.5, P(x) = 0.1
	scoreSafe := engine.CalculateTrustScore(9.5, 0.1)
	if scoreSafe < 8.0 {
		t.Errorf("Falha no Cenário Nominal: Esperado > 8.0, Recebido %f", scoreSafe)
	}

	// Cenário 2: Ataque "One Strike" (Histórico Perfeito, Risco Crítico)
	// A startup tem métricas ótimas (9.5), mas foi detectada uma fraude ativa (0.9)
	scoreRisky := engine.CalculateTrustScore(9.5, 0.9)
	
	// A Poison Pill deve derrubar a nota para perto de zero, ignorando o 9.5
	if scoreRisky > 2.0 {
		t.Errorf("FALHA DE SEGURANÇA: Poison Pill não ativou. Risco Alto (0.9) gerou score %f", scoreRisky)
	}
}

func TestSigmoidGate_Sensitivity(t *testing.T) {
	engine := NewTrustEngine()
	
	// Testando o ponto de corte (Threshold ~0.75)
	// Risco 0.76 deve penalizar severamente
	val := engine.sigmoidGate(0.76)
	if val > 0.5 {
		t.Errorf("Sensibilidade insuficiente: Gate em 0.76 deveria fechar, retornou %f", val)
	}
}
