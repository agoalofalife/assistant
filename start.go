package assistant

import (
	"github.com/googollee/go-socket.io"
	"log"
	"net/http"

	"time"
)

func Go() {
	server, err := socketio.NewServer(nil)
	if err != nil {
		log.Fatal(err)
	}

			//ps := New()
			//ps.allPs()

	server.On("connection", func(so socketio.Socket) {
		log.Println("on connection")
		go func() {
			for {
				ps := New()
				ps.allPs()
				so.Emit("listProcess", string(ps.toJson()))
				time.Sleep(10 * time.Second)
			}

		}()

	})
	server.On("error", func(so socketio.Socket, err error) {
		log.Println("error:", err)
	})

	http.Handle("/socket.io/", server)
	http.Handle("/", http.FileServer(http.Dir("./asset")))
	log.Println("Serving at localhost:7777...")
	log.Fatal(http.ListenAndServe(":7777", nil))
}
