package model

import (
	helper "ThaiLy/helpers"
	"time"
)

type Sex string

const (
	Nam  Sex = "Nam"
	Nữ   Sex = "Nữ"
	Khác Sex = "Khác"
)

// User đại diện cho bảng users trong cơ sở dữ liệu
type AccountDB struct {
	ID             int       `gorm:"primaryKey"`                                                             // Khóa chính
	FullName       string    `gorm:"column:fullName"`                                                        // Cột fullName
	Email          string    `gorm:"unique;column:email"`                                                    // Cột email và đảm bảo tính duy nhất
	Password       string    `gorm:"column:password"`                                                        // Cột password
	Status         string    `gorm:"column:status;size:16"`                                                  // Cột status
	Address        string    `gorm:"column:address"`                                                         // Cột adress
	Phone          string    `gorm:"column:phone"`                                                           // Cột phone
	Avatar         string    `gorm:"column:avatar"`                                                          // Cột avatar
	Sex            string    `gorm:"column:sex;default:'Khác"`                                               // Cột giới tính
	Birthday       time.Time `gorm:"column:birthday;type:date;default:CURRENT_TIMESTAMP"`                    // Cột birthday
	Deleted        bool      `gorm:"column:deleted;default:false"`                                           // Cột deleted
	CreatedAt      time.Time `gorm:"column:createdAt;default:CURRENT_TIMESTAMP"`                             // Cột createdAt
	UpdatedAt      time.Time `gorm:"column:updatedAt;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"` // Cột updatedAt
	RequestFriends []byte    `gorm:"column:requestFriends"`                                                  // Cột requestFriends (dạng JSON)
	AcceptFriends  []byte    `gorm:"column:acceptFriends"`                                                   // Cột acceptFriends (dạng JSON)
	StatusOnline   string    `gorm:"column:statusOnline"`                                                    // Cột statusOnline
	FriendList     []byte    `gorm:"column:friendList"`                                                      // Cột friendList (dạng JSON)
}

func (AccountDB) TableName() string {
	return "accounts"
}

func stringPointer(s string) *string {
	return &s
}

func (accountDB *AccountDB) ToAccount() (*Account, error) {
	// Nếu không hợp lệ, trả về chuỗi rỗng
	// Chuyển đổi từ AccountDB sang Account
	account := &Account{
		ID:       &accountDB.ID,
		FullName: &accountDB.FullName,
		Email:    &accountDB.Email,
		Address:  &accountDB.Address,
		Phone:    &accountDB.Phone,
		Avatar:   &accountDB.Avatar,
		Sex:      &accountDB.Sex,
		Birthday: &accountDB.Birthday,
	}

	token := helper.CreateJWT(accountDB.ID)
	account.Token = &token
	account.Code = stringPointer("success")
	account.Msg = stringPointer("lấy tài khoản thành công")

	return account, nil
}
