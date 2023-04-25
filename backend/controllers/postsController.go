package controllers

import (
	"net/http"

	"github.com/VolunteerOne/volunteer-one-app/backend/models"
	"github.com/VolunteerOne/volunteer-one-app/backend/service"
	"github.com/gin-gonic/gin"
)

type PostsController interface {
	CreatePost(c *gin.Context)
	DeletePost(c *gin.Context)
	EditPost(c *gin.Context)
	FindPost(c *gin.Context)
	AllPosts(c *gin.Context)
}

type postsController struct {
	postsService service.PostsService
}

func NewPostsController(s service.PostsService) PostsController {
	return postsController{
		postsService: s,
	}
}

var postsModel = new(models.Posts)

func (controller postsController) CreatePost(c *gin.Context) {
	var err error
	var body struct {
		Handle          string
		PostDescription string
	}
	err = c.Bind(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Request body is invalid",
		})
		return
	}

	object := models.Posts{
		Handle:          body.Handle,
		PostDescription: body.PostDescription,
	}

	result, err := controller.postsService.CreatePost(object)

	_ = result

	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Creation failed",
		})

		return
	}

	// Respond
	c.JSON(http.StatusOK, object)
}

func (controller postsController) DeletePost(c *gin.Context) {
	id := c.Param("id")

	result, err := controller.postsService.FindPost(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Could not retrieve object",
		})

		return
	}

	// Delete the object
	err1 := controller.postsService.DeletePost(result)

	if err1 != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Could not delete object",
		})

		return
	}

	// Respond
	c.JSON(http.StatusOK, gin.H{
		"message": "Object deleted successfully",
	})
}

func (controller postsController) EditPost(c *gin.Context) {
	id := c.Param("id")
	result, err1 := controller.postsService.FindPost(id)

	var err error
	var body struct {
		PostDescription string
	}
	err = c.Bind(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Request body is invalid",
		})
		return
	}
	if err1 != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Could not retrieve object",
		})

		return
	}

	result.PostDescription = body.PostDescription

	results1, err := controller.postsService.EditPost(result)

	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Creation failed",
		})

		return
	}

	// Respond
	c.JSON(http.StatusOK, results1)
}

func (controller postsController) FindPost(c *gin.Context) {
	id := c.Param("id")

	result, err1 := controller.postsService.FindPost(id)

	if err1 != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Could not retrieve object",
		})

		return
	}

	// Return the object
	c.JSON(http.StatusOK, result)
}

func (controller postsController) AllPosts(c *gin.Context) {
	// Get object from the database
	posts, err := controller.postsService.AllPosts()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Could not retrieve objects",
		})

		return
	}

	// Return the array of objects
	c.JSON(http.StatusOK, posts)
}

type CommentsController interface {
	CreateComment(c *gin.Context)
	DeleteComment(c *gin.Context)
	EditComment(c *gin.Context)
	FindComment(c *gin.Context)
	AllComments(c *gin.Context)
}

type commentsController struct {
	commentsService service.CommentsService
}

func NewCommentsController(s service.CommentsService) CommentsController {
	return commentsController{
		commentsService: s,
	}
}

var CommentsModel = new(models.Comments)

func (controller commentsController) CreateComment(c *gin.Context) {
	var err error
	var body struct {
		PostsID            uint
		Handle             string
		CommentDescription string
	}
	err = c.Bind(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Request body is invalid",
		})
		return
	}

	object := models.Comments{
		PostsID:            body.PostsID,
		Handle:             body.Handle,
		CommentDescription: body.CommentDescription,
	}

	result, err := controller.commentsService.CreateComment(object)

	_ = result

	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Creation failed",
		})

		return
	}

	// Respond
	c.JSON(http.StatusOK, object)
}

func (controller commentsController) DeleteComment(c *gin.Context) {
	id := c.Param("id")

	result, err := controller.commentsService.FindComment(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Could not retrieve object",
		})

		return
	}

	// Delete the object
	err1 := controller.commentsService.DeleteComment(result)

	if err1 != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Could not delete object",
		})

		return
	}

	// Respond
	c.JSON(http.StatusOK, gin.H{
		"message": "Object deleted successfully",
	})
}

func (controller commentsController) EditComment(c *gin.Context) {
	id := c.Param("id")
	result, err1 := controller.commentsService.FindComment(id)

	var err error
	var body struct {
		CommentDescription string
	}
	err = c.Bind(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Request body is invalid",
		})
		return
	}
	if err1 != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Could not retrieve object",
		})

		return
	}

	result.CommentDescription = body.CommentDescription

	results1, err := controller.commentsService.EditComment(result)

	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Creation failed",
		})

		return
	}

	// Respond
	c.JSON(http.StatusOK, results1)
}

func (controller commentsController) FindComment(c *gin.Context) {
	id := c.Param("id")

	result, err1 := controller.commentsService.FindComment(id)

	if err1 != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Could not retrieve object",
		})

		return
	}

	// Return the object
	c.JSON(http.StatusOK, result)
}

func (controller commentsController) AllComments(c *gin.Context) {
	// Get object from the database
	comments, err := controller.commentsService.AllComments()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Could not retrieve objects",
		})

		return
	}

	// Return the array of objects
	c.JSON(http.StatusOK, comments)
}

type LikesController interface {
	CreateLike(c *gin.Context)
	DeleteLike(c *gin.Context)
	FindLike(c *gin.Context)
	AllLikes(c *gin.Context)
	GetLikes(c *gin.Context)
}

type likesController struct {
	likesService service.LikesService
}

func NewLikesController(s service.LikesService) LikesController {
	return likesController{
		likesService: s,
	}
}

var LikesModel = new(models.Likes)

func (controller likesController) CreateLike(c *gin.Context) {
	var err error
	var body struct {
		Handle  string
		PostsID uint
	}
	err = c.Bind(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Request body is invalid",
		})
		return
	}

	object := models.Likes{
		Handle:  body.Handle,
		PostsID: body.PostsID,
	}

	result, err := controller.likesService.CreateLike(object)

	_ = result

	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Creation failed",
		})

		return
	}

	// Respond
	c.JSON(http.StatusOK, object)
}

func (controller likesController) DeleteLike(c *gin.Context) {
	id := c.Param("id")

	result, err := controller.likesService.FindLike(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Could not retrieve object",
		})

		return
	}

	// Delete the object
	err1 := controller.likesService.DeleteLike(result)

	if err1 != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Could not delete object",
		})

		return
	}

	// Respond
	c.JSON(http.StatusOK, gin.H{
		"message": "Object deleted successfully",
	})
}

func (controller likesController) FindLike(c *gin.Context) {
	id := c.Param("id")

	result, err1 := controller.likesService.FindLike(id)

	if err1 != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Could not retrieve object",
		})

		return
	}

	// Return the object
	c.JSON(http.StatusOK, result)
}

func (controller likesController) AllLikes(c *gin.Context) {
	// Get object from the database
	likes, err := controller.likesService.AllLikes()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Could not retrieve objects",
		})

		return
	}

	// Return the array of objects
	c.JSON(http.StatusOK, likes)
}

func (controller likesController) GetLikes(c *gin.Context) {
	id := c.Param("id")
	// Get object from the database
	likes, err := controller.likesService.GetLikes(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Could not retrieve objects",
		})

		return
	}

	// Return the array of objects
	c.JSON(http.StatusOK, likes)
}
