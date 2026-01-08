package risk

import (
	"math"
	"time"
)

// ReactiveEngine calcula os riscos voláteis e temporais.
type ReactiveEngine struct {
	DecayRate float64 // Taxa de decaimento diário (Entropy)
}

func NewReactiveEngine() *ReactiveEngine {
	return &ReactiveEngine{
		DecayRate: 0.05, // 5% de perda de confiança por dia de inatividade
	}
}

// CalculatePx combina o risco de inatividade com sinais de alerta externos.
// lastUpdate: Data da última validação de dados da startup.
// externalThreatLevel: 0.0 a 1.0 (Vindo de OSINT/Google News).
func (r *ReactiveEngine) CalculatePx(lastUpdate time.Time, externalThreatLevel float64) float64 {
	
	// 1. Cálculo de Entropia (Decaimento Temporal)
	// Quanto mais tempo sem update, maior o risco "passivo".
	hoursSinceUpdate := time.Since(lastUpdate).Hours()
	daysInactive := hoursSinceUpdate / 24.0
	
	// Fórmula de Decaimento Logístico limitado
	// Risco aumenta conforme os dias passam.
	temporalRisk := 1.0 - (1.0 / (1.0 + (r.DecayRate * daysInactive)))

	// 2. Fusão com Ameaças Externas (Active Threat)
	// Se houver uma notícia ruim (externalThreatLevel alto), isso domina.
	
	totalRisk := math.Max(temporalRisk, externalThreatLevel)

	// Clamp de segurança (Risco nunca > 1.0)
	return math.Min(totalRisk, 1.0)
}
