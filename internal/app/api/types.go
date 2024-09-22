package api

type (
	OrderList    []OrderListItem
	CustomerList []CustomerListItem

	OrderItem struct {
		ID     string `json:"id"`
		Status string `json:"status"`
	}
	CustomerItem struct {
		ID        string `json:"id"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Address   string `json:"address"`
	}

	OrderListItem struct {
		OrderItem
		Customer CustomerItem `json:"customer"`
	}
	CustomerListItem struct {
		CustomerItem
		Orders []OrderItem `json:"orders"`
	}

	CustomerOrdersItem struct {
		Customer CustomerItem `json:"customer"`
		Orders   OrderList    `json:"orders"`
	}

	UserInfoItem struct {
		ID        string `json:"id"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Email     string `json:"email"`
	}
)
