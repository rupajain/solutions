package main

import (
   // "bufio"
    "fmt"
   // "io"
    "io/ioutil"
   // "os"
   "encoding/json"
   "net/http"
 "bytes"
)

// Reading files requires checking most calls for errors.
// This helper will streamline our error checks below.
func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
	const vault_token="6adIAN8Zht18EtMxGjcIJgTq"
    // Perhaps the most basic file reading task is
    // slurping a file's entire contents into memory.
    dat, err := ioutil.ReadFile("/Users/rupa.j/Desktop/GO/src/practice/payload.json")
    check(err)
    fmt.Println(string(dat))
	json_bytes := []byte(dat)

fmt.Println(string(json_bytes)) 
   // f.Close()
   var m map[string]interface{}
  // err = json.Unmarshal(json_bytes, &m)
   b := new(bytes.Buffer)
   json.NewEncoder(b).Encode(m)
   fmt.Println(b)
   if err != nil {
	   panic(err)
   }


   request,err:=http.NewRequest("POST"," http://127.0.0.1:8200/v1/secret/data/baz/foo",b)
//request,err:=http.NewRequest("POST","http://127.0.0.1:8200/v1/secret/data/keyTest",bytes.NewBuffer(jsonValue))


request.Header.Set("X-Vault-Token",vault_token)
request.Header.Set("Content-Type","application/json")
//request.Body.Set("key1","1")

client:=http.Client{}
resp,err:=client.Do(request)

if err!=nil{
    fmt.Printf("the Http request failed with error %s \n", err)
}




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