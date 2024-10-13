package controllers

import (
	"coeffee/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

// HandleResponse handles the API response based on the type of the response object.
func HandleResponse(c *gin.Context, response interface{}) {

	switch res := response.(type) {
	case domain.SuccessResponse:
		c.JSON(http.StatusOK, res)
	case domain.LoginSucessResponse:
		c.JSON(http.StatusOK, res)
	case domain.ErrorResponse:
		c.JSON(http.StatusBadRequest, res)	
	default:
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: "Internal Server Error", Status: 500})
	}
}
