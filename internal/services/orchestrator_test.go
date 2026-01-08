package services

import (
	"context"
	"testing"
	"spezzaturalitev1/internal/rules"
)

func TestOrchestrator_ProcessVerification(t *testing.T) {
	orch := NewOrchestrator()
	ctx := context.Background()

	// Dados de entrada simulados
	input := rules.CapTableInput{
		DeclaredFounders: 2,
		DeclaredEquity:   []float64{50.0, 50.0},
	}

	// Executa o processo
	score, artifactID, err := orch.ProcessVerification(ctx, input)

	if err != nil {
		t.Fatalf("O orquestrador falhou: %v", err)
	}

	// Validações
	if score < 0 || score > 10 {
		t.Errorf("Score fora dos limites (0-10): %f", score)
	}

	if artifactID == "" {
		t.Error("Artefato de verificação não foi gerado")
	}

	t.Logf("Sucesso: Score %f, Artefato %s", score, artifactID)
}
