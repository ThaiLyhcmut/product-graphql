package model

import (
	"time"
)

type ProductDB struct {
	ID              int       `gorm:"primaryKey"`                   // Khóa chính
	CategoryID      int       `gorm:"column:categoryID"`            // Cột categoryID
	Title           string    `gorm:"column:title"`                 // Cột title
	Description     string    `gorm:"column:description"`           // Cột description
	Price           float64   `gorm:"column:price"`                 // Cột price
	DiscountPercent int       `gorm:"column:discountPercent"`       // Cột discountPercent
	Stock           int       `gorm:"column:stock"`                 // Cột stock
	Thumbnail       string    `gorm:"column:thumbnail"`             // Cột thumbnail
	Status          string    `gorm:"column:status"`                // Cột status
	Deleted         bool      `gorm:"column:deleted;default:false"` // Cột deleted
	Position        int       `gorm:"column:position"`              // Cột postion
	Slug            string    `gorm:"column:slug"`                  // Cột slug
	Featured        string    `gorm:"column:featured"`              // Cột featured
	CreatedAt       time.Time `gorm:"column:createdAt"`             // Cột createdAt
	UpdatedAt       time.Time `gorm:"column:updatedAt"`             // Cột updatedAt
	CreatedBy       int       `gorm:"column:createdBy"`             // Cột createdBy
	UpdatedBy       int       `gorm:"column:updatedBy"`             // Cột updatedBy
}

func (ProductDB) TableName() string {
	return "products"
}

func (productDB *ProductDB) ToProduct() *Product {

	return &Product{
		ID:              &(productDB.ID),
		Title:           &productDB.Title,
		Description:     &productDB.Description,
		Thumbnail:       &productDB.Thumbnail,
		Price:           &productDB.Price,
		DiscountPercent: &productDB.DiscountPercent,
		Stock:           &productDB.Stock,
		Status:          &productDB.Status,
		Position:        &productDB.Position,
		Slug:            &productDB.Slug,
		Featured:        &productDB.Featured,
	}
}
