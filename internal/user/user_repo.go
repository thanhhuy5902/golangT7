package user

import (
	"cakho.com/tudye/domain/user"
	"context"
	"database/sql"
	"github.com/ServiceWeaver/weaver"
	_ "github.com/go-sql-driver/mysql"
)

type UserRepository interface {
	Create(ctx context.Context, user user.User) error
	Update(ctx context.Context, user user.User) error
	Delete(ctx context.Context, id string) error
	FindById(ctx context.Context, id string) (user.User, error)
	FindAll(ctx context.Context) ([]user.User, error)
}

type userRepository struct {
	weaver.Implements[UserRepository]
	weaver.WithConfig[user.UserRepositoryConfig]
	db *sql.DB
}

func (u *userRepository) Create(ctx context.Context, user user.User) error {
	tx, err := u.db.Begin()
	if err != nil {
		return err
	}
	_, err = tx.Exec("INSERT INTO users (id, username, email) VALUES (?, ?, ?)", user.Id, user.Username, user.Email)
	if err != nil {
		_ = tx.Rollback()
	}
	return tx.Commit()
}

func (u *userRepository) Update(ctx context.Context, user user.User) error {
	tx, err := u.db.Begin()
	if err != nil {
		return err
	}
	_, err = tx.Exec("UPDATE users SET username= ?, email= ? WHERE id= ?", user.Username, user.Email, user.Id)
	if err != nil {
		_ = tx.Rollback()
	}
	return tx.Commit()
}

func (u *userRepository) Delete(ctx context.Context, id string) error {
	tx, err := u.db.Begin()
	if err != nil {
		return err
	}
	_, err = tx.Exec("DELETE FROM users WHERE id=?", id)
	if err != nil {
		_ = tx.Rollback()
	}
	return tx.Commit()
}

func (u *userRepository) FindById(ctx context.Context, id string) (user.User, error) {
	var user user.User
	err := u.db.QueryRow("SELECT id, username, email FROM users WHERE id= ?", id).Scan(user.Id, &user.Username, &user.Email)
	if err != nil {
		return user, err

	}
	return user, nil
}

func (u *userRepository) FindAll(ctx context.Context) ([]user.User, error) {
	rows, err := u.db.Query("SELECT id, username, email FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var users []user.User
	for rows.Next() {
		var user user.User
		err := rows.Scan(&user.Id, &user.Username, &user.Email)
		if err != nil {
			return nil, err
		}
		users = append(users, user)

	}
	return users, nil
}

func (u *userRepository) Init(ctx context.Context) error {
	db, err := sql.Open("mysql", u.Config().DsnUser)
	if err != nil {
		return err
	}
	u.db = db
	// Create table if not exists
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS users (id VARCHAR(36) PRIMARY KEY, username VARCHAR(255), email VARCHAR(255))")
	if err != nil {
		return err
	}
	return nil
}
