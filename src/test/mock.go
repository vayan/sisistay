package test

import "github.com/vayan/sisistay/src/model"

type OrderMockDB struct {
	FakeID           uint
	TakeError        error
	ListOrderResults []model.Order
}

func (o OrderMockDB) Migrate() {
}

func (o OrderMockDB) Create(order *model.Order) {
	order.ID = o.FakeID
	o.FakeID++
}

func (o OrderMockDB) Take(orderID uint) error {
	return o.TakeError
}

func (o OrderMockDB) List(page int, limit int) []model.Order {
	return o.ListOrderResults
}

type MockRouteFetcher struct {
	Distance int
	Error    error
}

func (rf MockRouteFetcher) GetDistance(coordinates model.Coordinates, to model.Coordinates) (int, error) {
	return rf.Distance, rf.Error
}
