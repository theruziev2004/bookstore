package service

import (
	"github.com/theruziev2004/bookstore/internal/app/model"
	"github.com/theruziev2004/bookstore/internal/app/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	UserRepository *repository.UserRepository
}

func NewUserService(r *repository.UserRepository) *UserService {
	return &UserService{
		UserRepository: r,
	}
}

func (s *UserService) Register(u *model.User) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(u.Password), 14)
	if err != nil {
		return err
	}
	u.Password = string(bytes)
	err = s.UserRepository.Add(u)
	return err

}

func (s *UserService) Auth(username, password string) (*model.User, error) {
	usr, err := s.UserRepository.Get(username)
	if err != nil {
		return nil, err
	}
	if err = bcrypt.CompareHashAndPassword([]byte(usr.Password), []byte(password)); err != nil {
		return nil, err
	}
	return usr, nil
}

func (s *UserService) Get(username string) (*model.User, error) {
	return s.UserRepository.Get(username)
}

func (s *UserService) ChangeBalance(uid int64, balance int64) error {
	return s.UserRepository.ChangeBalance(uid, balance)
}
