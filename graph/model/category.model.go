package model

import (
	"fmt"
	"time"
)

type CategoryDB struct {
	ID          int       `gorm:"primaryKey"`                   // Khóa chính
	Title       string    `gorm:"column:title"`                 // Cột title
	Description string    `gorm:"column:description"`           // Cột description
	Thumbnail   string    `gorm:"column:thumbnail"`             // Cột thumbnail
	Status      string    `gorm:"column:status"`                // Cột status
	Postion     int       `gorm:"column:postion"`               // Cột postion
	Deleted     bool      `gorm:"column:deleted;default:false"` // Cột deleted
	Slug        string    `gorm:"column:slug"`                  // Cột slug
	CreatedAt   time.Time `gorm:"column:createdAt"`             // Cột createdAt
	UpdatedAt   time.Time `gorm:"column:updatedAt"`             // Cột updatedAt
}

func (CategoryDB) TableName() string {
	return "categories"
}

// Hàm helper để tạo con trỏ từ giá trị bool
func boolPointer(value bool) *bool {
	return &value
}

func (categoryDB *CategoryDB) ToCategory() (*Category, error) {
	// Kiểm tra nếu categoryDB là nil
	if categoryDB == nil {
		return nil, fmt.Errorf("categoryDB is nil")
	}

	// Chuyển đổi từ CategoryDB sang Category
	category := &Category{
		ID:          &(categoryDB.ID),
		Title:       &(categoryDB.Title),
		Description: &(categoryDB.Description),
		Thumbnail:   &(categoryDB.Thumbnail),
		Status:      &(categoryDB.Status),
		Position:    &(categoryDB.Postion),
		Deleted:     &(categoryDB.Deleted),
		Slug:        &(categoryDB.Slug),
		// Product sẽ được gán giá trị nil, thêm logic xử lý nếu cần thiết
		Product: nil,
	}

	return category, nil
}
