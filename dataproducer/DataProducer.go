package dataproducer

import (
	"math/rand"
)

type DataProducer interface {
	GiveData() int
}

type RandomDataProducer struct {
}

func (r RandomDataProducer) GiveData() int {
	return rand.Intn(7)
}
