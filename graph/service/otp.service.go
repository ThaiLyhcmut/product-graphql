package service

import (
	"ThaiLy/configs"
	"ThaiLy/graph/model"
	"fmt"
)

type OTPService struct{}

func (o *OTPService) DeletedOTP(otp, email string) error {
	// var otpModel model.Otp
	result := configs.GetDB().Where("code = ? AND email = ?", otp, email).Delete(&model.OtpDB{})
	if result.Error != nil {
		return fmt.Errorf("mã OTP không hợp lệ")
	}

	// Trả về nil nếu không có lỗi
	return nil
}
