package storage

import (
	"database/sql"
	"errors"
	"fmt"
)

func AddPayments(tx *sql.Tx, p Payments, OrderUID string) error {
	query := `INSERT INTO payments ("transaction", "request_id", "currency", "provider", "amount", "payment_dt", "bank", "delivery_cost", "goods_total", "custom_fee", "order_uid") VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`
	_, err := tx.Exec(
		query,
		p.Transaction,
		p.Request_id,
		p.Currency,
		p.Provider,
		p.Amount,
		p.Payment_dt,
		p.Bank,
		p.Delivery_cost,
		p.Goods_total,
		p.Custom_fee,
		OrderUID,
	)
	if err != nil {
		return err
	}
	return nil
}

func GetPayments(db *sql.DB, OrderUID string) (*Payments, error) {
	query := "SELECT * FROM payments WHERE order_uid = $1"
	row := db.QueryRow(query, OrderUID)
	var pay Payments
	var uid string
	err := row.Scan(
		&uid,
		&pay.Transaction,
		&pay.Request_id,
		&pay.Currency,
		&pay.Provider,
		&pay.Amount,
		&pay.Payment_dt,
		&pay.Bank,
		&pay.Delivery_cost,
		&pay.Goods_total,
		&pay.Custom_fee,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("payment not found: %v", err)
		}
		return nil, fmt.Errorf("payment failed: %v", err)
	}
	return &pay, nil
}
