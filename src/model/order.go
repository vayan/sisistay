package model

import (
	"errors"

	"github.com/jinzhu/gorm"
)

type OrderStatus string

const (
	OrderUnassigned OrderStatus = "UNASSIGNED"
	OrderTaken      OrderStatus = "TAKEN"
)

type Order struct {
	ID             uint        `gorm:"primary_key" json:"id"`
	DistanceMeters int         `json:"distance"`
	Status         OrderStatus `json:"status"`
}

type OrderStorage interface {
	Migrate()
	Create(order *Order)
	Take(orderID uint) error
}

type OrderDatabase struct {
	Database *gorm.DB
}

func (o OrderDatabase) Migrate() {
	o.Database.AutoMigrate(Order{})
}

func (o OrderDatabase) Create(order *Order) {
	o.Database.Create(order)
}

func (o OrderDatabase) Take(orderID uint) error {
	// atomic operation because:
	// UPDATE "orders" SET "status" = 'TAKEN' WHERE "orders"."id" = xx AND ((status = 'UNASSIGNED'))

	result := o.Database.Model(Order{ID: orderID}).
		Where("status = ?", OrderUnassigned).
		Update("status", OrderTaken)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("nothing taken")
	}

	return nil
}
