# Exercícios do desafio técnico da Target

As soluçoes desse repositório foram desenvolvidas em Go 1.27, mas podem ser executadas em qualquer versão da linguagem.

Para executar as soluções basta executar o comando:

```bash
cd 'pasta do exercício'
go run main.go
```

## Exe 1

Considerando que o json abaixo tem registros de vendas de um time comercial, faça
um programa que leia os dados e calcule a comissão de cada vendedor, seguindo a
seguinte regra para cada venda:
- Vendas abaixo de R$100,00 não gera comissão
- Vendas abaixo de R$500,00 gera 1% de comissão
- A partir de R$500,00 gera 5% de comissão

## Exe 2

Faça um programa onde eu possa lançar movimentações de estoque dos produtos
que estão no json abaixo, dando entrada ou saída da mercadoria no meu depósito,
onde cada movimentação deve ter:
- Um número identificador único.
- Uma descrição para identificar o tipo da movimentação realizada
E que ao final da movimentação me retorne a qtde final do estoque do produto
movimentado.

## Exe 3

Faça um programa que a partir de um valor e de uma data de vencimento, calcule o
valor dos juros na data de hoje considerando que a multa seja de 2,5% ao dia.
