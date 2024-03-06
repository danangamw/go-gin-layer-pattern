package service

import (
	"go-gin/dto"
	"go-gin/entity"
	"go-gin/errorhandler"
	"go-gin/repository"

	"github.com/go-playground/validator/v10"
)

type PostService interface {
	Create(req *dto.PostRequest) error
}

type postService struct {
	repository repository.PostRepository
	validator  *validator.Validate
}

func NewPostService(r repository.PostRepository) *postService {
	return &postService{
		repository: r,
		validator:  validator.New(),
	}
}

func (s *postService) Create(req *dto.PostRequest) error {
	if err := s.validator.Struct(req); err != nil {
		return &errorhandler.BadRequestError{Message: err.Error()}
	}
	post := entity.Post{
		UserID: req.UserID,
		Tweet:  req.Tweet,
	}

	if req.Picture != nil {
		post.PictureUrl = &req.Picture.Filename
	}

	if err := s.repository.Create(&post); err != nil {
		return &errorhandler.InternalServerError{
			Message: err.Error(),
		}
	}

	return nil
}
