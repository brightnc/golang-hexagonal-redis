package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"goredis/internal/core/domains"
	"goredis/internal/core/ports"
	"goredis/mocks"
	"time"

	"github.com/go-redis/redis/v9"
	"gorm.io/gorm"
)

type productRepositoryRedis struct {
	db          *gorm.DB
	redisClient *redis.Client
}

func NewProductRepositoryRedis(db *gorm.DB, redisClient *redis.Client) ports.ProductRepository {
	db.AutoMigrate(&domains.Product{})
	err := mocks.MockData(db)
	if err != nil {
		panic(err)
	}
	return productRepositoryRedis{db: db, redisClient: redisClient}
}

func (r productRepositoryRedis) GetProduct() (products []domains.Product, err error) {
	key := "repository::GetProducts"
	//Redis Get
	productsJson, err := r.redisClient.Get(context.Background(), key).Result()
	if err == nil {
		err = json.Unmarshal([]byte(productsJson), &products)
		if err == nil {
			fmt.Println("Redis")
			return products, err
		}
	}

	// Database
	err = r.db.Order("quantity desc").Limit(30).Find(&products).Error
	if err != nil {
		return nil, err
	}

	//Redis Set
	dataByte, err := json.Marshal(products)
	if err != nil {
		return nil, err
	}
	dataStr := string(dataByte)
	err = r.redisClient.Set(context.Background(), key, dataStr, time.Second*10).Err()
	if err != nil {
		return nil, err
	}
	fmt.Println("Database")
	return products, nil
}
