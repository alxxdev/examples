package main

import (
	"log"
	"net/http"

	"golang.org/x/time/rate"
)

func main() {
	const addr = ":8080"
	mux := http.NewServeMux()
	mux.HandleFunc("/", HelloHandler)
	log.Print("Listening on", addr)
	http.ListenAndServe(addr, limit(mux))
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from server..."))
}

var limiter = rate.NewLimiter(1, 3)

func limit(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !limiter.Allow() {
			http.Error(w, http.StatusText(429), http.StatusTooManyRequests)
			return
		}
		next.ServeHTTP(w, r)
	})
}
