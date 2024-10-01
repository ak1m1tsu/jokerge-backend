package api

import (
	"github.com/ak1m1tsu/jokerge/internal/pkg/types"
	"github.com/gofiber/fiber/v2"
	zlog "github.com/rs/zerolog"
)

// ProductList возвращает спиок продуктов
//
//	@Summary	список продуктов
//	@Tags		products
//	@Security	BasicAuth
//	@Product	json
//	@Param		X-Request-ID	header		string	true	"ID запроса"
//	@Success	200				{object}	types.ProductListResponse
//	@Failure	500				{object}	types.APIResponse
//	@Router		/api/v1/product/list [get]
func (e *Env) ProductList(ctx *fiber.Ctx) error {
	products, err := e.Service().GetProducts(ctx.UserContext())
	if err != nil {
		zlog.Ctx(ctx.UserContext()).Error().Err(err).Msg("failed to get product list")
		return err
	}

	response := make(types.ProductListResponse, 0, len(products))

	for _, p := range products {
		response = append(response, types.ProductInfoResponse{
			ID:          p.ID,
			Name:        p.Name,
			Description: p.Description,
			Price:       p.Price,
		})
	}

	return ctx.JSON(response)
}

// ProductGet возвращает информацию о продукте
//
//	@Summary	информация о продукте
//	@Tags		products
//	@Security	BasicAuth
//	@Produce	json
//	@Param		product_id		path		string	true	"ID продукта"
//	@Param		X-Request-ID	header		string	true	"ID запроса"
//	@Success	200				{object}	types.ProductInfoResponse
//	@Failure	404				{object}	types.APIResponse
//	@Failure	500				{object}	types.APIResponse
//	@Router		/api/v1/product/{product_id} [get]
func (e *Env) ProductGet(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	product, err := e.Service().GetProductInfo(ctx.UserContext(), id)
	if err != nil {
		zlog.Ctx(ctx.UserContext()).Error().Err(err).Msgf("failed to get product by id %s", id)
		return err
	}

	if product == nil {
		return e.NotFound(ctx)
	}

	return ctx.JSON(types.ProductInfoResponse{
		ID:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
	})
}

// ProductCreate обработчик создания нового продукта
//
//	@Summary	создание нового продукта
//	@Tags		products
//	@Security	BasicAuth
//	@Accept		json
//	@Produce	json
//	@Param		body			body		types.ProductCreateBody	true	"Тело запроса"
//	@Param		X-Request-ID	header		string					true	"ID запроса"
//	@Success	200				{object}	types.ProductInfoResponse
//	@Failure	400				{object}	types.APIResponse
//	@Failure	500				{object}	types.APIResponse
//	@Router		/api/v1/product [post]
func (e *Env) ProductCreate(ctx *fiber.Ctx) error {
	var body types.ProductCreateBody
	if err := ctx.BodyParser(&body); err != nil {
		zlog.Ctx(ctx.UserContext()).Error().Err(err).Msg("failed to parse request body")
		return err
	}

	if err := body.Validate(); err != nil {
		zlog.Ctx(ctx.UserContext()).Error().Err(err).Msg("invalid body request")
		return err
	}

	zlog.Ctx(ctx.UserContext()).Info().Any("body", body).Msg("create new product")

	id, err := e.Service().CreateProduct(ctx.UserContext(), body)
	if err != nil {
		zlog.Ctx(ctx.UserContext()).Error().Err(err).Msg("failed to create product")
		return err
	}

	return ctx.JSON(types.ProductInfoResponse{
		ID:          id,
		Name:        body.Name,
		Description: body.Description,
		Price:       body.Price,
	})
}

// ProductUpdate обработчик обновления информации о продукте
//
//	@Summary	обновление информации о продукте
//	@Tags		products
//	@Security	BasicAuth
//	@Accept		json
//	@Produce	json
//	@Param		body			body		types.ProductUpdateBody	true	"Тело запроса"
//	@Param		X-Request-ID	header		string					true	"ID запроса"
//	@Success	200				{object}	types.APIResponse
//	@Failure	400				{object}	types.APIResponse
//	@Failure	404				{object}	types.APIResponse
//	@Failure	500				{object}	types.APIResponse
//	@Router		/api/v1/product/update [post]
func (e *Env) ProductUpdate(ctx *fiber.Ctx) error {
	var body types.ProductUpdateBody
	if err := ctx.BodyParser(&body); err != nil {
		zlog.Ctx(ctx.UserContext()).Error().Err(err).Msg("failed to parse request body")
		return err
	}

	if err := body.Validate(); err != nil {
		zlog.Ctx(ctx.UserContext()).Error().Err(err).Msg("invalid body request")
		return err
	}

	zlog.Ctx(ctx.UserContext()).Info().Any("body", body).Msg("update a product")

	if err := e.Service().UpdateProduct(ctx.UserContext(), body); err != nil {
		return err
	}

	return e.OK(ctx)
}
