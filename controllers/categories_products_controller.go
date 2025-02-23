package controllers

import (
	"PROJECT_TEST_GO/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type getProductData struct {
	CategoryId string `json:"categoryId"`
}

func GetListCategories(c *gin.Context) {
	categories := models.CategoriesAll()

	c.IndentedJSON(
		http.StatusOK,
		gin.H{"categories": categories},
	)
}

func GetListProducts(c *gin.Context) {
	var getProduct getProductData
	if err := c.BindJSON(&getProduct); err != nil {
		c.IndentedJSON(
			http.StatusInternalServerError,
			gin.H{"message": "Sorry Server Error"},
		)
		return
	}

	products := models.ProductsAll(getProduct.CategoryId)
	c.IndentedJSON(
		http.StatusOK,
		gin.H{"products": products},
	)
}

func GetProductById(c *gin.Context) {
	id := c.Param("id")

	product := models.ProductById(id)
	c.IndentedJSON(
		http.StatusOK,
		gin.H{"product": product},
	)
}
