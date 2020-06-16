package main

import (
	"flag"
	"log"
	"net/http"
)

func main()  {
	var serverAddr string
	flag.StringVar(&serverAddr, "l", "0.0.0.0:80", "listen addr")
	flag.Parse()

	log.Println("start")
	http.Handle("/hello", websocket.Handler(WebsocketConn))
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		//writer.Header().Set("Location", "https://google.com")
		writer.Write([]byte("Test OK"))
	})

	err := http.ListenAndServe(serverAddr,  nil)
	if err != nil {
		log.Println("Listen Error:", err.Error())
	}
}

func WebsocketConn(ws *websocket.Conn) {
	ws.Write([]byte("TestOk"))
	ws.Close()
}
