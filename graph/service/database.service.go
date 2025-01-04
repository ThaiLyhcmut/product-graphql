package service

import (
	"fmt"
	"log"
	"os"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// type Service interface {
// 	Created() (interface{}, error)
// 	Findone() (interface{}, error)
// 	Deleted() (interface{}, error)
// }

// Database struct để quản lý kết nối
type Database struct{}

// Singleton đảm bảo DB chỉ được khởi tạo một lần
var (
	once sync.Once
	db   *gorm.DB
)

// Khởi tạo Database
func (d *Database) InitDB() {
	once.Do(func() {
		// Lấy thông tin từ biến môi trường cho MySQL
		dbName := os.Getenv("MYSQL_NAME")
		dbUser := os.Getenv("MYSQL_USERNAME")
		dbPassword := os.Getenv("MYSQL_PASSWORD")
		dbHost := os.Getenv("MYSQL_HOST")

		// Chuỗi kết nối MySQL
		dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			dbUser, dbPassword, dbHost, dbName)

		// Kết nối tới MySQL
		var err error
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		})

		if err != nil {
			log.Fatalf("Kết nối database thất bại: %v", err)
		}
		fmt.Println("Kết nối database MySQL thành công!")
	})
}

// GetDB sẽ trả về kết nối cơ sở dữ liệu toàn cục
func GetDB() *gorm.DB {
	if db == nil {
		log.Fatal("Cơ sở dữ liệu chưa được khởi tạo")
	}
	return db
}
