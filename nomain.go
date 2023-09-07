package main

import (
	"github.com/nats-io/stan.go"
	"log"
)

func main() {
	channelName := "my-channel"
	clusterID := "test-cluster"        // ID кластера NATS Streaming
	clientID := "my-client3"           // ID клиента
	natsURL := "nats://localhost:4222" // URL сервера NATS Streaming

	sc, err := stan.Connect(clusterID, clientID, stan.NatsURL(natsURL))
	if err != nil {
		log.Fatalf("Ошибка подключения к серверу NATS Streaming: %v", err)
	}
	defer sc.Close()

	// Сообщение, которое вы хотите отправить
	message := []byte("Ваше сообщение здесь")

	// Отправка сообщения в канал
	err = sc.Publish(channelName, message)
	if err != nil {
		log.Fatalf("Ошибка отправки сообщения: %v", err)
	}

	log.Printf("Сообщение успешно отправлено в канал %s", channelName)
}
