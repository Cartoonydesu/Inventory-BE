package item

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

func SetRouter(r *gin.Engine, db *sql.DB) {
	h := &Handler{Db: db}
	r.GET("/api/v1/items", h.GetAllItems)
	r.POST("/api/v1/items", h.AddNewItem)
	r.DELETE("/api/v1/items/:key", h.DeleteItemById)
}
