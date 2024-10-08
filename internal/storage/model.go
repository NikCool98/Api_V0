package storage

type Order struct {
	OrderUID          string     `json:"order_uid"`
	TrackNumber       string     `json:"track_number"`
	Entry             string     `json:"entry"`
	Delivery          Deliveries `json:"deliveries"`
	Payment           Payments   `json:"payment"`
	Items             []Items    `json:"items"`
	Locale            string     `json:"locale"`
	InternalSignature string     `json:"internal_signature"`
	CustomerID        string     `json:"customer_id"`
	DeliveryService   string     `json:"delivery_service"`
	ShardKey          string     `json:"shard_key"`
	SmID              int        `json:"sm_id"`
	DateCreated       string     `json:"date_created"`
	OofShard          string     `json:"oof_shard"`
}

type Deliveries struct {
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Zip     string `json:"zip"`
	City    string `json:"city"`
	Address string `json:"address"`
	Region  string `json:"region"`
	Email   string `json:"email"`
}

type Payments struct {
	Transaction   string `json:"transaction"`
	Request_id    string `json:"request_id"`
	Currency      string `json:"currency"`
	Provider      string `json:"provider"`
	Amount        int    `json:"amount"`
	Payment_dt    int    `json:"payment_dt"`
	Bank          string `json:"bank"`
	Delivery_cost int    `json:"delivery_cost"`
	Goods_total   int    `json:"goods_total"`
	Custom_fee    int    `json:"custom_fee"`
}

type Items struct {
	Chrt_id      int    `json:"chrt_id"`
	Track_number string `json:"track_number"`
	Price        int    `json:"price"`
	Rid          string `json:"rid"`
	Name         string `json:"name"`
	Sale         int    `json:"sale"`
	Size         string `json:"size"`
	Total_price  int    `json:"total_price"`
	Nm_id        int    `json:"nm_id"`
	Brand        string `json:"brand"`
	Status       int    `json:"status"`
}
