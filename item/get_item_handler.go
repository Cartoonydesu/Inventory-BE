package item

import (
	"cartoonydesu/response"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetAllItems(c *gin.Context) {
	rows, err := h.Db.Query("SELECT itemId, ean, title, brand, amount, note, expiredDate FROM item;")
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse("error", "Query error"))
		return
	}
	defer rows.Close()
	var invs []Item
	for rows.Next() {
		var inv Item
		err := rows.Scan(&inv.ItemId, &inv.Ean, &inv.Title, &inv.Brand, &inv.Amount, &inv.Note, &inv.ExpiredDate)
		if err != nil {
			c.JSON(http.StatusBadRequest, response.ErrorResponse("error", "Can not extract data from database"))
			return
		}
		invs = append(invs, inv)
	}
	c.Header("Access-Control-Allow-Origin", "*")
	c.JSON(http.StatusOK, response.SuccessResponse("success", invs))
}
