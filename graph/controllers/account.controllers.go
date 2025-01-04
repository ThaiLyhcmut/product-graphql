package controller

import (
	"ThaiLy/graph/model"
	"ThaiLy/graph/service"
	"context"
	"fmt"
)

type AccountController struct {
	otpService     service.OTPService
	accountService service.AccountService
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
	return account, nil
}

func (a *AccountController) RegisterAccountController(accountInput model.RegisterAccountInput) (*model.Account, error) {
	if err := a.otpService.Deleted(accountInput.Otp, accountInput.Password); err != nil {
		return nil, err
	}
	result, err := a.accountService.Created(accountInput.FullName, accountInput.Email, accountInput.Password)
	if err != nil {
		return nil, err
	}
	account := result.(model.Account)
	// token := helper.CreateJWT(account.ID)
	// account.Token := token
	return &account, nil
}

func (a *AccountController) LoginAccountController(accountInput model.LoginAccountInput) (*model.Account, error) {
	result, err := a.accountService.FindOne(accountInput.Email, accountInput.Password)
	if err != nil {
		return nil, err
	}
	accountDB := result.(model.AccountDB)
	account, err := accountDB.ToAccount()
	if err != nil {
		return nil, err
	}
	return account, nil
}
