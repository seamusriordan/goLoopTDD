package dataproducer

import (
	"fmt"
	"math/rand"
	"time"
)

type DataProducer interface {
	GiveData() int
}

type RandomDataProducer struct {
}

func NewRandomDataProducer() RandomDataProducer {
	rand.Seed(time.Now().Unix())
	return RandomDataProducer{}
}

func (r RandomDataProducer) GiveData() int {
	data := rand.Intn(7)
	fmt.Printf("Data produced: %d\n", data)
	return data
}
