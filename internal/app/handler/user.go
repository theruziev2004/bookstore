package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/theruziev2004/bookstore/internal/app/model"
	"github.com/theruziev2004/bookstore/internal/app/service"
	"github.com/theruziev2004/bookstore/internal/pkg"
	"github.com/theruziev2004/bookstore/internal/pkg/auth"
)

type UserRestAPI struct {
	UserService *service.UserService
}

func NewUserRestAPI(s *service.UserService) *UserRestAPI {
	return &UserRestAPI{
		UserService: s,
	}
}

func (r *UserRestAPI) Register(writer http.ResponseWriter, request *http.Request) {
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		pkg.SendString(writer, http.StatusBadRequest, "bad request")
		return
	}

	var usr model.User

	if err = json.Unmarshal(body, &usr); err != nil {
		pkg.SendString(writer, http.StatusBadRequest, "json is invalid")
		return
	}

	err = r.UserService.Register(&usr)
	if err != nil {
		pkg.SendString(writer, http.StatusBadRequest, "failed to save")
		return
	}

	writer.WriteHeader(http.StatusOK)

}

func (r *UserRestAPI) Me(writer http.ResponseWriter, request *http.Request) {
	token := auth.ReadTokenFromContext(request.Context())
	user, err := r.UserService.Get(token.ID)
	if err != nil {
		pkg.SendString(writer, http.StatusBadRequest, err.Error())
	}
	userJson, err := json.Marshal(user)
	if err != nil {
		pkg.SendString(writer, http.StatusBadRequest, "failed to make book to json")
		return
	}

	pkg.SendString(writer, http.StatusOK, string(userJson))
}

func (r *UserRestAPI) Login(writer http.ResponseWriter, request *http.Request) {
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		pkg.SendString(writer, http.StatusBadRequest, "bad request")
		return
	}

	var login model.Login

	if err = json.Unmarshal(body, &login); err != nil {
		pkg.SendString(writer, http.StatusBadRequest, "json is invalid")
		return
	}

	savedUser, err := r.UserService.Auth(login.Username, login.Password)
	if err != nil {
		pkg.SendString(writer, http.StatusBadRequest, "incorrect username or password")
		return
	}

	// Create the Claims
	claims := &jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(10 * 24 * time.Hour)),
		ID:        savedUser.Username,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(auth.GetJWTSecret())
	if err != nil {
		pkg.SendString(writer, http.StatusBadRequest, "jwt error")
		return
	}
	fmt.Fprint(writer, ss)
}
