package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
)

func main() {

	url := "https://webmaniabr.com/api/1/nfe/sefaz/"

	req, _ := http.NewRequest("GET", url, nil)

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