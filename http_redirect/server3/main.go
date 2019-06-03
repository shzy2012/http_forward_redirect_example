package main

import "net/http"

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		//响应
		w.Write([]byte("i am server 3."))
	})
	http.ListenAndServe(":8083", nil)
}
