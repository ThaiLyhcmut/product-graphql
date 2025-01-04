package service

import (
	"ThaiLy/graph/model"
	"fmt"
)

type AccountService struct {
	database Database
}

func (a *AccountService) Created(FullName, Email, Password string) (interface{}, error) {
	account := model.AccountDB{
		FullName: FullName,
		Email:    Email,
		Password: Password,
		Status:   "active",
	}
	result := GetDB().Create(&account)
	if result.Error != nil {
		return nil, fmt.Errorf("xóa mã OTP không thành công")
	}
	return account, nil
}

func (a *AccountService) FindOne(Email, Password string) (interface{}, error) {
	var account model.AccountDB
	result := GetDB().Where(
		"email = ? AND password = ?", Email, Password,
	).First(&account)
	if result.Error != nil {
		return nil, fmt.Errorf("lấy tài khoảng không thành công")
	}
	return account, nil
}
