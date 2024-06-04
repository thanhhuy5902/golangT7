package category

import (
	"cakho.com/tudye/domain/product"
	"context"
	"database/sql"
	"github.com/ServiceWeaver/weaver"
	_ "github.com/go-sql-driver/mysql"
)

type CategoryRepository interface {
	Create(ctx context.Context, category product.Category) error
	Update(ctx context.Context, category product.Category) error
	Delete(ctx context.Context, id string) error
	FindById(ctx context.Context, id string) (product.Category, error)
	FindAll(ctx context.Context) ([]product.Category, error)
}

type categoryRepository struct {
	weaver.Implements[CategoryRepository]
	weaver.WithConfig[product.CategoryRepositoryConfig]
	db *sql.DB
}

func (c *categoryRepository) Create(ctx context.Context, category product.Category) error {
	tx, err := c.db.Begin()
	if err != nil {
		return err
	}
	_, err = tx.Exec("INSERT INTO categories (id, name) VALUES (?, ?)", category.Id, category.Name)
	if err != nil {
		_ = tx.Rollback()
	}
	return tx.Commit()
}

func (c *categoryRepository) Update(ctx context.Context, category product.Category) error {
	tx, err := c.db.Begin()
	if err != nil {
		return err
	}
	_, err = tx.Exec("UPDATE categories SET name= ? WHERE id= ?", category.Name, category.Id)
	if err != nil {
		_ = tx.Rollback()
	}
	return tx.Commit()
}

func (c *categoryRepository) Delete(ctx context.Context, id string) error {
	tx, err := c.db.Begin()
	if err != nil {
		return err
	}
	_, err = tx.Exec("DELETE FROM categories WHERE id=?", id)
	if err != nil {
		_ = tx.Rollback()
	}
	return tx.Commit()
}

func (c *categoryRepository) FindById(ctx context.Context, id string) (product.Category, error) {
	row := c.db.QueryRow("SELECT id, name FROM categories WHERE id=?", id)
	var category product.Category
	err := row.Scan(&category.Id, &category.Name)
	if err != nil {
		return product.Category{}, err
	}
	return category, nil
}

func (c *categoryRepository) FindAll(ctx context.Context) ([]product.Category, error) {
	rows, err := c.db.Query("SELECT id, name FROM categories")
	if err != nil {
		return nil, err

	}
	var categories []product.Category
	for rows.Next() {
		var category product.Category
		err = rows.Scan(&category.Id, &category.Name)
		if err != nil {
			return nil, err
		}
		categories = append(categories, category)

	}
	return categories, nil
}

func (i *categoryRepository) Init(ctx context.Context) error {
	db, err := sql.Open("mysql", i.Config().Dsn)
	if err != nil {
		return err
	}
	i.db = db
	// Create table if not exists
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS categories (id VARCHAR(36) PRIMARY KEY, name VARCHAR(255) NOT NULL)")
	if err != nil {
		return err
	}
	return nil
}
