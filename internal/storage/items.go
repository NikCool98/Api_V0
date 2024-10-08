package storage

import (
	"database/sql"
	"errors"
	"fmt"
)

func AddItem(tx *sql.Tx, i Items, OrderUID string) error {
	query := `INSERT INTO "items"("chrt_id", "track_number", "price", "rid", "name", "sale", "size", "total_price", "nm_id", "brand", "status", "order_uid") VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)`
	_, err := tx.Exec(
		query,
		i.Chrt_id,
		i.Track_number,
		i.Price,
		i.Rid,
		i.Name,
		i.Sale,
		i.Size,
		i.Total_price,
		i.Nm_id,
		i.Brand,
		i.Status,
		OrderUID,
	)
	if err != nil {
		return err
	}
	return nil
}

func AddItems(tx *sql.Tx, i []Items, OrderUID string) (err error) {
	for _, v := range i {
		err = AddItem(tx, v, OrderUID)
		if err != nil {
			return err
		}
	}
	return nil
}

func GetItems(db *sql.DB, OrderUID string) ([]Items, error) {
	query := "SELECT * FROM items WHERE order_uid = $1"
	rows, err := db.Query(query, OrderUID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("items not found: %v", err)
		}
		return nil, fmt.Errorf("items failed: %v", err)
	}
	var items []Items
	var uid string
	for rows.Next() {
		var item Items
		err := rows.Scan(
			&uid,
			&item.Chrt_id,
			&item.Track_number,
			&item.Price,
			&item.Rid,
			&item.Name,
			&item.Sale,
			&item.Size,
			&item.Total_price,
			&item.Nm_id,
			&item.Brand,
			&item.Status,
		)
		if err != nil {
			return nil, fmt.Errorf("Scan failed: %v", err)
		}
		items = append(items, item)
	}
	return items, nil
}
