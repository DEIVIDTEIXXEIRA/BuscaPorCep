# Busca por CEP

Busca por CEP, consome uma API externa fazendo uma pesquisa de endereço baseado no CEP fornecido.

# Descrição
O projeto "Busca por CEP" é um aplicativo de linha de comando Go que recupera detalhes de endereço com base em códigos postais usando a API ViaCEP. Ele permite que os usuários busquem e exibam informações de endereço precisas, como endereço, cidade e estado, consultando a API. 

O projeto demonstra como interagir com APIs, lidar com respostas JSON, testes e executar operações básicas de arquivo em Go. Serve como ponto de partida para o desenvolvimento de aplicações que envolvam funcionalidades de pesquisa de endereço e geolocalização.

# Utilização
 - Certifique-se de ter [Go](https://golang.org/)instalado em seu sistema.

- Clone o repositório do projeto em sua máquina local ou baixe os arquivos de código-fonte.

```shell
$ git clone git@github.com:DEIVIDTEIXXEIRA/BuscaPorCep.git 
```

- Abra um terminal ou prompt de comando e navegue até o diretório do projeto.

- Compile o projeto executando o seguinte comando:

```shell
$ go build main.go
```

- Agora você pode executar. Use o seguinte formato de comando, fornecendo um ou mais CEPs como argumentos:

```shell
$ ./main “cep” “cep” “cep”
```

- Além da resposta na linha de comando, será gerado um arquivo “txt” com as informações referente aos CEPs


# Contato
Para qualquer dúvida ou sugestão, sinta-se à vontade para entrar em contato comigo:
- **E-mail**: deividteixeira.go@gmail.com
- **GitHub**: https://github.com/DEIVIDTEIXXEIRA
- **Instagram**: @deivid_tx
