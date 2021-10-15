package service

import (
	"fmt"
	"time"

	"github.com/theruziev2004/bookstore/internal/app/model"
	"github.com/theruziev2004/bookstore/internal/app/repository"
)

type ShopService struct {
	TransactionRepository *repository.TransactionRepository
	UserService           *UserService
	BookService           *BookService
}

func NewShopService(r *repository.TransactionRepository, userService *UserService, bookService *BookService) *ShopService {
	return &ShopService{
		TransactionRepository: r,
		UserService:           userService,
		BookService:           bookService,
	}
}

func (s *ShopService) Buy(bookID int64, qty int64, username string) (*model.Transaction, error) {
	book, err := s.BookService.Get(bookID)
	if err != nil {
		return nil, err
	}
	user, err := s.UserService.Get(username)
	if err != nil {
		return nil, err
	}
	transaction := model.Transaction{
		BookID:   book.ID,
		UserID:   user.ID,
		Amount:   book.Price,
		Qty:      qty,
		Datetime: time.Now(),
	}

	if user.Balance < book.Price*qty {
		return nil, fmt.Errorf("not enough money")
	}

	err = s.TransactionRepository.Add(&transaction)
	if err != nil {
		return nil, err
	}
	if err = s.UserService.ChangeBalance(user.ID, -(book.Price * qty)); err != nil {
		return nil, err
	}
	return nil, nil
}
