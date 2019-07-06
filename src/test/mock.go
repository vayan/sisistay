package test

import "github.com/vayan/sisistay/src/model"

type OrderMockDB struct {
	FakeID uint
}

func (o OrderMockDB) Migrate() {

}

func (o OrderMockDB) Create(order *model.Order) {
	order.ID = o.FakeID
	o.FakeID++
}
