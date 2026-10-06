package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
)

// Faça um programa onde eu possa lançar movimentações de estoque dos produtos
// que estão no json abaixo, dando entrada ou saída da mercadoria no meu depósito,
// onde cada movimentação deve ter:
// - Um número identificador único.
// - Uma descrição para identificar o tipo da movimentação realizada
// E que ao final da movimentação me retorne a qtde final do estoque do produto
// movimentado.

type Produto struct {
    CodigoProduto int `json:"codigoProduto"`
    DescricaoProduto string `json:"descricaoProduto"`
    Quantidade int `json:"estoque"`
}

type DadosEstoque struct {
    Produtos []Produto `json:"estoque"`
}

func extrairDadosDoArquivo(arquivo string) ([]byte, error) {
    arq, err := os.Open(arquivo)

    if err != nil {
        return []byte{}, err
    }

    dados, err := io.ReadAll(arq)

    if err != nil {
        return []byte{}, err
    }

    return dados, nil
}

func main() {
    d, err := extrairDadosDoArquivo("./dados.json")

    if err != nil {
        fmt.Println(err)
        return
    }

    de := DadosEstoque{}

    err = json.Unmarshal(d, &de)

    if err != nil {
        fmt.Println(err)
        return
    }

    estoque := map[int]*Produto{}
    
    for _, produto := range de.Produtos {
        estoque[produto.CodigoProduto] = &produto
    }

    for {
        var tipoMovimentacao string
        var codigoProduto int
        var quantidadeProduto int

        fmt.Println("=== Para sair digite 'q' em 'Tipo de Movimentação' ===")
        fmt.Print("Tipo de Movimentação (Entrada/Saída): ")
        fmt.Scanln(&tipoMovimentacao)

        tipoMovimentacao = strings.ToLower(strings.Trim(tipoMovimentacao, " "))

        for {
            if tipoMovimentacao != "entrada" && tipoMovimentacao != "saída" && tipoMovimentacao != "q" {
                fmt.Printf("Tipo de movimentação '%s' inválido!\n", tipoMovimentacao)
                fmt.Print("Digite novamente o tipo de movimentação (Entrada/Saída): ")
                fmt.Scanln(&tipoMovimentacao)

                tipoMovimentacao = strings.ToLower(strings.Trim(tipoMovimentacao, " "))
            } else {
                break
            }
        }

        if tipoMovimentacao == "q" {
            break
        }

        fmt.Print("Código do Produto: ")
        fmt.Scanln(&codigoProduto)

        for {
            if _, ok := estoque[codigoProduto]; !ok {
                fmt.Printf("Código %d inválido!\n", codigoProduto)
                fmt.Print("Digite o código do produto novamente: ")
                fmt.Scanln(&codigoProduto)
            } else {
                break
            }
        }

        fmt.Print("Quantidade: ")
        fmt.Scanln(&quantidadeProduto)

        for {
            if quantidadeProduto <= 0 {
                fmt.Println("Quantidade inválida! A quantidade deve ser maior que 0!")
                fmt.Print("Digite a quantidade novamente: ")
                fmt.Scanln(&quantidadeProduto)
            } else {
                break
            }
        }

        produto := estoque[codigoProduto]

        if tipoMovimentacao == "entrada" {
            produto.Quantidade += quantidadeProduto
        } else {
            produto.Quantidade -= quantidadeProduto
        }

        fmt.Printf("Tipo de Movimentação: %s\n", tipoMovimentacao)
        fmt.Println("Produto Alterado: ")
        fmt.Printf("    - Código: %d\n", produto.CodigoProduto)
        fmt.Printf("    - Descrição: %s\n", produto.DescricaoProduto)
        fmt.Printf("    - Quantidade em estoque: %d\n", produto.Quantidade)
    }
}
