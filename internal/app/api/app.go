package api

import (
	"context"
	"database/sql"
	"errors"
	"net/http"
	"time"

	_ "github.com/ak1m1tsu/jokerge/api"
	"github.com/ak1m1tsu/jokerge/internal/pkg/middleware"
	"github.com/ak1m1tsu/jokerge/internal/pkg/service"
	"github.com/ak1m1tsu/jokerge/internal/pkg/types"
	"github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/redirect"
	"github.com/gofiber/swagger"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

//	@title		jokerge
//	@version	1.0

//	@host		localhost:8000
//	@BasePath	/

//	@securityDefinitions.basic	BasicAuth

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
			ServerHeader: "jokerge",
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

	env.app.Use(redirect.New(redirect.Config{
		Rules: map[string]string{
			"/": "/swagger",
		},
	}))
	env.app.Get("/swagger/*", swagger.New(swagger.Config{
		Title: "Jokerge API",
	}))

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
		Unauthorized: func(c *fiber.Ctx) error {
			return c.Status(http.StatusUnauthorized).JSON(types.APIResponse{Error: "unauthorized"})
		},
	}))

	v1.Route("/order", func(router fiber.Router) {
		router.Get("/list", env.OrderList)
		router.Get("/:id<int>", env.OrderGet)
		router.Post("/", env.OrderCreate)
		router.Post("/update", env.OrderUpdate)
	})

	v1.Route("/customer", func(router fiber.Router) {
		router.Get("/list", env.CustomerList)
		router.Get("/:id<guid>", env.CustomerGet)
		router.Post("/", env.CustomerCreate)
		router.Post("/update", env.CustomerUpdate)
	})

	v1.Route("/product", func(router fiber.Router) {
		router.Get("/list", env.ProductList)
		router.Get("/:id<guid>", env.ProductGet)
		router.Post("/", env.ProductCreate)
		router.Post("/update", env.ProductUpdate)
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
	return ctx.Status(http.StatusNotFound).JSON(types.APIResponse{Error: "not found"})
}

func (e *Env) OK(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(types.APIResponse{Message: "ok"})
}

func (e *Env) Authorizer(email, pass string) bool {
	_, ok, err := e.Service().ValidateUser(context.Background(), email, pass)
	if err != nil {
		log.Error().Err(err).Msg("failed to validate user")
		return false
	}

	if !ok {
		log.Error().Msgf("invalid user credentails: %s:%s", email, pass)
		return false
	}

	return true
}

func (e *Env) SeedData() error {
	var (
		err error
		ctx = context.Background()
	)

	e.Service().DB().RegisterModel((*types.OrderItemModel)(nil))

	tx, err := e.Service().DB().Begin()
	if err != nil {
		return err
	}

	if _, err = tx.NewCreateTable().IfNotExists().Model((*types.UserModel)(nil)).Exec(ctx); err != nil {
		return err
	}

	if _, err = tx.NewCreateTable().IfNotExists().Model((*types.ProductModel)(nil)).Exec(ctx); err != nil {
		return err
	}

	if _, err = tx.NewCreateTable().IfNotExists().Model((*types.CustomerModel)(nil)).Exec(ctx); err != nil {
		return err
	}

	if _, err = tx.NewCreateTable().IfNotExists().Model((*types.OrderModel)(nil)).Exec(ctx); err != nil {
		return err
	}

	if _, err = tx.NewCreateTable().IfNotExists().Model((*types.OrderItemModel)(nil)).Exec(ctx); err != nil {
		return err
	}

	user := &types.UserModel{
		ID:        uuid.NewString(),
		Email:     "admin",
		Password:  "admin",
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

	if _, err = tx.NewInsert().Model(&customers).Exec(ctx); err != nil {
		return err
	}

	orders := []types.OrderModel{
		{
			ID:         1,
			CustomerID: "c22946d7-991e-44a1-b0dc-6b775a34664c",
			Status:     types.OrderStatusActive,
			Price:      100,
			CreatedAt:  time.Now().Unix(),
		},
		{
			ID:         2,
			CustomerID: "c22946d7-991e-44a1-b0dc-6b775a34664c",
			Status:     types.OrderStatusCanceled,
			Price:      150,
			CreatedAt:  time.Now().Unix(),
		},
	}

	if _, err = tx.NewInsert().Model(&orders).Exec(ctx); err != nil {
		return err
	}

	products := []types.ProductModel{
		{
			ID:          "b94f9683-4b11-4eab-ac7a-1fc66e22ce6e",
			Name:        "Item 1",
			Description: "some description",
			Price:       100,
		},
	}

	if _, err := tx.NewInsert().Model(&products).Exec(ctx); err != nil {
		return err
	}

	orderItems := []types.OrderItemModel{
		{
			OrderID:   1,
			ProductID: "b94f9683-4b11-4eab-ac7a-1fc66e22ce6e",
			Count:     10,
		},
		{
			OrderID:   2,
			ProductID: "b94f9683-4b11-4eab-ac7a-1fc66e22ce6e",
			Count:     15,
		},
	}

	if _, err := tx.NewInsert().Model(&orderItems).Exec(ctx); err != nil {
		return err
	}

	return tx.Commit()
}

func HandleError(ctx *fiber.Ctx, err error) error {
	if errors.Is(sql.ErrNoRows, err) {
		return ctx.Status(http.StatusNotFound).JSON(types.APIResponse{Error: "not found"})
	}

	code := http.StatusInternalServerError

	var apiErr *fiber.Error
	if errors.As(err, &apiErr) {
		code = apiErr.Code
	}

	switch code {
	case http.StatusMethodNotAllowed:
		return ctx.Status(code).JSON(types.APIResponse{Error: "method not allowed"})
	case http.StatusBadRequest:
		return ctx.Status(code).JSON(types.APIResponse{Error: "bad request"})
	case http.StatusNotFound:
		return ctx.Status(code).JSON(types.APIResponse{Error: "not found"})
	default:
		return ctx.Status(code).JSON(types.APIResponse{Error: "internal"})
	}
}
