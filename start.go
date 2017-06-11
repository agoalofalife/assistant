package assistant

import (
	"github.com/googollee/go-socket.io"
	"log"
	"net/http"
)

func Go() {
	server, err := socketio.NewServer(nil)
	if err != nil {
		log.Fatal(err)
	}
	allPs()


	server.On("connection", func(so socketio.Socket) {
		log.Println("on connection")
		so.Emit("listProcess", "{test : test }")
	})
	server.On("error", func(so socketio.Socket, err error) {
		log.Println("error:", err)
	})

	http.Handle("/socket.io/", server)
	http.Handle("/", http.FileServer(http.Dir("./asset")))
	log.Println("Serving at localhost:7777...")
	log.Fatal(http.ListenAndServe(":7777", nil))
}
