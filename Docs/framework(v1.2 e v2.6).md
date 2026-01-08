### 1. O Framework Spezzatura (v1.2 e v2.6)
Este framework é a base para quantificar o **Risco de Admissão** (*Intake Risk*) de startups através de um modelo estritamente **multiplicativo**. A escolha de uma operação multiplicativa é estratégica: se qualquer vetor crítico de integridade for zero, a pontuação final colapsa a zero, funcionando como um **"fusível de admissão"** que impede que métricas de vaidade mascarem falhas graves de conformidade.

*   **Vetor de Completude (C):** Razão entre artefatos esperados e fornecidos.
*   **Vetor de Atividade (A):** Mede o "batimento digital" via APIs (ex: commits no GitHub).
*   **Vetor de Confiança ($T^2$):** Aplica uma **Lei de Potência** onde sinais de entidades de elite (como Y Combinator) são elevados ao quadrado, conferindo peso exponencial à legitimidade institucional.
*   **A Fórmula Base:** O produto bruto é definido como $p = C \times A \times (T^2) \times U \times R \times \hat{A}$.

### 2. Compressão Logarítmica ($\log_2$)
Para gerenciar a **variância massiva** entre métricas de startups (como receitas que variam de US$ 10 mil a US$ 10 milhões), o framework aplica um logaritmo de base 2.
*   **Normalização:** Transforma dados brutos caóticos em uma escala linear de 0 a 10.
*   **Anti-Gaming:** O crescimento logarítmico é mais lento à medida que os valores aumentam, o que impede que uma empresa infle sua nota artificialmente ao exagerar um único indicador.
*   **Mensuração de Crescimento:** Permite medir o progresso em "dobras" (*doublings*), refletindo com mais fidelidade a dinâmica de escala de empresas de tecnologia.

### 3. Algoritmos de Similaridade e Detecção de Fraude (SimHash)
O vetor de **Unicidade (U)** utiliza o algoritmo **SimHash** (um tipo de *Locality Sensitive Hashing*) para detectar fraudes sintéticas e reciclagem de modelos de negócio.
*   **Fingerprint Semântico:** Ao contrário de hashes tradicionais, o SimHash captura o conteúdo semântico, permitindo calcular a **distância de Hamming** entre documentos (como *pitch decks*).
*   **Identificação de Clones:** Se um material for identificado como uma "quase-duplicata" de modelos genéricos ou de golpes conhecidos, o score de unicidade cai drasticamente.

### 4. Componente Reativo e Decaimento Exponencial [$P(x)$]
Diferente da base histórica, o componente reativo captura a volatilidade imediata e sinais de alerta (*red flags*).
*   **Fórmula de Resfriamento:** Utiliza decaimento exponencial ($P(x) = severity \times e^{-\Delta t/\tau}$) para garantir que incidentes antigos percam peso gradualmente ("perdão algorítmico"), enquanto crises atuais mantêm a nota "quente" e arriscada.
*   **Booster de Severidade:** Se um alerta OSINT (Inteligência de Fontes Abertas) atinge uma severidade crítica ($> 0,8$), um multiplicador de 1,5x é acionado para derrubar o score instantaneamente.

### 5. A Porta Sigmoide (O Disjuntor Matemático)
A fusão final entre o histórico e o risco reativo é governada por uma **função Sigmoide**.
*   **Comportamento Não Linear:** Ela atua como um **disjuntor ou "Poison Pill"**. Quando o risco detectado ($P$) cruza um limite crítico, a Sigmoide "descarrega", forçando o multiplicador da nota histórica para zero.
*   **Regra de Um Golpe:** Essa lógica garante que uma única fraude confirmada anule um passado de sucesso, movendo a startup automaticamente para o veredito **BLOCK**.

### 6. Integridade e Imutabilidade (SHA-256 e Árvores de Merkle)
Para garantir que as decisões sejam à prova de falsificação, o sistema utiliza algoritmos de hashing criptográfico:
*   **Rationale Hash (SHA-256):** Gera um fingerprint digital imutável da justificativa da decisão, assegurando que o motivo de um veredito não possa ser alterado posteriormente.
*   **Agregação de Merkle:** Milhares de pontuações são agrupadas em **Árvores de Merkle**, onde apenas a raiz é ancorada em uma blockchain (ex: Polygon), permitindo auditoria de baixo custo e alta eficiência.

### 7. Algoritmos de Machine Learning no Setor Financeiro
Além do framework determinístico Spezzatura, as fontes mencionam o uso de algoritmos clássicos de ML e Deep Learning para gestão de risco de crédito e análise de sentimento:
*   **Naive Bayes e SVM:** Utilizados para classificação de risco e probabilidade de inadimplência baseada em frequência de palavras e hiperplanos de separação.
*   **RNNs e Transformers (BERT):** Aplicados em análises de sentimento para processar sequências temporais de dados e capturar nuances textuais em notícias e redes sociais.

Como analogia, esses algoritmos funcionam como a **Escala Richter** (via logaritmos) integrada a um **sensor de fumaça inteligente** (via Sigmoide e P(x)). Enquanto a escala Richter mantém números gigantescos de energia em uma faixa compreensível (0-10), o sensor de fumaça garante que, ao menor sinal de "incêndio" (fraude), a energia do sistema seja cortada instantaneamente, não importando a solidez da construção histórica do prédio.
