package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"

	"net/http"
)

const vault_token = "pvIcmSYezrOd2JRx4wdGr2ST"

func main() {
	http.HandleFunc("/postdata", postdata)
	http.HandleFunc("/getdata", getData)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}

}

func postdata(w http.ResponseWriter, r *http.Request) {
	jsonData := []byte(`{
	        "options": {
	            "version": "2"
	        },
	        "data": {
	            "foo": "bar",
	            "zip": "zap"
	          }
	      }`)

	request, err := http.NewRequest("POST", "http://127.0.0.1:8200/v1/secret/data/baz", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatal(err)
	}
	request.Header.Set("X-Vault-Token", vault_token)
	request.Header.Set("Content-Type", "application/json")

	client := http.Client{}
	resp, err := client.Do(request)

	if err != nil {
		log.Fatal(err)
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Header ", resp.Header)
	if data != nil {
		fmt.Println("Data : ", string(data))
	} else {
		fmt.Println("error to get data : ", err)
	}

	defer resp.Body.Close()
}

func getData(w http.ResponseWriter, r *http.Request) {

	request, err := http.NewRequest("Get", " http://127.0.0.1:8200/v1/secret/data/baz/foo", nil)
	request.Header.Set("X-Vault-Token", vault_token)
	if err != nil {
		log.Fatal(err)
	}
	client := http.Client{}
	resp, err := client.Do(request)

	if err != nil {
		log.Fatal(err)
	}
	data, err := ioutil.ReadAll(resp.Body)
	fmt.Println("Header ", resp.Header)
	if err != nil {
		log.Fatal(err)
	}
	if data != nil {
		fmt.Println("Data : ", string(data))
	} else {
		log.Fatal(err)
	}

	defer resp.Body.Close()
}
