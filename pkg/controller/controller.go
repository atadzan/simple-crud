package controller

import (
	"github.com/atadzan/simple-crud/config"
	"github.com/atadzan/simple-crud/pkg/repository"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/swagger"
)

type Controller struct {
	repo       repository.Repo
	authConfig config.Authorization
}

func New(repo repository.Repo, authCfg config.Authorization) *Controller {
	return &Controller{
		repo:       repo,
		authConfig: authCfg,
	}
}

func (ctl *Controller) InitRoutes() (app *fiber.App) {
	app = fiber.New(
		fiber.Config{
			AppName:           "Simple CRUD app",
			BodyLimit:         100 * 1024 * 1024,
			EnablePrintRoutes: true,
			StrictRouting:     true,
		})

	app.Use(
		cors.New(cors.Config{
			AllowOrigins:     "*",
			AllowCredentials: false,
			AllowHeaders:     "Origin, Content-Length, Content-Type, User-Agent, Referrer, Host, Token, Authorization",
			AllowMethods:     "GET, POST, PUT, DELETE",
		}),
		logger.New(),
	)

	app.Get("/swagger/*", swagger.HandlerDefault)
	v1 := app.Group("/v1")
	{
		v1.Post("/register", ctl.register)
		v1.Post("/signIn", ctl.signIn)

		v1.Get("/genres", ctl.getBookGenres)

		books := v1.Group("/books")
		{
			books.Get("/", ctl.getBooks)
			books.Get("/search", ctl.searchBook)
			books.Get("/:id", ctl.getBookById)
		}
		authorized := books.Group("/")
		{
			authorized.Post("/", ctl.createBook)
			authorized.Patch("/:id", ctl.updateBook)
			authorized.Delete("/:id", ctl.deleteBook)
			authorized.Get("/img/:path/download", ctl.downloadBookIMG)
		}
	}
	return
}
