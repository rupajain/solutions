package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"io/ioutil"
)

func main() {
	const vault_token="6adIAN8Zht18EtMxGjcIJgTq"
	// jsonBlob:=[]byte(`{
	// 	"options": {
	// 		"version": "2"
	// 	},
	// 	"data": {
	// 		"foo": "bar",
	// 		"zip": "zap"
	// 	  }
	//   }`)


// 	var jsonBlob = []byte(`[
//     {"Name": "Platypus", "Order": "Monotremata"},
//     {"Name": "Quoll",    "Order": "Dasyuromorphia"}
// ]`)

	//var b bytes.Buffer
	// b := *bytes.NewBuffer(make([]byte, 20))
	// b.Write(jsonBlob)
	//b.String()
	locJson, _:= json.Marshal(`{loc{Lat: 35.14326, Lon: -116.104}}`)
	  
	//buf,err :=json.Marshal(b)
	//fmt.Printf("ddddddd     %s  \n",buf)
	request,_:=http.NewRequest("POST","http://127.0.0.1:8200/v1/secret/data/",bytes.NewBuffer(locJson))


	request.Header.Set("X-Vault-Token",vault_token)
request.Header.Set("Content-Type","application/json")


client:=http.Client{}
resp,err:=client.Do(request)
	// var dat interface{}
	// err = json.Unmarshal(b.Bytes(), &dat)
	// if err != nil {
	// 	fmt.Println("error:", err)
	// }
	// fmt.Printf("%+v", dat)


	if err!=nil{
		fmt.Printf("the Http request failed with error %s \n", err)
	}
		data, err:=ioutil.ReadAll(resp.Body)
		fmt.Println("Header ",resp.Header)
	   if data!=nil{
		   fmt.Println("Data : ",string(data))
	   }else{
		   fmt.Println("error to get data : ",err)
	   }
	
	   
	   defer resp.Body.Close()
}