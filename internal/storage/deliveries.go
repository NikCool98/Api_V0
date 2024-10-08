package storage

import (
	"database/sql"
	"errors"
	"fmt"
)

func GetDeliveries(db *sql.DB, OrderUID string) (*Deliveries, error) {
	query := "SELECT * FROM deliveries WHERE order_uid = $1"

	row := db.QueryRow(query, OrderUID)

	var d Deliveries
	var uid string
	err := row.Scan(
		&uid,
		&d.Name,
		&d.Phone,
		&d.Zip,
		&d.City,
		&d.Address,
		&d.Region,
		&d.Email,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("get delivery failed: %w", err)
		}
		return nil, fmt.Errorf("get delivery failed: %w", err)
	}

	return &d, nil
}

func AddDeliveries(tx *sql.Tx, d Deliveries, OrderUID string) error {
	query := `INSERT INTO deliveries ("name", "phone", "zip", "city", "address", "region", "email", "order_uid") VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`
	_, err := tx.Exec(
		query,
		d.Name,
		d.Phone,
		d.Zip,
		d.City,
		d.Address,
		d.Region,
		d.Email,
		OrderUID,
	)
	if err != nil {
		return err
	}
	return nil
}
