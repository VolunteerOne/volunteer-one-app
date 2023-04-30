package repository

import (
	"errors"

	"github.com/VolunteerOne/volunteer-one-app/backend/models"
	"gorm.io/gorm"
)

type PostsRepository interface {
	CreatePost(post models.Posts) (models.Posts, error)
	DeletePost(post models.Posts) error
	EditPost(post models.Posts) (models.Posts, error)
	FindPost(id string) (models.Posts, error)
	AllPosts() ([]models.Posts, error)
}
type CommentsRepository interface {
	CreateComment(Comment models.Comments) (models.Comments, error)
	DeleteComment(Comment models.Comments) error
	EditComment(Comment models.Comments) (models.Comments, error)
	FindComment(id string) (models.Comments, error)
	AllComments() ([]models.Comments, error)
}

type LikesRepository interface {
	CreateLike(Like models.Likes) (models.Likes, error)
	DeleteLike(Like models.Likes) error
	FindLike(id string) (models.Likes, error)
	AllLikes() ([]models.Likes, error)
	GetLikes(id string) (int64, error)
}

type postsRepository struct {
	DB *gorm.DB
}

func NewPostsRepository(db *gorm.DB) PostsRepository {
	return postsRepository{
		DB: db,
	}
}

type commentsRepository struct {
	DB *gorm.DB
}

func NewCommentsRepository(db *gorm.DB) CommentsRepository {
	return commentsRepository{
		DB: db,
	}
}

type likesRepository struct {
	DB *gorm.DB
}

func NewLikesRepository(db *gorm.DB) LikesRepository {
	return likesRepository{
		DB: db,
	}
}

func (r postsRepository) CreatePost(post models.Posts) (models.Posts, error) {

	err := r.DB.Create(&post).Error

	return post, err
}

func (r postsRepository) DeletePost(post models.Posts) error {

	result := r.DB.Where("ID = ?", post.ID).Delete(&post)
	if result.Error != nil {
		return errors.New("could not delete post")
	}

	return nil
}

func (r postsRepository) EditPost(post models.Posts) (models.Posts, error) {

	result := r.DB.Save(&post)
	if result.Error != nil {
		return models.Posts{}, errors.New("could not update post")
	}

	return post, nil
}

func (r postsRepository) FindPost(id string) (models.Posts, error) {
	var post models.Posts
	result := r.DB.Find(&post, id)

	if result.Error != nil {
		return models.Posts{}, errors.New("could not retrive post")
	}
	return post, nil
}

func (r postsRepository) AllPosts() ([]models.Posts, error) {
	var posts []models.Posts
	result := r.DB.Find(&posts)

	if result.Error != nil {
		return []models.Posts{}, errors.New("could not retrive post")
	}
	return posts, nil
}

func (r commentsRepository) CreateComment(comment models.Comments) (models.Comments, error) {

	err := r.DB.Create(&comment).Error

	return comment, err
}

func (r commentsRepository) DeleteComment(comment models.Comments) error {

	result := r.DB.Delete(&comment)
	if result.Error != nil {
		return errors.New("could not delete comment")
	}

	return nil
}

func (r commentsRepository) EditComment(comment models.Comments) (models.Comments, error) {

	result := r.DB.Save(&comment)
	if result.Error != nil {
		return models.Comments{}, errors.New("could not update comment")
	}

	return comment, nil
}

func (r commentsRepository) FindComment(id string) (models.Comments, error) {
	var comment models.Comments
	result := r.DB.Find(&comment, id)

	if result.Error != nil {
		return models.Comments{}, errors.New("could not retrive comment")
	}
	return comment, nil
}

func (r commentsRepository) AllComments() ([]models.Comments, error) {
	var comments []models.Comments
	result := r.DB.Find(&comments)

	if result.Error != nil {
		return []models.Comments{}, errors.New("could not retrive comment")
	}
	return comments, nil
}

func (r likesRepository) CreateLike(like models.Likes) (models.Likes, error) {

	err := r.DB.Create(&like).Error

	return like, err
}

func (r likesRepository) DeleteLike(like models.Likes) error {

	result := r.DB.Delete(&like)
	if result.Error != nil {
		return errors.New("could not delete like")
	}

	return nil
}

func (r likesRepository) FindLike(id string) (models.Likes, error) {
	var likes models.Likes
	result := r.DB.Find(&likes, id)

	if result.Error != nil {
		return models.Likes{}, errors.New("could not retrive like")
	}
	return likes, nil
}

func (r likesRepository) AllLikes() ([]models.Likes, error) {
	var likes []models.Likes
	result := r.DB.Find(&likes)

	if result.Error != nil {
		return []models.Likes{}, errors.New("could not retrive like")
	}
	return likes, nil
}

func (r likesRepository) GetLikes(id string) (int64, error) {
	var likes models.Likes
	var count int64
	result := r.DB.Where("PostID=?", id).Find(&likes).Count(&count)

	if result.Error != nil {
		return count, errors.New("could not retrive like")
	}
	return count, nil
}
