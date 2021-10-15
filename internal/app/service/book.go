package service

import (
	"github.com/theruziev2004/bookstore/internal/app/model"
	"github.com/theruziev2004/bookstore/internal/app/repository"
)

type BookService struct {
	BookRepository *repository.BookRepository
}

func NewBookService(r *repository.BookRepository) *BookService {
	return &BookService{BookRepository: r}
}

func (s *BookService) Add(b *model.Book) (*model.Book, error) {
	return s.BookRepository.Add(b)
}

func (s *BookService) Update(b *model.Book) (*model.Book, error) {
	return s.BookRepository.Update(b)
}

func (s *BookService) Get(bookID int64) (*model.Book, error) {
	return s.BookRepository.Get(bookID)
}

func (s *BookService) All() ([]*model.Book, error) {
	return s.BookRepository.All()
}
