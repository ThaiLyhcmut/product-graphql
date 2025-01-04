package service

import (
	"ThaiLy/graph/model"
	"fmt"
)

type OTPService struct {
	database Database
}

func (o *OTPService) Deleted(otp, email string) error {
	// var otpModel model.Otp
	result := GetDB().Where("code = ? AND email = ?", otp, email).Delete(&model.OtpDB{})
	if result.Error != nil {
		return fmt.Errorf("mã OTP không hợp lệ")
	}

	// Trả về nil nếu không có lỗi
	return nil
}
