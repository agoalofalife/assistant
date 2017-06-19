package models

import (
	"github.com/boltdb/bolt"
	"fmt"
	"encoding/json"
	"encoding/binary"
)

const tableName = "Task"
type Task struct {
	ID int `json:"id"`
	Name string `json:"name"`
	CommandConsole string `json:"commandConsole"`
	TimeOut string `json:"timeOut"`
	TimeStart string `json:"timeStart"`
	db *bolt.DB
}

func (task *Task) CreateTask() error {
	return task.db.Update(func(tx *bolt.Tx) error {

		b := tx.Bucket([]byte(tableName))
		id, _ := b.NextSequence()
		task.ID = int(id)

		buf, err := json.Marshal(task)
		if err != nil {
			return err
		}

		return b.Put(itob(task.ID), buf)
	})
}

func (task *Task) Find(Id int)  error {
	return task.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("MyBucket"))
		v := b.Get(itob(Id))
		fmt.Printf("The answer is: %s\n", v)
		return nil
	})
}
func NewTask(db *bolt.DB) (task *Task, err error)  {
	err = db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucket([]byte(tableName))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		return nil
	})
	task = &Task{db:db}
	return task, err
}

// itob returns an 8-byte big endian representation of v.
func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}