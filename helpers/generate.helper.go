package helper

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	ID                   int // Email được lưu trong token
	jwt.RegisteredClaims     // Thêm các trường chuẩn như exp, iat
}

func CreateJWT(id int) string {
	claims := Claims{
		ID: id, // Sử dụng email ở đây
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)), // Hết hạn sau 1 ngày
		}, // Thời gian hết hạn của token
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwtSecret := []byte(os.Getenv("JWT_SECRET"))
	tokenString, err := token.SignedString(jwtSecret) // Tạo token
	if err != nil {
		fmt.Println("Error creating token:", err)
		return ""
	}
	return tokenString
}

func ParseJWT(tokenString string) (*Claims, error) {
	jwtSecret := os.Getenv("JWT_SECRET") // Lấy secret từ .env

	// Parse token với claims
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		// Kiểm tra phương thức ký
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("phương thức ký không hợp lệ")
		}
		return []byte(jwtSecret), nil
	})

	if err != nil {
		return nil, err
	}
	// Kiểm tra nếu token hợp lệ và ép kiểu claims
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, errors.New("token không hợp lệ")
	}
}

func RandomNumber(length int) string {
	const characters = "0123456789"
	result := make([]byte, length)
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		result[i] = characters[random.Intn(len(characters))]
	}
	return string(result)
}
