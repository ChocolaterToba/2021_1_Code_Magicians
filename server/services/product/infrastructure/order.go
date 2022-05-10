package repository

import (
	"context"
	"pinterest/services/product/domain"

	"github.com/jackc/pgx/v4"
)

func (repo *ProductRepo) CreateOrder(ctx context.Context, order domain.Order) (id uint64, err error) {
	tx, err := repo.postgresDB.Begin(ctx)
	if err != nil {
		return 0, domain.TransactionBeginError
	}
	defer tx.Rollback(ctx)

	createOrderQuery := `INSERT INTO orders (user_id, items, created_at, total_price, 
						 pick_up, delivery_address, payment_method, call_needed, status)
						 VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
						 RETURNING id`

	row := tx.QueryRow(ctx, createOrderQuery,
		order.UserID, cusjsonb(order.Items), order.CreatedAt, order.TotalPrice,
		order.PickUp, order.DeliveryAddress, order.PaymentMethod, order.CallNeeded, order.Status)
	err = row.Scan(&id)
	if err != nil {
		return 0, err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return 0, domain.TransactionCommitError
	}
	return id, nil
}

func (repo *ProductRepo) GetOrderByID(ctx context.Context, id uint64) (order domain.Order, err error) {
	tx, err := repo.postgresDB.Begin(ctx)
	if err != nil {
		return domain.Order{}, domain.TransactionBeginError
	}
	defer tx.Rollback(ctx)

	getOrderByIDQuery := `SELECT id, user_id, items, created_at, total_price, pick_up, 
						  delivery_address, payment_method, call_needed, status
						  FROM orders
						  WHERE id = $1`

	row := tx.QueryRow(ctx, getOrderByIDQuery, id)

	dbItems := make(cusjsonb)
	err = row.Scan(&order.Id, &order.UserID, &dbItems, &order.CreatedAt, &order.TotalPrice, &order.PickUp,
		&order.DeliveryAddress, &order.PaymentMethod, &order.CallNeeded, &order.Status)
	order.Items = dbItems
	if err != nil {
		if err == pgx.ErrNoRows {
			return domain.Order{}, domain.CartNotFoundError
		}

		return domain.Order{}, err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return domain.Order{}, domain.TransactionCommitError
	}
	return order, nil
}

func (repo *ProductRepo) GetOrdersByUserID(ctx context.Context, userID uint64) (orders []domain.Order, err error) {
	tx, err := repo.postgresDB.Begin(ctx)
	if err != nil {
		return nil, domain.TransactionBeginError
	}
	defer tx.Rollback(ctx)

	getOrdersByUserIDQuery := `SELECT id, user_id, items, created_at, total_price, pick_up, 
							   delivery_address, payment_method, call_needed, status
							   FROM orders
							   WHERE user_id = $1`

	rows, err := tx.Query(ctx, getOrdersByUserIDQuery, userID)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var order domain.Order
		dbItems := make(cusjsonb)
		err = rows.Scan(&order.Id, &order.UserID, &dbItems, &order.CreatedAt, &order.TotalPrice, &order.PickUp,
			&order.DeliveryAddress, &order.PaymentMethod, &order.CallNeeded, &order.Status)
		order.Items = dbItems
		if err != nil {
			return nil, err
		}

		orders = append(orders, order)
	}

	err = tx.Commit(ctx)
	if err != nil {
		return nil, domain.TransactionCommitError
	}
	return orders, nil
}
