package main

import (
	"github.com/NickCool98/Api_V0/handlers"
	config2 "github.com/NickCool98/Api_V0/internal/config"
	"github.com/NickCool98/Api_V0/internal/storage"
	"github.com/NickCool98/Api_V0/service"
	"go.uber.org/zap"
	"log"
	"os"
	"os/signal"
)

var (
	cfgPath = "config/local.yaml"
)

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatal(err)
	}
	defer logger.Sync()

	logger.Info("Starting application")

	cfg := config2.MustLoad(cfgPath)

	//connect to database
	storageRep, err := storage.ConnectBD(cfg)
	if err != nil {
		log.Fatalf("Failed to connect database: %v", err)
		return
	}
	defer storageRep.DB.Close()
	logger.Info(
		"DB connected successfully",
		zap.String("host", cfg.DB.Host),
		zap.String("port", cfg.DB.Port),
		zap.String("db", cfg.DB.Name),
		zap.String("user", cfg.DB.User),
	)

	cache := storage.NewCache()

	//Get orders from database and pull in cache
	orders, err := storageRep.GetOrders()
	if err != nil {
		logger.Error("Failed to refill cache from database", zap.Error(err))
		return
	}
	for _, order := range orders {
		cache.SaveOrder(order)
		logger.Info("Cached", zap.String("order_uid", order.OrderUID))
	}

	start := handlers.New(cfgPath, cache)
	go func() {
		start.Launch()
	}()
	logger.Info(
		"Server is starting",
		zap.String("host", cfg.HTTPServer.Address),
		zap.String("port", cfg.HTTPServer.Port),
	)
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, os.Interrupt)
	//connect consumer
	err = service.Subscribe(cache, storageRep, *logger, sigchan)
	if err != nil {
		logger.Fatal("consumer error:", zap.Error(err))
	}
	logger.Info("Application shutdown")
}
