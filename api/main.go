package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {

	http.HandleFunc("/hostname", func(w http.ResponseWriter, r *http.Request) {
		hostname, _ := os.Hostname()
		log.Println("serving from: " + hostname)
		time.Sleep(time.Second * 5)
		fmt.Fprint(w, "hello from "+hostname)
	})

	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}
