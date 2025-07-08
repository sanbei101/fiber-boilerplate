package controller

import (
	"strconv"

	"github.com/efectn/fiber-boilerplate/app/module/article/service"
	"github.com/efectn/fiber-boilerplate/utils/response"
	"github.com/gofiber/fiber/v2"
)

// Requests & responses for ArticleController & ArticleService
type ArticleRequest struct {
	Title   string `json:"title" form:"title" validate:"required,max=255"`
	Content string `json:"content" form:"content" validate:"required"`
}

type ArticleController struct {
	articleService *service.ArticleService
}

type IArticleController interface {
	Index(c *fiber.Ctx) error
	Show(c *fiber.Ctx) error
	Store(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
	Destroy(c *fiber.Ctx) error
}

func NewArticleController(articleService *service.ArticleService) *ArticleController {
	return &ArticleController{
		articleService: articleService,
	}
}

func (con *ArticleController) Index(c *fiber.Ctx) error {
	articles, err := con.articleService.GetArticles()
	if err != nil {
		return response.NewError(
			fiber.StatusInternalServerError,
			"Failed to retrieve articles",
		)
	}

	return response.Resp(c, response.Response{
		Messages: response.Messages{"Article list retreived successfully!"},
		Data:     articles,
	})
}

func (con *ArticleController) Show(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return response.NewError(
			fiber.StatusBadRequest,
			"Invalid article ID",
		)
	}

	article, err := con.articleService.GetArticleByID(uint(id))
	if err != nil {
		return response.NewError(
			fiber.StatusNotFound,
			"Article not found",
		)
	}

	return response.Resp(c, response.Response{
		Messages: response.Messages{"The article retrieved successfully!"},
		Data:     article,
	})
}

func (con *ArticleController) Store(c *fiber.Ctx) error {
	req := new(ArticleRequest)
	if err := response.ParseAndValidate(c, req); err != nil {
		return response.NewError(
			fiber.StatusBadRequest,
			"Invalid input data",
		)
	}

	article, err := con.articleService.CreateArticle(req.Title, req.Content)
	if err != nil {
		return response.NewError(
			fiber.StatusInternalServerError,
			"Failed to create article",
		)
	}

	return response.Resp(c, response.Response{
		Messages: response.Messages{"The article was created successfully!"},
		Data:     article,
	})
}

func (con *ArticleController) Update(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return response.NewError(
			fiber.StatusBadRequest,
			"Invalid article ID",
		)
	}

	req := new(ArticleRequest)
	if err := response.ParseAndValidate(c, req); err != nil {
		return response.NewError(
			fiber.StatusBadRequest,
			"Invalid input data",
		)
	}

	article, err := con.articleService.UpdateArticle(uint(id), req.Title, req.Content)
	if err != nil {
		return response.NewError(
			fiber.StatusInternalServerError,
			"Failed to update article",
		)
	}

	return response.Resp(c, response.Response{
		Messages: response.Messages{"The article was updated successfully!"},
		Data:     article,
	})
}

func (con *ArticleController) Destroy(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return response.NewError(
			fiber.StatusBadRequest,
			"Invalid article ID",
		)
	}

	if err = con.articleService.DeleteArticle(uint(id)); err != nil {
		return response.NewError(
			fiber.StatusInternalServerError,
			"Failed to delete article",
		)
	}

	return response.Resp(c, response.Response{
		Messages: response.Messages{"The article was deleted successfully!"},
	})
}
