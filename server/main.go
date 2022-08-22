package main

import (
	"fmt"
	"net/http"
	"time"

	// import throttle package
	"github.com/galuhpradipta/go-throttle/server/throttle"
)

const (
	port   = "8080"
	window = time.Second * 1
	limit  = 10
)

func main() {
	store := make(map[string]throttle.Limiter)
	throttleHandler := throttle.New(window, store)
	_ = throttleHandler

	http.HandleFunc("/", helloHandler)
	fmt.Println("Listening on port ", port)
	http.ListenAndServe(":"+port, nil)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world!"))
}
