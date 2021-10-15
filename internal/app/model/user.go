package model

type User struct {
	ID          int64  `json:"id"`
	Username    string `json:"username"`
	Password    string `json:"-"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	PhoneNumber int64  `json:"phone_number"`
	Balance     int64  `json:"balance"`
	Status      string `json:"status"`
}

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
