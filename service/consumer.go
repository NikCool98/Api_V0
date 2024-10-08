package service

import (
	"encoding/json"
	"fmt"
	"github.com/NickCool98/Api_V0/internal/storage"
	"os"

	"github.com/IBM/sarama"
	"go.uber.org/zap"
)

func ConnectConsumer(brokers []string) (sarama.Consumer, error) {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true

	return sarama.NewConsumer(brokers, config)
}

func Subscribe(c *storage.OrderCache, db *storage.OrdersRepository, logger zap.Logger, sigchan <-chan os.Signal) error {
	subject := "orders"

	work, err := ConnectConsumer([]string{"localhost:9092"})
	if err != nil {
		return fmt.Errorf("connection failed: %v", err)
	}

	consumer, err := work.ConsumePartition(subject, 0, sarama.OffsetNewest)
	if err != nil {
		return fmt.Errorf("consume partition failed: %v", err)
	}
	ch := make(chan struct{})
	go func() {
		for {
			select {
			case err := <-consumer.Errors():
				logger.Error("consuming error:", zap.Error(err))
			case msg := <-consumer.Messages():
				var order storage.Order
				if err := json.Unmarshal(msg.Value, &order); err != nil {
					logger.Error("unmarshal failed", zap.Error(err))
					continue
				}
				if _, found := c.GetOrd(order.OrderUID); found {
					logger.Info("order exist, didn't add")
					continue
				}
				if err := db.AddOrder(order); err != nil {
					logger.Error("failed to save order to DB", zap.Error(err))
					continue
				}
				c.SaveOrder(order)
				logger.Info("Consumed", zap.String("order_uid", order.OrderUID))
			case <-sigchan:
				ch <- struct{}{}
			}
		}
	}()
	logger.Info("Consumer successfully subscribed Kafka!")
	<-ch
	close(ch)
	if err := work.Close(); err != nil {
		return err
	}
	return nil
}
