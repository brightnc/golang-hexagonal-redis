package protocol

import (
	"goredis/database"
	"goredis/internal/core/ports"
	"goredis/internal/core/services"
	"goredis/internal/repository"
)

var app *application

type application struct {
	svr ports.CatalogService
}

func init() {
	db := database.InitDatabase()
	redis := database.InitRedis()
	_ = redis
	productRepo := repository.NewProductRepositoryRedis(db, redis)
	productSrv := services.NewCatalogService(productRepo)
	app = &application{
		svr: productSrv,
	}
}
