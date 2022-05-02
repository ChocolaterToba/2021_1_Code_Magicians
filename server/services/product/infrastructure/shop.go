package repository

import (
	"context"
	"pinterest/services/product/domain"

	"github.com/jackc/pgx/v4"
	"github.com/lib/pq"
)

func (repo *ProductRepo) CreateShop(ctx context.Context, shop domain.Shop) (id uint64, err error) {
	tx, err := repo.postgresDB.Begin(ctx)
	if err != nil {
		return 0, domain.TransactionBeginError
	}
	defer tx.Rollback(ctx)

	createShopQuery := `INSERT INTO shops (title, description, manager_ids)
						VALUES ($1, $2, $3)
						RETURNING id`

	row := tx.QueryRow(ctx, createShopQuery, shop.Title, shop.Description, pq.Array(shop.ManagerIDs))
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

func (repo *ProductRepo) UpdateShop(ctx context.Context, shop domain.Shop) (err error) {
	tx, err := repo.postgresDB.Begin(ctx)
	if err != nil {
		return domain.TransactionBeginError
	}
	defer tx.Rollback(ctx)

	updateShopQuery := `UPDATE shops
						SET title = $2, description = $3, manager_ids = $4
						WHERE id = $1`

	result, err := tx.Exec(ctx, updateShopQuery, shop.Id, shop.Title, shop.Description, pq.Array(shop.ManagerIDs))
	if err != nil {
		return err
	}

	if result.RowsAffected() != 1 {
		return domain.ShopNotFoundError
	}

	err = tx.Commit(ctx)
	if err != nil {
		return domain.TransactionCommitError
	}
	return nil
}

func (repo *ProductRepo) GetShopByID(ctx context.Context, id uint64) (shop domain.Shop, err error) {
	tx, err := repo.postgresDB.Begin(ctx)
	if err != nil {
		return domain.Shop{}, domain.TransactionBeginError
	}
	defer tx.Rollback(ctx)

	getShopByIDQuery := `SELECT id, title, description, manager_ids
						 FROM shops
						 WHERE id = $1`

	var managerIDs pq.Int64Array

	row := tx.QueryRow(ctx, getShopByIDQuery, id)
	err = row.Scan(&shop.Id, &shop.Title, &shop.Description, &managerIDs)
	if err != nil {
		if err == pgx.ErrNoRows {
			return domain.Shop{}, domain.ShopNotFoundError
		}

		return domain.Shop{}, err
	}

	shop.ManagerIDs = make([]uint64, 0, len(managerIDs))
	for _, id := range managerIDs {
		shop.ManagerIDs = append(shop.ManagerIDs, uint64(id))
	}

	err = tx.Commit(ctx)
	if err != nil {
		return domain.Shop{}, domain.TransactionCommitError
	}
	return shop, nil
}
