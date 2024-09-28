package api

import (
	"github.com/ak1m1tsu/jokerge/internal/pkg/types"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
)

// OrderList возвращает список заказов
//
//	@Summary	список заказов
//	@Tags		orders
//	@Produce	json
//	@Param		X-Request-ID	header		string	true	"ID запроса"
//	@Success	200				{object}	types.OrderListResponse
//	@Failure	500				{object}	types.APIResponse
//	@Router		/api/v1/order/list [get]
func (e *Env) OrderList(ctx *fiber.Ctx) error {
	orders, err := e.Service().GetOrders(ctx.Context())
	if err != nil {
		zerolog.Ctx(ctx.UserContext()).Error().Err(err).Msg("failed to get orders")
		return err
	}

	response := make(types.OrderListResponse, 0, len(orders))

	for _, o := range orders {
		converted := types.OrderInfoResponse{
			ID:        o.ID,
			Status:    o.Status.String(),
			Price:     o.CalculatePrice(),
			CreatedAt: o.CreatedAt.String(),
			Customer: types.OrderInfoCustomerResponse{
				ID:        o.Customer.ID,
				FirstName: o.Customer.FirstName,
				LastName:  o.Customer.LastName,
				Address:   o.Customer.Address,
			},
			Products: make([]types.OrderInfoProductResponse, 0, len(o.Products)),
		}

		for _, p := range o.Products {
			converted.Products = append(converted.Products, types.OrderInfoProductResponse{
				ID:          p.ID,
				Name:        p.Name,
				Description: p.Description,
				Price:       p.Price,
				Count:       p.Count,
			})
		}

		response = append(response, converted)
	}

	return ctx.JSON(response)
}

// OrderGet информация о заказе
//
//	@Summary	информация о заказе
//	@Tags		orders
//	@Produce	json
//	@Param		order_id		path		int		true	"ID заказа"
//	@Param		X-Request-ID	header		string	true	"ID запроса"
//	@Success	200				{object}	types.OrderInfoResponse
//	@Failure	404				{object}	types.APIResponse
//	@Failure	500				{object}	types.APIResponse
//	@Router		/api/v1/order/{order_id} [get]
func (e *Env) OrderGet(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return err
	}

	order, err := e.Service().GetOrderInfo(ctx.Context(), id)
	if err != nil {
		zerolog.Ctx(ctx.UserContext()).Error().Err(err).Msgf("failed to get order by id %d", id)
		return err
	}

	if order == nil {
		return e.NotFound(ctx)
	}

	response := types.OrderInfoResponse{
		ID:        order.ID,
		Status:    order.Status.String(),
		Price:     order.CalculatePrice(),
		CreatedAt: order.CreatedAt.String(),
		Customer: types.OrderInfoCustomerResponse{
			ID:        order.Customer.ID,
			FirstName: order.Customer.FirstName,
			LastName:  order.Customer.LastName,
			Address:   order.Customer.Address,
		},
		Products: make([]types.OrderInfoProductResponse, 0, len(order.Products)),
	}

	for _, p := range order.Products {
		response.Products = append(response.Products, types.OrderInfoProductResponse{
			ID:          p.ID,
			Name:        p.Name,
			Description: p.Description,
			Price:       p.Price,
			Count:       p.Count,
		})
	}

	return ctx.JSON(response)
}

// OrderCreate создаение заказа
//
//	@Summary	создаение заказа
//	@Tags		orders
//	@Accept		json
//	@Produce	json
//	@Param		X-Request-ID	header		string	true	"ID запроса"
//	@Success	200				{object}	types.APIResponse
//	@Failure	400				{object}	types.APIResponse
//	@Failure	500				{object}	types.APIResponse
//	@Router		/api/v1/order [post]
func (e *Env) OrderCreate(ctx *fiber.Ctx) error {
	return e.OK(ctx)
}

// OrderUpdate обновление инфорамации заказа
//
//	@Summary
//	@Tags		orders
//	@Accept		json
//	@Produce	json
//	@Param		X-Request-ID	header		string	true	"ID запроса"
//	@Success	200				{object}	types.APIResponse
//	@Failure	400				{object}	types.APIResponse
//	@Failure	500				{object}	types.APIResponse
//	@Router		/api/v1/order/update [post]
func (e *Env) OrderUpdate(ctx *fiber.Ctx) error {
	return e.OK(ctx)
}
