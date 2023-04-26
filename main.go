package main

import (
	"fmt"
	"net/http"
	"log"
	"html/template"
)

func router() {

	tmpl, _ := template.ParseFiles("static/index.html")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		Data := struct {
			Name string
			Lang string
		} {
			"Hello world",
			"Go",
		}

		tmpl.Execute(w, Data)
	})
}

func run_server() {
	port := ":8000"
	fmt.Printf("Running on %s\n", port);
	err := http.ListenAndServe(port, nil);

	if (err != nil) {
		log.Fatal(err)
	}

}

func main() {
	router()
	run_server()
}