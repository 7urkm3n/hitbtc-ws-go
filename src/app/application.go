package app

import (
	"fmt"
	serve "net/http"
	"orange-currency/src/domain/currency"
	"orange-currency/src/http"
	"orange-currency/src/repository/cache"

	"github.com/go-chi/chi/v5"
)

var (
	router = chi.NewRouter()
)

func StartApplication() {
	// using repo as cache!
	cache := cache.NewRepository()
	atHandler := http.NewHandler(currency.NewService(cache))

	router.Get("/currency/{symbol}", atHandler.GetBySymbol)
	router.Get("/currency/all", atHandler.All)

	fmt.Println("Serve port 3000")
	serve.ListenAndServe(":3000", router)
}
