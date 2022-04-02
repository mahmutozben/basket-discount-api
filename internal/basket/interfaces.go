package basket

import "github.com/google/uuid"

type IBasketRepository interface {
	Create(customer CustomerBasket) error
	Delete(basketId uuid.UUID) error
	Update(entity CustomerBasket) error
	Get(basketId uuid.UUID) (CustomerBasket, error)
}

type Service interface {
}
