package item

import (
	"cartoonydesu/response"
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
)

type NewItem struct {
	Ean         string `json:"ean" binding:"required"`
	Title       string `json:"title" binding:"required"`
	Brand       string `json:"brand"`
	Amount      int    `json:"amount" binding:"required"`
	Note        string `json:"note"`
	ExpiredDate string `json:"expiredDate"`
}

func (h *Handler) AddNewItem(c *gin.Context) {
	var i NewItem
	err := c.BindJSON(&i)
	if err != nil {
		log.Print(err)
		c.JSON(http.StatusBadRequest, response.ErrorResponse("error", "Can not extract data from json"))
		return
	}
	stmt, err := h.Db.Prepare("INSERT INTO item(ean, title, brand, amount, note, expiredDate) VALUES ($1, $2, $3, $4, $5, $6)")
	if err != nil {
		log.Print(err)
		c.JSON(http.StatusBadRequest, response.ErrorResponse("error", "Statement error"))
		return
	}
	defer stmt.Close()
	if _, err := stmt.Exec(i.Ean, i.Title, i.Brand, i.Amount, i.Note, i.ExpiredDate); err != nil {
		log.Print(err)
		c.JSON(http.StatusBadRequest, response.ErrorResponse("error", "Object already exists"))
		return
	}
	c.JSON(http.StatusOK, response.SuccessResponse("success", i))
}
