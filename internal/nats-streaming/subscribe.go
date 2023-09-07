package nats_streaming

import (
	"fmt"
	"github.com/nats-io/stan.go"
	"log"
	"sync"
)

func Subscribe(channelName, clientID, clusterID, natsURL string) <-chan *stan.Msg {
	msgCh := make(chan *stan.Msg)

	nc, err := stan.Connect(clusterID, clientID, stan.NatsURL(natsURL))
	if err != nil {
		log.Fatalf("Ошибка подключения к серверу NATS Streaming: %v", err)
	}
	defer nc.Close()

	subscription, err := nc.Subscribe(channelName, func(msg *stan.Msg) {
		msgCh <- msg
	})
	if err != nil {
		log.Fatalf("Ошибка подписки на канал: %v", err)
	}

	fmt.Printf("Подписка на канал %s. Ожидание сообщений...\n", channelName)

	// Создайте WaitGroup для ожидания завершения работы
	var wg sync.WaitGroup
	wg.Add(1)

	// Запустите горутину для graceful shutdown
	go func() {
		defer wg.Done()

		// Ожидание завершения работы (можете добавить логику завершения, например, по сигналу OS)
		select {}
	}()

	// Завершение работы и отписка при получении сигнала (например, Ctrl+C)
	go func() {
		defer subscription.Close()
		<-someSignalChannel // Замените на канал для сигнала завершения
	}()

	// Ожидание завершения работы
	wg.Wait()

	return msgCh
}
