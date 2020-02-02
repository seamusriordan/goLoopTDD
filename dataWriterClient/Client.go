package dataWriterClient

import "fmt"

type Client interface {
	Connect(url string)
	SendData(data int)
}

type StubClient struct {
	Url    string
	NCalls int64
}

func (s *StubClient) Connect(url string) {
	fmt.Printf("Connected to %s\n", url)
}

func (s *StubClient) SendData(data int) {
	fmt.Printf("Data sent: %d\n", data)
}
