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
//	@BasePath	/

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

	tx, err := e.Service().DB().Begin()
	if err != nil {
		return err
	}
	if _, err = tx.NewCreateTable().Model((*types.UserModel)(nil)).Exec(ctx); err != nil {
		return err
	}

	if _, err = tx.NewCreateTable().Model((*types.ProductModel)(nil)).Exec(ctx); err != nil {
		return err
	}

	if _, err = tx.NewCreateTable().Model((*types.CustomerModel)(nil)).Exec(ctx); err != nil {
		return err
	}

	if _, err = tx.NewCreateTable().Model((*types.OrderModel)(nil)).Exec(ctx); err != nil {
		return err
	}

	if _, err = tx.NewCreateTable().Model((*types.OrderItemModel)(nil)).Exec(ctx); err != nil {
		return err
	}

	user := &types.UserModel{
		ID:        uuid.NewString(),
		Email:     "admin@admin.com",
		Password:  "SuperPassword",
		FirstName: "Иван",
		LastName:  "Иванов",
	}
	if _, err = tx.NewInsert().Model(user).Exec(ctx); err != nil {
		return err
	}

	customers := []*types.CustomerModel{
		{
			ID:        "c22946d7-991e-44a1-b0dc-6b775a34664c",
			FirstName: "Модест",
			LastName:  "Пономарёв",
			Address:   "121248, г. Приозерное, ул. Пяловская, дом 40, квартира 30",
		},
		{
			ID:        "4c8a6931-04cb-4c93-9a90-e7a29342aba2",
			FirstName: "Мстислав",
			LastName:  "Рощин",
			Address:   "396522, г. Надеждино, ул. Жилой поселок 2-й Заельцовского Бора тер, дом 41, квартира 96",
		},
		{
			ID:        "7dfa4d19-d4c8-4596-bcb7-249e0de951bc",
			FirstName: "Инесса",
			LastName:  "Химченко",
			Address:   "164505, г. Дятьково, ул. Белоостровская (Выборгский), дом 59, квартира 24",
		},
	}

	for _, customer := range customers {
		if _, err = tx.NewInsert().Model(customer).Exec(ctx); err != nil {
			return err
		}
	}

	return tx.Commit()
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
