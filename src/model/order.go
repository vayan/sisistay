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
