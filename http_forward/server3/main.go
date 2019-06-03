package main

import (
	"log"
	"net/http"
)

func main() {

	log.Println("listen on:8083")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("response from  server 3.")
		w.Write([]byte("i am server 3."))
	})
	http.ListenAndServe(":8083", nil)
}
