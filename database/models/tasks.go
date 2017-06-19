package models

import (
	"github.com/boltdb/bolt"
	"fmt"
	"encoding/json"
	"strconv"
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

		return b.Put([]byte(strconv.Itoa(task.ID)), buf)
	})
}


func (task *Task) All () {
	task.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(tableName))
		c := b.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			fmt.Printf("key=%s, value=%s\n", k, v)
		}
		return nil
	})
}

func (task *Task) Find(Id int)  ( stringJson []byte, err error) {
	err =  task.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(tableName))
		stringJson = b.Get([]byte(strconv.Itoa(Id)))
		fmt.Printf("The value is: %s\n", stringJson)
		return nil
	})

	return stringJson, err
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
