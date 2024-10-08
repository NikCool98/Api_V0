package example

import (
	"encoding/json"
	"github.com/NickCool98/Api_V0/internal/storage"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"net/http"
	"time"
)

var Validate = validator.New()

func WriteJson(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application-json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

func WriteErr(w http.ResponseWriter, status int, err error) {
	WriteJson(w, status, map[string]string{"error": err.Error()})
}

func GenerateOrder() storage.Orders {
	return storage.Orders{
		Order_uid:    uuid.NewString(),
		Track_number: "WBILMTESTTRACK",
		Entry:        "WBIL",
		Delivery: storage.Deliveries{
			Name:    "Test Testov",
			Phone:   "+9720000000",
			Zip:     "2639809",
			City:    "Kiryat Mozkin",
			Address: "Ploshad Mira 15",
			Region:  "Kraiot",
			Email:   "test@gmail.com",
		},
		Payment: storage.Payments{
			Transaction:   uuid.NewString(),
			Request_id:    "",
			Currency:      "RUB",
			Provider:      "wbpay",
			Amount:        534,
			Payment_dt:    9239123,
			Delivery_cost: 1500,
			Bank:          "alpha",
			Goods_total:   317,
			Custom_fee:    0,
		},
		Items: []storage.Items{
			{
				Chrt_id:      9934930,
				Track_number: "WBILMTESTTRACK",
				Price:        453,
				Rid:          "ab4219087a764ae0btest",
				Name:         "Maracas",
				Sale:         30,
				Size:         "0",
				Total_price:  332,
				Nm_id:        2389212,
				Brand:        "Vivienne Sabo",
				Status:       202,
			},
		},
		Locale:            "ru",
		InternalSignature: "",
		CustomerID:        "test",
		DeliveryService:   "meest",
		ShardKey:          "9",
		SmID:              99,
		DateCreated:       time.Now(),
		OofShard:          "1",
	}
}
