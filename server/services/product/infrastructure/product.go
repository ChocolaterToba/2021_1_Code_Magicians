package repository

import (
	"context"
	"fmt"
	"pinterest/services/product/domain"

	"github.com/huandu/go-sqlbuilder"
	"github.com/jackc/pgx/v4"
	"github.com/lib/pq"
)

func (repo *ProductRepo) CreateProduct(ctx context.Context, product domain.Product) (id uint64, err error) {
	tx, err := repo.postgresDB.Begin(ctx)
	if err != nil {
		return 0, domain.TransactionBeginError
	}
	defer tx.Rollback(ctx)

	createProductQuery := `INSERT INTO products (title, description, price, availability, 
						   assembly_time, parts_amount, size, category, shop_id)
						   VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
						   RETURNING id`

	row := tx.QueryRow(ctx, createProductQuery,
		product.Title, product.Description, product.Price, product.Availability, product.AssemblyTime,
		product.PartsAmount, product.Size, product.Category, product.ShopId)
	err = row.Scan(&id)
	if err != nil {
		// TODO: check if shop exists?
		return 0, err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return 0, domain.TransactionCommitError
	}
	return id, nil
}

func (repo *ProductRepo) UpdateProduct(ctx context.Context, product domain.Product) (err error) {
	tx, err := repo.postgresDB.Begin(ctx)
	if err != nil {
		return domain.TransactionBeginError
	}
	defer tx.Rollback(ctx)

	updateProductQuery := `UPDATE products
						   SET title = $2, description = $3, price = $4, availability =$5, 
						   assembly_time = $6, parts_amount = $7, size = $8, category = $9, shop_id = $10
						   WHERE id = $1`

	result, err := tx.Exec(ctx, updateProductQuery,
		product.Id, product.Title, product.Description, product.Price, product.Availability,
		product.AssemblyTime, product.PartsAmount, product.Size, product.Category, product.ShopId)
	if err != nil {
		return err
	}

	if result.RowsAffected() != 1 {
		return domain.ProductNotFoundError
	}

	err = tx.Commit(ctx)
	if err != nil {
		return domain.TransactionCommitError
	}
	return nil
}

func (repo *ProductRepo) GetProductByID(ctx context.Context, id uint64) (product domain.Product, err error) {
	tx, err := repo.postgresDB.Begin(ctx)
	if err != nil {
		return domain.Product{}, domain.TransactionBeginError
	}
	defer tx.Rollback(ctx)

	getProductByIDQuery := `SELECT id, title, description, price, availability, assembly_time, 
							parts_amount, rating, size, category, image_links, video_link, shop_id
							FROM products
							WHERE id = $1`

	row := tx.QueryRow(ctx, getProductByIDQuery, id)
	err = row.Scan(&product.Id, &product.Title, &product.Description, &product.Price, &product.Availability,
		&product.AssemblyTime, &product.PartsAmount, &product.Rating, &product.Size, &product.Category,
		pq.Array(&product.ImageLinks), &product.VideoLink, &product.ShopId)
	if err != nil {
		if err == pgx.ErrNoRows {
			return domain.Product{}, domain.ProductNotFoundError
		}

		return domain.Product{}, err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return domain.Product{}, domain.TransactionCommitError
	}
	return product, nil
}

func (repo *ProductRepo) GetProducts(ctx context.Context, pageOffset uint64, pageSize uint64, category string) (products []domain.Product, err error) {
	tx, err := repo.postgresDB.Begin(ctx)
	if err != nil {
		return nil, domain.TransactionBeginError
	}
	defer tx.Rollback(ctx)

	sb := sqlbuilder.Select("id", "title", "description", "price", "availability", "assembly_time",
		"parts_amount", "rating", "size", "category", "image_links", "video_link", "shop_id").
		From("products").
		Desc().OrderBy("id").
		Limit(int(pageSize)).Offset(int(pageSize * pageOffset))

	if category != "" {
		fmt.Println(category)
		sb.Where(sb.Equal("category", category))
	}

	query, args := sb.BuildWithFlavor(sqlbuilder.PostgreSQL)

	rows, err := tx.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var product domain.Product
		err = rows.Scan(&product.Id, &product.Title, &product.Description, &product.Price, &product.Availability,
			&product.AssemblyTime, &product.PartsAmount, &product.Rating, &product.Size, &product.Category,
			pq.Array(&product.ImageLinks), &product.VideoLink, &product.ShopId)
		if err != nil {
			return nil, err
		}

		products = append(products, product)
	}

	err = tx.Commit(ctx)
	if err != nil {
		return nil, domain.TransactionCommitError
	}
	return products, nil
}
