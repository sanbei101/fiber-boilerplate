package article

import (
	"github.com/efectn/fiber-boilerplate/app/module/article/controller"
	"github.com/efectn/fiber-boilerplate/app/module/article/service"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
)

type ArticleRouter struct {
	App        fiber.Router
	Controller *controller.ArticleController
}

// Register bulkly
var NewArticleModule = fx.Options(
	// Register Service
	fx.Provide(service.NewArticleService),

	// Register Controller
	fx.Provide(controller.NewArticleController),

	// Register Router
	fx.Provide(NewArticleRouter),
)

// Router methods
func NewArticleRouter(fiber *fiber.App, articleController *controller.ArticleController) *ArticleRouter {
	return &ArticleRouter{
		App:        fiber,
		Controller: articleController,
	}
}

func (r *ArticleRouter) RegisterArticleRoutes() {
	// Define routes
	r.App.Route("/articles", func(router fiber.Router) {
		router.Get("/", r.Controller.Index)
		router.Get("/:id", r.Controller.Show)
		router.Post("/", r.Controller.Store)
		router.Patch("/:id", r.Controller.Update)
		router.Delete("/:id", r.Controller.Destroy)
	})
}
