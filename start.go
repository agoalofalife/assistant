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

	ps := New()
	ps.allPs()
	//log.Println(string(ps.toJson()))
	//os.Exit(2)
	server.On("connection", func(so socketio.Socket) {
		log.Println("on connection")
		so.Emit("listProcess", string(ps.toJson()))
	})
	server.On("error", func(so socketio.Socket, err error) {
		log.Println("error:", err)
	})

	http.Handle("/socket.io/", server)
	http.Handle("/", http.FileServer(http.Dir("./asset")))
	log.Println("Serving at localhost:7777...")
	log.Fatal(http.ListenAndServe(":7777", nil))
}
