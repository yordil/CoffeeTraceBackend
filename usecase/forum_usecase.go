package usecase

import (
	"coeffee/domain"
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BlogUseCase struct {
	BlogRepository		domain.BlogRepository
	ReplyRepository    domain.ReplyRepository
	UserRepository	 domain.UserRepository
	contextTimeout  	time.Duration
}

func NewBlogUseCase(blog domain.BlogRepository, reply domain.ReplyRepository ,user domain.UserRepository ,  timeout time.Duration ) domain.BlogUseCase {
	return &BlogUseCase{
		BlogRepository: blog,
		ReplyRepository: reply,
		UserRepository: user,
		contextTimeout: timeout,
	}
}

func (uc *BlogUseCase) CreateBlog(blog domain.Blog) (domain.SuccessResponse, domain.ErrorResponse) {
	
	ctx , cancel := context.WithTimeout(context.Background(), uc.contextTimeout)
	blog.IsResource = false
	defer cancel()	

	if blog.Title == "" || blog.Description == "" {
		return domain.SuccessResponse{}, domain.ErrorResponse{Message: "Title or Description cannot be empty", Status: 400}
	}
	blog.ID = primitive.NewObjectID()
	
	
	// get author name and profile image url
	// get author name and profile image url

	user, err := uc.UserRepository.GetUserByID(ctx, blog.UserID)

	if err != nil {
		return domain.SuccessResponse{}, domain.ErrorResponse{Message: "Failed to create blog", Status: 500}
	}
	
	blog.Name = user.Name
	blog.Image = user.Image
	blog.CreatedAt = time.Now()
	_, err  = uc.BlogRepository.CreateBlog(ctx, blog)
	
	if err != nil {
		return domain.SuccessResponse{}, domain.ErrorResponse{Message: "Failed to create blog", Status: 500}
	}


	if err != nil {
		return domain.SuccessResponse{}, domain.ErrorResponse{Message: "Failed to create blog", Status: 500}
	}

	return domain.SuccessResponse{Message: "Blog created successfully" , Status: 200}, domain.ErrorResponse{}
}

//  get all blog implementation

func (uc *BlogUseCase) GetAllBlog(tag, limit, page string) ([]domain.Blog, error) {
	ctx , cancel := context.WithTimeout(context.Background(), uc.contextTimeout)
	defer cancel()
	return uc.BlogRepository.GetAllBlog(ctx , tag, limit, page)
}

// get blog by id


func (uc *BlogUseCase) GetBlogByID(id string) (domain.Blog, error) {
	ctx , cancel := context.WithTimeout(context.Background(), uc.contextTimeout)
	defer cancel()
	return uc.BlogRepository.GetBlogByID(ctx, id)
}

func (uc *BlogUseCase) CreateReply(reply domain.Reply) (domain.Reply, error) {
	ctx , cancel := context.WithTimeout(context.Background(), uc.contextTimeout)
	defer cancel()
	reply.ID = primitive.NewObjectID()	

	user, err := uc.UserRepository.GetUserByID(ctx, reply.UserID)

	if err != nil {
		return domain.Reply{}, err
	}
	fmt.Println(user)
	reply.Name = user.Name
	reply.Image = user.Image
	reply.CreatedAt = time.Now()

	return uc.ReplyRepository.CreateReply(ctx, reply)
}


func (uc *BlogUseCase) GetReply(blogID string) ([]domain.Reply, error) {
	ctx , cancel := context.WithTimeout(context.Background(), uc.contextTimeout)
	defer cancel()

	return uc.ReplyRepository.GetReply(ctx, blogID)
}


func (uc *BlogUseCase) CreateResouce(blog domain.Blog) (domain.SuccessResponse, domain.ErrorResponse) {


	blog.IsResource = true

	blog.ID = primitive.NewObjectID()
	ctx , cancel := context.WithTimeout(context.Background(), uc.contextTimeout)

	defer cancel()


	_, err := uc.BlogRepository.CreateResource(ctx, blog)

	if err.Message != "" {
		return domain.SuccessResponse{}, domain.ErrorResponse{Message: "Failed to create resource", Status: 500}
	}

	return domain.SuccessResponse{Message: "Resource created successfully" , Status: 200}, domain.ErrorResponse{}


}

func (uc *BlogUseCase) GetResource() ([]domain.Blog, error) {
	ctx , cancel := context.WithTimeout(context.Background(), uc.contextTimeout)
	defer cancel()

	return uc.BlogRepository.GetResource(ctx)
}

