package main

import (
	"fmt"
	"strings"
	"net/http"
	"io/ioutil"
)

func main() {

	url := "https://webmaniabr.com/api/1/nfe/complementar/"

	payload := strings.NewReader("{\n\t\"nfe_referenciada\": \"00000000000000000000000000000000000000000000\",\n\t\"operacao\": 1,\n\t\"natureza_operacao\": \"Natureza da operação\",\n\t\"ambiente\": 2,\n\t\"cliente\": {\n\t\t\"cpf\": \"000.000.000-00\",\n\t\t\"nome_completo\": \"Nome completo\",\n\t\t\"endereco\": \"Av. Brg. Faria Lima\",\n\t\t\"complemento\": \"Escritorio\",\n\t\t\"numero\": 1000,\n\t\t\"bairro\": \"Itaim Bibi\",\n\t\t\"cidade\": \"São Paulo\",\n\t\t\"uf\": \"SP\",\n\t\t\"cep\": \"00000-000\",\n\t\t\"telefone\": \"(00) 0000-0000\",\n\t\t\"email\": \"nome@email.com\"\n\t},\n\t\"produtos\": [{\n\t\t\"nome\": \"Nome do produto\",\n\t\t\"codigo\": \"nome-do-produto\",\n\t\t\"ncm\": \"6109.10.00\",\n\t\t\"cest\": \"28.038.00\",\n\t\t\"quantidade\": 1,\n\t\t\"unidade\": \"UN\",\n\t\t\"origem\": 0,\n\t\t\"subtotal\": \"29.90\",\n\t\t\"total\": \"29.90\",\n\t\t\"impostos\": {\n\t\t\t\"icms\": {\n\t\t\t\t\"codigo_cfop\": \"6.102\",\n\t\t\t\t\"situacao_tributaria\": \"102\"\n\t\t\t}\n\t\t}\n\t}]\n}")

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