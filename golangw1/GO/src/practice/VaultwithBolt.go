package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"

	"github.com/boltdb/bolt"
)

const token = "r62VB44FedLUXMaglTKwWNsp"

func main() {

	url := "http://127.0.0.1:8200/v1/secret/data/baz"
	fmt.Println("URL", url)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("error....", err)
	}
	req.Header.Set("X-Vault-Token", token)

	client := http.Client{}
	resp, err := client.Do(req)
	output, err := httputil.DumpRequest(req, true)
	if err != nil {
		log.Fatal("error....", err)
	}
	fmt.Println("request format is...", string(output))
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("error....", err)
	}
	fmt.Println(string(data))

	db, err := bolt.Open("my.db", 0600, nil)
	if err != nil {
		log.Fatal("error....", err)
	}
	defer db.Close()
	db.Update(func(tx *bolt.Tx) error {

		var j = []byte(string(data))

		b, err := tx.CreateBucketIfNotExists([]byte("MyBucket"))
		b = tx.Bucket([]byte("MyBucket"))
		if err != nil {
			log.Fatal("error....", err)
		}
		c := make(map[string]interface{})

		err = json.Unmarshal(j, &c)

		if err != nil {
			log.Fatal("error....", err)
		}

		// a string slice to hold the keys
		k := make([]string, len(c))

		// iteration counter
		i := 0
		mapdata := make(map[string]interface{})

		// copy c's keys into k
		for s, _ := range c {
			k[i] = s
			if k[i] == "data" {
				fmt.Printf("data............%#v\n", c["data"])
				mapdata["data"] = c["data"]
				d := mapdata["data"].(map[string]interface{})
				k1 := d["data"].(map[string]interface{})
				j := 0
				keys := make([]string, len(k1))
				stringvalueo := make([]string, len(k1))
				for x, _ := range k1 {
					keys[j] = x
					fmt.Printf("....keys....%#v\n", keys[j])
					fmt.Printf("....values....%#v\n", k1[keys[j]])
					stringvalueo[j] = k1[keys[j]].(string)
					err = b.Put([]byte(keys[j]), []byte(stringvalueo[j]))
					j++

					fmt.Printf("value of j..%v", j)
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

			return err

		})

		defer resp.Body.Close()

		return err

	})
}
