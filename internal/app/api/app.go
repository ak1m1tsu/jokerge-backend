package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

//	@title		jokerge
//	@version	1.0

//	@host		localhost:8000
//	@BasePath	/api/v1

//	@externalDocs.description	OpenAPI
//	@externalDocs.url			https://swagger.io/resources/open-api/

type Env struct {
	app *fiber.App
}

func New() *Env {
	env := Env{
		app: fiber.New(),
	}

	api := env.app.Group("/api").Use(requestid.New(), logger.New())

	v1 := api.Group("/v1")

	v1.Route("/order", func(router fiber.Router) {
		router.Get("/list", env.OrderList)
		router.Get("/:id<guid>", env.OrderGet)
		router.Post("/", env.OrderCreate)
		router.Post("/update", env.OrderUpdate)
	})

	v1.Route("/customer", func(router fiber.Router) {
		router.Get("/list", env.CustomerList)
		router.Get("/:id<guid>", env.CustomerGet)
		router.Post("/", env.CustomerCreate)
	})

	return &env
}

func (e *Env) Run() error {
	return e.app.Listen(":8000")
}
