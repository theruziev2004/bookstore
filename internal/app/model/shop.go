package model

type BuyBook struct {
	BookID int64 `json:"book_id"`
	Qty    int64 `json:"qty"`
}
