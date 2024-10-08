package storage

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/NickCool98/Api_V0/internal/config"
	_ "github.com/lib/pq"
)

type OrdersRepository struct {
	DB *sql.DB
}

func ConnectBD(cfg *config.Config) (*OrdersRepository, error) {
	//Строка подключения к БД
	connString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DB.Host,
		cfg.DB.Port,
		cfg.DB.User,
		cfg.DB.Password,
		cfg.DB.Name,
	)
	db, err := sql.Open(cfg.DB.Schema, connString)
	if err != nil {
		return nil, err
	}
	ordersDB := &OrdersRepository{
		DB: db,
	}
	return ordersDB, nil

}

func (q *OrdersRepository) AddOrder(o Order) error {
	//Begin starts a transaction
	tx, err := q.DB.Begin()
	if err != nil {
		return fmt.Errorf("failed transaction: %v", err)
	}
	query := `INSERT INTO orders("order_uid", "track_number","entry", "locale", "internal_signature", "customer_id", "delivery_service", "shardkey", "sm_id", "date_created", "oof_shard") 
	VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`
	_, err = q.DB.Exec(
		query,
		o.OrderUID,
		o.TrackNumber,
		o.Entry,
		o.Locale,
		o.InternalSignature,
		o.CustomerID,
		o.DeliveryService,
		o.ShardKey,
		o.SmID,
		o.DateCreated,
		o.OofShard,
	)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("Insert failed: %v", err)
	}
	err = AddDeliveries(tx, o.Delivery, o.OrderUID)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("Delivery failed insert: %v", err)
	}
	err = AddPayments(tx, o.Payment, o.OrderUID)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("payment failed: %v", err)
	}
	err = AddItems(tx, o.Items, o.OrderUID)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("Insert items failed: %v", err)
	}
	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("failed transaction: %v", err)
	}
	return nil
}
func (o *OrdersRepository) GetOrder(OrderUID string) (*Order, error) {
	query := "SELECT * FROM orders WHERE order_uid = $1"
	row := o.DB.QueryRow(query, OrderUID)
	var or Order
	err := row.Scan(
		&or.OrderUID,
		&or.TrackNumber,
		&or.Entry,
		&or.Locale,
		&or.InternalSignature,
		&or.CustomerID,
		&or.DeliveryService,
		&or.ShardKey,
		&or.SmID,
		&or.DateCreated,
		&or.OofShard,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to get orders: %v", err)
	}
	payment, err := GetPayments(o.DB, OrderUID)
	if err != nil {
		return nil, err
	}
	or.Payment = *payment

	delivery, err := GetDeliveries(o.DB, OrderUID)
	if err != nil {
		return nil, err
	}
	or.Delivery = *delivery

	items, err := GetItems(o.DB, OrderUID)
	if err != nil {
		return nil, err
	}
	or.Items = items

	return &or, nil
}

func (o *OrdersRepository) GetOrders() ([]Order, error) {
	query := "SELECT * FROM orders"
	rows, err := o.DB.Query(query)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to get orders: %v", err)
	}
	defer rows.Close()

	var orders []Order
	for rows.Next() {
		var order Order
		err := rows.Scan(&order.OrderUID, &order.TrackNumber, &order.Entry, &order.Locale, &order.InternalSignature, &order.CustomerID, &order.DeliveryService, &order.ShardKey, &order.SmID, &order.DateCreated, &order.OofShard)
		if err != nil {
			return nil, fmt.Errorf("failed to scan order row: %v", err)
		}

		delivery, err := GetDeliveries(o.DB, order.OrderUID)
		if err != nil {
			return nil, fmt.Errorf("failed to get delivery for order %s: %v", order.OrderUID, err)
		}
		order.Delivery = *delivery

		payment, err := GetPayments(o.DB, order.OrderUID)
		if err != nil {
			return nil, fmt.Errorf("failed to get payment for order %s: %v", order.OrderUID, err)
		}
		order.Payment = *payment

		items, err := GetItems(o.DB, order.OrderUID)
		if err != nil {
			return nil, fmt.Errorf("failed to get items for order %s: %v", order.OrderUID, err)
		}
		order.Items = items

		orders = append(orders, order)
	}

	return orders, nil
}
