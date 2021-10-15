package repository

import (
	"context"

	"github.com/theruziev2004/bookstore/internal/app/model"
	"github.com/theruziev2004/bookstore/internal/pkg/db"
)

type TransactionRepository struct {
	Db *db.Database
}

func NewTransactionRepository(db *db.Database) *TransactionRepository {
	return &TransactionRepository{Db: db}
}

func (s *TransactionRepository) Add(t *model.Transaction) error {
	insertSql := "insert into transactions (book_id, user_id, amount, qty) values ($1, $2, $3, $4);"
	_, err := s.Db.Konnect.Exec(context.Background(), insertSql, t.BookID, t.UserID, t.Amount, t.Qty)
	return err
}
