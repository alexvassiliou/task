package db

import (
	"time"

	"go.etcd.io/bbolt"
)

var taskBucket = []byte("tasks")
var db *bbolt.DB

// Task for the todo list
type Task struct {
	Key   int
	Value string
}

// Init initialises a database connection
func Init(dbPath string) error {
	var err error
	db, err := bbolt.Open(dbPath, 0666, &bbolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return err
	}
	defer db.Close()
	return db.Update(func(tx *bbolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(taskBucket)
		return err
	})
}
