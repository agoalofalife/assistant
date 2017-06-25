package assistant

import (
	"github.com/googollee/go-socket.io"
	"log"
	"net/http"
	"os"
	"github.com/agoalofalife/assistant/database"
	"github.com/agoalofalife/assistant/database/models"
	"time"
)

const  (
	DATABASE_NAME = "Assistant.db"
)
func Go() {
	server, err := socketio.NewServer(nil)
	if err != nil {
		log.Fatal(err)
	}

	// kill process
	server.On(KILL_PS, func(pid int) bool {
		process, _ := os.FindProcess(pid)
		process.Kill()
		return true
	})

	// list queues
	server.On(LIST_QUEUES, func() (json string){
		db,_ := database.Open(DATABASE_NAME)
		defer db.Close()
		task, _ := models.NewTask(db)
		task.All()
		bt,_ := task.ToJson()
		return string(bt)
	})

	// create new task
	server.On(CREATE_QUEUES, func(strJson []byte) bool {
		log.Println("event server ")
		//db,_ := database.Open(DATABASE_NAME)
		//defer db.Close()
		//task, _ := models.NewTask(db)
		//json.Unmarshal(strJson, task)
		//
		//log.Println("unmarshal : ", task.Name
		return true
	})

	server.On(UPDATE_QUEUE, func(strJson []byte) bool {
		log.Println("event server : update")
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
