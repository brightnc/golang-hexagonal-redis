package protocol

import "goredis/repository"

var app *application

type application struct {
	svr *service.CustomerService
}

func init() {
	custRepository := repository.NewProductRepositoryDB()
	app = &application{
		svr: &custService,
	}
}
