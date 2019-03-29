package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {

	http.HandleFunc("/hostname", func(w http.ResponseWriter, r *http.Request) {
		hostname, _ := os.Hostname()
		fmt.Fprint(w, "hello from "+hostname)
	})
	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}
