package repository

import (
	"context"
	"pinterest/services/product/domain"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/lib/pq"
)

type ProductRepoInterface interface {
	CreateShop(ctx context.Context, shop domain.Shop) (id uint64, err error)
	UpdateShop(ctx context.Context, shop domain.Shop) (err error)
	GetShopByID(ctx context.Context, id uint64) (shop domain.Shop, err error)
}

type ProductRepo struct {
	postgresDB *pgxpool.Pool
}

func NewProductRepo(postgresDB *pgxpool.Pool) *ProductRepo {
	return &ProductRepo{postgresDB: postgresDB}
}

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

	row := tx.QueryRow(ctx, getShopByIDQuery, id)
	err = row.Scan(&shop.Id, &shop.Title, &shop.Description, pq.Array(&shop.ManagerIDs))
	if err != nil {
		if err == pgx.ErrNoRows {
			return domain.Shop{}, domain.ShopNotFoundError
		}

		return domain.Shop{}, err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return domain.Shop{}, domain.TransactionCommitError
	}
	return shop, nil
}
