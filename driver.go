package main

import (
	"./dataWriterClient"
	"./dataproducer"
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hiiii")
}

func Program(client dataWriterClient.Client, dataProducer dataproducer.DataProducer) {
	client.Connect("someplace")
	for true {
		client.SendData(dataProducer.GiveData())
		time.Sleep(0)
	}
}
