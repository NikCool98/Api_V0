package service

import (
	"encoding/json"
	"fmt"
	"github.com/IBM/sarama"
	"github.com/NickCool98/Api_V0/internal/config"
	"github.com/NickCool98/Api_V0/internal/storage"
	"github.com/NickCool98/Api_V0/volume"
	"log"
	"strconv"
)

var (
	cfgPath = "config/local.yaml"
)

func main() {
	cfg := config.MustLoad(cfgPath)
	subject := "orders"
	ordersRep, err := storage.ConnectBD(cfg)
	if err != nil {
		log.Fatal("Connection to DB failed", err)
	}
	defer ordersRep.DB.Close()
	orders, err := ordersRep.GetOrders()
	if err != nil {
		log.Fatalf("Failed to get orders from database: %v", err)
		return
	}
	for {
		log.Println("Type 's' to generate order")
		log.Println("Type 'c' to send copy")
		log.Println("Type 'exit' to quit")
		var input string
		var orderJSON []byte
		fmt.Scanln(&input)

		if input == "exit" {
			fmt.Println("Exiting the program...")
			break
		}
		if input == "s" {
			orderGenerated := volume.GenerateOrder()
			orderJSON, err = json.Marshal(orderGenerated)
			if err != nil {
				log.Printf("Failed to convert order to JSON: %s", err)
				continue
			}
		}
		if input == "c" {
			log.Println("Choose 1 of old orders:")
			for i := 0; i < len(orders); i++ {
				fmt.Println(i, orders[i].OrderUID)
			}
			var indstr string
			fmt.Scanln(&indstr)
			ind, err := strconv.Atoi(indstr)
			if err != nil {
				log.Println("Entered is not a number!")
				continue
			}
			if ind < 0 || ind > len(orders) {
				log.Println("Entered number isn't in range of orders!")
				continue
			}
			orderJSON, err = json.Marshal(orders[ind])
			if err != nil {
				log.Printf("Failed to convert order to JSON: %s", err)
				continue
			}
		}

		err = PushOrderToQueue(subject, orderJSON)
		if err != nil {
			log.Printf("Failed to send message to Kafka: %s", err)
			continue
		}

		log.Printf("Successfully sent order")
	}
}

func ConnecttoProducer(brok []string) (sarama.SyncProducer, error) {
	cfg := sarama.NewConfig()
	cfg.Producer.Return.Successes = true
	cfg.Producer.RequiredAcks = sarama.WaitForAll

	return sarama.NewSyncProducer(brok, cfg)
}

func PushOrderToQueue(topic string, message []byte) error {
	brokers := []string{"localhost:9092"}

	producer, err := ConnecttoProducer(brokers)
	if err != nil {
		return err
	}
	defer producer.Close()

	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(message),
	}

	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		return err
	}

	log.Printf("Order is stored in topic(%s)/partition(%d)/offset(%d)\n",
		topic,
		partition,
		offset,
	)

	return nil
}
