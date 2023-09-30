package services

type OrderService interface {
	CreateOrder(itemsCount int) (packs map[int]int, err error)
}
