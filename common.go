package main

import "log"

type Configuration struct {
	AMQPConnectionURL string
}

type AddTask struct {
	Number1 int
	Number2 int
	Number3 int
}

var Config = Configuration{
	AMQPConnectionURL: "amqp://guest:guest@localhost:5672/",
}

func handleError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
