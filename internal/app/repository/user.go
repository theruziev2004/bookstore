package repository

import (
	"context"

	"github.com/theruziev2004/bookstore/internal/app/model"
	"github.com/theruziev2004/bookstore/internal/pkg/db"
)

type UserRepository struct {
	Db *db.Database
}

func NewUserRepository(db *db.Database) *UserRepository {
	return &UserRepository{
		db,
	}
}

func (r *UserRepository) Add(u *model.User) error {
	insertSql := "insert into users (username, password, first_name, last_name, phone_number, balance, status) values ($1, $2, $3, $4, $5, $6, $7);"
	_, err := r.Db.Konnect.Exec(context.Background(), insertSql, u.Username, u.Password, u.FirstName, u.LastName, u.PhoneNumber, u.Balance, u.Status)
	return err
}

func (r *UserRepository) Get(username string) (*model.User, error) {
	selectSql := "SELECT id, username, password, balance FROM users WHERE username=$1;"
	var user model.User
	err := r.Db.Konnect.QueryRow(context.Background(), selectSql, username).Scan(&user.ID, &user.Username, &user.Password, &user.Balance)
	return &user, err
}

func (r *UserRepository) ChangeBalance(uid int64, balance int64) error {
	updateSql := "UPDATE users SET balance=balance + $1 WHERE id=$2;"
	_, err := r.Db.Konnect.Exec(context.Background(), updateSql, balance, uid)
	if err != nil {
		return err
	}
	return nil
}
