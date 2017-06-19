package assistant

import (
	"github.com/googollee/go-socket.io"
	"log"
	"net/http"
	"time"
	"os"
	"github.com/boltdb/bolt"
	"fmt"
	"github.com/agoalofalife/assistant/database"
)

func Go() {
	server, err := socketio.NewServer(nil)
	if err != nil {
		log.Fatal(err)
	}

			//ps := New()
			//ps.allPs()

	// kill process
	server.On(KILL_PS, func(pid int) bool {
		process,_ := os.FindProcess(pid)
		process.Kill()
		return true
	})

	server.On(CONNECT, func(so socketio.Socket) {
		log.Println("on connection")
		go func() {
			for {
				ps := New()
				ps.allPs()
				so.Emit("listProcess", string(ps.toJson()))
				time.Sleep(time.Second)
			}

		}()

	})
	server.On(ERROR, func(so socketio.Socket, err error) {
		log.Println("error:", err)
	})

	http.Handle("/socket.io/", server)
	http.Handle("/", http.FileServer(http.Dir("./asset")))
	log.Println("Serving at localhost:7777...")
	log.Fatal(http.ListenAndServe(":7777", nil))
}


func SoTestBd() {

	db, err := 	database.Open("my.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucket([]byte("MyBucket"))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		return nil
	})
	db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("MyBucket"))
		err := b.Put([]byte("answer"), []byte("42"))
		return err
	})
	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("MyBucket"))
		v := b.Get([]byte("answer"))
		fmt.Printf("The answer is: %s\n", v)
		return nil
	})
	db.View(func(tx *bolt.Tx) error {
		// Assume bucket exists and has keys
		b := tx.Bucket([]byte("MyBucket"))

		c := b.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			fmt.Printf("key=%s, value=%s\n", k, v)
		}

		return nil
	})
}