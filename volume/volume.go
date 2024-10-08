package volume

import (
	"fmt"
	"github.com/NickCool98/Api_V0/internal/storage"
	"math/rand"
	"time"
)

func GenerateOrder() storage.Order {
	delivery := storage.Deliveries{
		Name:    randomString(10),
		Phone:   randomPhone(),
		Zip:     randomZip(),
		City:    randomString(8),
		Address: randomString(15),
		Region:  randomString(8),
		Email:   randomString(5),
	}

	item := storage.Items{
		Chrt_id:      rand.Intn(1000),
		Track_number: randomString(10),
		Price:        rand.Intn(1000),
		Rid:          randomString(6),
		Name:         randomString(10),
		Sale:         rand.Intn(100),
		Size:         randomSize(),
		Total_price:  rand.Intn(1000),
		Nm_id:        rand.Intn(1000),
		Brand:        randomString(8),
		Status:       rand.Intn(5),
	}

	currencies := []string{"USD", "RUB", "EUR"}
	currency := currencies[rand.Intn(len(currencies))]

	payment := storage.Payments{
		Transaction:   randomString(10),
		Request_id:    randomString(8),
		Currency:      currency,
		Provider:      randomString(6),
		Amount:        rand.Intn(10000),
		Payment_dt:    int(time.Now().Unix()),
		Bank:          randomString(6),
		Delivery_cost: rand.Intn(500),
		Goods_total:   rand.Intn(10000),
		Custom_fee:    rand.Intn(100),
	}

	localies := []string{"en", "ru"}
	locale := localies[rand.Intn(len(localies))]
	order := storage.Order{
		OrderUID:          randomString(12),
		TrackNumber:       randomString(10),
		Entry:             randomString(5),
		Delivery:          delivery,
		Payment:           payment,
		Items:             []storage.Items{item},
		Locale:            locale,
		InternalSignature: randomString(8),
		CustomerID:        randomString(8),
		DeliveryService:   randomString(5),
		ShardKey:          randomString(5),
		SmID:              rand.Intn(100),
		DateCreated:       time.Now().Format("2006-01-02"),
		OofShard:          randomString(4),
	}
	return order
}

func randomString(length int) string {
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	res := make([]byte, length)
	for i := range res {
		res[i] = charset[rand.Intn(len(charset))]
	}
	return string(res)
}

// random size
func randomSize() string {
	sizes := []string{"S", "M", "L", "XL", "XXL"}
	return sizes[rand.Intn(len(sizes))]
}

// random zip
func randomZip() string {
	return fmt.Sprintf("%05d", rand.Intn(100000))
}

// random phone
func randomPhone() string {
	return fmt.Sprintf("+1%010d", rand.Int63n(10000000000))
}
