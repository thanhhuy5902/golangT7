package categoryItem

import (
	"cakho.com/tudye/domain/product"
	"context"
	"database/sql"
	"github.com/ServiceWeaver/weaver"

	_ "github.com/go-sql-driver/mysql"
)

type CategoryItemRepository interface {
	Create(ctx context.Context, categoryItem product.CategoryItem) error
	Update(ctx context.Context, categoryItem product.CategoryItem) error
	Delete(ctx context.Context, id string) error
	FindById(ctx context.Context, id string) (product.CategoryItem, error)
	FindAll(ctx context.Context) ([]product.CategoryItem, error)
	FindByCategoryId(ctx context.Context, categoryId string) ([]product.CategoryItem, error)
}

type categoryItemRepository struct {
	weaver.Implements[CategoryItemRepository]
	weaver.WithConfig[product.CategoryItemRepositoryConfig]
	db *sql.DB
}

func (i *categoryItemRepository) Create(ctx context.Context, categoryItem product.CategoryItem) error {
	tx, err := i.db.Begin()
	if err != nil {
		return err
	}
	_, err = tx.Exec("INSERT INTO category_items (id, category_id, item_id) VALUES (?, ?, ?)", categoryItem.Id, categoryItem.CategoryId, categoryItem.ItemId)
	if err != nil {
		_ = tx.Rollback()
	}
	return tx.Commit()
}

func (i *categoryItemRepository) Update(ctx context.Context, categoryItem product.CategoryItem) error {
	tx, err := i.db.Begin()
	if err != nil {
		return err
	}
	_, err = tx.Exec("UPDATE category_items SET category_id= ?, item_id= ? WHERE id= ?", categoryItem.CategoryId, categoryItem.ItemId, categoryItem.Id)
	if err != nil {
		_ = tx.Rollback()
	}
	return tx.Commit()
}

func (i *categoryItemRepository) Delete(ctx context.Context, id string) error {
	tx, err := i.db.Begin()
	if err != nil {
		return err
	}
	_, err = tx.Exec("DELETE FROM category_items WHERE id=?", id)
	if err != nil {
		_ = tx.Rollback()
	}
	return tx.Commit()
}

func (i *categoryItemRepository) FindById(ctx context.Context, id string) (product.CategoryItem, error) {
	var categoryItem product.CategoryItem
	err := i.db.QueryRow("SELECT id, category_id, item_id FROM category_items WHERE id=?", id).Scan(&categoryItem.Id, &categoryItem.CategoryId, &categoryItem.ItemId)
	if err != nil {
		return categoryItem, err
	}
	return categoryItem, nil
}

func (i *categoryItemRepository) FindAll(ctx context.Context) ([]product.CategoryItem, error) {
	rows, err := i.db.Query("SELECT id, category_id, item_id FROM category_items")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categoryItems []product.CategoryItem
	for rows.Next() {
		var categoryItem product.CategoryItem
		err := rows.Scan(&categoryItem.Id, &categoryItem.CategoryId, &categoryItem.ItemId)
		if err != nil {
			return nil, err
		}
		categoryItems = append(categoryItems, categoryItem)

	}
	return categoryItems, nil
}

func (i *categoryItemRepository) FindByCategoryId(ctx context.Context, categoryId string) ([]product.CategoryItem, error) {
	rows, err := i.db.Query("SELECT id, category_id, item_id FROM category_items WHERE category_id=?", categoryId)
	if err != nil {
		return nil, err
	}
	var categoryItems []product.CategoryItem
	for rows.Next() {
		var categoryItem product.CategoryItem
		err = rows.Scan(&categoryItem.Id, &categoryItem.CategoryId, &categoryItem.ItemId)
		if err != nil {
			return nil, err
		}
		categoryItems = append(categoryItems, categoryItem)
	}
	return categoryItems, nil
}

func (i *categoryItemRepository) Init(ctx context.Context) error {
	db, err := sql.Open("mysql", i.Config().Dsn)
	if err != nil {
		return err
	}
	i.db = db
	// Create table if not exists
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS category_items (id VARCHAR(36) PRIMARY KEY, category_id VARCHAR(36), item_id VARCHAR(36))")
	if err != nil {
		return err
	}
	return nil
}
