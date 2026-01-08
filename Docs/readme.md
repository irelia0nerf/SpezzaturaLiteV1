<!-- README.md -->

<div align="center">
  <img
    src="https://placehold.co/1200x300/1e293b/ffffff?text=CRESCIMENTO+E+DECAIMENTO+EXPONENCIAL"
    alt="Crescimento e Decaimento Exponencial"
    width="100%"
  />
</div>

<div align="center">

# 📈 Entendendo Crescimento e Decaimento Exponencial

![status](https://img.shields.io/badge/status-stable-16a34a?style=flat-square)
![versão](https://img.shields.io/badge/version-1.0.0-2563eb?style=flat-square)
![licença](https://img.shields.io/badge/license-CC--BY--4.0-7c3aed?style=flat-square)

</div>

> [!NOTE]
> Este material explica **mudanças rápidas por multiplicação repetida** (crescimento e decaimento exponencial), com fórmulas, exemplos e um comparativo final.

---

## 📚 Sumário

- [✨ Introdução](#-introdução)
- [🚀 Crescimento Exponencial](#-crescimento-exponencial)
  - [Fórmula](#fórmula)
  - [Variáveis](#variáveis)
  - [Dica do Professor](#dica-do-professor)
- [📉 Decaimento Exponencial](#-decaimento-exponencial)
  - [Fórmula](#fórmula-1)
  - [Variáveis](#variáveis-1)
  - [Dica do Professor](#dica-do-professor-1)
  - [Exemplo Prático: Depreciação de um Computador](#exemplo-prático-depreciação-de-um-computador)
- [🆚 Resumo e Principais Diferenças](#-resumo-e-principais-diferenças)
- [🧠 Como Identificar no Enunciado](#-como-identificar-no-enunciado)
- [🧪 Exercícios Rápidos](#-exercícios-rápidos)
- [✅ Checklist de Estudo](#-checklist-de-estudo)
- [🏁 Conclusão](#-conclusão)

---

## ✨ Introdução

Bem-vindo ao fascinante mundo do **crescimento** e **decaimento exponencial**!  
Esses conceitos descrevem situações em que uma quantidade muda ao longo do tempo por **multiplicação repetida**, em vez de aumentar/diminuir por soma/subtração constante (como no crescimento linear).

Você vê isso no dia a dia em fenômenos como:

- 📉 **Depreciação** de eletrônicos e carros
- 🦠 **Crescimento** de populações (bactérias, usuários de um app)
- 💰 **Juros compostos**
- ⚛️ **Decaimento radioativo**

> [!TIP]
> Se a frase do problema diz “**aumenta x% ao ano**” ou “**cai x% a cada mês**”, acende o alerta: **exponencial**.

---

## 🚀 Crescimento Exponencial

### O que é?

O crescimento exponencial acontece quando uma quantidade **aumenta por um fator constante** a cada intervalo de tempo.  
Em vez de “somar sempre a mesma coisa”, você **multiplica sempre pelo mesmo fator** (maior que 1).

### Fórmula

\[
y = a(1 + r)^x
\]

### Variáveis

| Variável | O que significa                                                                 |
|---------:|----------------------------------------------------------------------------------|
| `y`      | Valor final (após o crescimento)                                                |
| `a`      | Valor inicial (ponto de partida)                                                |
| `r`      | Taxa de crescimento (em **decimal**)                                            |
| `x`      | Número de intervalos de tempo (anos, meses, dias etc.)                           |

O termo **(1 + r)** é o **fator de crescimento**.

#### Dica do Professor

> [!TIP]
> Converta porcentagem para decimal:  
> `25% = 0,25` • `8% = 0,08` • `120% = 1,20`

---

## 📉 Decaimento Exponencial

### O que é?

O decaimento exponencial segue a mesma lógica, mas agora o fator multiplicativo é **menor que 1**.  
Isso provoca uma queda consistente ao longo do tempo — como uma “divisão repetida”.

### Fórmula

\[
y = a(1 - r)^x
\]

### Variáveis

| Variável | O que significa                                                                 |
|---------:|----------------------------------------------------------------------------------|
| `y`      | Valor final (após diminuir)                                                     |
| `a`      | Valor inicial                                                                   |
| `r`      | Taxa de decaimento (em **decimal**)                                             |
| `x`      | Número de intervalos de tempo                                                   |

O termo **(1 − r)** é o **fator de decaimento** (fica entre 0 e 1 quando `0 < r < 1`).

#### Dica do Professor

> [!WARNING]
> Se a taxa for `r = 1` (100%), o valor vai para **zero** em 1 passo:  
> `y = a(1 - 1)^x = a(0)^x`.

---

### Exemplo Prático: Depreciação de um Computador

Você compra um computador por **$1.600** e ele perde **50% do valor a cada ano**.

- Taxa de decaimento: `r = 0,5`
- Fator de decaimento: `(1 - r) = 0,5 = 1/2`

Equação:

\[
y = 1600\left(\frac{1}{2}\right)^x
\]

Evolução ano a ano:

1. Após 1 ano: `1600 × 1/2 = 800`
2. Após 2 anos: `800 × 1/2 = 400`
3. Após 3 anos: `400 × 1/2 = 200`
4. Após 4 anos: `200 × 1/2 = 100`

> [!NOTE]
> Perceba o padrão: **cada ano multiplica por 1/2** — não é uma subtração fixa.

---

## 🆚 Resumo e Principais Diferenças

| Aspecto                | Crescimento Exponencial                     | Decaimento Exponencial                        |
|------------------------|----------------------------------------------|-----------------------------------------------|
| Fórmula                | `y = a(1 + r)^x`                             | `y = a(1 - r)^x`                              |
| Fator de mudança       | `(1 + r) > 1`                                | `0 < (1 - r) < 1`                             |
| Direção típica do gráfico | Curva **sobe** da esquerda para a direita | Curva **desce** da esquerda para a direita   |
| Linguagem comum        | “aumenta x% por período”                     | “diminui x% por período”                      |

---

## 🧭 Como Identificar no Enunciado

Use este mini-guia:

- ✅ “**Dobra** a cada período” → multiplicar por `2`
- ✅ “Cresce **20%** ao ano” → fator `(1 + 0,20) = 1,20`
- ✅ “Cai **15%** ao mês” → fator `(1 - 0,15) = 0,85`
- ✅ “Metade a cada período” → multiplicar por `1/2`
- ⚠️ “Diminui **10 unidades** ao mês” → isso é **linear**, não exponencial

---

## 🧪 Exercícios Rápidos

1) Uma população começa com `a = 500` e cresce `10%` ao ano. Qual a expressão?  
**Resposta esperada:** `y = 500(1,10)^x`

2) Um produto custa `a = 200` e sofre desconto de `30%` ao mês. Quanto vale após 3 meses?  
**Resposta esperada:** `y = 200(0,70)^3`

3) Um valor dobra a cada mês. Qual é o fator e a expressão?  
**Resposta esperada:** fator `2`, então `y = a(2)^x`

> [!TIP]
> Se quiser conferir rapidamente valores, faça uma tabelinha com `x = 0, 1, 2, 3...` e observe o padrão.

---

## ✅ Checklist de Estudo

- [ ] Sei converter `%` para decimal (`r`)
- [ ] Sei escolher entre `(1 + r)` e `(1 - r)`
- [ ] Sei identificar quando o problema é linear (soma constante) vs exponencial (multiplicação constante)
- [ ] Consigo montar a fórmula e interpretar `a`, `r` e `x`
- [ ] Consigo explicar o exemplo de depreciação do computador

---

## 🏁 Conclusão

Você agora entende o “código” por trás de mudanças rápidas: **multiplicação repetida**.  
Com as fórmulas de crescimento e decaimento exponencial, você consegue modelar situações reais como depreciação, juros compostos, crescimento populacional e muito mais.

> [!TIP]
> Quanto mais você praticar lendo enunciados e montando o fator correto, mais “automático” isso fica.

---

### 📌 Diagrama mental (Mermaid)

```mermaid
flowchart TB
  A["Enunciado"] --> B{"Muda por soma/subtração fixa?"}
  B -- "Sim" --> C["Linear (não exponencial)"]
  B -- "Não" --> D{"Muda por porcentagem/fator?"}
  D -- "Aumenta" --> E["Crescimento: y = a(1+r)^x"]
  D -- "Diminui" --> F["Decaimento: y = a(1-r)^x"]
