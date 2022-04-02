package basket

import "github.com/google/uuid"

type (
	CustomerBasket struct {
		Id         uuid.UUID
		CustomerId uuid.UUID
		Items      *[]BasketItem
	}

	BasketItem struct {
		ProductCode string
		ProductName string
		MerchantId  uuid.UUID
		UnitPrice   float64
		Quantity    int
	}
)
