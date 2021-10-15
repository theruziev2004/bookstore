package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/theruziev2004/bookstore/internal/pkg"

	"github.com/theruziev2004/bookstore/internal/app/model"
	"github.com/theruziev2004/bookstore/internal/app/service"
	"github.com/theruziev2004/bookstore/internal/pkg/auth"
)

type ShopRestAPI struct {
	ShopService *service.ShopService
}

func NewShopRestAPI(s *service.ShopService) *ShopRestAPI {
	return &ShopRestAPI{
		ShopService: s,
	}
}

func (r *ShopRestAPI) Buy(writer http.ResponseWriter, request *http.Request) {
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		pkg.SendString(writer, http.StatusBadRequest, "bad request")
		return
	}

	var buyBook model.BuyBook

	if err = json.Unmarshal(body, &buyBook); err != nil {
		pkg.SendString(writer, http.StatusBadRequest, "json is invalid")
		return
	}
	claim := auth.ReadTokenFromContext(request.Context())

	_, err = r.ShopService.Buy(buyBook.BookID, buyBook.Qty, claim.ID)
	if err != nil {
		pkg.SendString(writer, http.StatusBadRequest, err.Error())
		return
	}

	writer.WriteHeader(http.StatusOK)
	fmt.Fprint(writer, "you buy the book")
}
