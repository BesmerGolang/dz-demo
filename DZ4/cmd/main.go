package main

import (
	"dz4/configs"
	"dz4/internal/product"
	"dz4/pkg/db"
	"net/http"
)

func main() {
	config := configs.LoadConfig()
	db := db.NewDb(config)
	router := http.NewServeMux()
	prodRepo := product.NewProductrRepo(db)
	product.NewProdHandler(router, product.ProductHandlerDeps{
		ProdRepo: prodRepo,
	})
	server := http.Server{Addr: ":8081", Handler: router}
	server.ListenAndServe()
}
