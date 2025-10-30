package main

import (
	"fmt"
	"log"

	"github.com/google/uuid"
	"go.etcd.io/bbolt"
)

func main() {
	db, err := bbolt.Open(".db", 0666, nil)
	if err != nil {
		log.Fatal(err)
	}

	user := map[string]string {
		"name" : "john",
		"age": "19",
	}	// int, string, []byte, float, ...

	db.Update(func(tx *bbolt.Tx ) error { 
		b, err := tx.CreateBucket([]byte("users"))
		if err != nil {
			return err
		}

		id := uuid.New()
		for k, v := range user {
			if err := b.Put([]byte(k), []byte(v)); err != nil {
				return err
			}
		}
		if err := b.Put([]byte("id"), []byte(id.String())); err != nil {
			return err
		}
		return nil
	})

	newUser := make(map[string]string)

	if err := db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte("users"))
		if  b == nil {
			return fmt.Errorf("bucket (%s) not found", "user")
		}

		b.ForEach(func(k, v []byte) error {
			newUser[string(k)] = string(v)
			return nil
		})


		return nil
	}); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Hello, world!")
	fmt.Println(newUser)
}