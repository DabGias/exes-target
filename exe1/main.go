package main

import (
	"encoding/json/v2"
	"fmt"
	"io"
	"os"
)

// Considerando que o json abaixo tem registros de vendas de um time comercial, faça
// um programa que leia os dados e calcule a comissão de cada vendedor, seguindo a
// seguinte regra para cada venda:
// - Vendas abaixo de R$100,00 não gera comissão
// - Vendas abaixo de R$500,00 gera 1% de comissão
// - A partir de R$500,00 gera 5% de comissão

type Venda struct {
    Vendedor string `json:"vendedor"`
    Valor float64 `json:"valor"`
}

type Dados struct {
    Vendas []Venda `json:"vendas"`
}

func extrairDadosDoArquivo(arquivo string) ([]byte, error) {
    arq, err := os.Open(arquivo)

    if err != nil {
        return []byte{}, err
    }

    defer arq.Close()
    
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
    
    dados := Dados{}
    
    err = json.Unmarshal(d, &dados)

    if err != nil {
        fmt.Println(err)
        return
    }

    comissoes := map[string]float64{}

    const VALOR_MINIMO float64 = 100.00
    const TAXA_MINIMA float64 = 0.01
    
    const VALOR_NIVEL_1 float64 = 500.00
    const TAXA_NIVEL_1 float64 = 0.05

    for _, venda := range dados.Vendas {
        vendedor := venda.Vendedor
        valor := venda.Valor
        
        valorComissao := 0.0
        taxa := 0.0
        
        if valor > VALOR_MINIMO && valor < VALOR_NIVEL_1 {
            taxa = TAXA_MINIMA
        } else if valor > VALOR_NIVEL_1 {
            taxa = TAXA_NIVEL_1
        }

        valorComissao += (valor * taxa)

        if comissao, ok := comissoes[vendedor]; ok {
            comissao += valorComissao
        } else {
            comissoes[vendedor] = valorComissao
        }
    }

    for vendedor, comissao := range comissoes {
        fmt.Printf("%s: %.2f\n", vendedor, comissao)
    }
}