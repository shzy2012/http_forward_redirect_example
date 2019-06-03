package main

import (
	"log"
	"net/http"
)

func main() {

	log.Println("listen on:8081")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("response from  server 1.")
		w.Write([]byte("i am server 1."))
	})
	http.ListenAndServe(":8081", nil)
}
