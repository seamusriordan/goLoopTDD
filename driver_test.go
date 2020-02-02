package main

import (
	"./dataWriterClient"
	"./dataproducer"
	"fmt"
	"testing"
	"time"
)

func Test_program_calls_connect_once(t *testing.T) {
	client := MockClient{}
	dataProducer := dataproducer.RandomDataProducer{}

	go Program(&client, &dataProducer)
	time.Sleep(1 * time.Microsecond)

	if client.nConnectCalls != 1 {
		t.Errorf("connect not called")
	}
}

func Test_program_calls_GiveData_AtLeastOnce(t *testing.T) {
	client := dataWriterClient.StubClient{}
	dataProducer := MockDataProducer{}

	go Program(&client, &dataProducer)
	time.Sleep(1 * time.Microsecond)

	if dataProducer.nGiveDataCalls == 0 {
		t.Errorf("GiveData never called")
	}
}

func Test_program_calls_SendData_WithDataProducerValue1(t *testing.T) {
	client := MockClient{}
	valueToSend := 1
	dataProducer := MockDataProducer{value: valueToSend}

	go Program(&client, &dataProducer)
	time.Sleep(1 * time.Microsecond)

	if client.dataSent != valueToSend {
		t.Errorf("SendData not called with correct value")
	}
}

func Test_program_calls_SendData_WithDataProducerValue4(t *testing.T) {
	client := MockClient{}
	valueToSend := 4
	dataProducer := MockDataProducer{value: valueToSend}

	go Program(&client, &dataProducer)
	time.Sleep(1 * time.Microsecond)

	if client.dataSent != valueToSend {
		t.Errorf("SendData not called with correct value")
	}
}

func Test_program_calls_SendData_more_times_after_10ms(t *testing.T) {
	client := MockClient{}
	dataProducer := MockDataProducer{}

	go Program(&client, &dataProducer)
	time.Sleep(1 * time.Microsecond)
	nCallsBefore := client.nSendDataCalls
	time.Sleep(10 * time.Millisecond)
	nCallsAfter := client.nSendDataCalls

	fmt.Printf("calls %d -> %d\n", nCallsBefore, nCallsAfter)
	if nCallsBefore == nCallsAfter {
		t.Errorf("SendData not called repeatedly")
	}
}

func Test_program_calls_SendData_more_times_from_10ms_to_100ms(t *testing.T) {
	client := MockClient{}
	dataProducer := MockDataProducer{}

	go Program(&client, &dataProducer)
	time.Sleep(10 * time.Millisecond)
	nCallsBefore := client.nSendDataCalls

	time.Sleep(100 * time.Millisecond)
	nCallsAfter := client.nSendDataCalls

	fmt.Printf("calls %d -> %d\n", nCallsBefore, nCallsAfter)
	if nCallsBefore == nCallsAfter {
		t.Errorf("SendData not called repeatedly")
	}
}

type MockClient struct {
	nConnectCalls  int
	nSendDataCalls int
	dataSent       int
}

func (m *MockClient) Connect(_ string) {
	m.nConnectCalls++
}

func (m *MockClient) SendData(data int) {
	m.dataSent = data
	m.nSendDataCalls++
}

type MockDataProducer struct {
	nGiveDataCalls int
	value          int
}

func (m *MockDataProducer) GiveData() int {
	m.nGiveDataCalls++
	return m.value
}
