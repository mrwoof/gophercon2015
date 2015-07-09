package main

import (
	"encoding/json"
	"log"
	"sync"

	"github.com/bitly/go-nsq"
)

var (
	nsqlookupds = []string{"127.0.0.1:4161"}
	topic       = "test"
	channel     = "my_consumer_channel"
)

type RedirectMessage struct {
	Url        string `json:"url"`
	Timestamp  int    `json:"ts"`
	UserAgent  string `json:"ua"`
	IPAddr     string `json:"ip"`
}

func MessageHandler(message *nsq.Message) error {
	log.Printf("Got a message: \"%s\"", message.Body)

	var redirMsg RedirectMessage
	err := json.Unmarshal(message.Body, &redirMsg)
	if err != nil {
		log.Printf("Unable to decode json, skipping '%s'", message.Body)
		return nil
	}

	log.Printf("Got a good message: %#v", redirMsg)
	log.Printf("url: %s", redirMsg.Url)
	log.Printf("timestamp: %d", redirMsg.Timestamp)

	return nil
}

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(1)

	config := nsq.NewConfig()
	consumer, err := nsq.NewConsumer(topic, channel, config)
	if err != nil {
		log.Fatal(err)
	}

	consumer.AddHandler(nsq.HandlerFunc(MessageHandler))

	err = consumer.ConnectToNSQLookupds(nsqlookupds)
	if err != nil {
		log.Panic("Could not connect")
	}
	wg.Wait()
}
