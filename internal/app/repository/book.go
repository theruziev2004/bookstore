package repository

import (
	"context"

	"github.com/theruziev2004/bookstore/internal/app/model"
	"github.com/theruziev2004/bookstore/internal/pkg/db"
)

type BookRepository struct {
	Db *db.Database
}

func NewBookRepository(db *db.Database) *BookRepository {
	return &BookRepository{Db: db}
}

func (r *BookRepository) Add(b *model.Book) (*model.Book, error) {
	insertSql := "INSERT INTO books (title, author, genre, price, description) VALUES ($1, $2, $3, $4, $5);"
	_, err := r.Db.Konnect.Exec(context.Background(), insertSql, b.Title, b.Author, b.Genre, b.Price, b.Description)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (r *BookRepository) Update(b *model.Book) (*model.Book, error) {
	updateSql := "UPDATE books SET title=$1, author=$2, genre=$3, price=$4, description=$5 WHERE id=$6;"
	_, err := r.Db.Konnect.Exec(context.Background(), updateSql, b.Title, b.Author, b.Genre, b.Price, b.Description, b.ID)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (r *BookRepository) Get(bookID int64) (*model.Book, error) {
	selectSql := "SELECT id, title, author, genre, price, description FROM books WHERE id=$1;"
	var book model.Book
	err := r.Db.Konnect.QueryRow(context.Background(), selectSql, bookID).Scan(&book.ID, &book.Title, &book.Author, &book.Genre, &book.Price, &book.Description)
	return &book, err
}

func (r *BookRepository) All() ([]*model.Book, error) {
	selectSql := "SELECT id, genre, price, title, author, description FROM books;"
	row, err := r.Db.Konnect.Query(context.Background(), selectSql)
	if err != nil {
		return nil, err
	}
	book := make([]*model.Book, 0)
	for row.Next() {
		var b model.Book
		if err = row.Scan(&b.ID, &b.Genre, &b.Price, &b.Title, &b.Author, &b.Description); err != nil {
			return nil, err
		}
		book = append(book, &b)
	}
	return book, nil
}
