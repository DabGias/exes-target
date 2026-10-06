package main

import (
	"fmt"
	"time"
)

// Faça um programa que a partir de um valor e de uma data de vencimento, calcule o
// valor dos juros na data de hoje considerando que a multa seja de 2,5% ao dia.

func main() {
    var valor float64
    var data string

    fmt.Print("Digite o valor: ")
    fmt.Scanln(&valor)

    for {
        if valor <= 0 {
            fmt.Println("Valor inválido! O valor deve ser maior que 0!")
            fmt.Print("Digite o valor: ")
            fmt.Scanln(&valor)
        } else {
            break
        }
    }

    fmt.Print("Digite a data de vencimento [aaaa-mm-dd]: ")
    fmt.Scanln(&data)

    hoje := time.Now()
    dataVencimento, err := time.Parse(time.DateOnly, data)

    for {
        if err != nil {
            fmt.Println("Data inválida!")
            fmt.Print("Digite a data de vencimento [aaaa-mm-dd]: ")
            fmt.Scanln(&data)

            dataVencimento, err = time.Parse(time.DateOnly, data)
        } else if hoje.Sub(dataVencimento) < 0 {
            fmt.Println("Data inválida!")
            fmt.Print("Digite a data de vencimento [aaaa-mm-dd]: ")
            fmt.Scanln(&data)

            dataVencimento, err = time.Parse(time.DateOnly, data)
        } else {
            break
        }
    }

    diasPosVencimento := int(hoje.Sub(dataVencimento) / 86400000000000)
    juros := (valor * 0.025) * float64(diasPosVencimento)

    fmt.Printf("Juros: %.2f\n", juros)
}