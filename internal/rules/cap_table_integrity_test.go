package rules

import (
	"context"
	"testing"
)

func TestEvaluateCapTableConsistency(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		name           string
		input          CapTableInput
		sovereign      SovereignData
		expectedScore  float64
		expectError    bool
	}{
		{
			name: "Consistência Perfeita",
			input: CapTableInput{DeclaredFounders: 2},
			sovereign: SovereignData{DetectedTechnicalContributors: 2, CorporateRegistryFounders: 2},
			expectedScore: 10.0,
			expectError:   false,
		},
		{
			name: "Inflação de Time (Diz 4, Tem 2)",
			input: CapTableInput{DeclaredFounders: 4},
			sovereign: SovereignData{DetectedTechnicalContributors: 2, CorporateRegistryFounders: 2},
			expectedScore: 5.0, // 2/4 * 10
			expectError:   false,
		},
		{
			name: "Startup Fantasma (Diz 3, Nada Detectado)",
			input: CapTableInput{DeclaredFounders: 3},
			sovereign: SovereignData{DetectedTechnicalContributors: 0, CorporateRegistryFounders: 0},
			expectedScore: 0.0,
			expectError:   false,
		},
		{
			name: "Erro de Input (Zero Declarado)",
			input: CapTableInput{DeclaredFounders: 0},
			sovereign: SovereignData{DetectedTechnicalContributors: 1},
			expectedScore: 0.0,
			expectError:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			score, err := EvaluateCapTableConsistency(ctx, tt.input, tt.sovereign)
			
			if tt.expectError {
				if err == nil {
					t.Errorf("Esperava erro, mas recebeu nil")
				}
			} else {
				if err != nil {
					t.Errorf("Erro inesperado: %v", err)
				}
				if score != tt.expectedScore {
					t.Errorf("Score incorreto. Esperado: %f, Recebido: %f", tt.expectedScore, score)
				}
			}
		})
	}
}
