package order

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

func canPurchase(db *sql.DB, albumID, quantity int) (bool, error) {
	var enough bool
	sqlEnough := "select (quantity >= ?) from album where id=?"
	if err := db.QueryRow(sqlEnough,
		quantity, albumID).Scan(&enough); err != nil {
		if err == sql.ErrNoRows {
			return false, fmt.Errorf("canPurchase %d: unknown album", albumID)
		}
		return false, fmt.Errorf("canPurchase %d: %v", albumID, err)
	}
	return enough, nil
}

// CreateOrder create an order for an album
// and return the new order ID
func CreateOrder(db *sql.DB, ctx context.Context, albumID, quantity, customerID int) (orderID int64, err error) {
	// create a helper function for preparing failure results
	fail := func(err error) (int64, error) {
		return 0, fmt.Errorf("CreateOrder: %v", err)
	}

	// get a Tx for making transaction requests
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return fail(err)
	}
	// defer a rollback in case anything fails
	defer tx.Rollback()

	// confirm that album inventory is enough for the order
	enough, err := canPurchase(db, albumID, quantity)
	if err != nil {
		return fail(err)
	}
	if !enough {
		return fail(fmt.Errorf("not enough inventory"))
	}

	// create a new row in the album_order table
	sqlInsert := "insert into album_order " +
		"(album_id, customer_id, quantity, date) " +
		"values (?, ?, ?, ?)"
	result, err := tx.ExecContext(ctx, sqlInsert,
		albumID, customerID, quantity, time.Now())
	if err != nil {
		return fail(err)
	}
	// get the ID of the order item just created
	orderID, err = result.LastInsertId()
	if err != nil {
		return fail(err)
	}

	// 这里确实会回滚
	// 更新失败，order表也不会有记录
	// update the album inventory to remove the quantity in the order
	sqlUpdate := "update album set quantity=quantity-? where id=?"
	_, err = tx.ExecContext(ctx, sqlUpdate,
		quantity, "albumID")
	if err != nil {
		return fail(err)
	}

	// commit the transaction
	if err = tx.Commit(); err != nil {
		return fail(err)
	}

	// all done
	return
}
