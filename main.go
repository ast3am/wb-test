package main

import (
	"context"
	"fmt"
	"github.com/example/internal/cache"
	"github.com/example/internal/db"
	"github.com/example/internal/nats_streaming"
	"github.com/example/internal/writer"
)

func main() {
	// Настройте соединение с сервером NATS Streaming
	channelName := "my-channel"
	clusterID := "test-cluster"        // ID кластера NATS Streaming
	clientID := "my-client2"           // ID клиента
	natsURL := "nats://localhost:4222" // URL сервера NATS Streaming

	usernameDB := "postgres"
	passwordDB := "password"
	hostDB := "localhost"
	portDB := "5432"
	database := "WB_db"

	/*del := models.Delivery{
		"Test Testov",
		"+9721111111",
		"2639809",
		"Kiryat Mozkin",
		"Ploshad Mira 15",
		"Kraiot",
		"test@gmail.com",
	}

	items := models.Items{
		99,
		"WBILMTESTTRACK",
		453,
		"ab4319087a764ae0btest",
		"Mascaras",
		30,
		"0",
		317,
		2389232,
		"Vivienne Sabo",
		202,
	}

	pay := models.Payment{
		"b563feb7b2b84b6test",
		"1",
		"USD",
		"wbpay",
		1817,
		1637907627,
		"alpha",
		1500,
		317,
		0,
	}

	layout := "2006-01-02 15:04:05 -0700 MST"
	dateString := "2021-11-26 06:22:19 +0000 UTC"

	t, err := time.Parse(layout, dateString)

	order := models.Order{
		"3",
		"WBILMTESTTRACK",
		"WBIL1",
		"en",
		"test",
		"test",
		"meest",
		"9",
		99,
		t,
		"1",
	}

	orders := models.Orders{
		order,
		del,
		pay,
		items,
	}*/

	ctx := context.Background()
	testDB, err := db.NewClient(ctx, usernameDB, passwordDB, hostDB, portDB, database)
	if err != nil {
		fmt.Printf("Ошибка подключения: %s\n", err)
	}
	println("im work")

	/*err = testDB.InsertOrders(ctx, orders)
	if err != nil {
		fmt.Printf("Ошибка записи: %s\n", err)
	}*/

	/*result, err := testDB.GetOrders(ctx)
	for _, k := range result {
		fmt.Printf("%+v\n\n", k)
	}*/

	cacheMap := cache.CacheInit()

	write := writer.NewWriter(testDB, cacheMap)
	err = write.GetCacheOnStart(ctx)
	if err != nil {
		fmt.Printf("Ошибка при заполнении кеша: %s\n", err)
	}

	/*for i := 1; i < 4; i++ {
		str := strconv.Itoa(i)
		fmt.Printf("%+v\n\n", cacheMap.GetOrdersById(str))
	}*/

	natsCon := nats_streaming.InitNats(write)

	err = natsCon.Connect(clusterID, clientID, natsURL)
	if err != nil {
		fmt.Printf("Ошибка  коннекта nats: %s\n", err)
	}
	defer natsCon.Close()

	err = natsCon.Subscribe(channelName)
	if err != nil {
		fmt.Printf("Ошибка подписки nats: %s\n", err)
	}

	select {}
}
