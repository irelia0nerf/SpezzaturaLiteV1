# Resumo Técnico: Arquitetura de Privacidade da Spezzatura Lite

## 1. O Dilema: Dados Tóxicos no Ecossistema de Venture Capital

Aceleradoras e fundos enfrentam o paradoxo **Volume vs. Rigor**. O processamento de mais de 1.000 inscrições por ciclo, frente às 118 horas necessárias para uma *due diligence* completa por startup, força a coleta massiva de **Informações Pessoais Identificáveis (PII)**. O armazenamento dessas PII de startups rejeitadas cria o inevitável risco de **"dados tóxicos"**.

Este *Paradoxo da Retenção* gera dois riscos primários:

*   **Risco Legal/Financeiro:** Violações de LGPD ou GDPR resultam em multas que ameaçam a continuidade do negócio.
*   **Dano Reputacional:** A quebra de confiança no mercado, crucial para atrair empreendedores de ponta.

A API Spezzatura Lite resolve esta crise eliminando a necessidade de retenção de dados, agindo como uma solução cirúrgica para o problema na origem.

## 2. A Solução: Motor de Verificação *Stateless*

A Spezzatura Lite é um **Motor de Verificação Sem Estado** (*Stateless Verification Engine*) fundamentado na filosofia de **"Confiança pela Física"** (*Trust by Physics*). A confiança é um resultado determinístico da arquitetura, não apenas uma promessa.

A fundação de segurança é sustentada por dois pilares:

1.  **Design *Stateless* (Sem Estado)**
2.  **Mecanismo de *Crypto-Shredding* (Destruição Criptográfica)**

Juntos, eles garantem a obtenção de *insights* sem armazenamento permanente, neutralizando os "dados tóxicos".

## 3. Pilar 1: Arquitetura *Stateless* (*Compute-and-Forget*)

A arquitetura sem estado opera pelo princípio **"Processar e Esquecer"** (*Compute-and-Forget*). Cada transação é um evento isolado. Para cada análise de risco, uma infraestrutura segura é criada e, imediatamente após o uso, é completamente destruída.

### Fluxo de Requisição e Destruição de Infraestrutura

| Passo | Ação com os Dados | Ação com a Infraestrutura |
| :---: | :---: | :---: |
| **1. Requisição** | Recebimento dos dados para processamento. | Criação instantânea de contêiner isolado (ex: **AWS Fargate** / **Google Cloud Run** com **gVisor sandboxes**). |
| **2. Processamento** | Análise e cálculo do *score* ocorrem **inteiramente na memória RAM**. | O contêiner executa a lógica de negócio em um ambiente efêmero. |
| **3. Resposta** | O resultado (*score*) é enviado ao usuário. | N/A |
| **4. Finalização** | **Nenhum dado é gravado em disco** de forma permanente. | O contêiner e toda a memória RAM utilizada são **completamente destruídos**. |

## 4. Pilar 2: *Crypto-Shredding* (Irrecuperabilidade por Design)

O *Crypto-Shredding* é a garantia final de irrecuperabilidade. A ideia é simples: em vez de apagar dados (um processo incerto), **apaga-se a chave criptográfica que os torna legíveis.**

O processo rigoroso:

1.  **Criptografia Imediata:** Dados recebidos são criptografados com uma chave simétrica (`K_sym`) única e temporária.
2.  **Processamento Seguro:** Um Hardware Security Module (HSM) desembrulha `K_sym`. Os dados são **descriptografados momentaneamente na memória RAM isolada** do contêiner para análise.
3.  **Destruição da Chave:** Após o envio da resposta, um comando é emitido para **deletar permanentemente a chave** (`K_sym`).

Sem a chave, os dados criptografados persistentes tornam-se **matematicamente irrecuperáveis**, equivalendo a ruído digital aleatório.

## 5. Conformidade LGPD/GDPR: O Direito ao Esquecimento Automático

A arquitetura implementa o **"Direito ao Esquecimento"** (*Right to Erasure*) por *design*, tornando a conformidade o **comportamento padrão** do sistema:

*   A natureza **stateless** e o **Crypto-Shredding** garantem o cumprimento automático dos princípios de **minimização de dados** e **limitação de armazenamento**.
*   A destruição da chave é uma ação **instantânea e verificável**, garantindo que nenhuma PII seja retida além do tempo estritamente necessário para a transação.

## 6. Conclusão: Higiene Algorítmica por Design

A abordagem Spezzatura Lite redefine a gestão de dados na *due diligence*. Ao projetar o sistema para minimizar a exposição e maximizar a privacidade, a conformidade legal torna-se uma característica intrínseca e automática.

Isso é a **higiene algorítmica** de fato: a maneira mais fácil de operar o sistema é também a mais segura e compatível.

> **Premissa Estratégica:** Não ajudamos a encontrar a agulha no palheiro; garantimos que a agulha não esteja enferrujada.
