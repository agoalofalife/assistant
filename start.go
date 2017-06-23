package assistant

import (
	"github.com/googollee/go-socket.io"
	"log"
	"net/http"
	"os"
	"time"
	"github.com/agoalofalife/assistant/database"
	"github.com/agoalofalife/assistant/database/models"
)

const  (
	DATABASE_NAME = "Assistant.db"
)
func Go() {
	server, err := socketio.NewServer(nil)
	if err != nil {
		log.Fatal(err)
	}

	//ps := New()
	//ps.allPs()
	// create new task

	// list queues
	server.On(LIST_QUEUES, func() (json string){
		db,_ := database.Open(DATABASE_NAME)
		defer db.Close()
		task, _ := models.NewTask(db)
		task.All()
		bt,_ := task.ToJson()
		return string(bt)
	})


	server.On(CREATE_QUEUES, func(s int) bool {
		log.Println("How are you???????")
		log.Println("json create task", s)
		return true
	})

	// kill process
	server.On(KILL_PS, func(pid int) bool {
		process, _ := os.FindProcess(pid)
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
