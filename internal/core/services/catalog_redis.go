package services

import (
	"context"
	"encoding/json"
	"fmt"
	"goredis/internal/core/domains"
	"goredis/internal/core/ports"
	"time"

	"github.com/go-redis/redis/v9"
)

type catalogServiceRedis struct {
	productsRepo ports.ProductRepository
	redisClient  *redis.Client
}

func NewCatalogServiceRedis(productRepo ports.ProductRepository, redisClient *redis.Client) ports.CatalogService {
	return catalogServiceRedis{
		productsRepo: productRepo,
		redisClient:  redisClient,
	}
}

func (r catalogServiceRedis) GetProducts() (products []domains.ProductSrv, err error) {
	key := "service::GetProducts"
	//Redis Get
	productsJson, err := r.redisClient.Get(context.Background(), key).Result()
	if err == nil {
		err = json.Unmarshal([]byte(productsJson), &products)
		if err == nil {
			fmt.Println("Redis")
			return products, nil
		}
	}

	//Repository

	productsDB, err := r.productsRepo.GetProduct()
	if err != nil {
		return nil, err
	}
	for _, p := range productsDB {
		products = append(products, domains.ProductSrv{
			ID:       p.ID,
			Name:     p.Name,
			Quantity: p.Quantity,
		})
	}

	//Redis Set
	dataByte, err := json.Marshal(products)
	if err == nil {
		dataStr := string(dataByte)
		r.redisClient.Set(context.Background(), key, dataStr, time.Second*10).Err()
	}
	fmt.Println("database")

	return products, nil
}
