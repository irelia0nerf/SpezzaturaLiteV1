### Arquitetura de Governança de Decisão: Spezzatura Lite (TrustScore API) v2.6

Este documento detalha o *framework* **Spezzatura Lite**, uma engine de verificação sem estado (*stateless*) projetada para mitigar o **Risco de Admissão** (*Intake Risk*) em ecossistemas de capital de risco e infraestruturas de nuvem governadas. O sistema opera sob o paradigma de **"Confiança pela Física"**, substituindo heurísticas subjetivas por um motor matemático auditável.

---

#### 1. Pilar 1: Inteligência Determinística ($T^2$) - O Motor Histórico

O núcleo analítico fundamenta-se em um modelo estritamente **multiplicativo**, atuando como um "fusível" de admissão. Se qualquer vetor de integridade for nulo, a saída do produto bruto ($p$) será zero, impedindo que métricas de vaidade compensem falhas de conformidade.

**A Fórmula Base (v2.6):**
$$p = C \times A \times (T^2) \times U \times R \times \hat{A}$$

*   **Vetor C (Completude):** Razão entre artefatos esperados e fornecidos (P&L, *Cap Table*, registros legais).
*   **Vetor A (Atividade):** Telemetria do "batimento cardíaco digital" via integrações com a **GitHub REST API** (frequência de *commits*, *pull requests*) e tráfego de rede.
*   **Vetor $T^2$ (Confiança ao Quadrado):** Aplicação de uma **Lei de Potência** para sinais de entidades de elite (Y Combinator, Stripe Atlas). O peso é exponencial para refletir a validação institucional prévia.
*   **Vetor U (Unicidade):** Implementação do algoritmo **SimHash** (LSH) para detectar "startups sintéticas". Ele calcula a distância de Hamming entre o material submetido e modelos genéricos ou de fraudes conhecidas.
*   **Vetor R (Reputação):** Análise de sentimento e consistência biográfica via **OSINT**.
*   **Vetor $\hat{A}$ (Ânimo):** Processamento de Linguagem Natural (NLP) para medir a intencionalidade estratégica do fundador.

**Compressão Logarítmica:**
O produto $p$ é processado via **$\log_2$** para normalizar a variância massiva de métricas entre startups (ex: receitas de $\$10k$ vs. $\$10M$). O logaritmo garante estabilidade numérica e atua como barreira *anti-gaming*. No *framework* v2.6, este valor é mapeado linearmente para uma escala de **0 a 10** ($score\_hist$).

---

#### 2. Pilar 2: Termômetro de Risco Reativo [$P(x)$]

Diferente do histórico, o componente reativo monitora a volatilidade imediata e o decaimento da integridade.

*   **Decaimento Exponencial (Cooling):** A confiança "esfria" na ausência de novos sinais. Aplica-se $P_{new} = P_{prev} \times e^{-\Delta t/\tau}$, onde $\tau$ representa a meia-vida do risco.
*   **Injeção de Risco (Heat Up):** Sinais de alerta (processos judiciais, inatividade súbita no GitHub) recebem severidade de 0 a 1. Eventos com severidade $> 0.8$ acionam um **Booster de 1.5x**, forçando a queda imediata do score.

---

#### 3. Pilar 3: Fusão e a Porta Sigmoide (O Disjuntor)

A convergência entre o histórico e o reativo é mediada pela **Função Sigmoide**, que opera como um **disjuntor matemático** (*Poison Pill*).

**Lógica de Fusão:**
$$z = \alpha \cdot score\_hist - \beta \cdot (10 \cdot P(x))$$
$$Score\_final = 10 \cdot \sigma(z) \quad \text{onde} \quad \sigma(z) = \frac{1}{1 + e^{-z}}$$

A natureza não linear da Sigmoide garante a **"Regra de Um Golpe"**: um risco reativo crítico (P próximo de 1) "descarrega" a função, colapsando a nota final para zero, independentemente da solidez histórica.

**Protocolos de Veredito Automático:**
*   **BLOCK ($\leq 4.5$):** Negação imediata por risco crítico ou inconsistência física grave.
*   **REVIEW (4.6 a 5.5):** Situação ambígua requerendo intervenção humana do *Success Manager*.
*   **ALLOW ($> 5.5$):** Integridade validada; aprovação e provisionamento automático.

---

#### 4. Stack Tecnológica e Fluxo End-to-End (GCP Native)

O Spezzatura Lite utiliza uma arquitetura **Stateless de Zero-Persistência**, eliminando a responsabilidade fiduciária sobre "Dados Tóxicos" (PII/IP).

1.  **Ingestão e Blindagem:** O payload (PDF/JSON) entra via **FastAPI** protegido pelo **Google Cloud Armor**.
2.  **Execução Confiável (TEE):** O processamento ocorre em memória RAM criptografada no **Google Cloud Confidential Space** (AMD SEV-SNP). Nem o Google nem a FoundLab acessam o conteúdo bruto.
3.  **Processamento e Enriquecimento:**
    *   **Vertex AI Model Registry:** O motor matemático executa a inferência de risco.
    *   **BigQuery Remote Functions:** Acionam o **Cloud Run** (isolado via **gVisor**) para coletar sinais externos em tempo real.
4.  **Crypto-Shredding:** Após a geração do score, a chave de criptografia AES-256 da sessão (DEK) é permanentemente deletada do **Cloud KMS**. Os dados originais tornam-se matematicamente irrecuperáveis, satisfazendo o "Direito ao Esquecimento" da LGPD.
5.  **Auditabilidade Forense:**
    *   **Rationale Hash:** Um fingerprint **SHA-256** da justificativa da decisão é gerado para garantir imutabilidade.
    *   **Blockchain Anchoring:** O hash da decisão é ancorado em árvores de Merkle na rede **Polygon**, criando um carimbo de tempo público e à prova de falsificação.

---

