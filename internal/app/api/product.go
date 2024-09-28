package api

import (
	"github.com/ak1m1tsu/jokerge/internal/pkg/types"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
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
//	@Router		/api/v1/produuct/list [get]
func (e *Env) ProductList(ctx *fiber.Ctx) error {
	products, err := e.Service().GetProducts(ctx.UserContext())
	if err != nil {
		zerolog.Ctx(ctx.UserContext()).Error().Err(err).Msg("failed to get product list")
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
		zerolog.Ctx(ctx.UserContext()).Error().Err(err).Msgf("failed to get product by id %s", id)
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
