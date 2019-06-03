package main

import "net/http"

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		//重定向
		http.Redirect(w, r, "http://localhost:8083", 308)
	})
	http.ListenAndServe(":8082", nil)
}
