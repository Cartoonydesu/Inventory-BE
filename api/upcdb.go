package api

import (
	"cartoonydesu/response"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
}

type Item struct {
	EAN   string `json:"ean"`
	Title string `json:"title"`
	Brand string `json:"brand"`
}

type UPC struct {
	Code  string `json:"code"`
	Total int    `json:"total"`
	Items []Item `json:"items"`
}

func (h *Handler) GetDataFromUPC(c *gin.Context) {
	ean := c.Query("ean")
	if ean == "" {
		c.JSON(http.StatusBadRequest, response.ErrorResponse("error", "Requested EAN id not found"))
		return
	}
	res, err := http.Get(fmt.Sprintf("https://api.upcitemdb.com/prod/trial/lookup?upc=%v", ean))
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse("error", "Cannot get data from fetch"))
		return
	}
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, "Can not read fetched data")
		return
	}
	// json.Decoder(res)
	var upc UPC
	err = json.Unmarshal(data, &upc)
	if err != nil {
		c.JSON(http.StatusBadRequest, "Can not convert data to JSON")
		return
	}
	var result gin.H
	if upc.Code == "OK" && upc.Items == nil {
		c.JSON(http.StatusNotFound, response.ErrorResponse("error", "EAN for item not found"))
		return
	}
	if len(upc.Items) > 0 {
		item := upc.Items[0]
		result = gin.H{
			"ean":   item.EAN,
			"title": item.Title,
			"brand": item.Brand,
		}
	}
	if upc.Code != "OK" {
		c.JSON(http.StatusBadRequest, response.ErrorResponse("error", upc.Code))
		return
	}
	c.JSON(http.StatusOK, response.SuccessResponse("success", result))
}
