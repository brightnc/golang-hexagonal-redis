package mocks

import (
	"fmt"
	"goredis/internal/core/domains"
	"math/rand"
	"time"

	"gorm.io/gorm"
)

func MockData(db *gorm.DB) error {
	var count int64
	db.Model(&domains.Product{}).Count(&count)
	if count > 0 {
		return nil
	}
	seed := rand.NewSource(time.Now().UnixNano())
	random := rand.New(seed)
	products := []domains.Product{}
	for i := 0; i < 5000; i++ {
		products = append(products, domains.Product{
			Name:     fmt.Sprintf("Product %v", i+1),
			Quantity: random.Intn(100),
		})
	}
	return db.Create(&products).Error

}
