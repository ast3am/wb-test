package main

import (
	"fmt"
	"github.com/example/internal/nats-streaming"
)

func main() {
	// Настройте соединение с сервером NATS Streaming
	channelName := "my-channel"
	clusterID := "test-cluster"        // ID кластера NATS Streaming
	clientID := "my-client2"           // ID клиента
	natsURL := "nats://localhost:4222" // URL сервера NATS Streaming

	msgCh := nats_streaming.Subscribe(channelName, clientID, clusterID, natsURL)
	fmt.Printf("Получено сообщение: %s\n", msgCh)
	select {}
}
