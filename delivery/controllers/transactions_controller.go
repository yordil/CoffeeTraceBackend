package controllers

import (
	"coeffee/domain"

	"github.com/gin-gonic/gin"
)

type transactionController struct {
	TransactionUsecase domain.TransactionsUseCase
}

func NewTransactionController(tc domain.TransactionsUseCase) *transactionController {
	return &transactionController{
		TransactionUsecase: tc,
	}
}


func (tc *transactionController) CreateTransaction(c *gin.Context) {
	var transaction domain.Transactions
	c.BindJSON(&transaction)

	userId := c.GetString("user_id")

	
	transaction.MerchantID = userId

	response, err := tc.TransactionUsecase.CreateTransaction(transaction)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON( 200, gin.H{
		"message": "Transaction created successfully",
		"data": response,
	})
}


func (tc *transactionController) GetMyTransactions(c *gin.Context) {
	// userId := c.("user_id")
	userId := c.GetString("user_id")
	Role := c.GetString("role")
	response, err := tc.TransactionUsecase.GetTransactions(userId, Role)
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