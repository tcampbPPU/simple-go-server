package models

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name      string
	Code      string
	Price     uint
	Foo       string `gorm:"-"` // Ignored
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

func (p *Product) TableName() string {
	return "products"
}

func (p *Product) AfterFind(tx *gorm.DB) (err error) {

	p.Foo = "bar"

	return
}
