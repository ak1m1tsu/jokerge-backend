package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
)

// OrderList возвращает список заказов
//
//	@Summary	список заказов
//	@Tags		orders
//	@Produce	json
//	@Param		X-Request-ID	header		string	true	"ID запроса"
//	@Success	200				{object}	OrderList
//	@Failure	500				{object}	Response
//	@Router		/api/v1/order/list [get]
func (e *Env) OrderList(ctx *fiber.Ctx) error {
	orders, err := e.Service().GetOrders(ctx.Context())
	if err != nil {
		zerolog.Ctx(ctx.UserContext()).Error().Err(err).Msg("failed to get orders")
		return err
	}

	customers, err := e.Service().GetCustomers(ctx.Context())
	if err != nil {
		zerolog.Ctx(ctx.UserContext()).Error().Err(err).Msg("failed to get customers")
		return err
	}

	result := FillOrderList(orders, customers)

	return ctx.JSON(result)
}

// OrderGet информация о заказе
//
//	@Summary	информация о заказе
//	@Tags		orders
//	@Produce	json
//	@Param		order_id		path		string	true	"ID заказа"
//	@Param		X-Request-ID	header		string	true	"ID запроса"
//	@Success	200				{object}	OrderListItem
//	@Failure	404				{object}	Response
//	@Failure	500				{object}	Response
//	@Router		/api/v1/order/{order_id} [get]
func (e *Env) OrderGet(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	order, err := e.Service().GetOrderInfo(ctx.Context(), id)
	if err != nil {
		zerolog.Ctx(ctx.UserContext()).Error().Err(err).Msgf("failed to get order by id %s", id)
		return err
	}

	customer, err := e.Service().GetCustomerInfo(ctx.Context(), order.CustomerID)
	if err != nil {
		zerolog.Ctx(ctx.UserContext()).Error().Err(err).Msgf("failed to get customer by id %s", order.CustomerID)
		return err
	}

	result := FillOrderListItem(order, customer)

	return ctx.JSON(result)
}

// OrderCreate создаение заказа
//
//	@Summary	создаение заказа
//	@Tags		orders
//	@Accept		json
//	@Produce	json
//	@Param		X-Request-ID	header		string	true	"ID запроса"
//	@Success	200				{object}	Response
//	@Failure	400				{object}	Response
//	@Failure	500				{object}	Response
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
//	@Success	200				{object}	Response
//	@Failure	400				{object}	Response
//	@Failure	500				{object}	Response
//	@Router		/api/v1/order/update [post]
func (e *Env) OrderUpdate(ctx *fiber.Ctx) error {
	return e.OK(ctx)
}
