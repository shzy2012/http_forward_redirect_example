package main

import (
	"log"
	"net/http"
)

func main() {

	log.Println("listen on:8082")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("response from  server 2.")
		w.Write([]byte("i am server 2."))
	})

	err := http.ListenAndServe(":8082", nil)
	if err != nil {
		log.Panicln(err)
	}
}
