package item

import (
	"cartoonydesu/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetAllItems(c *gin.Context) {
	rows, err := h.Db.Query("SELECT itemId, ean, title, brand, amount, note, expiredDate FROM item")
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse("error", "Query error"))
		return
	}
	defer rows.Close()
	var items []Item
	for rows.Next() {
		var i Item
		err := rows.Scan(&i.ItemId, &i.Ean, &i.Title, &i.Brand, &i.Amount, &i.Note, &i.ExpiredDate)
		if err != nil {
			c.JSON(http.StatusBadRequest, response.ErrorResponse("error", err.Error()))
			return
		}
		items = append(items, i)
	}
	// c.Header("Access-Control-Allow-Origin", "*")
	c.JSON(http.StatusOK, response.SuccessResponse("success", items))
}

func (h *Handler) getItemById(key string) (Item, error) {
	row := h.Db.QueryRow("SELECT itemId, ean, title, brand, amount, note, expiredDate FROM item WHERE itemId = $1", key)
	var i Item
	err := row.Scan(&i.ItemId, &i.Ean, &i.Title, &i.Brand, &i.Amount, &i.Note, &i.ExpiredDate)
	if err != nil {
		return Item{}, err
	}
	return i, nil
}
