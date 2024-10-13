package controllers

import (
	"coeffee/domain"
	"fmt"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserUsecase domain.UserUseCase
}


func NewUserController(userUsecase domain.UserUseCase) *UserController {
	return &UserController{UserUsecase: userUsecase}
}


func(uc *UserController) CreateAccount(c *gin.Context) {
	var user domain.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	result,err := uc.UserUsecase.CreateAccount(user)
	if err.Message != "" {
		c.JSON(400, gin.H{
			"status": err.Status,
			"message": err.Message,
		})
		return
	}
	c.JSON(200, gin.H{"data": result})
}




func(uc *UserController) Login(c *gin.Context) {
	var user domain.User
	
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	result,err := uc.UserUsecase.Login(user)
	if err.Message != "" {
		c.JSON(400, gin.H{
			"status": err.Status,
			"message": err.Message,
	})
		return
	}

	c.JSON(200, gin.H{"data": result})
}

func(uc *UserController) GetByID(c *gin.Context) {
	id := c.Param("id")

	result,err := uc.UserUsecase.GetByID(id)

	if err.Message != "" {
		c.JSON(400, gin.H{
			"status": err.Status,
			"message": err.Message,
		})
		return
	}

	c.JSON(200, gin.H{"data": result})
}

func (uc *UserController) UpdateProfile(c *gin.Context) {
	
	var user domain.User
	
	userid := c.GetString("user_id")
	
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(userid , )
	result, err := uc.UserUsecase.UpdateProfile(userid, user)

	if  err.Message != "" {
		c.JSON(400, gin.H{
			"status":  err.Status,
			"message": err.Message,
		})
		return
	}

	c.JSON(200, gin.H{"data": result})
}

func (uc *UserController) AddDriver(c *gin.Context) {
	var user domain.User
	

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	result, err := uc.UserUsecase.AddDriver(user)

	if err.Message != "" {
		c.JSON(400, gin.H{
			"status": err.Status,
			"message": err.Message,
		})
		return
	}

	c.JSON(200, gin.H{"data": result})

}


func (uc *UserController) GetAllUser(c *gin.Context) {
	result, err := uc.UserUsecase.GetAllUser()
	if err != nil {
		c.JSON(400, gin.H{"error" : err.Error()})
	}

	c.IndentedJSON(200, gin.H{"result": result})

}

func (uc *UserController) GetMe(c *gin.Context) {
	userID := c.GetString("user_id")
	fmt.Println(userID , "********************************")
	result , err := uc.UserUsecase.GetUserByID(userID)

	
	if err != nil {
		c.JSON(400, gin.H{
			"status": 400 ,
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{"user": result})

	
}


// // create a user 
// func (uc *UserController) CreateUser( c *gin.Context) {
// 	var user domain.UserRegister
	
// 	if err := c.ShouldBindJSON(&user); err != nil {
// 		c.JSON(400, gin.H{"error": err.Error()})
// 		return
// 	}

// 	result := uc.UserUsecase.CreateUser(user)

// 	HandleResponse(c, result)

// }

// // login a user

// func (uc *UserController) Login(c *gin.Context) {
// 	var user domain.UserLogin

// 	if err := c.ShouldBindJSON(&user); err != nil {
// 		c.JSON(400, gin.H{"error": err.Error()})
// 		return
// 	}

// 	result := uc.UserUsecase.Login(user)

// 	HandleResponse(c, result)
// }

// // create a farmer

// func (uc *UserController) CreateFarmer(c *gin.Context) {
// 	var user domain.UserRegister

// 	// get the current logged in user id from the context

// 	adminID := c.GetString("userID")

// 	if err := c.ShouldBindJSON(&user); err != nil {
// 		c.JSON(400, gin.H{"error": err.Error()})
// 		return
// 	}

// 	result := uc.UserUsecase.CreateFarmer(adminID , user)

// 	HandleResponse(c, result)
// }

// // // create a driver

// func (uc *UserController) CreateDriver(c *gin.Context) {
// 	var user domain.UserRegister

// 	adminID := c.GetString("userID")
// 	if err := c.ShouldBindJSON(&user); err != nil {
// 		c.JSON(400, gin.H{"error": err.Error()})
// 		return
// 	}

// 	result := uc.UserUsecase.CreateDriver(adminID, user)

// 	HandleResponse(c, result)
// }



// func (uc *UserController) GetAllUser(c *gin.Context) {
// 	AdminID := c.GetString("userID")

// 	// check if the user is an admin before getting all the users
// 	result := uc.UserUsecase.GetAllUser(AdminID)

// 	HandleResponse(c, result)


// }

// func (uc *UserController) GetUserByID(c *gin.Context) {
// 	id := c.Param("id")

// 	result := uc.UserUsecase.GetUserByID(id)

// 	HandleResponse(c, result)
// }



