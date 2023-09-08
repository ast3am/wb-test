package nats_streaming

import (
	"github.com/nats-io/stan.go"
	"log"
)

type Writer interface {
	Write(text []byte) error
}

type Nats struct {
	NatsConnection stan.Conn
	writer         Writer
}

func InitNats(w Writer) *Nats {
	Nats := Nats{
		writer: w,
	}
	return &Nats
}

func (n *Nats) Connect(clusterID, clientID, natsURL string) error {
	NatsConnection, err := stan.Connect(clusterID, clientID, stan.NatsURL(natsURL))
	if err != nil {
		return err
	}
	n.NatsConnection = NatsConnection
	return err
}

func (n *Nats) Close() {
	n.NatsConnection.Close()
}

func (n *Nats) Subscribe(channelName string) error {
	stanHandler := func(msg *stan.Msg) {
		err := n.writer.Write(msg.Data)
		if err != nil {
			log.Printf("Ошибка при записи: %v", err)
		}
	}

	_, err := n.NatsConnection.Subscribe(channelName, stanHandler, stan.DurableName(channelName))
	if err != nil {
		return err
	}

	return nil
}
