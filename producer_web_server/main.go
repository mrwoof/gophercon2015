package main

import (
	"fmt"
	"log"
	"net/http"

	//"github.com/bitly/go-nsq"
)

func redirect(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() // parse arguments, you have to call this by yourself
	url := r.Form["url"][0]
	fmt.Fprintf(w, "I would redirect you to \"%s\" if I were so inclined", url)
}

func main() {
	http.HandleFunc("/redirect", redirect)
	log.Printf("Listening...")
	err := http.ListenAndServe(":9090", nil) // set listen port
	if err != nil {
		log.Fatal("Error starting server: ", err)
	}
}
