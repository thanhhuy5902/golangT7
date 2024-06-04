package item

import (
	"cakho.com/tudye/domain/product"
	"context"
	"database/sql"
	"github.com/ServiceWeaver/weaver"
	_ "github.com/go-sql-driver/mysql"
)

type ItemRepository interface {
	Create(ctx context.Context, item product.Item) error
	Update(ctx context.Context, item product.Item) error
	Delete(ctx context.Context, id string) error
	FindById(ctx context.Context, id string) (product.Item, error)
	FindAll(ctx context.Context) ([]product.Item, error)
}

type itemRepository struct {
	weaver.Implements[ItemRepository]
	weaver.WithConfig[product.ItemRepositoryConfig]
	db *sql.DB
}

func (i *itemRepository) Create(ctx context.Context, item product.Item) error {
	tx, err := i.db.Begin()
	if err != nil {
		return err
	}
	_, err = tx.Exec("INSERT INTO items (id, name, description, photo) VALUES (?, ?, ?, ?)", item.Id, item.Name, item.Description, item.Photo)
	if err != nil {
		_ = tx.Rollback()
	}
	return tx.Commit()

}

func (i *itemRepository) Update(ctx context.Context, item product.Item) error {
	tx, err := i.db.Begin()
	if err != nil {
		return err
	}
	_, err = tx.Exec("UPDATE items SET name= ?, description= ? , photo= ? WHERE id= ?", item.Name, item.Description, item.Photo, item.Id)
	if err != nil {
		_ = tx.Rollback()
	}
	return tx.Commit()
}

func (i *itemRepository) Delete(ctx context.Context, id string) error {
	tx, err := i.db.Begin()
	if err != nil {
		return err
	}
	_, err = tx.Exec("DELETE FROM items WHERE id= ?", id)
	if err != nil {
		_ = tx.Rollback()
	}
	return tx.Commit()
}

func (i *itemRepository) FindById(ctx context.Context, id string) (product.Item, error) {
	var item product.Item
	err := i.db.QueryRow("SELECT id, name, description, photo FROM items WHERE id= ?", id).Scan(item.Id, &item.Name, &item.Description, &item.Photo)
	if err != nil {
		return product.Item{}, err
	}
	return item, nil
}

func (i *itemRepository) FindAll(ctx context.Context) ([]product.Item, error) {
	rows, err := i.db.Query("SELECT id, name, description, photo FROM items")
	if err != nil {
		return nil, err
	}
	var items []product.Item
	for rows.Next() {
		var item product.Item
		err := rows.Scan(&item.Id, &item.Name, &item.Description, &item.Photo)
		if err != nil {
			return nil, err
		}
		items = append(items, item)

	}
	return items, nil
}

func (i *itemRepository) Init(ctx context.Context) error {
	db, err := sql.Open("mysql", i.Config().Dsn)
	if err != nil {
		return err
	}
	i.db = db
	// Create table if not exists
	_, err = i.db.Exec("CREATE TABLE IF NOT EXISTS items (id VARCHAR(36) PRIMARY KEY, name VARCHAR(255), description TEXT, photo TEXT)")
	if err != nil {
		return err
	}
	return nil
}
