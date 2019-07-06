package model

import "github.com/jinzhu/gorm"

type Status int

const (
	Unassigned Status = iota + 1
	Assigned
)

type Order struct {
	gorm.Model
	Distance int
	Status
}

type OrderStorage interface {
	Migrate()
}

type OrderDatabase struct {
	Database *gorm.DB
}

func (o OrderDatabase) Migrate() {
	o.Database.AutoMigrate(Order{})
}
