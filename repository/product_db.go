package repository

import (
	"fmt"
	"goredis/internal/core/domains"
	"goredis/internal/core/ports"
	"math/rand"
	"time"

	"gorm.io/gorm"
)

type productRepositoryDB struct {
	db *gorm.DB
}

func NewProductRepositoryDB(db *gorm.DB) ports.ProductRepository {
	db.AutoMigrate(&domains.Product{})
	mockData(db)
	return productRepositoryDB{db}
}

func mockData(db *gorm.DB) error {
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

func (r productRepositoryDB) GetProduct() (products []domains.Product, err error) {
	err = r.db.Order("quantity desc").Limit(30).Find(&products).Error
	return products, err
}
