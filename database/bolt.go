package database

import (
	"github.com/boltdb/bolt"
	"log"
	"time"
)

// open database bolt
func Open(databaseName string) (*bolt.DB, error) {
	db, err := bolt.Open(databaseName, 0600,  &bolt.Options{Timeout: 1 * time.Second})

	if err != nil {
		log.Fatal(err)
	}

	return db, err
}
