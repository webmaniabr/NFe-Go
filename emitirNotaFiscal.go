package main

import (
	"fmt"
	"strings"
	"net/http"
	"io/ioutil"
)

func main() {

	url := "https://webmaniabr.com/api/1/nfe/emissao/"

	payload := strings.NewReader("{\n\t\"ID\": 1137,\n\t\"url_notificacao\": \"https://webmaniabr.com/retorno.php\",\n\t\"operacao\": 1,\n\t\"natureza_operacao\": \"Venda de produção do estabelecimento\",\n\t\"modelo\": 1,\n\t\"finalidade\": 1,\n\t\"ambiente\": 2,\n\t\"cliente\": {\n\t\t\"cpf\": \"000.000.000-00\",\n\t\t\"nome_completo\": \"Miguel Pereira da Silva\",\n\t\t\"endereco\": \"Av. Anita Garibaldi\",\n\t\t\"complemento\": \"Sala 809 Royal\",\n\t\t\"numero\": 850,\n\t\t\"bairro\": \"Ahú\",\n\t\t\"cidade\": \"Curitiba\",\n\t\t\"uf\": \"PR\",\n\t\t\"cep\": \"80540-180\",\n\t\t\"telefone\": \"(41) 4063-9102\",\n\t\t\"email\": \"suporte@webmaniabr.com\"\n\t},\n\t\"produtos\": [{\n\t\t\"nome\": \"Camisetas Night Run\",\n\t\t\"sku\": \"camiseta-night-run\",\n\t\t\"ean\": \"0789602015376\",\n\t\t\"ncm\": \"6109.10.00\",\n\t\t\"cest\": \"28.038.00\",\n\t\t\"cnpj_produtor\": \"11.290.027/0001-82\",\n\t\t\"quantidade\": 3,\n\t\t\"unidade\": \"UN\",\n\t\t\"peso\": \"0.800\",\n\t\t\"origem\": 0,\n\t\t\"subtotal\": \"44.90\",\n\t\t\"total\": \"134.70\",\n\t\t\"classe_imposto\": \"REF2892\"\n\t}, {\n\t\t\"nome\": \"Camisetas 10 Milhas\",\n\t\t\"sku\": \"camisetas-10-milhas\",\n\t\t\"ean\": \"0789602015376\",\n\t\t\"ncm\": \"6109.10.00\",\n\t\t\"cest\": \"28.038.00\",\n\t\t\"cnpj_produtor\": \"11.290.027/0001-82\",\n\t\t\"quantidade\": \"1\",\n\t\t\"unidade\": \"UN\",\n\t\t\"peso\": \"0.200\",\n\t\t\"origem\": 0,\n\t\t\"subtotal\": \"29.90\",\n\t\t\"total\": \"29.90\",\n\t\t\"classe_imposto\": \"REF2892\"\n\t}],\n\t\"pedido\": {\n\t\t\"pagamento\": 0,\n\t\t\"presenca\": 2,\n\t\t\"modalidade_frete\": 0,\n\t\t\"frete\": \"12.56\",\n\t\t\"desconto\": \"10.00\",\n\t\t\"total\": \"174.60\"\n\t},\n\t\"transporte\": {\n\t\t\"cnpj\": \"11.290.027/0001-82\",\n\t\t\"razao_social\": \"Transportes LTDA\",\n\t\t\"ie\": \"123.456.789.123\",\n\t\t\"endereco\": \"Av. Anita Garibaldi\",\n\t\t\"uf\": \"PR\",\n\t\t\"cidade\": \"Curitiba\",\n\t\t\"cep\": \"80540-180\"\n\t}\n}")

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