package database

import (
	"github.com/boltdb/bolt"
	"log"
)

// open database bolt
func Open(databaseName string) (*bolt.DB, error){
	db, err := bolt.Open(databaseName, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	return db,err
}
