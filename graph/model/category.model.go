package model

import (
	"fmt"
	"time"
)

type CategoryDB struct {
	ID          string    `gorm:"primaryKey"`                   // Khóa chính
	Title       string    `gorm:"column:title"`                 // Cột title
	Description string    `gorm:"column:description"`           // Cột description
	Thumbnail   string    `gorm:"column:thumbnail"`             // Cột thumbnail
	Status      string    `gorm:"column:status"`                // Cột status
	Postion     string    `gorm:"column:postion"`               // Cột postion
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
		ID:          stringPointer(categoryDB.ID),
		Title:       stringPointer(categoryDB.Title),
		Description: stringPointer(categoryDB.Description),
		Thumbnail:   stringPointer(categoryDB.Thumbnail),
		Status:      stringPointer(categoryDB.Status),
		Position:    stringPointer(categoryDB.Postion),
		Deleted:     boolPointer(categoryDB.Deleted),
		Slug:        stringPointer(categoryDB.Slug),
		// Product sẽ được gán giá trị nil, thêm logic xử lý nếu cần thiết
		Product: nil,
	}

	return category, nil
}
