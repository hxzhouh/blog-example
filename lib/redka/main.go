package main

import (
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"github.com/nalgeon/redka"
	"log"
)

func main() {
	// Open or create the data.db file.
	db, err := redka.Open("data.db", nil)
	if err != nil {
		log.Fatal(err)
	}
	defer func(db *redka.DB) {
		err = db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(db)
	//err = db.Str().Set("foo", "bar")
	//if err != nil {
	//	log.Fatal(err)
	//}
	v, err := db.Str().Get("foo")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(v)
}
