package artifacts

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"
)

// Generator cuida da criação de provas imutáveis.
type Generator struct{}

func NewGenerator() *Generator {
	return &Generator{}
}

// CreateProof gera um ID único baseado no hash dos dados + score + timestamp.
// Isso permite verificar no futuro se aquele score foi gerado para aqueles dados,
// sem precisar armazenar os dados originais no banco.
func (g *Generator) CreateProof(inputHash string, score float64) string {
	timestamp := time.Now().UTC().UnixNano()
	
	// Payload do artefato: InputHash + Score + Time
	data := fmt.Sprintf("%s:%.4f:%d", inputHash, score, timestamp)
	
	// Hash SHA-256 do artefato
	hash := sha256.Sum256([]byte(data))
	proofHash := hex.EncodeToString(hash[:])

	// Formato: veritas_v1_[short_hash]
	return fmt.Sprintf("veritas_v1_%s", proofHash[:16])
}

// GenerateInputFingerprint cria um hash dos dados de entrada para correlação.
// Isso é o que chamamos de "Blind Index".
func GenerateInputFingerprint(founders int, equity []float64) string {
	raw := fmt.Sprintf("%d:%v", founders, equity)
	hash := sha256.Sum256([]byte(raw))
	return hex.EncodeToString(hash[:])
}
