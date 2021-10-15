package model

import "time"

type Transaction struct {
	ID       int64
	BookID   int64
	UserID   int64
	Amount   int64
	Qty      int64
	Datetime time.Time
}
