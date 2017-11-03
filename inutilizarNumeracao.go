/*
JSON request:

{
  "sequencia": "101-109",
  "motivo": "Inutilização por problemas técnicos.",
  "ambiente": "1",
  "serie": "99",
  "modelo": "1"
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

	url := "https://webmaniabr.com/api/1/nfe/inutilizar/"

	payload := strings.NewReader("{\"sequencia\":\"101-109\",\"motivo\":\"Inutilização por problemas técnicos.\",\"ambiente\":\"1\",\"serie\":\"99\",\"modelo\":\"1\"}")

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
