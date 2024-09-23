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
//	@Param		X-Request-ID	header		string	true	"ID запроса"
//	@Success	200				{object}	[]CustomerItem
//	@Failure	500				{object}	Response
//	@Router		/api/v1/customer/list [get]
func (e *Env) CustomerList(ctx *fiber.Ctx) error {
	customers, err := e.Service().GetCustomers(ctx.Context())
	if err != nil {
		zerolog.Ctx(ctx.UserContext()).Error().Err(err).Msg("failed to get customer list")
		return err
	}

	return ctx.JSON(customers)
}

// CustomerGet возвращает информацию о клиенте
//
//	@Summary	информация о клиенте
//	@Tags		customers
//	@Produce	json
//	@Accept		json
//	@Param		customer_id		path		string	true	"ID клиента"
//	@Param		X-Request-ID	header		string	true	"ID запроса"
//	@Success	200				{object}	CustomerItem
//	@Failure	404				{object}	Response
//	@Failure	500				{object}	Response
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

	return ctx.JSON(customer)
}

// CustomerCreate создает клиента
//
//	@Summary	создание нового клиента
//	@Tags		customers
//	@Accept		json
//	@Produce	json
//	@Param		X-Request-ID	header		string	true	"ID запроса"
//	@Success	200				{object}	Response
//	@Failure	400				{object}	Response
//	@Failure	500				{object}	Response
//	@Router		/api/v1/customer [post]
func (e *Env) CustomerCreate(ctx *fiber.Ctx) error {
	return e.OK(ctx)
}
