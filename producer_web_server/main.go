package main

import (
	"fmt"
	"log"
	"encoding/json"
	"net/http"
	"time"

	"github.com/bitly/go-nsq"
)

var (
	outgoing_topic = "test"
	producer *nsq.Producer
	nsqdAddr = "127.0.0.1:4150"
)

type RedirectMessage struct {
	Url       string `json:"url"`
	Timestamp int    `json:"ts"`
	UserAgent string `json:"ua"`
	IPAddr    string `json:"ip"`
}

func redirect(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() // parse arguments, you have to call this by yourself
	url := r.Form["url"][0]
	fmt.Fprintf(w, "I would redirect you to \"%s\" if I were so inclined", url)

	redirMsg := &RedirectMessage{
		Url: url,
		Timestamp: int(time.Now().Unix()),
		UserAgent: "Mozilla",
		IPAddr: "10.10.10.10",
	}

	msg, err := json.Marshal(redirMsg)
	if err != nil {
		log.Printf("Error mashalling msg to JSON bytes: %#v", redirMsg)
	}
	err = producer.Publish(outgoing_topic, msg)
	if err != nil {
		log.Printf("Error publishing to NSQ %s", err)
	}
}

func main() {
	var err error
	producer, err = nsq.NewProducer(nsqdAddr, nsq.NewConfig())
	if err != nil {
		log.Fatal("Error creating NSQ producer: ", err)
	}

	http.HandleFunc("/redirect", redirect)
	log.Printf("Listening...")
	err = http.ListenAndServe(":9090", nil) // set listen port
	if err != nil {
		log.Fatal("Error starting server: ", err)
	}
}
