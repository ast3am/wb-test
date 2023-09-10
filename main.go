package main

import (
	"context"
	"github.com/example/api"
	"github.com/example/internal/cache"
	"github.com/example/internal/config"
	"github.com/example/internal/db"
	"github.com/example/internal/nats_streaming"
	"github.com/example/internal/writer"
	"github.com/example/pkg/logging"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	cfg := config.GetConfig("config.yml")
	log := logging.GetLogger(cfg.LogLevel)

	testDB, err := db.NewClient(ctx, cfg, log)
	if err != nil {
		log.FatalMsg("", err)
	}

	cacheMap := cache.CacheInit(log)

	write := writer.NewWriter(testDB, cacheMap, log)
	err = write.GetCacheOnStart(ctx)
	if err != nil {
		log.FatalMsg("сache generation error", err)
	}

	natsCon := nats_streaming.InitNats(write, log)
	err = natsCon.Connect(cfg)
	if err != nil {
		log.FatalMsg("connect error", err)
	}
	defer natsCon.Close()

	err = natsCon.Subscribe(cfg.NatsConfig.ChannelName)
	if err != nil {
		log.FatalMsg("subscribe error", err)
	}

	router := mux.NewRouter()
	handler := api.NewHandler(cacheMap, log)
	handler.Register(router)
	http.Handle("/", router)

	shutdownSignal := make(chan os.Signal, 1)
	signal.Notify(shutdownSignal, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-shutdownSignal
		log.InfoMsgf("termination signal received: %v", sig)
		natsCon.Close()
		testDB.Close(ctx)
		time.Sleep(2 * time.Second)
		os.Exit(0)
	}()

	log.InfoMsg("service start")

	err = http.ListenAndServe(":"+cfg.ListenPort, router)
	if err != nil {
		log.FatalMsg("r", err)
	}
}
