package types

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type (
	APIResponse struct {
		Error   string `json:"error,omitempty"`
		Message string `json:"message,omitempty"`
	}

	ValidateUserCredentialsResponse struct {
		ID        string `json:"id"`
		Email     string `json:"email"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
	}

	OrderListResponse []OrderInfoResponse
	OrderInfoResponse struct {
		ID        int                        `json:"id"`
		Status    string                     `json:"status"`
		Price     int                        `json:"price"`
		CreatedAt string                     `json:"created_at"`
		Customer  OrderInfoCustomerResponse  `json:"customer"`
		Products  []OrderInfoProductResponse `json:"products"`
	}
	OrderCreateResponse struct {
		ID int `json:"id"`
	}
	OrderInfoCustomerResponse struct {
		ID        string `json:"id"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Address   string `json:"address"`
	}
	OrderInfoProductResponse struct {
		ID          string `json:"id"`
		Name        string `json:"name"`
		Description string `json:"description"`
		Price       int    `json:"price"`
		Count       int    `json:"count"`
	}

	CustomerListResponse []CustomerInfoResponse
	CustomerInfoResponse struct {
		ID        string                      `json:"id"`
		FirstName string                      `json:"first_name"`
		LastName  string                      `json:"last_name"`
		Address   string                      `json:"address"`
		Orders    []CustomerInfoOrderResponse `json:"orders"`
	}
	CustomerInfoOrderResponse struct {
		ID        int    `json:"id"`
		Status    string `json:"status"`
		Price     int    `json:"price"`
		CreatedAt string `json:"created_at"`
	}

	ProductListResponse []ProductInfoResponse
	ProductInfoResponse struct {
		ID          string `json:"id"`
		Name        string `json:"name"`
		Description string `json:"description"`
		Price       int    `json:"price"`
	}
)

type ProductCreateBody struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
}

func (b ProductCreateBody) Validate() error {
	if b.Name == "" {
		return fiber.ErrBadRequest
	}

	if len(b.Name) >= 20 {
		return fiber.ErrBadRequest
	}

	if b.Description != "" && len(b.Description) > 100 {
		return fiber.ErrBadRequest
	}

	if b.Price <= 0 {
		return fiber.ErrBadRequest
	}

	return nil
}

type ProductUpdateBody struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
}

func (b ProductUpdateBody) Validate() error {
	if b.ID == "" {
		return fiber.ErrBadRequest
	}

	if _, err := uuid.Parse(b.ID); err != nil {
		return fiber.ErrBadRequest
	}

	if b.Name != "" && len(b.Name) >= 20 {
		return fiber.ErrBadRequest
	}

	if b.Description != "" && len(b.Description) > 100 {
		return fiber.ErrBadRequest
	}

	if b.Price < 0 {
		return fiber.ErrBadGateway
	}

	return nil
}

type CustomerCreateBody struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Address   string `json:"address"`
}

func (b CustomerCreateBody) Validate() error {
	if b.FirstName == "" || b.LastName == "" || b.Address == "" {
		return fiber.ErrBadRequest
	}

	if len(b.FirstName) > 16 {
		return fiber.ErrBadRequest
	}

	if len(b.LastName) > 16 {
		return fiber.ErrBadRequest
	}

	if len(b.Address) > 100 {
		return fiber.ErrBadRequest
	}

	return nil
}

type CustomerUpdateBody struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Address   string `json:"address"`
}

func (b CustomerUpdateBody) Validate() error {
	if b.ID == "" {
		return fiber.ErrBadRequest
	}

	if _, err := uuid.Parse(b.ID); err != nil {
		return fiber.ErrBadRequest
	}

	if b.FirstName != "" && len(b.FirstName) > 16 {
		return fiber.ErrBadRequest
	}

	if b.LastName != "" && len(b.LastName) > 16 {
		return fiber.ErrBadRequest
	}

	if b.Address != "" && len(b.Address) > 100 {
		return fiber.ErrBadRequest
	}

	return nil
}

type CreateOrderBody struct {
	CustomerID string         `json:"customer_id"`
	Products   map[string]int `json:"products"`
}

func (b CreateOrderBody) Validate() error {
	if b.CustomerID == "" {
		return fiber.ErrBadRequest
	}

	if _, err := uuid.Parse(b.CustomerID); err != nil {
		return fiber.ErrBadRequest
	}

	for id, count := range b.Products {
		if _, err := uuid.Parse(id); err != nil {
			delete(b.Products, id)
			continue
		}

		if count <= 0 {
			delete(b.Products, id)
		}
	}

	if len(b.Products) == 0 {
		return fiber.ErrBadRequest
	}

	return nil
}
