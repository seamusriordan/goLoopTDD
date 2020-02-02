package main

import (
	"./dataWriterClient"
	"./dataproducer"
	"time"
)

func main() {
	dataProducer := dataproducer.NewRandomDataProducer()
	client := dataWriterClient.StubClient{}
	Program(&client, &dataProducer)
}

func Program(client dataWriterClient.Client, dataProducer dataproducer.DataProducer) {
	client.Connect("http://localhost")
	for true {
		client.SendData(dataProducer.GiveData())
		time.Sleep(0)
	}
}
