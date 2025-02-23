package controller

import (
	"ThaiLy/graph/model"
	"ThaiLy/graph/service"
)

type CategoryController struct {
	categoryService service.CategoryService
}

func (c *CategoryController) GetCategoryController(ID *string) ([]*model.Category, error) {
	if ID == nil {
		result, err := c.categoryService.FindAllCategory()
		if err != nil {
			return nil, err
		}
		categoryDBs := result.([]model.CategoryDB)
		var categories []*model.Category
		for _, categoryDB := range categoryDBs {
			category, err := categoryDB.ToCategory()
			if err != nil {
				return nil, err
			}
			categories = append(categories, category)
		}
		return categories, nil
	}
	result, err := c.categoryService.FindCategoryByID(*ID)
	if err != nil {
		return nil, err
	}
	categoryDB := result.(model.CategoryDB)
	var categories []*model.Category
	category, err := categoryDB.ToCategory()
	if err != nil {
		return nil, err
	}
	categories = append(categories, category)
	return categories, nil
}
