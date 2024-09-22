package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
)

// CustomerList возвращает список клиентов
//
//	@Summary	список клиентов
//	@Tags		customers
//	@Produce	json
//	@Success	200	{object}	[]CustomerItem
//	@Failure	500 {object}	Response
//	@Router		/api/v1/customer/list [get]
func (e *Env) CustomerList(ctx *fiber.Ctx) error {
	customers, err := e.Service().GetCustomers(ctx.Context())
	if err != nil {
		zerolog.Ctx(ctx.Context()).Error().Err(err).Msg("failed to get customer list")
		return err
	}

	orders, err := e.Service().GetOrders((ctx.Context()))
	if err != nil {
		zerolog.Ctx(ctx.Context()).Error().Err(err).Msg("failed to get order list")
		return err
	}

	result := FillCustomerList(customers, orders)

	return ctx.JSON(result)
}

// CustomerGet возвращает информацию о клиенте
//
//	@Summary	информация о клиенте
//	@Tags		customers
//	@Produce	json
//	@Accept		json
//	@Success	200	{object}	CustomerItem
//	@Failure	404 {object}	Response
//	@Failure	500 {object}	Response
//	@Router		/api/v1/customer/{customer_id} [get]
func (e *Env) CustomerGet(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	customer, err := e.Service().GetCustomerInfo(ctx.Context(), id)
	if err != nil {
		zerolog.Ctx(ctx.Context()).Error().Err(err).Msgf("failed to get customer by id %s", id)
		return err
	}

	if customer == nil {
		return e.NotFound(ctx)
	}

	orders, err := e.Service().GetCustomerOrders(ctx.Context(), id)
	if err != nil {
		zerolog.Ctx(ctx.Context()).Error().Err(err).Msgf("failed to get customer %s orders", id)
		return err
	}

	result := FillCustomerListItem(customer, orders)

	return ctx.JSON(result)
}

// CustomerCreate создает клиента
//
//	@Summary	создание нового клиента
//	@Tags		customers
//	@Accept		json
//	@Produce	json
//	@Success	200	{object}	Response
//	@Failure	400 {object}	Response
//	@Failure	500 {object}	Response
//	@Router		/api/v1/customer [post]
func (e *Env) CustomerCreate(ctx *fiber.Ctx) error {
	return e.OK(ctx)
}
