package controller

import (
	"ThaiLy/graph/model"
	"ThaiLy/graph/service"
	"ThaiLy/middlewares"
	"context"
	"fmt"
)

func stringPointer(s string) *string {
	return &s
}

type AccountController struct {
	otpService     service.OTPService
	accountService service.AccountService
}

func (a *AccountController) GetAccountController(ctx context.Context) (*model.Account, error) {
	result := ctx.Value(middlewares.AccountKey)
	if result == nil {
		return nil, fmt.Errorf("vui lòng đăng nhập")
	}
	accountDB, ok := result.(model.AccountDB)
	if !ok {
		return nil, fmt.Errorf("tạo tài khoản không thành công")
	}
	account, err := accountDB.ToAccount()
	if err != nil {
		return nil, fmt.Errorf("tạo tài khoản không thành công")
	}
	return account, nil
}

func (a *AccountController) RegisterAccountController(accountInput model.RegisterAccountInput) (*model.Account, error) {
	if err := a.otpService.DeletedOtp(accountInput.Otp, accountInput.Email); err != nil {
		return nil, err
	}
	result, err := a.accountService.CreateAccount(accountInput.FullName, accountInput.Email, accountInput.Password)
	if err != nil {
		return nil, err
	}
	accountDB, ok := result.(model.AccountDB)
	if !ok {
		return nil, fmt.Errorf("tạo tài khoảng không thành công")
	}
	account, err := accountDB.ToAccount()
	if err != nil {
		return nil, err
	}
	return account, nil
}

func (a *AccountController) LoginAccountController(accountInput model.LoginAccountInput) (*model.Account, error) {
	result, err := a.accountService.FindAccountByEmailPassword(accountInput.Email, accountInput.Password)
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

func (a *AccountController) UpdateAccountController(ctx context.Context, accountInput model.UpdateAccountInput) (*model.Account, error) {
	myFunc := "UpdateAccountController"
	result := ctx.Value(middlewares.AccountKey)
	if result == nil {
		return nil, fmt.Errorf("vui lòng đăng nhập %s", myFunc)
	}
	accountDB, ok := result.(model.AccountDB)
	if !ok {
		return nil, fmt.Errorf("lấy thông tin không thành công %s", myFunc)
	}
	_, err := a.accountService.UpdateAccountByID(accountDB, accountInput)
	if err != nil {
		return nil, err
	}
	result, err = a.accountService.FindAccountByEmailPassword(accountDB.Email, accountDB.Password)
	accountDB, ok = result.(model.AccountDB)
	if !ok {
		return nil, fmt.Errorf("lấy thông tin không thành công %s", myFunc)
	}
	account, err := accountDB.ToAccount()
	if err != nil {
		return nil, fmt.Errorf("lay thong tin khon thanh cong %s", myFunc)
	}
	return account, nil
}

func (a *AccountController) CreateOtpController(Email string) (*model.Otp, error) {
	err := a.accountService.CreateOtpByEmail(Email)
	if err != nil {
		return nil, err
	}
	otp := &model.Otp{
		Code: stringPointer("success"),
		Msg:  stringPointer("Tạo otp thành công"),
	}
	return otp, nil

}
