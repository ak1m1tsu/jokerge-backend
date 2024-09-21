package admin

import (
	"github.com/ak1m1tsu/jokerge/internal/pkg/service"
	"github.com/gofiber/fiber/v2"
)

// @title admin
// @version 1.0

// @host localhost:8080
// @BasePath /api/v1

//	@externalDocs.description	OpenAPI
//	@externalDocs.url			https://swagger.io/resources/open-api/

type Env struct {
	app *fiber.App
	srv *service.Service
}

func New() *Env {
	env := &Env{
		app: fiber.New(),
		srv: service.New(),
	}

	return env
}

func (e *Env) Run() error {
	return e.app.Listen(":8080")
}

func (e *Env) Service() *service.Service {
	return e.srv
}
