package main

import (
//"bytes"
"fmt"
"io/ioutil"
"net/http"
"encoding/json"
"github.com/boltdb/bolt"
"log"
)

const token = "4JG3lzJlYB0uKs3TIWCFvA7s"

func main() {
db, err := bolt.Open("my.db", 0600, nil)
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()
    db.Update(fupdate(db))
}
func fupdate(tx *bolt.Tx) error {
	b, err := tx.CreateBucketIfNotExists([]byte("MyBucket"))
	b = tx.Bucket([]byte("MyBucket"))
	if err != nil {
		log.Fatal(err)
	}

url := "http://127.0.0.1:8200/v1/secret/data/baz/foo"
fmt.Println("URL", url)

//var jsonStr = []byte(`{"title":"Buy cheese and bread for breakfast,"}`)

req, err := http.NewRequest("GET", url,nil)

req.Header.Set("X-Vault-Token", token)
req.Header.Set("Content-Type", "application/json")

client := http.Client{}
resp, err := client.Do(req)

if err != nil {
panic(err)
} else {
data, _ := ioutil.ReadAll(resp.Body)

fmt.Println(string(data))

fmt.Println(string(data[0]))

var j=[]byte(string(data))

c := make(map[string]interface{})


e := json.Unmarshal(j, &c)


if e != nil {
panic(e)
}

// a string slice to hold the keys
k := make([]string, len(c))

// iteration counter
i := 0
mapdata:= make(map[string]interface{})

// copy c's keys into k
for s, _ := range c {
k[i] = s
if k[i]=="data"{
// fmt.Printf("data............%#v\n", c["data"])
mapdata["data"]=c["data"]
d:=mapdata["data"].(map[string]interface{})
k1:=d["data"].(map[string]interface{})
j:=0
keys := make([]string, len(k1))
stringvalueo := make([]string, len(k1))
for x,_ :=range k1 {
keys[j]=x
fmt.Printf("....keys....%#v\n",keys[j])
fmt.Printf("....values....%#v\n",k1[keys[j]])
stringvalueo[j]=k1[keys[j]].(string)
err = b.Put([]byte(keys[j]),[]byte(stringvalueo[j]))
j++
fmt.Printf("value of j..%v",j)
}

if err != nil {
return fmt.Errorf("create bucket: %s", err)
}
//fmt.Printf("data............%#v\n",d["data"].(map[string]interface{}))
//fmt.Printf("data............%#v\n", mapdata["data"].(map[string]interface{}))
}

i++
}



db.View(func(tx *bolt.Tx) error {
b := tx.Bucket([]byte("MyBucket"))

c := b.Cursor()

for k, v := c.First(); k != nil; k, v = c.Next() {
fmt.Printf("key=%s, value=%s\n", k, v)
}

return nil

})

defer resp.Body.Close()
}
return nil
	
}