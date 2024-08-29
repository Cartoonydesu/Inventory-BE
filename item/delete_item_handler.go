package item

import (
	"cartoonydesu/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) DeleteItemById(c *gin.Context) {
	p := c.Param("key")
	_, err := h.getItemById(p)
	if err != nil {
		c.JSON(http.StatusNotFound, response.ErrorResponse("error", "Item not found"))
		return
	}
	stmt, err := h.Db.Prepare("DELETE FROM item WHERE itemId = $1")
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse("error", "Statement error"))
		return
	}
	defer stmt.Close()
	_, err = stmt.Exec(p)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse("error", "Can not delete item"))
		return
	}
	c.JSON(http.StatusOK, response.SuccessResponse("success", "Delete successful"))
}