package service

import (
	"ThaiLy/graph/model"
	helper "ThaiLy/helpers"
	"fmt"
	"time"
)

type AccountService struct{}

func (a *AccountService) CreateAccount(FullName, Email, Password string) (interface{}, error) {
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

func (a *AccountService) FindAccountByEmailPassword(Email, Password string) (interface{}, error) {
	var account model.AccountDB
	result := GetDB().Where(
		"email = ? AND password = ?", Email, Password,
	).First(&account).Limit(1)
	if result.Error != nil {
		return nil, fmt.Errorf("lấy tài khoảng không thành công")
	}
	return account, nil
}

func (a *AccountService) UpdateAccountByID(accountDB model.AccountDB, accountInput model.UpdateAccountInput) (interface{}, error) {
	result := GetDB().Model(&accountDB).Updates(accountInput)
	if result.Error != nil {
		return nil, fmt.Errorf("cập nhật không thành công")
	}
	return nil, nil
}

func (a *AccountService) CreateOtpByEmail(Email string) error {
	otp := helper.RandomNumber(6)
	otpModel := model.OtpDB{
		Email:     Email,
		Code:      otp,
		ExpiredAt: time.Now().Add(time.Minute * 5),
	}
	result := GetDB().Create(&otpModel)
	if result.Error != nil {
		return fmt.Errorf("tao otp khong thanh cong")
	}
	err := helper.SendMail(Email, "Mã OTP", otp)
	if err != nil {
		return fmt.Errorf("tao mail khong thanh cong")
	}
	return nil
}
