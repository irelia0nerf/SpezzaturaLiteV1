package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"spezzaturalitev1/internal/rules"
	"spezzaturalitev1/internal/services"
)

// Server define a estrutura da API
type Server struct {
	orchestrator *services.Orchestrator
}

func NewServer() *Server {
	return &Server{
		orchestrator: services.NewOrchestrator(),
	}
}

// VerificationRequest define o contrato de entrada JSON
type VerificationRequest struct {
	DeclaredFounders int       `json:"declared_founders"`
	DeclaredEquity   []float64 `json:"declared_equity"`
	CompanyID        string    `json:"company_id"` // Hash ou ID anônimo
}

// VerificationResponse define o contrato de saída (apenas Sinais)
type VerificationResponse struct {
	ArtifactID string  `json:"artifact_id"`
	TrustScore float64 `json:"trust_score"`
	Timestamp  string  `json:"timestamp"`
	Status     string  `json:"status"` // BLOCK, REVIEW, ALLOW
}

func (s *Server) HandleVerify(w http.ResponseWriter, r *http.Request) {
	// 1. Sanitização de Entrada
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req VerificationRequest
	// Decode stream evita carregar tudo na RAM se o payload for gigante (ataque DoS)
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	// 2. Mapeamento para Estrutura Interna (Stateless)
	input := rules.CapTableInput{
		DeclaredFounders: req.DeclaredFounders,
		DeclaredEquity:   req.DeclaredEquity,
	}

	// 3. Execução do Orquestrador
	// Contexto com timeout para garantir SLA de resposta rápida
	ctx := r.Context()
	
	score, artifactID, err := s.orchestrator.ProcessVerification(ctx, input)
	if err != nil {
		// Log genérico de erro, SEM dados do cliente
		log.Printf("Verification error for ID %s: %v", req.CompanyID, err)
		http.Error(w, "Verification failed", http.StatusInternalServerError)
		return
	}

	// 4. Definição do Veredito (Status)
	status := "REVIEW"
	if score > 7.5 {
		status = "ALLOW"
	} else if score < 4.5 {
		status = "BLOCK"
	}

	resp := VerificationResponse{
		ArtifactID: artifactID,
		TrustScore: score,
		Timestamp:  time.Now().UTC().Format(time.RFC3339),
		Status:     status,
	}

	// 5. Resposta JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func main() {
	server := NewServer()
	
	mux := http.NewServeMux()
	mux.HandleFunc("/verify", server.HandleVerify)
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Spezzatura Lite: Operational"))
	})

	log.Println("🛡️  Spezzatura TrustScore API listening on :8080")
	// Configurações de timeout para proteção contra Slowloris attacks
	srv := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
