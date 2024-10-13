package controllers

import (
	"coeffee/domain"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	ProductUsecase domain.ProductUseCase
}
// create a product

func (pc *ProductController) CreateProduct(c *gin.Context) {
	var product domain.Product

	// get the id of the current logged in user from the context
	userID := c.GetString("user_id")
	name := c.GetString("name")


	product.FarmerID = userID
	product.FarmerName = name

	
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	suss, err := pc.ProductUsecase.CreateProduct(userID, product)

	if err.Message != "" {
		c.JSON(err.Status, gin.H{"error": err.Message})
		return
	}
		
	c.JSON(suss.Status, gin.H{"message": suss.Message})
}


// get all product
func (pc *ProductController) GetAllProduct(c *gin.Context) {
	limiter := c.Query("limit")
	page := c.Query("page")
	name := c.Query("name")

	if limiter == "" {
		limiter = "10"
	}

	if page == "" {
		page = "1"
	}
	
	products, err := pc.ProductUsecase.GetAllProduct(name, limiter, page)

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"products": products})
}
