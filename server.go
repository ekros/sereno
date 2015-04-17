package main

import (
	"fmt"
	"log"
	"net/http"
)

type Hello struct{}

func (h Hello) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request) {
	fmt.Fprint(w, "<b>Hello!</b>")
}

func main() {
	var h Hello
	err := http.ListenAndServe("localhost:4000", h)
	if err != nil {
		log.Fatal(err)
	}

  http.Handle("/ping", string("pong!"))
}
