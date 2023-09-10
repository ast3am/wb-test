package nats_streaming

import (
	"github.com/example/internal/models"
	"github.com/nats-io/stan.go"
)

type Writer interface {
	Write(text []byte) error
}

type logger interface {
	InfoMsgf(format string, v ...interface{})
	DebugMsg(msg string)
	ErrorMsg(msg string, err error)
}

type Nats struct {
	NatsConnection stan.Conn
	writer         Writer
	log            logger
}

func InitNats(w Writer, log logger) *Nats {
	Nats := Nats{
		writer: w,
		log:    log,
	}
	return &Nats
}

func (n *Nats) Connect(cfg *models.Config) error {
	natsURL := "nats://" + cfg.NatsConfig.Host + ":" + cfg.NatsConfig.Port
	NatsConnection, err := stan.Connect(cfg.NatsConfig.ClusterID, cfg.NatsConfig.ClientID, stan.NatsURL(natsURL))
	if err != nil {
		return err
	}
	n.log.DebugMsg("connect to nuts is OK")
	n.NatsConnection = NatsConnection
	return nil
}

func (n *Nats) Close() {
	err := n.NatsConnection.Close()
	if err != nil {
		n.log.ErrorMsg("error closing connection to nuts", err)
	} else {
		n.log.DebugMsg("closing connection to nuts is OK")
	}
}

func (n *Nats) Subscribe(channelName string) error {
	stanHandler := func(msg *stan.Msg) {
		err := n.writer.Write(msg.Data)
		if err != nil {
			n.log.ErrorMsg("", err)
		}
	}

	_, err := n.NatsConnection.Subscribe(channelName, stanHandler, stan.DurableName(channelName))
	if err != nil {
		return err
	}
	n.log.DebugMsg("subscribe to channel " + channelName + " is OK")
	return nil
}
