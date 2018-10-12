package main

import (
	"fmt"
	"log"

	"github.com/boltdb/bolt"
)

func main() {
	// Open the my.db data file in your current directory.
	// It will be created if it doesn't exist.

	var key, value string

	db, err := bolt.Open("my.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	db.Update(func(tx *bolt.Tx) error {
		//b, err := tx.CreateBucket([]byte("MyBucket"))
		b := tx.Bucket([]byte("MyBucket"))

		for i := 0; i < 3; i++ {
			fmt.Print("Enter Key ........")
			fmt.Scanln(&key)
			fmt.Print("Enter  Value....")
			fmt.Scanln(&value)
			err = b.Put([]byte(key), []byte(value))

			fmt.Println("key..." + key + "......" + "value.. " + value)
		}
		// err = b.Put([]byte("answer1"), []byte("45"))
		// err = b.Put([]byte("answer2"), []byte("46"))
		// err = b.Put([]byte("answer3"), []byte("47"))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}

		return err
	})

	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("MyBucket"))
		// v := b.Get([]byte("answer"))
		// fmt.Printf("The answer is: %s\n", v)
		// v = b.Get([]byte("answer1"))
		// fmt.Printf("The answer1 is: %s\n", v)

		c := b.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			fmt.Printf("key=%s, value=%s\n", k, v)
		}

		return nil
		return nil
	})
	defer func() { _ = db.Close() }()
}
