package dataWriterClient

type Client interface {
	Connect(url string)
	SendData(data int)
}

type StubClient struct {
	Url    string
	NCalls int64
}

func (s *StubClient) Connect(_ string) {
}

func (s *StubClient) SendData(_ int) {
}
