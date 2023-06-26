package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Viacep struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	DDD         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func main() {
	for _, cep := range os.Args[1:] {
		requisicao, erro := http.Get("http://viacep.com.br/ws/" + cep + "/json/")
		if erro != nil {
			fmt.Fprintf(os.Stderr, "Erro ao fazer a requisição, %v\n", erro)
		}
		defer requisicao.Body.Close()

		resposta, erro := io.ReadAll(requisicao.Body)
		if erro != nil {
			fmt.Fprintf(os.Stderr, "Erro ao ler a resposta, %s\n", erro)
		}

		var dados Viacep
		erro = json.Unmarshal(resposta, &dados)
		if erro != nil {
			fmt.Fprintf(os.Stderr, "Erro ao converter a resposta, %s\n", erro)
		}
		fmt.Println(dados)

		file, erro := os.Create("Informações do local.txt")
		if erro != nil {
			fmt.Fprintf(os.Stderr, "Erro ao criar o arquivo %s\n", erro)
		}
		defer file.Close()

		_, erro = file.WriteString(fmt.Sprintf("Aqui estão os dados do endereço referente ao CEP %s:\n"+
			"Endereço:%s %s\nCidade:%s\nUF:%s", dados.Cep, dados.Logradouro,
			dados.Complemento, dados.Localidade, dados.Uf))
	}
}
