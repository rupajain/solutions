package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	//"encoding/json"
	"bytes"
)

func main() {
	const vault_token = "1lqKwGE4HlIGrd7zx6y87rp1"

	// dat, err := ioutil.ReadFile("/payload.json")
	// check(err)
	// fmt.Print(string(dat))

	// f, err := os.Open("/payload.json")
	// check(err)

	//    jsonData :=map[string] interface{}{`{
	//     "options": {
	//         "version": "2"
	//     },
	//     "data": {
	//         "foo": "bar",
	//         "zip": "zap"
	//       }
	//   }`}
	// jsonValue,err :=json.Marshal(jsonData)

	// jsonString:=byte[](`{"keyTest":"ValueTest"}`)

	//    jsonData:=`{
	//         "options": {
	//             "version": "2"
	//         },
	//         "data": {
	//             "foo": "bar",
	//             "zip": "zap"
	//           }
	//       }`

	// jsonValue,err :=json.Unmarshal(jsonData)
	//       fmt.Println("........."+jsonData+"........"+string(jsonValue))
	//request,err:=http.NewRequest("Get","http://127.0.0.1:8200/v1/secret/list=true",nil)
	var jsonStr = []byte(`{"title":"Buy cheese and bread for breakfast,"}`)

	request, err := http.NewRequest("Get", " http://127.0.0.1:8200/v1/secret/data/foo", bytes.NewBuffer(jsonStr))
	//request,err:=http.NewRequest("POST","http://127.0.0.1:8200/v1/secret/data/baz/foo",bytes.NewBuffer(jsonValue))
	//request,err:=http.NewRequest("POST","http://127.0.0.1:8200/v1/secret/data/keyTest",bytes.NewBuffer(jsonValue))

	request.Header.Set("X-Vault-Token", vault_token)
	request.Header.Set("Content-Type", "application/json")
	//request.Body.Set("key1","1")

	client := http.Client{}
	resp, err := client.Do(request)

	if err != nil {
		fmt.Printf("the Http request failed with error %s \n", err)
	}

	data, err := ioutil.ReadAll(resp.Body)
	fmt.Println("Header ", resp.Header)
	if data != nil {
		fmt.Println("Data : ", string(data))
	} else {
		fmt.Println("error to get data : ", err)
	}

	defer resp.Body.Close()
}
