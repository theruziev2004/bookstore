package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/theruziev2004/bookstore/internal/app/model"
	"github.com/theruziev2004/bookstore/internal/app/service"
	"github.com/theruziev2004/bookstore/internal/pkg"
)

type BookRestAPI struct {
	BookService *service.BookService
}

func NewBookRestAPI(s *service.BookService) *BookRestAPI {
	return &BookRestAPI{
		BookService: s,
	}
}

func (s *BookRestAPI) Add(writer http.ResponseWriter, request *http.Request) {
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		pkg.SendString(writer, http.StatusBadRequest, "bad request")
		return
	}

	book := model.Book{}

	if err = json.Unmarshal(body, &book); err != nil {
		pkg.SendString(writer, http.StatusBadRequest, "json is invalid")
		return
	}
	_, err = s.BookService.Add(&book)
	if err != nil {
		pkg.SendString(writer, http.StatusBadRequest, "failed to save")
		return
	}

	pkg.SendString(writer, http.StatusCreated, "ok")
}

func (s *BookRestAPI) Update(writer http.ResponseWriter, request *http.Request) {
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		pkg.SendString(writer, http.StatusBadRequest, "bad request")
		return
	}

	book := model.Book{}

	if err = json.Unmarshal(body, &book); err != nil {
		pkg.SendString(writer, http.StatusBadRequest, "json is invalid")
		return
	}
	_, err = s.BookService.Update(&book)
	if err != nil {
		pkg.SendString(writer, http.StatusBadRequest, "failed to save")
		return
	}

	pkg.SendString(writer, http.StatusOK, "ok")
}

func (s *BookRestAPI) Get(writer http.ResponseWriter, request *http.Request) {
	bookParam, ok := request.URL.Query()["id"]
	if !ok {
		pkg.SendString(writer, http.StatusBadRequest, "bad request")
		return
	}
	bookID, err := strconv.ParseInt(bookParam[0], 10, 64)
	if err != nil {
		pkg.SendString(writer, http.StatusBadRequest, "failed to save")
		return
	}

	book, err := s.BookService.Get(bookID)
	if err != nil {
		pkg.SendString(writer, http.StatusBadRequest, "failed to get book")
		return
	}

	bookJson, err := json.Marshal(book)
	if err != nil {
		pkg.SendString(writer, http.StatusBadRequest, "failed to make book to json")
		return
	}

	pkg.SendString(writer, http.StatusOK, string(bookJson))
}

func (s *BookRestAPI) All(writer http.ResponseWriter, request *http.Request) {
	book, err := s.BookService.All()
	if err != nil {
		pkg.SendString(writer, http.StatusBadRequest, "failed to all book")
		return
	}

	bookJson, err := json.Marshal(book)
	if err != nil {
		pkg.SendString(writer, http.StatusBadRequest, "failed to make book to json")
		return
	}

	writer.WriteHeader(http.StatusOK)
	fmt.Fprint(writer, string(bookJson))

}
