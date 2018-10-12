package main

import(
	"fmt"
	"net/http"
	"io/ioutil"
	"log"
    "bytes"
	"encoding/json"

)


func main(){
	
	
	//response ,err :=http.Get("http://127.0.0.1:8200/v1/sys/init")

	const vault_token="7CB7XsDZ5iy9MIGM2NszfLJY"
log.Fatal(http.ListenAndServe(":8082",nil))
//http.Handle("/",ListKeys)
//http.HeadersetHeader()


	//log.Fatal(http.ListenAndServe(":8200",nil))
	response ,err :=http.Get("http://127.0.0.1:8200/v1/secret/foo")
	response.setHeader("X-Vault-Token",vault_token)
	response.setHeader("content-type","application/json")

	//response.Header.Set("Content-Type", "application/json")

	//response.Header.Set("X-Vault-Token",vault_token)
	//response.Header.Add()
	
	if err!=nil {
		fmt.Printf("the Http Request failed with error  %s \n", err)
	}
		data, _:=ioutil.ReadAll(response.Body)
		fmt.Println(string(data))
	
	jsonData :=map[string] string{"firstname":"Rupa","lastName":"jain"}
	jsonValue,_ :=json.Marshal(jsonData)
	
	request, _ :=http.NewRequest("POST","http://127.0.0.1:8200/v1/data/secret/foo",bytes.NewBuffer(jsonValue))

	//request, err:=http.NewRequest("GET","http://127.0.0.1:8200/v1/secret/foo",bytes.NewBuffer(jsonValue))
	request.Header.Set("Content-Type","application/json")
	request.Header.Set("key","7CB7XsDZ5iy9MIGM2NszfLJY")
	client:= &http.Client{}
	response,err =client.Do(request)
	response,err =http.Post("http://127.0.0.1:8200/v1/secret/foo","application/json",bytes.NewBuffer(jsonValue))

	
}