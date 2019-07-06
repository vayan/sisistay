package model

import "github.com/jinzhu/gorm"

type OrderStatus string

const (
	OrderUnassigned OrderStatus = "UNASSIGNED"
	OrderAssigned   OrderStatus = "ASSIGNED"
)

type Order struct {
	ID             uint        `gorm:"primary_key" json:"id"`
	DistanceMeters int         `json:"distance"`
	Status         OrderStatus `json:"status"`
}

type OrderStorage interface {
	Migrate()
	Create(order *Order)
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
