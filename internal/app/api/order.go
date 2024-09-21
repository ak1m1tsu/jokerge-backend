package api

import (
	"time"

	"github.com/ak1m1tsu/jokerge/internal/pkg/types"
	"github.com/gofiber/fiber/v2"
)

// OrderList возвращает список заказов
//
//	@Summary	список заказов
//	@Tags		orders
//	@Produce	json
//	@Success	200	{object}	OrderList
//	@Router		/order/list [get]
func (e *Env) OrderList(ctx *fiber.Ctx) error {
	orders := []types.Order{
		{
			ID:         "87aa13fc-81f8-4712-bb38-4415eb2e294d",
			CustomerID: "c22946d7-991e-44a1-b0dc-6b775a34664c",
			Status:     types.OrderStatusActive,
			CreatedAt:  time.Now(),
		},
		{
			ID:         "f9593790-4948-4cb4-8ca0-1510bb3ef116",
			CustomerID: "c22946d7-991e-44a1-b0dc-6b775a34664c",
			Status:     types.OrderStatusCompleted,
			CreatedAt:  time.Now(),
		},
		{
			ID:         "bd512907-11ec-402d-a05a-7c9d147e23ef",
			CustomerID: "7dfa4d19-d4c8-4596-bcb7-249e0de951bc",
			Status:     types.OrderStatusActive,
			CreatedAt:  time.Now(),
		},
		{
			ID:         "b35248c2-559c-4bb1-8d43-9b1aaa0df8f4",
			CustomerID: "4c8a6931-04cb-4c93-9a90-e7a29342aba2",
			Status:     types.OrderStatusActive,
			CreatedAt:  time.Now(),
		},
		{
			ID:         "2bb7d14a-d98a-41ce-a0bc-764798672776",
			CustomerID: "7dfa4d19-d4c8-4596-bcb7-249e0de951bc",
			Status:     types.OrderStatusCanceled,
			CreatedAt:  time.Now(),
		},
		{
			ID:         "cfd373d6-4708-461d-892a-327b6dae27d7",
			CustomerID: "4c8a6931-04cb-4c93-9a90-e7a29342aba2",
			Status:     types.OrderStatusActive,
			CreatedAt:  time.Now(),
		},
	}
	customers := []types.Customer{
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

	return ctx.JSON(FillOrderList(orders, customers))
}

// OrderGet информация о заказе
//
//	@Summary	информация о заказе
//	@Tags		orders
//	@Produce	json
//	@Success	200	{object}	OrderListItem
//	@Router		/order/{order_id} [get]
func (e *Env) OrderGet(ctx *fiber.Ctx) error {
	return ctx.JSON(FillOrderListItem(
		types.Order{
			ID:         "f9593790-4948-4cb4-8ca0-1510bb3ef116",
			CustomerID: "c22946d7-991e-44a1-b0dc-6b775a34664c",
			Status:     types.OrderStatusCompleted,
			CreatedAt:  time.Now(),
		},
		types.Customer{
			ID:        "c22946d7-991e-44a1-b0dc-6b775a34664c",
			FirstName: "Модест",
			LastName:  "Пономарёв",
			Address:   "121248, г. Приозерное, ул. Пяловская, дом 40, квартира 30",
		},
	))
}

// OrderCreate создаение заказа
//
//	@Summary	создаение заказа
//	@Tags		orders
//	@Accept		json
//	@Produce	json
//	@Success	200	{object}	map[string]string
//	@Router		/order/ [post]
func (e *Env) OrderCreate(ctx *fiber.Ctx) error {
	return ctx.JSON(map[string]string{
		"status": "OK",
	})
}

// OrderUpdate обновление инфорамации заказа
//
//	@Summary
//	@Tags		orders
//	@Accept		json
//	@Produce	json
//	@Success	200	{object}	map[string]string
//	@Router		/order/update [post]
func (e *Env) OrderUpdate(ctx *fiber.Ctx) error {
	return ctx.JSON(map[string]string{
		"status": "OK",
	})
}
