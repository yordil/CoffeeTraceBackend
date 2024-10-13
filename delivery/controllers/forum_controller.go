package controllers

import (
	"coeffee/domain"
	"fmt"

	"github.com/gin-gonic/gin"
)

type ForumController struct {
	ForumUsecase domain.BlogUseCase
}

func NewForumController(bc domain.BlogUseCase) *ForumController {
	return &ForumController{
		ForumUsecase: bc,
	}
}

// CreateOrder
func (fc *ForumController) CreatePost(c *gin.Context) {
	var blog domain.Blog

	// get the id of the current logged in user from the context
	userID := c.GetString("user_id")

	if err := c.ShouldBindJSON(&blog); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	blog.UserID = userID

	suss, err := fc.ForumUsecase.CreateBlog(blog)

	if err.Message != "" {
		c.JSON(err.Status, gin.H{"error": err.Message})
		return
	}

	c.JSON(suss.Status, gin.H{"message": suss.Message})

	return 


}


// get blog by id

func (fc *ForumController) GetBlogByID(c *gin.Context) { 
	id := c.Param("id")

	blog, err := fc.ForumUsecase.GetBlogByID(id)

	if err != nil {
		c.JSON(400, gin.H{"error":"Failed to get blog"})
		return
	}

	c.JSON(200, gin.H{"blog": blog})
	

}

// get all blog

func (fc *ForumController) GetAllBlog(c *gin.Context) {

	limiter := c.Query("limit")
	page := c.Query("page")
	tag := c.Query("tag")

	if limiter == "" {
		limiter = "10"
	}

	if page == "" {
		page = "1"
	}

	blogs, err := fc.ForumUsecase.GetAllBlog(tag, limiter, page)

	if err != nil {
		c.JSON(400, gin.H{"error":"Failed to get all blogs"})
		return
	}

	c.JSON(200, gin.H{"blogs": blogs})

	return
}

// CreateReply
func (fc *ForumController) CreateReply(c *gin.Context) {
	userId := c.GetString("user_id")
	blogID := c.Param("id")

	// bind the request body to the reply struct
	var reply domain.Reply

	if err := c.ShouldBindJSON(&reply); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	reply.BlogID = blogID
	reply.UserID = userId
	reply, err := fc.ForumUsecase.CreateReply(reply)

	if err != nil {
		c.JSON(400, gin.H{"error":"Failed to create reply"})
		return
	}

	c.JSON(200, gin.H{"reply": reply})

	return
}

// GetReply

func (fc *ForumController) GetReply(c *gin.Context) {
	blogID := c.Param("id")
	fmt.Println(blogID)
	replies, err := fc.ForumUsecase.GetReply(blogID)

	if err != nil {
		c.JSON(400, gin.H{"error":"Failed to get replies"})
		return
	}

	c.JSON(200, gin.H{"replies": replies})

	return
}



// CreateResource

func (fc *ForumController) CreateResource(c *gin.Context) {

	var blog domain.Blog	

	// get the id of the current logged in user from the context
	userID := c.GetString("user_id")

	// should bind 
	if err := c.ShouldBindJSON(&blog); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	blog.UserID = userID

	suss, err := fc.ForumUsecase.CreateResouce(blog)

	if err.Message != "" {
		c.JSON(err.Status, gin.H{"error": err.Message})
		return
	}

	c.JSON(suss.Status, gin.H{"message": suss.Message})

	return


}


// GetResource

func (fc *ForumController) GetAllResource(c *gin.Context) {
	blog, err := fc.ForumUsecase.GetResource()

	if err != nil {
		c.JSON(400, gin.H{"error":"Failed to get resource"})
		return
	}

	c.JSON(200, gin.H{"blog": blog})



}