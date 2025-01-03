package controller

import (
	"ThaiLy/graph/model"
	"ThaiLy/graph/service"
	helper "ThaiLy/helpers"
	"context"
	"fmt"
)

type AccountController struct {
	*service.OTPService
	*service.AccountService
}

func (a *AccountController) GetAccountController(ctx context.Context) (interface{}, error) {
	accountInterface := ctx.Value("account")
	if accountInterface == nil {
		return nil, fmt.Errorf("vui lòng đăng nhập")
	}
	account, ok := accountInterface.(model.Account)
	if !ok {
		return nil, fmt.Errorf("tạo tài khoản không thành công")
	}
	return &account, nil
}

func (a *AccountController) RegisterAccountController(accountInput model.RegisterAccountInput) (interface{}, error) {
	if err := a.OTPService.DeletedOTP(accountInput.Otp, accountInput.Password); err != nil {
		return nil, err
	}
	result, err := a.AccountService.CreateAccount(accountInput.FullName, accountInput.Email, accountInput.Password)
	if err != nil {
		return nil, err
	}
	account := result.(model.AccountDB)
	token := helper.CreateJWT(account.ID)

	return token, nil
}
