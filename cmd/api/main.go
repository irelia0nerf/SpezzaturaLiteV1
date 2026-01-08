package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/foundlab/spezzatura-lite/internal/core"
	"github.com/foundlab/spezzatura-lite/internal/rules"
	"github.com/foundlab/spezzatura-lite/internal/service"
)

type EvaluateRequest struct {
	CorrelationID string            `json:"correlation_id"`
	AsOfUTC       time.Time         `json:"as_of_utc"`
	Claims        map[string]string `json:"claims"`
	References    map[string][]struct {
		Value  string `json:"value"`
		Source string `json:"source"`
	} `json:"references"`
}

type EvaluateResponse struct {
	Result service.EvaluationResult `json:"result"`
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/healthz", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	})

	mux.HandleFunc("/evaluate", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}

		defer r.Body.Close()
		var req EvaluateRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "invalid json", http.StatusBadRequest)
			return
		}

		// Build core input
		coreClaims := []core.Claim{}
		for k, v := range req.Claims {
			// Expect keys in form namespace.key
			parts := splitKey(k)
			if len(parts) != 2 {
				http.Error(w, "invalid claim key format", http.StatusBadRequest)
				return
			}
			coreClaims = append(coreClaims, core.Claim{Namespace: parts[0], Key: parts[1], Value: v})
		}

		coreRefs := []core.Reference{}
		ruleRefs := map[string][]rules.ReferenceRecord{}
		for k, recs := range req.References {
			parts := splitKey(k)
			if len(parts) != 2 {
				http.Error(w, "invalid reference key format", http.StatusBadRequest)
				return
			}
			for _, rr := range recs {
				coreRefs = append(coreRefs, core.Reference{
					Namespace: parts[0],
					Key:       parts[1],
					Value:     rr.Value,
					Source:    rr.Source,
				})
				ruleRefs[k] = append(ruleRefs[k], rules.ReferenceRecord{
					Value:  rr.Value,
					Source: rr.Source,
				})
			}
		}

		coreIn := core.Input{
			CorrelationID: req.CorrelationID,
			AsOfUTC:       req.AsOfUTC,
			Claims:        coreClaims,
			References:    coreRefs,
		}

		ruleIn := rules.Input{
			Claims:     req.Claims,
			References: ruleRefs,
		}

		res, err := service.Evaluate(coreIn, ruleIn)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnprocessableEntity)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(EvaluateResponse{Result: res})
	})

	log.Println("api listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}

func splitKey(k string) []string {
	out := []string{}
	cur := ""
	for _, c := range k {
		if c == '.' {
			out = append(out, cur)
			cur = ""
		} else {
			cur += string(c)
		}
	}
	out = append(out, cur)
	return out
}
