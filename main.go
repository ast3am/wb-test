package main

import (
	"context"
	"fmt"
	"github.com/example/api"
	"github.com/example/internal/cache"
	"github.com/example/internal/db"
	"github.com/example/internal/nats_streaming"
	"github.com/example/internal/writer"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
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

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	testDB, err := db.NewClient(ctx, usernameDB, passwordDB, hostDB, portDB, database)
	if err != nil {
		fmt.Printf("Ошибка подключения: %s\n", err)
		return
	}
	println("im work")

	cacheMap := cache.CacheInit()

	write := writer.NewWriter(testDB, cacheMap)
	err = write.GetCacheOnStart(ctx)
	if err != nil {
		fmt.Printf("Ошибка при заполнении кеша: %s\n", err)
		return
	}

	natsCon := nats_streaming.InitNats(write)

	err = natsCon.Connect(clusterID, clientID, natsURL)
	if err != nil {
		fmt.Printf("Ошибка  коннекта nats: %s\n", err)
		return
	}
	defer natsCon.Close()

	err = natsCon.Subscribe(channelName)
	if err != nil {
		fmt.Printf("Ошибка подписки nats: %s\n", err)
		return
	}

	router := mux.NewRouter()
	handler := api.NewHandler(cacheMap)
	handler.Register(router)
	http.Handle("/", router)

	shutdownSignal := make(chan os.Signal, 1)
	signal.Notify(shutdownSignal, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-shutdownSignal
		fmt.Printf("Получен сигнал завершения: %v\n", sig)
		natsCon.Close()
		testDB.Close(ctx)
		time.Sleep(2 * time.Second)
		os.Exit(0)
	}()

	err = http.ListenAndServe(":8080", router)
	if err != nil {
		return
	}

}
