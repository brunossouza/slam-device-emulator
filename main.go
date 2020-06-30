package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	var equipamento = "televisão"

	response, err := http.Post("http://localhost:8080/api/v1/devices/registry", "application/json", bytes.NewBufferString(equipamento))
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	bodyString := string(body)
	fmt.Println("Token:\t", bodyString)
	fmt.Println("done")
}
