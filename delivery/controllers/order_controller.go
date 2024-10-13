package controllers

import (
	"coeffee/domain"

	"github.com/gin-gonic/gin"
)

type OrderController struct {
	OrderUseCase domain.OrderUseCase
}

func NewOrderController(oc domain.OrderUseCase) *OrderController {
	return &OrderController{
		OrderUseCase: oc,
	}
}

func (oc *OrderController) CreateOrder(c *gin.Context) {
	var order domain.Order
	c.BindJSON(&order)

	userId := c.GetString("user_id")
	merchant_name := c.GetString("name")


	
	order.MerchantID = userId
	order.MerchantName = merchant_name

	response, err := oc.OrderUseCase.CreateOrder(order)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON( 200, gin.H{
		"message": "Order created successfully",
		"data": response,
	})
}


func (oc *OrderController) GetMyOrders(c *gin.Context) {
	// userId := c.("user_id")
	userId := c.GetString("user_id")
	Role := c.GetString("role")
	response, err := oc.OrderUseCase.GetMyOrders(userId, Role)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"data": response,
	})
}


func (oc *OrderController) AcceptOrderFarmer(c *gin.Context) {
	orderID := c.Param("id")
	Role := c.GetString("role")
	response, err := oc.OrderUseCase.AcceptOrderFarmer(orderID, Role)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Order accepted successfully",
		"data": response,
	})
}

func (oc *OrderController) RejectOrderFarmer(c *gin.Context) {
	orderID := c.Param("id")
	Role := c.GetString("role")
	response, err := oc.OrderUseCase.RejectOrderFarmer(orderID, Role)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Order rejected successfully",
		"data": response,
	})
}

func (oc *OrderController) AcceptOrderDriver(c *gin.Context) {
	orderID := c.Param("id")
	Role := c.GetString("role")
	response, err := oc.OrderUseCase.AcceptOrderDriver(orderID, Role)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Order accepted successfully",
		"data": response,
	})
}

func (oc *OrderController) RejectOrderDriver(c *gin.Context) {
	orderID := c.Param("id")
	Role := c.GetString("role")
	response, err := oc.OrderUseCase.RejectOrderDriver(orderID, Role)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Order rejected successfully",
		"data": response,
	})
}


func (oc *OrderController) PickupOrderDriver(c *gin.Context) {
	orderID := c.Param("id")
	Role := c.GetString("role")
	response, err := oc.OrderUseCase.PickupOrderDriver(orderID, Role)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Order picked up successfully",
		"data": response,
	})
}


func (oc *OrderController) PayFarmer(c *gin.Context) {
	orderID := c.Param("id")
	Role := c.GetString("role")
	response, err := oc.OrderUseCase.PayFarmer(orderID, Role)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Payment successful",
		"data": response,
	})
}

func (oc *OrderController) PayDriver(c *gin.Context) {
	orderID := c.Param("id")
	Role := c.GetString("role")
	response, err := oc.OrderUseCase.PayDriver(orderID, Role)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Payment successful",
		"data": response,
	})
}


func (oc *OrderController) UpdateDestinations(c *gin.Context) {
	orderID := c.Param("id")
	Role := c.GetString("role")
	var destination domain.Destinations
	c.BindJSON(&destination)

	response, err := oc.OrderUseCase.UpdateDestinations(orderID, Role, destination)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Destinations updated successfully",
		"data": response,
	})
}