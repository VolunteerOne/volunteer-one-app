package service

import (
	"github.com/VolunteerOne/volunteer-one-app/backend/models"
	"github.com/VolunteerOne/volunteer-one-app/backend/repository"
)

type PostsService interface {
	CreatePost(post models.Posts) (models.Posts, error)
	DeletePost(post models.Posts) error
	EditPost(post models.Posts) (models.Posts, error)
	FindPost(id string) (models.Posts, error)
	AllPosts() ([]models.Posts, error)
}

type postsService struct {
	postsRepository repository.PostsRepository
}

func NewPostsService(r repository.PostsRepository) PostsService {
	return postsService{
		postsRepository: r,
	}
}

type CommentsService interface {
	CreateComment(comment models.Comments) (models.Comments, error)
	DeleteComment(comment models.Comments) error
	EditComment(comment models.Comments) (models.Comments, error)
	FindComment(id string) (models.Comments, error)
	AllComments() ([]models.Comments, error)
}

type commentsService struct {
	commentsRepository repository.CommentsRepository
}

func NewCommentsService(r repository.CommentsRepository) CommentsService {
	return commentsService{
		commentsRepository: r,
	}
}

type LikesService interface {
	CreateLike(like models.Likes) (models.Likes, error)
	DeleteLike(like models.Likes) error
	FindLike(id string) (models.Likes, error)
	AllLikes() ([]models.Likes, error)
}

type likesService struct {
	likesRepository repository.LikesRepository
}

func NewLikesService(r repository.LikesRepository) LikesService {
	return likesService{
		likesRepository: r,
	}
}

func (f postsService) CreatePost(post models.Posts) (models.Posts, error) {
	return f.postsRepository.CreatePost(post)
}

func (f postsService) DeletePost(post models.Posts) error {
	return f.postsRepository.DeletePost(post)
}

func (f postsService) EditPost(post models.Posts) (models.Posts, error) {
	return f.postsRepository.EditPost(post)
}

func (f postsService) FindPost(id string) (models.Posts, error) {
	return f.postsRepository.FindPost(id)
}

func (f postsService) AllPosts() ([]models.Posts, error) {
	return f.postsRepository.AllPosts()
}

func (f commentsService) CreateComment(comment models.Comments) (models.Comments, error) {
	return f.commentsRepository.CreateComment(comment)
}

func (f commentsService) DeleteComment(comment models.Comments) error {
	return f.commentsRepository.DeleteComment(comment)
}

func (f commentsService) EditComment(comment models.Comments) (models.Comments, error) {
	return f.commentsRepository.EditComment(comment)
}

func (f commentsService) FindComment(id string) (models.Comments, error) {
	return f.commentsRepository.FindComment(id)
}

func (f commentsService) AllComments() ([]models.Comments, error) {
	return f.commentsRepository.AllComments()
}

func (f likesService) CreateLike(like models.Likes) (models.Likes, error) {
	return f.likesRepository.CreateLike(like)
}

func (f likesService) DeleteLike(like models.Likes) error {
	return f.likesRepository.DeleteLike(like)
}

func (f likesService) FindLike(id string) (models.Likes, error) {
	return f.likesRepository.FindLike(id)
}

func (f likesService) AllLikes() ([]models.Likes, error) {
	return f.likesRepository.AllLikes()
}
