package repository

import (
	"goredis/internal/core/domains"
	"goredis/internal/core/ports"
	"goredis/mocks"

	"gorm.io/gorm"
)

type productRepositoryDB struct {
	db *gorm.DB
}

func NewProductRepositoryDB(db *gorm.DB) ports.ProductRepository {
	db.AutoMigrate(&domains.Product{})
	err := mocks.MockData(db)
	if err != nil {
		panic(err)
	}
	return productRepositoryDB{db}
}

func (r productRepositoryDB) GetProduct() (products []domains.Product, err error) {
	err = r.db.Order("quantity desc").Limit(30).Find(&products).Error
	return products, err
}
