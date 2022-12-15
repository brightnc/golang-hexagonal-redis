package database

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDatabase() *gorm.DB {
	dsn := fmt.Sprintf("root:%v@tcp(localhost:%v)/productsdb", os.Getenv("PASSWORDDB"), os.Getenv("PORT"))
	dial := mysql.Open(dsn)
	db, err := gorm.Open(dial, &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}
