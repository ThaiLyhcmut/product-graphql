package model

import (
	"fmt"
	"time"
)

type ProductDB struct {
	ID              string    `gorm:"primaryKey"`                   // Khóa chính
	CategoryID      string    `gorm:"column:categoryID"`            // Cột categoryID
	Title           string    `gorm:"column:title"`                 // Cột title
	Description     string    `gorm:"column:description"`           // Cột description
	Price           float32   `gorm:"column:price"`                 // Cột price
	DiscountPercent int       `gorm:"column:discountPercent"`       // Cột discountPercent
	Stock           int       `gorm:"column:stock"`                 // Cột stock
	Thumbnail       string    `gorm:"column:thumbnail"`             // Cột thumbnail
	Status          string    `gorm:"column:status"`                // Cột status
	Deleted         bool      `gorm:"column:deleted;default:false"` // Cột deleted
	Postion         string    `gorm:"column:postion"`               // Cột postion
	Slug            string    `gorm:"column:slug"`                  // Cột slug
	Featured        bool      `gorm:"column:featured"`              // Cột featured
	CreatedAt       time.Time `gorm:"column:createdAt"`             // Cột createdAt
	UpdatedAt       time.Time `gorm:"column:updatedAt"`             // Cột updatedAt
	CreatedBy       int       `gorm:"column:createdBy"`             // Cột createdBy
	UpdatedBy       int       `gorm:"column:updatedBy"`             // Cột updatedBy
}

func (ProductDB) TableName() string {
	return "products"
}

func (productDB *ProductDB) ToProduct() *Product {
	// Chuyển đổi các giá trị phù hợp từ ProductDB sang Product
	price := float64(productDB.Price)
	discountPercent := fmt.Sprintf("%d", productDB.DiscountPercent)
	stock := fmt.Sprintf("%d", productDB.Stock)
	featured := fmt.Sprintf("%t", productDB.Featured)

	return &Product{
		ID:              &productDB.ID,
		Title:           &productDB.Title,
		Description:     &productDB.Description,
		Thumbnail:       &productDB.Thumbnail,
		Price:           &price,
		DiscountPercent: &discountPercent,
		Stock:           &stock,
		Status:          &productDB.Status,
		Position:        &productDB.Postion,
		Slug:            &productDB.Slug,
		Featured:        &featured,
	}
}
