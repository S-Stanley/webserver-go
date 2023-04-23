package main

import (
	// "fmt"
	"net/http"
	"log"
)

func router() {
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintf(w, "Hello world!")
	// })

	fs := http.FileServer(http.Dir("./static/"))

	// http.Handle("/", http.StripPrefix("/static/", fs))
	http.Handle("/", fs)

}

func run_server() {
	err := http.ListenAndServe(":8000", nil);

	if (err != nil) {
		log.Fatal(err)
	}
}

func main() {
	router()
	run_server()
}