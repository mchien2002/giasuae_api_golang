package services

import (
	"giasuaeapi/src/entities"
	"giasuaeapi/src/repositories"
)

type PostService interface {
	InsertPost(p *entities.Post) error
	UpdatePost(p *entities.Post) error
	DeletePost(id int) error
	FindAllPost() []entities.Post
	FindByID(id int) entities.Post
}
type postService struct {
	PostRepository repositories.PostRepository
}

// DeletePost implements PostService
func (svc *postService) DeletePost(id int) error {
	return svc.PostRepository.DeletePost(id)
}

// FindAllPost implements PostService
func (svc *postService) FindAllPost() []entities.Post {
	return svc.PostRepository.FindAllPost()
}

// FindByID implements PostService
func (svc *postService) FindByID(id int) entities.Post {
	return svc.PostRepository.FindByID(id)
}

// InsertPost implements PostService
func (svc *postService) InsertPost(p *entities.Post) error {
	return svc.PostRepository.InsertPost(p)
}

// UpdatePost implements PostService
func (svc *postService) UpdatePost(p *entities.Post) error {
	return svc.PostRepository.UpdatePost(p)
}

func NewPostService(PostRepo repositories.PostRepository) PostService {
	return &postService{
		PostRepository: PostRepo,
	}
}
