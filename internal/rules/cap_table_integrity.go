package rules

import (
	"context"
	"fmt"
	"math"
)

// CapTableInput define os dados declarados pelo fundador.
type CapTableInput struct {
	DeclaredFounders int
	DeclaredEquity   []float64 // Ex: [50.0, 50.0]
}

// SovereignData simula dados vindos do BigQuery/Fontes Autoritativas.
type SovereignData struct {
	DetectedTechnicalContributors int // Devs com commits/atividade real
	CorporateRegistryFounders     int // Sócios no contrato social
}

// EvaluateCapTableConsistency calcula a divergência entre narrativa e sinal.
func EvaluateCapTableConsistency(ctx context.Context, input CapTableInput, sovereign SovereignData) (float64, error) {
	
	// 1. Verificação de Existência Física (Sinal Soberano)
	if sovereign.DetectedTechnicalContributors == 0 && sovereign.CorporateRegistryFounders == 0 {
		// ALERTA VERMELHO: Startup fantasma.
		return 0.0, nil
	}

	// 2. Cálculo da Divergência de Fundadores
	// Penaliza se declarar MAIS fundadores do que os detectados (risco de "laranjas" ou inflação de time).
	// Não penaliza severamente se declarar MENOS (pode ser erro clerical).
	
	var consistencyRatio float64
	
	if input.DeclaredFounders > 0 {
		// Usamos o maior sinal soberano disponível (Técnico ou Legal)
		bestSignal := math.Max(float64(sovereign.DetectedTechnicalContributors), float64(sovereign.CorporateRegistryFounders))
		
		consistencyRatio = bestSignal / float64(input.DeclaredFounders)
	} else {
		return 0.0, fmt.Errorf("invalid input: declared founders must be > 0")
	}

	// 3. Normalização do Score (0 a 10)
	// Se ratio >= 1.0 (Detectou tudo ou mais), Score = 10.
	// Se ratio < 1.0, o score cai proporcionalmente.
	
	score := math.Min(consistencyRatio, 1.0) * 10.0
	
	return score, nil
}
