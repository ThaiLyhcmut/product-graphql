package service

import (
	"ThaiLy/configs"
	"ThaiLy/graph/model"
	"fmt"
)

type AccountService struct{}

func (a *AccountService) CreateAccount(FullName, Email, Password string) (interface{}, error) {
	account := model.AccountDB{
		FullName: FullName,
		Email:    Email,
		Password: Password,
		Status:   "active",
	}
	result := configs.GetDB().Create(&account)
	if result.Error != nil {
		return nil, fmt.Errorf("xóa mã OTP không thành công")
	}
	return account, nil
}
