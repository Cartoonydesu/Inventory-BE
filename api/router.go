package api

import (
	"github.com/gin-gonic/gin"
	// "github.com/swaggo/swag/example/basic/api"
)

func SetRouter(r *gin.Engine) {
	h := &Handler{}
	r.GET("/api/v1/upc-barcode-reader", h.GetDataFromUPC)
}
