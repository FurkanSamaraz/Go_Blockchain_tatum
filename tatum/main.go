package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {

	http.HandleFunc("/", tatum)
	http.ListenAndServe(":8080", nil)

}

func tatum(w http.ResponseWriter, r *http.Request) {

	url := "https://api-eu1.tatum.io/v3/nft/deploy"

	payload := strings.NewReader("{\"chain\":\"ETH\",\"name\":\"My ERC721\",\"symbol\":\"ERC_SYMBOL\",\"fromPrivateKey\":\"0x05e150c73f1920ec14caa1e0b6aa09940899678051a78542840c2668ce5080c2\",\"provenance\":false,\"cashback\":false,\"publicMint\":true,\"nonce\":0,\"fee\":{\"gasLimit\":\"40000\",\"gasPrice\":\"20\"}}")

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("content-type", "application/json")
	req.Header.Add("x-testnet-type", "SOME_STRING_VALUE")
	req.Header.Add("x-api-key", "REPLACE_KEY_VALUE")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

	fmt.Fprintf(w, string(body))

}
