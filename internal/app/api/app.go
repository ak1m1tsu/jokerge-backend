package api

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/ak1m1tsu/jokerge/internal/pkg/middleware"
	"github.com/ak1m1tsu/jokerge/internal/pkg/service"
	"github.com/ak1m1tsu/jokerge/internal/pkg/types"
	"github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/google/uuid"
)

//	@title		jokerge
//	@version	1.0

//	@host		localhost:8000
//	@BasePath	/api/v1

//	@externalDocs.description	OpenAPI
//	@externalDocs.url			https://swagger.io/resources/open-api/

type Env struct {
	app *fiber.App
	srv *service.Service
}

func New() (*Env, error) {
	srv, err := service.New()
	if err != nil {
		return nil, err
	}

	env := &Env{
		app: fiber.New(fiber.Config{
			ServerHeader: "X-Server",
			ReadTimeout:  time.Second * 5,
			WriteTimeout: time.Second * 5,
			IdleTimeout:  time.Second * 30,
			ErrorHandler: HandleError,
			JSONEncoder:  sonic.Marshal,
			JSONDecoder:  sonic.Unmarshal,
		}),
		srv: srv,
	}

	if err = env.SeedData(); err != nil {
		return nil, err
	}

	api := env.app.Group("/api")

	api.Use(middleware.RequestID())
	api.Use(middleware.Logger())
	api.Use(cors.New())

	auth := api.Group("/auth")
	auth.Post("/", env.ValidateUserCredentials)

	v1 := api.Group("/v1")
	v1.Use(basicauth.New(basicauth.Config{
		Realm:      "Forbidden",
		Authorizer: env.Authorizer,
	}))

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

	return env, nil
}

func (e *Env) Run() error {
	return e.app.Listen(":8000")
}

func (e *Env) Service() *service.Service {
	return e.srv
}

func (e *Env) NotFound(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusNotFound).JSON(Response{Error: "not found"})
}

func (e *Env) OK(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(Response{Status: "ok"})
}

func (e *Env) Authorizer(email, pass string) bool {
	return true
}

func (e *Env) SeedData() error {
	var (
		err error
		ctx = context.Background()
	)

	if _, err = e.Service().DB().NewCreateTable().Model((*types.User)(nil)).Exec(ctx); err != nil {
		return err
	}

	user := &types.User{
		ID:        uuid.NewString(),
		Email:     "admin@admin.com",
		Password:  "SuperPassword",
		FirstName: "Иван",
		LastName:  "Иванов",
	}
	if _, err = e.Service().DB().NewInsert().Model(user).Exec(ctx); err != nil {
		return err
	}

	return nil
}

func HandleError(ctx *fiber.Ctx, err error) error {
	code := http.StatusInternalServerError

	var e *fiber.Error
	if errors.As(err, &e) {
		code = e.Code
	}

	switch code {
	case http.StatusMethodNotAllowed:
		return ctx.Status(code).JSON(Response{Error: "method not allowed"})
	case http.StatusNotFound:
		return ctx.Status(code).JSON(Response{Error: "not found"})
	default:
		return ctx.Status(code).JSON(Response{Error: "internal"})
	}
}
