package risk

import (
	"math"
	"time"
)

// RiskEventType define a categoria do sinal de alerta.
type RiskEventType string

const (
	RiskTypeLegal      RiskEventType = "LEGAL"      // Processos, Compliance
	RiskTypeTechnical  RiskEventType = "TECHNICAL"  // Downtime, Code Churn bizarro
	RiskTypeReputation RiskEventType = "REPUTATION" // Escândalos, Founder exiting
	RiskTypeVolatility RiskEventType = "VOLATILITY" // Mudanças frequentes de dados
)

// RiskEvent representa um evento capturado por OSINT ou telemetria.
type RiskEvent struct {
	Type      RiskEventType
	Severity  float64   // 0.0 (Irrelevante) a 1.0 (Catastrófico/Poison Pill)
	Timestamp time.Time // Quando o evento ocorreu
}

// ReactiveEngine calcula a entropia e a volatilidade.
type ReactiveEngine struct {
	DecayHalfLifeDays float64 // Dias para a confiança cair pela metade sem updates
}

func NewReactiveEngine() *ReactiveEngine {
	return &ReactiveEngine{
		// Configuração agressiva: após 90 dias sem update, a confiança cai 50%.
		DecayHalfLifeDays: 90.0,
	}
}

// CalculatePx executa a fusão de riscos Temporais, Eventuais e de Volatilidade.
// Retorna um valor entre 0.0 (Risco Zero) e 1.0 (Risco Total/Bloqueio).
func (r *ReactiveEngine) CalculatePx(lastUpdate time.Time, events []RiskEvent) float64 {
	
	// 1. Risco Temporal (Entropia)
	// Usa decaimento exponencial baseado na meia-vida.
	// Fórmula: P_time = 1 - (0.5)^(days / half_life)
	daysSinceUpdate := time.Since(lastUpdate).Hours() / 24.0
	temporalRisk := 1.0 - math.Pow(0.5, daysSinceUpdate/r.DecayHalfLifeDays)

	// 2. Risco de Eventos (Sinais Ativos)
	// Processa a lista de "Red Flags" detectadas.
	var eventRiskSum float64
	var maxSeverity float64

	for _, event := range events {
		// Decaimento da relevância do evento (Notícias velhas importam menos)
		eventAgeHours := time.Since(event.Timestamp).Hours()
		// Eventos perdem 10% de relevância a cada 30 dias (heurística)
		relevance := math.Max(0, 1.0-(eventAgeHours/(24*30)*0.1))
		
		impact := event.Severity * relevance
		
		// Acumulador de risco (soma amortecida)
		eventRiskSum += impact
		
		// Rastreia o evento mais severo para o "One Strike"
		if impact > maxSeverity {
			maxSeverity = impact
		}
	}

	// Normaliza a soma de eventos (Curva assintótica para nunca passar de 1.0)
	// Risco de Eventos Agregado
	aggregateEventRisk := 1.0 - math.Exp(-eventRiskSum)

	// 3. Mecanismo de Disjuntor (Poison Pill Override)
	// Se houver UM evento com severidade > 0.85 (ex: Fraude Confirmada),
	// o risco salta imediatamente para 1.0, ignorando o resto.
	if maxSeverity > 0.85 {
		return 1.0
	}

	// 4. Fusão Final (Máximo entre Temporal e Eventos)
	// O sistema é pessimista: assume o pior cenário entre o silêncio e o ruído.
	totalRisk := math.Max(temporalRisk, aggregateEventRisk)

	return math.Min(totalRisk, 1.0)
}

// CalculateVolatility mede a instabilidade dos dados.
// changesCount: quantas vezes o cap table mudou nos últimos 30 dias.
func (r *ReactiveEngine) CalculateVolatility(changesCount int) float64 {
	// Se mudou mais de 3 vezes em um mês, é altamente suspeito.
	// Sigmóide simples centrada em 4 mudanças.
	k := 1.0 // inclinação
	x0 := 4.0 // ponto médio
	
	return 1.0 / (1.0 + math.Exp(-k*(float64(changesCount)-x0)))
}
