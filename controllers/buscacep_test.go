package controllers_test

import (
	"buscaporcep/controllers"
	"buscaporcep/modelos"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestConsultarViaCEP(t *testing.T) {
	// Servidor HTTP falso local para simular a resposta da ViaCEP
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Cabeçalho Content-Type para JSON
		w.Header().Set("Content-Type", "application/json")

		// Uma resposta simulada em JSON
		respostaJSON := `{
			"cep": "01001-000",
			"logradouro": "Praça da Sé",
			"localidade": "São Paulo",
			"uf": "SP"
		}`

		// A resposta no corpo da resposta HTTP
		io.WriteString(w, respostaJSON)
	}))
	defer server.Close()

	// Função que esta sendo testada
	resultado, erro := controllers.ConsultarViaCEP("01001-000")
	if erro != nil {
		t.Errorf("Erro inesperado durante a consulta: %v", erro)
		return
	}

	if resultado.Cep != "01001-000" {
		t.Errorf("CEP incorreto. Esperado: 01001-000, Obtido: %s", resultado.Cep)
	}

}

func TestEscreverArquivo(t *testing.T) {
	// Cria um diretório temporário para o teste
	diretorioTemporario, erro := ioutil.TempDir("", "test")
	if erro != nil {
		t.Fatalf("Erro ao criar diretório temporário: %v", erro)
	}
	defer os.RemoveAll(diretorioTemporario)

	// Dados de exemplo
	dados := &modelos.Viacep{
		Cep:         "01001-000",
		Logradouro:  "Praça da Sé",
		Complemento: "",
		Localidade:  "São Paulo",
		Uf:          "SP",
	}

	// Caminho completo do arquivo temporário
	caminhoArquivo := diretorioTemporario + "/dados.txt"

	// Função a ser testada
	erro = controllers.EscreverArquivo(dados, caminhoArquivo)
	if erro != nil {
		t.Errorf("Erro inesperado ao escrever o arquivo: %v", erro)
		return
	}

	// Le o conteúdo do arquivo
	conteudo, erro := ioutil.ReadFile(caminhoArquivo)
	if erro != nil {
		t.Errorf("Erro ao ler o arquivo: %v", erro)
		return
	}

	// Verificar se o conteúdo do arquivo é o esperado
	conteudoEsperado := "Aqui estão os dados do endereço referente ao CEP 01001-000:\n" +
		"Endereço:Praça da Sé \nCidade:São Paulo\nUF:SP"

	if string(conteudo) != conteudoEsperado {
		t.Errorf("Conteúdo do arquivo não corresponde ao esperado.\nEsperado:\n%s\nObtido:\n%s", conteudoEsperado, string(conteudo))
	}
}

func TestCriarArquivo(t *testing.T) {
	// Nome do arquivo de teste
	nomeArquivo := "arquivo_de_teste.txt"

	//função a ser testada
	arquivo, erro := controllers.CriarArquivo(nomeArquivo)
	if erro != nil {
		t.Errorf("Erro inesperado ao criar o arquivo: %v", erro)
		return
	}
	defer arquivo.Close()

	// Verifica se o arquivo foi criado
	if _, err := os.Stat(nomeArquivo); os.IsNotExist(err) {
		t.Errorf("O arquivo não foi criado.")
	}

	// Feche o arquivo explicitamente
	if erro := arquivo.Close(); erro != nil {
		t.Errorf("Erro ao fechar o arquivo: %v", erro)
		return
	}

	// Apaga o arquivo de teste após o teste
	if err := os.Remove(nomeArquivo); err != nil {
		t.Errorf("Erro ao excluir o arquivo de teste: %v", err)
	}
}
