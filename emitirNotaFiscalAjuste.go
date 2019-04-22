package main

import (
	"fmt"
	"strings"
	"net/http"
	"io/ioutil"
)

func main() {

	url := "https://webmaniabr.com/api/1/nfe/ajuste/"

	payload := strings.NewReader("{\n\"operacao\": 1,\n\"natureza_operacao\": \"CREDITO ICMS S/ ESTOQUE\", \n\"codigo_cfop\": 2.949,\n\"valor_icms\": 1000.00,\n\"ambiente\": 2,\n\"cliente\": {\n    \"cpf\" : \"000.000.000-00\",\n    \"nome_completo\" : \"Nome completo\",\n    \"endereco\" : \"Av. Brg. Faria Lima\",\n    \"complemento\" : \"Escritório\",\n    \"numero\" : 1000,\n    \"bairro\" : \"Itaim Bibi\",\n    \"cidade\" : \"São Paulo\",\n    \"uf\" : \"SP\",\n    \"cep\" : \"00000-000\",\n    \"telefone\" : \"(00) 0000-0000\",\n    \"email\" : \"nome@email.com\"\n  }\n}\n")

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