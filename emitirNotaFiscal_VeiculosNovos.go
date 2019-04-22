package main

import (
	"fmt"
	"strings"
	"net/http"
	"io/ioutil"
)

func main() {

	url := "https://webmaniabr.com/api/1/nfe/emissao/"

	payload := strings.NewReader("{\n  \"ID\": 1137,\n  \"url_notificacao\": \"http://meudominio.com/retorno.php\",\n  \"operacao\": 1,\n  \"natureza_operacao\": \"Venda de produção do estabelecimento\",\n  \"modelo\": 1,\n  \"finalidade\": 1,\n  \"ambiente\": 2,\n  \"cliente\": {\n     \"cpf\": \"000.000.000-00\",\n     \"nome_completo\": \"Nome do Cliente\",\n     \"endereco\": \"Av. Brg. Faria Lima\",\n     \"complemento\": \"Escritório\",\n     \"numero\": 1000,\n     \"bairro\": \"Itaim Bibi\",\n     \"cidade\": \"São Paulo\",\n     \"uf\": \"SP\",\n     \"cep\": \"00000-000\",\n     \"telefone\": \"(00) 0000-0000\",\n     \"email\": \"nome@email.com\"\n  },\n  \"produtos\": [\n     {\n        \"nome\": \"Nome do produto\",\n        \"codigo\": \"nome-do-produto\",\n        \"ncm\": \"6109.10.00\",\n        \"cest\": \"28.038.00\",\n        \"quantidade\": 3,\n        \"unidade\": \"UN\",\n        \"peso\": \"0.800\",\n        \"origem\": 0,\n        \"subtotal\": \"44.90\",\n        \"total\": \"134.70\",\n        \"classe_imposto\": \"REF2892\",\n        \"veiculos_novos\": {\n           \"tipo_operacao\": \"1\",\n           \"chassi\": \"00000000000000000\",\n           \"cor\": \"A1\",\n           \"cor_descricao\": \"Descrição da Cor\",\n           \"cv\": \"100\",\n           \"cilindrada\": \"1000\",\n           \"peso_liquido\": \"2.5000\",\n           \"peso_bruto\": \"2.5000\",\n           \"serie\": \"00001\",\n           \"tipo_combustivel\": \"16\",\n           \"numero_motor\": \"00001\",\n           \"cmt\": \"2.000\",\n           \"dist\": \"150\",\n           \"ano_modelo\": \"2019\",\n           \"ano_fabricacao\": \"2018\",\n           \"tipo_pintura\": \"Tipo de Pintura\",\n           \"tipo_veiculo\": \"06\",\n           \"especie_veiculo\": \"1\",\n           \"vin\": \"N\",\n           \"condicao_veiculo\": \"1\",\n           \"marca_modelo\": \"00001\",\n           \"cor_denatran\": \"04\",\n           \"lotacao\": \"5\",\n           \"restricao\": \"0\"\n        }\n     }\n  ],\n  \"pedido\": {\n     \"pagamento\": 0,\n     \"presenca\": 2,\n     \"modalidade_frete\": 0,\n     \"frete\": \"12.56\",\n     \"desconto\": \"10.00\",\n     \"total\": \"174.60\"\n  }\n}\n")

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("Cache-Control", "no-cache")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("X-Consumer-Key", "SEU_CONSUMER_KEY")
	req.Header.Add("X-Consumer-Secret", "SEU_CONSUMER_SECRET")
	req.Header.Add("X-Access-Token", "SEU_ACCESS_TOKEN")
	req.Header.Add("X-Access-Token-Secret", "SEU_ACCESS_TOKEN_SECRET")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

}