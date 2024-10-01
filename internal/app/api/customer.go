package api

import (
	"github.com/ak1m1tsu/jokerge/internal/pkg/types"
	"github.com/gofiber/fiber/v2"
	zlog "github.com/rs/zerolog"
)

// CustomerList возвращает список клиентов
//
//	@Summary	список клиентов
//	@Tags		customers
//	@Security	BasicAuth
//	@Produce	json
//	@Param		X-Request-ID	header		string	true	"ID запроса"
//	@Success	200				{object}	types.CustomerListResponse
//	@Failure	500				{object}	types.APIResponse
//	@Router		/api/v1/customer/list [get]
func (e *Env) CustomerList(ctx *fiber.Ctx) error {
	customers, err := e.Service().GetCustomers(ctx.UserContext())
	if err != nil {
		zlog.Ctx(ctx.UserContext()).Error().Err(err).Msg("failed to get customer list")
		return err
	}

	response := make(types.CustomerListResponse, 0, len(customers))

	for _, c := range customers {
		converted := types.CustomerInfoResponse{
			ID:        c.ID,
			FirstName: c.FirstName,
			LastName:  c.LastName,
			Address:   c.Address,
			Orders:    make([]types.CustomerInfoOrderResponse, 0, len(c.Orders)),
		}

		for _, o := range c.Orders {
			converted.Orders = append(converted.Orders, types.CustomerInfoOrderResponse{
				ID:        o.ID,
				Status:    o.Status.String(),
				Price:     o.Price,
				CreatedAt: o.CreatedAt.String(),
			})
		}

		response = append(response, converted)

	}

	return ctx.JSON(response)
}

// CustomerGet возвращает информацию о клиенте
//
//	@Summary	информация о клиенте
//	@Tags		customers
//	@Security	BasicAuth
//	@Produce	json
//	@Accept		json
//	@Param		customer_id		path		string	true	"ID клиента"
//	@Param		X-Request-ID	header		string	true	"ID запроса"
//	@Success	200				{object}	types.CustomerInfoResponse
//	@Failure	404				{object}	types.APIResponse
//	@Failure	500				{object}	types.APIResponse
//	@Router		/api/v1/customer/{customer_id} [get]
func (e *Env) CustomerGet(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	customer, err := e.Service().GetCustomerInfo(ctx.UserContext(), id)
	if err != nil {
		zlog.Ctx(ctx.UserContext()).Error().Err(err).Msgf("failed to get customer by id %s", id)
		return err
	}

	if customer == nil {
		return e.NotFound(ctx)
	}

	response := types.CustomerInfoResponse{
		ID:        customer.ID,
		FirstName: customer.FirstName,
		LastName:  customer.LastName,
		Address:   customer.Address,
		Orders:    make([]types.CustomerInfoOrderResponse, 0),
	}

	for _, o := range customer.Orders {
		response.Orders = append(response.Orders, types.CustomerInfoOrderResponse{
			ID:        o.ID,
			Status:    o.Status.String(),
			Price:     o.Price,
			CreatedAt: o.CreatedAt.String(),
		})
	}

	return ctx.JSON(response)
}

// CustomerCreate создает клиента
//
//	@Summary	создание нового клиента
//	@Tags		customers
//	@Security	BasicAuth
//	@Accept		json
//	@Produce	json
//	@Param		body			body		types.CustomerCreateBody	true	"Тело запроса"
//	@Param		X-Request-ID	header		string						true	"ID запроса"
//	@Success	200				{object}	types.CustomerInfoResponse
//	@Failure	400				{object}	types.APIResponse
//	@Failure	500				{object}	types.APIResponse
//	@Router		/api/v1/customer [post]
func (e *Env) CustomerCreate(ctx *fiber.Ctx) error {
	var body types.CustomerCreateBody
	if err := ctx.BodyParser(&body); err != nil {
		zlog.Ctx(ctx.UserContext()).Error().Err(err).Msg("failed to parse request body")
		return err
	}

	if err := body.Validate(); err != nil {
		zlog.Ctx(ctx.UserContext()).Error().Err(err).Msg("invalid body request")
		return err
	}

	zlog.Ctx(ctx.UserContext()).Info().Any("body", body).Msg("create new customer")

	id, err := e.Service().CreateCustomer(ctx.UserContext(), body)
	if err != nil {
		zlog.Ctx(ctx.UserContext()).Error().Err(err).Msg("failed to create customer")
		return err
	}

	return ctx.JSON(types.CustomerInfoResponse{
		ID:        id,
		FirstName: body.FirstName,
		LastName:  body.LastName,
		Address:   body.Address,
		Orders:    make([]types.CustomerInfoOrderResponse, 0),
	})
}

// CustomerUpdate обработчик обновления информации о клиенте
//
//	@Summary	обновление информации о клиенте
//	@Tags		customers
//	@Security	BasicAuth
//	@Accept		json
//	@Produce	json
//	@Param		body			body		types.CustomerUpdateBody	true	"Тело запроса"
//	@Param		X-Request-ID	header		string						true	"ID запроса"
//	@Success	200				{object}	types.APIResponse
//	@Failure	400				{object}	types.APIResponse
//	@Failure	404				{object}	types.APIResponse
//	@Failure	500				{object}	types.APIResponse
//	@Router		/api/v1/customer/update [post]
func (e *Env) CustomerUpdate(ctx *fiber.Ctx) error {
	var body types.CustomerUpdateBody
	if err := ctx.BodyParser(&body); err != nil {
		zlog.Ctx(ctx.UserContext()).Error().Err(err).Msg("failed to parse request body")
		return err
	}

	if err := body.Validate(); err != nil {
		zlog.Ctx(ctx.UserContext()).Error().Any("body", body).Err(err).Msg("invalid body request")
		return err
	}

	if err := e.Service().UpdateCustomer(ctx.UserContext(), body); err != nil {
		zlog.Ctx(ctx.UserContext()).Error().Err(err).Msg("failed to update customer")
		return err
	}

	return e.OK(ctx)
}
