package item

import (
	"database/sql"
	"time"
)

type Item struct {
	ItemId      string    `json:"itemId"`
	Ean         string    `json:"ean"`
	Title       string    `json:"title"`
	Brand       string    `json:"brand"`
	Amount      int       `json:"amount"`
	Note        string    `json:"note"`
	ExpiredDate time.Time `json:"expiredDate"`
}

type Handler struct {
	Db *sql.DB
}
