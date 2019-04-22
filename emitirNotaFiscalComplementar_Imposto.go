package main

import (
	"fmt"
	"strings"
	"net/http"
	"io/ioutil"
)

func main() {

	url := "https://webmaniabr.com/api/1/nfe/complementar/"

	payload := strings.NewReader("{\n    \"nfe_referenciada\": \"00000000000000000000000000000000000000000000\",\n    \"operacao\": 1,\n    \"natureza_operacao\": \"Natureza da operação\",\n    \"ambiente\": 2,\n    \"cliente\": {\n        \"cpf\": \"000.000.000-00\",\n        \"nome_completo\": \"Nome completo\",\n        \"endereco\": \"Av. Brg. Faria Lima\",\n        \"complemento\": \"Escritorio\",\n        \"numero\": 1000,\n        \"bairro\": \"Itaim Bibi\",\n        \"cidade\": \"São Paulo\",\n        \"uf\": \"SP\",\n        \"cep\": \"00000-000\",\n        \"telefone\": \"(00) 0000-0000\",\n        \"email\": \"nome@email.com\"\n    },\n    \"produtos\": [\n        {\n            \"nome\": \"Nome do produto\",\n            \"codigo\": \"nome-do-produto\",\n            \"ncm\": \"6109.10.00\",\n            \"cest\": \"28.038.00\",\n            \"quantidade\": 1,\n            \"unidade\": \"UN\",\n            \"peso\": \"0.500\",\n            \"origem\": 0,\n            \"subtotal\": \"29.90\",\n            \"total\": \"29.90\",\n            \"tributos_federais\": \"10.00\",\n            \"tributos_estaduais\": \"10.00\",\n            \"impostos\": {\n                \"icms\": {\n                    \"codigo_cfop\": \"6.102\",\n                    \"situacao_tributaria\": \"102\"\n                },\n                \"ipi\": {\n                    \"situacao_tributaria\": \"99\",\n                    \"codigo_enquadramento\": \"999\",\n                    \"aliquota\": \"0.00\"\n                },\n                \"pis\": {\n                    \"situacao_tributaria\": \"99\",\n                    \"aliquota\": \"0.00\"\n                },\n                \"cofins\": {\n                    \"situacao_tributaria\": \"99\",\n                    \"aliquota\": \"0.00\"\n                }\n            }\n        }\n    ]\n}")

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