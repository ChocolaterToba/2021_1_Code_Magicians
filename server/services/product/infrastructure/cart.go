package repository

import (
	"context"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"pinterest/services/product/domain"
	"strconv"

	"github.com/huandu/go-sqlbuilder"
	"github.com/jackc/pgx/v4"
)

func (repo *ProductRepo) CreateCart(ctx context.Context, userID uint64) (id uint64, err error) {
	tx, err := repo.postgresDB.Begin(ctx)
	if err != nil {
		return 0, domain.TransactionBeginError
	}
	defer tx.Rollback(ctx)

	createCartQuery := `INSERT INTO carts (user_id)
						VALUES ($1)
						RETURNING id`

	row := tx.QueryRow(ctx, createCartQuery, userID)
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

func (repo *ProductRepo) GetCart(ctx context.Context, userID uint64) (cart map[uint64]uint64, err error) {
	tx, err := repo.postgresDB.Begin(ctx)
	if err != nil {
		return nil, domain.TransactionBeginError
	}
	defer tx.Rollback(ctx)

	sb := sqlbuilder.Select("product_ids").
		From("carts")

	sb.Where(sb.Equal("user_id", userID))

	query, args := sb.BuildWithFlavor(sqlbuilder.PostgreSQL)

	row := tx.QueryRow(ctx, query, args...)

	dbCart := make(cusjsonb)
	err = row.Scan(&dbCart)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, domain.CartNotFoundError
		}

		return nil, err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return nil, domain.TransactionCommitError
	}
	return dbCart, nil
}

func (repo *ProductRepo) UpdateCart(ctx context.Context, userID uint64, cart map[uint64]uint64) (err error) {
	tx, err := repo.postgresDB.Begin(ctx)
	if err != nil {
		return domain.TransactionBeginError
	}
	defer tx.Rollback(ctx)

	updateCartQuery := `UPDATE carts
						SET product_ids = $2
						WHERE user_id = $1`

	result, err := tx.Exec(ctx, updateCartQuery, userID, cusjsonb(cart))
	if err != nil {
		return err
	}

	if result.RowsAffected() != 1 {
		return domain.CartNotFoundError
	}

	err = tx.Commit(ctx)
	if err != nil {
		return domain.TransactionCommitError
	}
	return nil
}

type cusjsonb map[uint64]uint64

// Returns the JSON-encoded representation
func (a cusjsonb) Value() (driver.Value, error) {
	// Convert to map[string]float32 from map[int]float32
	x := make(map[string]uint64)
	for k, v := range a {
		x[strconv.FormatUint(k, 10)] = v
	}
	// Marshal into json
	return json.Marshal(x)
}

// Decodes a JSON-encoded value
func (a *cusjsonb) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	// Unmarshal from json to map[string]float32
	x := make(map[string]uint64)
	if err := json.Unmarshal(b, &x); err != nil {
		return err
	}
	// Convert to map[int]float32 from  map[string]float32
	*a = make(cusjsonb, len(x))
	for k, v := range x {
		if ki, err := strconv.ParseUint(k, 10, 64); err != nil {
			return err
		} else {
			(*a)[ki] = v
		}
	}
	return nil
}
