package controllers

import (
	"buscaporcep/modelos"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

// consultarViaCEP faz uma requisição para a api retornando as informações referente ao cep
func ConsultarViaCEP(cep string) (*modelos.Viacep, error) {
	requisicao, erro := http.Get("http://viacep.com.br/ws/" + cep + "/json/")
	if erro != nil {
		return nil, erro
	}
	defer requisicao.Body.Close()

	resposta, erro := io.ReadAll(requisicao.Body)
	if erro != nil {
		return nil, erro
	}

	var dados modelos.Viacep
	erro = json.Unmarshal(resposta, &dados)
	if erro != nil {
		return nil, erro
	}

	return &dados, nil
}

//EscreverArquivo escreve as informações recebidas 
func EscreverArquivo(dados *modelos.Viacep, nomeArquivo string) error {
	arquivo, erro := CriarArquivo(nomeArquivo)
	if erro != nil {
		return erro
	}
	defer arquivo.Close()

	_, erro = arquivo.WriteString(fmt.Sprintf("Aqui estão os dados do endereço referente ao CEP %s:\n"+
		"Endereço:%s %s\nCidade:%s\nUF:%s", dados.Cep, dados.Logradouro,
		dados.Complemento, dados.Localidade, dados.Uf))
	if erro != nil {
		return erro
	}

	return nil
}

//CriarArquivo cria o arquivo com as informações recebidas
func CriarArquivo(nomeArquivo string) (*os.File, error) {
	arquivo, erro := os.Create(nomeArquivo)
	if erro != nil {
		return nil, erro
	}
	return arquivo, nil
}

