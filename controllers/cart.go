package controllers

import (
	"PROJECT_TEST_GO/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AddCart struct {
	ProductId string `json:"productId"`
}

func AddToCart(c *gin.Context) {
	user, _ := c.Get("user")

	userId := user.(models.User).Email
	var addCart AddCart
	if err := c.BindJSON(&addCart); err != nil {
		c.IndentedJSON(
			http.StatusInternalServerError,
			gin.H{"message": "Sorry Server Error"},
		)
		return
	}
	productId := addCart.ProductId

	cart := models.AddToCart(userId, productId)

	c.IndentedJSON(
		http.StatusOK,
		gin.H{"cart": cart},
	)
}

func DeleteFromCart(c *gin.Context) {
	user, _ := c.Get("user")

	userId := user.(models.User).Email
	productId := c.Param("id")

	cart := models.DeletCartByProductId(userId, productId)

	c.IndentedJSON(
		http.StatusOK,
		gin.H{"cart": cart},
	)
}
