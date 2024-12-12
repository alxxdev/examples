package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	const addr = ":8080"
	http.HandleFunc("/", HelloHanler)
	log.Println("Listening", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}

func HelloHanler(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintf(w, "Hello from server\n")
}
