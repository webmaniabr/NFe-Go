/*
JSON request:

{
   "chave":"45150819652219000198550990000000011442380343",
   "motivo":"Cancelamento por motivos administrativos.",
   "ambiente":"1"
}
*/

package main

import (
	"fmt"
	"strings"
	"net/http"
	"io/ioutil"
)

func main() {

	url := "https://webmaniabr.com/api/1/nfe/cancelar/"

	payload := strings.NewReader("{\"chave\":\"45150819652219000198550990000000011442380343\",\"motivo\":\"Cancelamento por motivos administrativos.\",\"ambiente\":\"1\"}")

	req, _ := http.NewRequest("PUT", url, payload)

	req.Header.Add("cache-control", "no-cache")
	req.Header.Add("content-type", "application/json")
	req.Header.Add("x-consumer-key", "SEU_CONSUMER_KEY")
	req.Header.Add("x-consumer-secret", "SEU_CONSUMER_SECRET")
	req.Header.Add("x-access-token", "SEU_ACCESS_TOKEN")
	req.Header.Add("x-access-token-secret", "SEU_ACCESS_TOKEN_SECRET")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

}
