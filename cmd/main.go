package main

import (
	"fmt"
	"log"

	"github.com/dipendra-mule/go-real-time-database/database"
)

func main() {
	db, err := database.New()
	if err != nil {
		log.Fatal(err)
	}

	user := map[string]string{
		"name": "john",
		"age":  "19",
	}

	// int, string, []byte, float, ...

	id, err := db.Insert("users", user)
	if err != nil {
		log.Fatal(err)
	}
	// coll, err := db.CreateCollection("users")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	fmt.Printf("%+v\n", id)
}
