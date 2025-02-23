package service

import (
	"ThaiLy/graph/model"
	"fmt"
)

type CategoryService struct{}

func (a *CategoryService) FindAllCategory() (interface{}, error) {
	var categoriesDB []model.CategoryDB
	result := GetDB().Find(&categoriesDB)
	if result.Error != nil {
		return nil, fmt.Errorf("lỗi thi lấy ra category")
	}
	return categoriesDB, nil
}

func (a *CategoryService) FindCategoryByID(ID string) (interface{}, error) {
	var category model.CategoryDB
	result := GetDB().Where("id = ?", ID).First(&category)
	if result.Error != nil {
		return nil, fmt.Errorf("lỗi khi lấy chi tiết category")
	}
	return category, nil
}
