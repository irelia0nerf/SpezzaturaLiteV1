<!-- AUDIT-SPEZZATURA.md -->

<div align="center">
  <img
    src="https://placehold.co/1200x300/1e293b/ffffff?text=AUDIT-READY%3A+SPEZZATURA+FRAMEWORK+v2.6"
    alt="Audit-Ready: Spezzatura Framework v2.6"
    width="100%"
  />
</div>

<div align="center">

# 🧠 Especificação Técnica de Governança Determinística  
## Spezzatura Framework v1.2 / v2.6 + TrustScore API (FoundLab)

![status](https://img.shields.io/badge/status-audit--ready-16a34a?style=flat-square)
![framework](https://img.shields.io/badge/framework-Spezzatura-2563eb?style=flat-square)
![version](https://img.shields.io/badge/version-2.6-7c3aed?style=flat-square)
![api](https://img.shields.io/badge/api-TrustScore-0ea5e9?style=flat-square)
![license](https://img.shields.io/badge/license-internal--spec-334155?style=flat-square)

</div>

> [!NOTE]
> Este documento formaliza, de modo **determinístico** e **auditável**, a “fórmula completa” do Spezzatura (v1.2/v2.6) e seu acoplamento ao circuito reativo **P(x)** com **porta sigmoide** (Sigmoid Gate), conforme o contexto fornecido. [1][2]

---

## 📚 Sumário

- [1. Abstract e Ontologia do Modelo](#1-abstract-e-ontologia-do-modelo)
- [2. Escopo, Convenções e Invariantes](#2-escopo-convenções-e-invariantes)
- [3. Visão Geral do Pipeline](#3-visão-geral-do-pipeline)
- [4. Seção 1 — O Núcleo Multiplicativo (Historic Vector)](#4-seção-1--o-núcleo-multiplicativo-historic-vector)
  - [4.1 Fórmula Canônica (v1.2/v2.6)](#41-fórmula-canônica-v12v26)
  - [4.2 Ontologia dos Vetores: C, A, T, U, R, Â](#42-ontologia-dos-vetores-c-a-t-u-r-â)
  - [4.3 Propriedades Matemáticas do Produto](#43-propriedades-matemáticas-do-produto)
- [5. Seção 2 — Regularização por Compressão Logarítmica](#5-seção-2--regularização-por-compressão-logarítmica)
  - [5.1 raw = log₂(p)](#51-raw--log₂p)
  - [5.2 Normalização v2.6 para Escala 0–10](#52-normalização-v26-para-escala-010)
  - [5.3 Racional de LOG_FLOOR_BITS (anti -∞)](#53-racional-de-log_floor_bits-anti--∞)
- [6. Seção 3 — Volatilidade Reativa P(x) (Poison Pill)](#6-seção-3--volatilidade-reativa-px-poison-pill)
  - [6.1 Decaimento Exponencial e “Meia-vida” τ](#61-decaimento-exponencial-e-meia-vida-τ)
  - [6.2 Semântica de severity e Δt](#62-semântica-de-severity-e-Δt)
- [7. Seção 4 — Disjuntor Sigmoide (Fusion / Sigmoid Gate)](#7-seção-4--disjuntor-sigmoide-fusion--sigmoid-gate)
  - [7.1 z = α·score_hist − β·(10·P(x))](#71-z--αscore_hist--β10px)
  - [7.2 Score_final = 10·σ(z)](#72-score_final--10σz)
  - [7.3 Interpretação: Limiar, Saturação e Robustez](#73-interpretação-limiar-saturação-e-robustez)
- [8. Seção 5 — Veredito e Classificação](#8-seção-5--veredito-e-classificação)
- [9. Stack & Auditability](#9-stack--auditability)
  - [9.1 Artefatos, Linhagem de Dados e Reprodutibilidade](#91-artefatos-linhagem-de-dados-e-reprodutibilidade)
  - [9.2 Rationale Hash (SHA-256) e Imutabilidade](#92-rationale-hash-sha-256-e-imutabilidade)
  - [9.3 Governança de Parâmetros (α, β, τ, LOG_FLOOR_BITS)](#93-governança-de-parâmetros-α-β-τ-log_floor_bits)
- [10. Referência de Implementação em L11 (Audit Mode)](#10-referência-de-implementação-em-l11-audit-mode)
- [11. Apêndice](#11-apêndice)
  - [A. Tabelas de Parametrização](#a-tabelas-de-parametrização)
  - [B. Casos de Teste Determinísticos](#b-casos-de-teste-determinísticos)
  - [C. Glossário](#c-glossário)
- [Referências](#referências)

---

## 1. Abstract e Ontologia do Modelo

O Spezzatura Framework é um mecanismo de decisão **determinístico** destinado a reduzir o “risco de admissão” (intake risk) por meio de uma ontologia de sinais mensuráveis, evitando heurísticas qualitativas pouco auditáveis. [1]  
Sua premissa central é tratar integridade/qualidade de um ativo como propriedade inferível a partir de telemetria digital e consistência documental, codificando tais evidências em vetores normalizados e combinados por uma álgebra explícita. [1][2]

Em termos epistemológicos:  
- o **núcleo histórico** (Historic Vector) materializa “estado acumulado” de evidências estáveis; [1]  
- o **componente reativo** P(x) representa “energia de risco” com decaimento temporal (perdão algorítmico) e injeções de severidade; [2]  
- o **disjuntor sigmoide** atua como porta não linear, impondo saturação, limiar e comportamento de “circuit breaker” quando risco reativo supera tolerâncias definidas pela governança. [2]

> [!TIP]
> Para auditoria: pense em três camadas separáveis — **(i)** medição/normalização de evidências, **(ii)** composição determinística, **(iii)** trilha imutável (hash) do racional. [1][2]

---

## 2. Escopo, Convenções e Invariantes

### 2.1 Escopo
Este documento cobre:  
1) fórmula multiplicativa canônica (v1.2/v2.6), 2) compressão log₂, 3) normalização v2.6, 4) P(x) reativo, 5) fusão com sigmoide, 6) veredito, 7) auditabilidade (hash SHA-256). [1][2]

### 2.2 Convenções Numéricas (Domínios)
- Vetores históricos **C, A, T, U, R, Â** são assumidos em **[0, 1]** após normalização. [1]  
- O produto bruto **p** reside em **[0, 1]** (quando todos os vetores ∈ [0,1]). [1]  
- A saída histórica normalizada **score_hist** reside em **[0, 10]** (após clamp). [1]  
- P(x) tipicamente reside em **[0, 1]** (por construção de severity e decaimento). [2]  
- Score_final reside em **(0, 10)** pela sigmoide, com saturação próxima aos extremos conforme z → ±∞. [2]

> [!WARNING]
> Se qualquer vetor histórico for 0, então **p = 0** e a estabilidade numérica exige um piso (LOG_FLOOR_BITS) para evitar **log₂(0) = −∞** na etapa logarítmica. [1]

### 2.3 Invariantes Auditáveis (Propriedades que Devem Sempre Valer)
1) **Determinismo forte**: mesma entrada + mesmos parâmetros ⇒ mesma saída. [1][2]  
2) **Monotonicidade local** do núcleo: aumentar qualquer vetor (mantidos os demais) não reduz p. [1]  
3) **Fusível multiplicativo**: qualquer componente crítico nulo colapsa p. [1]  
4) **Limitabilidade**: score_hist e Score_final permanecem em escala controlada (0–10). [1][2]  
5) **Decaimento temporal**: mantendo severity fixa, P(x) diminui monotonicamente com Δt. [2]

---

## 3. Visão Geral do Pipeline

```mermaid
flowchart LR
  A[Coleta de Evidências] --> B[Normalização Vetorial]
  B --> C[Produto Histórico p]
  C --> D[Compressão log2 + score_hist]
  A --> E[Eventos Reativos]
  E --> F[P(x) = severity * e^(-Δt/τ)]
  D --> G[Fusão: z = α*score_hist - β*(10*P(x))]
  F --> G
  G --> H[Score_final = 10 * σ(z)]
  H --> I[Veredito: BLOCK / REVIEW / ALLOW]
  H --> J[Rationale -> SHA-256 -> trilha imutável]
```

Cada bloco deve produzir artefatos auditáveis: entradas normalizadas, parâmetros efetivos, intermediários (p, raw, score_hist, P(x), z) e saída final com racional textual e hash. [1][2]

---

## 4. Seção 1 — O Núcleo Multiplicativo (Historic Vector)

### 4.1 Fórmula Canônica (v1.2/v2.6)

A forma canônica do produto histórico é:

\[
p = C \times A \times (T^2) \times U \times R \times \hat{A}
\]  

Onde:  
- **C** = Completeness (Completude/Conformidade)  
- **A** = Activity (Atividade)  
- **T** = Trust (Confiança) com **lei de potência** \(T^2\)  
- **U** = Uniqueness (Unicidade)  
- **R** = Reputation (Reputação)  
- **Â** = Animus/Intent (Ânimo/Intencionalidade) [1]

A escolha multiplicativa impõe um “invariante de integridade”: o sistema não permite compensação total de falhas estruturais por um único sinal excepcional (anti-“score doping”). [1]

---

### 4.2 Ontologia dos Vetores: C, A, T, U, R, Â

A seguir, uma especificação auditável de alto nível (sem prescrever um fornecedor específico), com foco em observabilidade, entropia e invariantes de normalização. [1][2]

| Vetor | Nome | Intuição Ontológica | Exemplo de Evidência | Normalização (exemplo) | Risco que mitiga |
|------:|------|----------------------|-----------------------|-------------------------|------------------|
| C | Completeness | “massa documental” mínima para identidade operacional | Cap table, IP assignment, P&L, políticas | razão entregues/esperados, clamp [0,1] | fraude por ausência de lastro |
| A | Activity | “batimento digital” e continuidade operacional | commits, releases, domínio, cadência | função de recência+frequência → [0,1] | empresas “fantasma” |
| T | Trust | “selo institucional” como lei de potência | validações e sinais de entidades | T ∈ [0,1], elevar ao quadrado | validação fraca vs forte |
| U | Uniqueness | “distância semântica” contra clones | similaridade de pitch vs base | 1 − sim(pitch,base) → [0,1] | startups sintéticas |
| R | Reputation | “consistência externa” e OSINT | histórico público, consistência biográfica | score de consistência → [0,1] | inconsistência/“time-travel” |
| Â | Animus | “intencionalidade” e clareza estratégica | NLP em narrativa executiva | classificador calibrado → [0,1] | ambiguidade oportunista |

**Observação de auditoria**: a implementação concreta pode variar, porém **o contrato** deve garantir: (i) mapeamento determinístico para [0,1], (ii) versionamento do extrator, (iii) provas de entrada (links, hashes de documentos, timestamps), (iv) logs de normalização. [1][2]

#### C — Completeness (Completude/Conformidade)
C mede a completude relativa de um conjunto mínimo de artefatos exigidos para um determinado estágio/fluxo de admissão. [1]  
Ontologicamente, C reduz a entropia de incerteza documental: quanto maior a cobertura de evidências, menor o espaço de hipóteses compatíveis com “identidade falsa”. [1]

#### A — Activity (Atividade)
A quantifica sinais de continuidade: presença de evolução temporal e cadência operacional. [1]  
Em auditoria, A deve ser derivada de séries temporais com janela definida e regras de recência explícitas, evitando ambiguidade semântica do tipo “atividade recente” sem unidade temporal. [1]

#### T² — Trust com Lei de Potência
O termo \(T^2\) implementa uma lei de potência deliberada: sinais de confiança institucional são **não lineares** (incrementos próximos de 1 carregam mais “energia de credibilidade” do que incrementos próximos de 0). [1]  
Isso força uma geometria em que “confiança alta” não é apenas “um pouco melhor” — ela domina proporcionalmente de forma controlada, desde que não haja colapso por vetores nulos. [1]

#### U — Uniqueness (Unicidade)
U modela a distância contra padrões genéricos, clones e fraudes por template, preferindo representações robustas (ex.: hashing semântico) com parâmetros fixados e reprodutíveis. [1]  
Em termos de entropia, U busca preservar diversidade informacional: conteúdos muito próximos de um conjunto base têm baixa entropia diferencial e, portanto, menor credibilidade de originalidade. [1]

#### R — Reputation (Reputação)
R sintetiza consistência pública e integridade de narrativas observáveis, priorizando verificabilidade (fontes externas, coerência cronológica, sinais de contradição). [1]  
Para auditoria: R exige logs de evidências e critérios de consistência, evitando arbitrariedade interpretativa. [1][2]

#### Â — Animus/Intent (Ânimo/Intencionalidade)
Â representa intencionalidade: clareza estratégica, consistência retórica e ausência de padrões típicos de “venda vazia”, idealmente por extratores NLP versionados e calibrados. [1]  
Â é particularmente sensível a drift de modelos; por isso, a governança deve registrar versão do modelo e dataset de calibração (ou ao menos seu hash). [1][2]

---

### 4.3 Propriedades Matemáticas do Produto

**(i) Fusível multiplicativo**  
Se algum vetor crítico é 0, o produto colapsa: \(p = 0\). [1]  
Este comportamento é desejado quando “ausência” em vetores essenciais equivale a invalidar o ativo, impedindo compensação por sinais cosméticos. [1]

**(ii) Monotonicidade**  
Para vetores em [0,1], \(p\) é monotônico em cada coordenada (mantidas as demais), logo a melhoria de um sinal não piora o produto histórico. [1]

**(iii) Interpretação informacional (entropia)**  
Como produto de fatores normalizados, \(p\) pode ser lido como uma “probabilidade composta” sob hipótese de independência operacional aproximada, o que é útil como heurística auditável (não como verdade ontológica). [1]  
A compressão logarítmica (Seção 2) torna essa leitura ainda mais estável numericamente. [1]

---

## 5. Seção 2 — Regularização por Compressão Logarítmica

### 5.1 raw = log₂(p)

A compressão logarítmica é definida como:

\[
raw = \log_2(p)
\]

A base 2 possui interpretação direta em “dobras” (doublings), o que é útil para expressar variações multiplicativas em escala comparável. [1]  
Como \(p \in (0,1]\), então \(raw \in (-\infty, 0]\); daí a necessidade de um piso para auditoria e estabilidade numérica. [1]

> [!NOTE]
> A função log₂ transforma produto em soma: \(\log_2(\prod_i v_i) = \sum_i \log_2(v_i)\).  
> Isso preserva o caráter multiplicativo, mas reduz “explosões de escala” e melhora interpretabilidade. [1]

---

### 5.2 Normalização v2.6 para Escala 0–10

A normalização auditável (conforme contexto fornecido) é:

\[
score_{hist} = 10 \times \left(1 + \frac{\log_2(p)}{LOG\_FLOOR\_BITS}\right)
\]

Com recomendação operacional de **clamp** para manter limites:

\[
score_{hist} = \mathrm{clamp}\left(10 \times \left(1 + \frac{\log_2(p')}{LOG\_FLOOR\_BITS}\right), 0, 10 \right)
\]  

onde:

\[
p' = \max(p, 2^{-LOG\_FLOOR\_BITS})
\]

Esse mapeamento cria uma bijeção prática (por partes) entre \(p\) e um score histórico em 0–10, evitando valores negativos e evitando \(-\infty\) quando \(p=0\). [1]

**Leitura geométrica**:  
- \(p = 1 \Rightarrow \log_2(p)=0 \Rightarrow score_{hist}=10\). [1]  
- \(p = 2^{-LOG\_FLOOR\_BITS} \Rightarrow \log_2(p)=-LOG\_FLOOR\_BITS \Rightarrow score_{hist}=0\). [1]

---

### 5.3 Racional de LOG_FLOOR_BITS (anti -∞)

LOG_FLOOR_BITS define um piso de precisão e um limite inferior operacional para \(p\), servindo como invariante de estabilidade numérica: nenhuma execução pode produzir \(raw = -\infty\). [1]  
Isso é essencial para auditoria e reexecução determinística, pois registros com \(-\infty\) tendem a quebrar pipelines e mascarar a causa raiz (foi um vetor nulo? foi underflow? foi dado faltante?). [1]

> [!WARNING]
> Para auditoria rigorosa, a execução deve registrar **(a)** p original, **(b)** p’, **(c)** LOG_FLOOR_BITS efetivo, **(d)** raw, **(e)** score_hist. [1]

---

## 6. Seção 3 — Volatilidade Reativa P(x) (Poison Pill)

O componente reativo captura risco “agora”, isto é, eventos recentes ou incidentes com severidade que não devem ser diluídos por histórico robusto. [2]  
O modelo assume que risco “esfria” com o tempo quando não há novos eventos, formalizado por decaimento exponencial. [2]

### 6.1 Decaimento Exponencial e “Meia-vida” τ

\[
P(x) = severity \times e^{-\Delta t/\tau}
\]

- \(severity\) ∈ [0,1] representa magnitude do evento. [2]  
- \(\Delta t\) é o tempo desde o último evento relevante (mesma unidade de \(\tau\)). [2]  
- \(\tau\) é constante de tempo (meia-vida operacional) controlada por governança. [2]

**Propriedade auditável**: para severidade fixa, \(\frac{dP}{d\Delta t} < 0\), garantindo monotonicidade decrescente (perdão algorítmico). [2]

### 6.2 Semântica de severity e Δt

Para evitar arbitrariedade, severity deve derivar de uma tabela ou função versionada (por tipo de incidente), mantendo rastreabilidade do mapeamento “evento → severidade”. [2]  
Da mesma forma, Δt deve ser calculado a partir de timestamps observáveis e registrados (ex.: evento OSINT detectado, alerta de conformidade, anomalia operacional). [2]

> [!TIP]
> Auditoria recomendada: registrar “event_id”, “event_type”, “event_timestamp”, “severity_model_version”, “τ_policy_id”, “Δt_computation_unit”. [2]

---

## 7. Seção 4 — Disjuntor Sigmoide (Fusion / Sigmoid Gate)

A fusão integra estabilidade histórica e risco reativo por meio de um circuito não linear que evita duas patologias:  
1) **excesso de confiança** (histórico alto mascarando risco atual), e 2) **excesso de penalização** (evento leve destruindo score de forma irreversível). [2]

### 7.1 z = α·score_hist − β·(10·P(x))

\[
z = \alpha \cdot score_{hist} - \beta \cdot (10 \cdot P(x))
\]

- \(\alpha\) pondera a “força institucional” do histórico. [2]  
- \(\beta\) pondera a “aversão ao risco” reativo. [2]  
- O termo \(10 \cdot P(x)\) coloca o risco reativo em escala comparável a score_hist (0–10). [2]

> [!WARNING]
> \(\alpha\) e \(\beta\) são **parâmetros de política**, não “constantes universais”; devem ser versionados, assinados e auditáveis, pois governam o limiar de aprovação. [2]

### 7.2 Score_final = 10·σ(z)

\[
Score_{final} = 10 \cdot \sigma(z)
\qquad\text{onde}\qquad
\sigma(z)=\frac{1}{1+e^{-z}}
\]

A sigmoide produz saturação suave:  
- se \(z \gg 0\), \(\sigma(z)\to 1\) e Score_final aproxima 10; [2]  
- se \(z \ll 0\), \(\sigma(z)\to 0\) e Score_final aproxima 0. [2]

### 7.3 Interpretação: Limiar, Saturação e Robustez

A sigmoide age como disjuntor: pequenas variações perto do limiar podem alterar significativamente a saída, mas fora do limiar a saída satura e reduz sensibilidade a ruído marginal. [2]  
Isso é um compromisso “físico” entre estabilidade e responsividade, especialmente útil quando sinais reativos têm alta variância e podem conter falsos positivos. [2]

**Nota de auditoria (invariante)**: registrar o valor de \(z\) é obrigatório, pois ele é a “tensão do circuito” onde a decisão realmente ocorre. [2]

---

## 8. Seção 5 — Veredito e Classificação

A classificação final (conforme contexto fornecido) é:

- **BLOCK** se \(Score_{final} \le 4.5\)  
- **REVIEW** se \(4.6 \le Score_{final} \le 5.5\)  
- **ALLOW** se \(Score_{final} > 5.5\) [2]

> [!NOTE]
> Os thresholds definem a “fronteira institucional” do risco e devem ser tratados como política governada (mudanças exigem versionamento e justificativa). [2]

---

## 9. Stack & Auditability

### 9.1 Artefatos, Linhagem de Dados e Reprodutibilidade

Para auditoria forense, o sistema deve produzir um pacote mínimo de evidências:

1) **Inputs normalizados**: C, A, T, U, R, Â, severity, Δt, τ, α, β, LOG_FLOOR_BITS. [1][2]  
2) **Intermediários**: p, p’, raw, score_hist, P(x), z, Score_final. [1][2]  
3) **Provas de evidência**: hashes de documentos, IDs de eventos, timestamps, referências de coleta. [1][2]  
4) **Metadados de versão**: versão do framework (v1.2/v2.6), versões de extratores, versão de modelos NLP, versão de política. [1][2]

Sem esses itens, o sistema se torna “não reexecutável”, perdendo o caráter determinístico sob auditoria. [1][2]

### 9.2 Rationale Hash (SHA-256) e Imutabilidade

O racional textual (“por que a decisão foi tomada”) deve ser serializado de forma canônica e hasheado:

\[
rationale\_hash = SHA\text{-}256(UTF\text{-}8(rationale))
\]

A finalidade é impedir mutação pós-fato: qualquer alteração no racional muda o digest, preservando integridade de registro. [2]  
Em arquiteturas com ancoragem em lote (ex.: árvores de Merkle), o hash pode ser agregado em um commit público para prova temporal, sem expor o conteúdo do racional. [2]

> [!TIP]
> Auditoria recomendada: usar JSON canônico (ordem de chaves estável, normalização de whitespace) antes do hash, para evitar divergências não semânticas. [2]

### 9.3 Governança de Parâmetros (α, β, τ, LOG_FLOOR_BITS)

Parâmetros são “lei” do sistema — mudá-los muda o universo decisório. [2]  
Portanto, deve existir um mecanismo de governança:

- **Parameter Set ID** (ex.: `policy_2026-01`)  
- **Assinatura** (ex.: chave institucional)  
- **Change log** com justificativa, impacto e janela de validade  
- **Replays**: reexecução retroativa sob política da época, preservando determinismo histórico. [2]

---

## 10. Referência de Implementação em L11 (Audit Mode)

Abaixo, uma referência de implementação em **L11** (estilo pseudo-formal), com ênfase em rastreabilidade de intermediários. [1][2]

```l11
module FoundLab.Spezzatura.V2_6

type Float01  = float where 0.0 <= _ <= 1.0
type Score10  = float where 0.0 <= _ <= 10.0
type BitsPos  = int   where _ > 0

record HistoricVectors {
  C: Float01,
  A: Float01,
  T: Float01,
  U: Float01,
  R: Float01,
  Â: Float01
}

record ReactiveSignal {
  severity: Float01,
  delta_t: float,  // same unit as tau
  tau: float       // > 0
}

record Policy {
  alpha: float,         // >= 0
  beta: float,          // >= 0
  log_floor_bits: BitsPos
}

record Intermediates {
  p: float,
  p_floor: float,
  raw_log2: float,
  score_hist: Score10,
  P_x: Float01,
  z: float,
  score_final: Score10
}

fn clamp(x: float, lo: float, hi: float) -> float {
  if x < lo then lo else if x > hi then hi else x
}

fn log2(x: float) -> float { builtin.log(x) / builtin.log(2.0) }

fn sigmoid(z: float) -> float { 1.0 / (1.0 + builtin.exp(-z)) }

fn compute_historic(v: HistoricVectors, policy: Policy) -> (float, float, float, Score10) {
  let p  = v.C * v.A * (v.T * v.T) * v.U * v.R * v.Â
  let p_floor = max(p, pow(2.0, -policy.log_floor_bits))
  let raw = log2(p_floor)
  let score = 10.0 * (1.0 + (raw / policy.log_floor_bits))
  let score_hist = clamp(score, 0.0, 10.0) as Score10
  return (p, p_floor, raw, score_hist)
}

fn compute_reactive(r: ReactiveSignal) -> Float01 {
  // P(x) = severity * e^(-Δt/τ)
  let P = r.severity * builtin.exp(-(r.delta_t / r.tau))
  return clamp(P, 0.0, 1.0) as Float01
}

fn compute_final(score_hist: Score10, P_x: Float01, policy: Policy) -> (float, Score10) {
  let z = policy.alpha * score_hist - policy.beta * (10.0 * P_x)
  let s = 10.0 * sigmoid(z)
  let score_final = clamp(s, 0.0, 10.0) as Score10
  return (z, score_final)
}

fn spezzatura_score(v: HistoricVectors, r: ReactiveSignal, policy: Policy) -> Intermediates {
  let (p, p_floor, raw, score_hist) = compute_historic(v, policy)
  let P_x = compute_reactive(r)
  let (z, score_final) = compute_final(score_hist, P_x, policy)
  return Intermediates{ p, p_floor, raw_log2: raw, score_hist, P_x, z, score_final }
}
```

**Notas de auditoria**:  
- `compute_historic` e `compute_reactive` devem registrar parâmetros e unidades (Δt/τ). [1][2]  
- A função `log2` deve ser estável e testada contra underflow; o piso `p_floor` é obrigatório. [1]  
- A saída deve ser acompanhada de um `rationale_hash` (Seção 9.2). [2]

---

## 11. Apêndice

### A. Tabelas de Parametrização

| Parâmetro | Papel | Domínio recomendado | Observação de Governança |
|----------:|------|---------------------|---------------------------|
| LOG_FLOOR_BITS | piso anti -∞ | inteiro > 0 | define menor p representável [1] |
| α | peso do histórico | ≥ 0 | política de confiança institucional [2] |
| β | peso do risco | ≥ 0 | política de aversão ao risco [2] |
| τ | meia-vida do risco | > 0 | por categoria de evento (tabela) [2] |

### B. Casos de Teste Determinísticos

1) **Colapso por vetor nulo**: se C=0 então p=0 e p’ = 2^{-LOG_FLOOR_BITS}. [1]  
2) **Sem risco reativo**: se P(x)=0 então z = α·score_hist e Score_final = 10·σ(α·score_hist). [2]  
3) **Risco máximo recente**: se severity=1 e Δt=0 então P(x)=1, logo z = α·score_hist − β·10. [2]  
4) **Perdão temporal**: aumentar Δt reduz P(x) monotonicamente para 0. [2]

### C. Glossário

- **Determinístico**: mesma entrada produz mesma saída, sem estado oculto. [1][2]  
- **Entropia (informacional)**: medida de incerteza; mais evidência reduz o espaço de hipóteses. [1]  
- **Invariante**: propriedade que deve permanecer verdadeira em todas as execuções válidas. [1][2]  
- **Disjuntor (circuit breaker)**: mecanismo que desliga/satura a saída quando risco excede um limiar. [2]  
- **Poison Pill**: componente reativo que impede que histórico “mascare” incidentes atuais. [2]

---

## Referências

- **[1]** Spezzatura Framework v1.2/v2.6 — Fórmula multiplicativa, compressão log₂ e normalização (contexto fornecido).  
- **[2]** TrustScore API / FoundLab — Componente reativo P(x), Sigmoid Gate (fusão), classificação BLOCK/REVIEW/ALLOW e trilha SHA-256 (contexto fornecido).
