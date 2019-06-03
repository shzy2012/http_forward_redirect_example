package main

import (
	"log"
	"math/rand"
	"net/http"
	"net/http/httputil"
	"net/url"
)

//NewMultipleHostsReverseProxy 转发
func NewMultipleHostsReverseProxy(targets []*url.URL) *httputil.ReverseProxy {
	director := func(req *http.Request) {
		target := targets[rand.Int()%len(targets)]
		req.URL.Scheme = target.Scheme
		req.URL.Host = target.Host
		req.URL.Path = target.Path
	}

	return &httputil.ReverseProxy{Director: director}
}

func main() {

	log.Println("proxy listen on:8080")
	proxy := NewMultipleHostsReverseProxy([]*url.URL{
		{
			Scheme: "http",
			Host:   "localhost:8083",
		},
		{
			Scheme: "http",
			Host:   "localhost:8082",
		},
		{
			Scheme: "http",
			Host:   "localhost:8081",
		},
	})
	log.Fatal(http.ListenAndServe(":8080", proxy))
}
