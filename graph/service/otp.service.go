package service

import (
	"ThaiLy/graph/model"
	"fmt"
)

type OTPService struct{}

func (o *OTPService) DeletedOtp(otp, email string) error {
	// var otpModel model.Otp
	fmt.Println(otp, email)
	result := GetDB().Where("code = ? AND email = ?", otp, email).Delete(&model.OtpDB{})
	if result.Error != nil {
		return fmt.Errorf("mã OTP không hợp lệ")
	}
	// Kiểm tra nếu không có bản ghi nào bị xóa
	if result.RowsAffected == 0 {
		return fmt.Errorf("không tìm thấy mã OTP hoặc email để xóa")
	}

	// Trả về nil nếu không có lỗi
	return nil
}
