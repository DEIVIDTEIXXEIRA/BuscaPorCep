package main

import (
	"fmt"
	"os"
	"buscaporcep/controllers"
)

func main() {
	for _, cep := range os.Args[1:] {
		dados, erro := controllers.ConsultarViaCEP(cep)
		if erro != nil {
			fmt.Fprintf(os.Stderr, "Erro ao consultar o CEP %s: %v\n", cep, erro)
			continue
		}

		erro = controllers.EscreverArquivo(dados, "Informacoes_do_local.txt")
		if erro != nil {
			fmt.Fprintf(os.Stderr, "Erro ao escrever os dados do CEP %s no arquivo: %v\n", cep, erro)
		}
	}
}
