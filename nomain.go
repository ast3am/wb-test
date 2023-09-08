package main

import (
	"github.com/nats-io/stan.go"
	"log"
)

func main() {

	json := `{
		"order_uid": "7",
		"track_number": "WBILMTESTTRACK",
		"entry": "WBILl",
		"delivery": {
			"name": "Test Testov",
			"phone": "+79211552101",
			"zip": "2639809",
			"city": "Kiryat Mozkin",
			"address": "Ploshad Mira 15",
			"region": "Kraiot",
			"email": "test@gmail.com"
		},
		"payment": {
			"transaction": "b563feb7b2b84b6test",
			"request_id": "hhjjh",
			"currency": "USD",
			"provider": "wbpay",
			"amount": 1817,
			"payment_dt": 1637907627,
			"bank": "alpha",
			"delivery_cost": 1500,
			"goods_total": 317,
			"custom_fee": 0
		},
		"items": [
			{
				"chrt_id": 9934922,
				"track_number": "WBILMTESTTRACK",
				"price": 453,
				"rid": "ab4319087a764ae0btest",
				"name": "Mascaras",
				"sale": 30,
				"size": "0",
				"total_price": 317,
				"nm_id": 2389232,
				"brand": "Vivienne Sabo",
				"status": 202
			}
		],
		"locale": "en",
		"internal_signature": "sad",
		"customer_id": "test",
		"delivery_service": "meest",
		"shardkey": "9",
		"sm_id": 99,
		"date_created": "2021-11-26T06:22:19Z",
		"oof_shard": "1"
	}`

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
	message := []byte(json)

	// Отправка сообщения в канал
	err = sc.Publish(channelName, message)
	if err != nil {
		log.Fatalf("Ошибка отправки сообщения: %v", err)
	}

	log.Printf("Сообщение успешно отправлено в канал %s", channelName)
}
