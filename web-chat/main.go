package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

func main() {
	options, err := getOptions()

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("This is the web chat app, running on host %v:%v\n", *options.ip, *options.port)

	portStr := strconv.Itoa(*options.port)
	if *options.isHost {
		fmt.Println("You are the host")
		go hostWS(":" + portStr)
		connect(*options.ip+":"+portStr, *options.username)
	} else {
		fmt.Println("You are not the host")
		connect(*options.ip+":"+portStr, *options.username)
	}
}

func hostWS(port string) {
	host := newHost()
	go host.run()
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(host, w, r)
	})

	server := &http.Server{
		Addr:              port,
		ReadHeaderTimeout: 3 * time.Second,
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
