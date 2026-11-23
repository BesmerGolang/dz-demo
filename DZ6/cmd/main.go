package main

import (
	"dz4/configs"
	"dz4/internal/product"
	"dz4/pkg/db"
	"dz4/pkg/middleware"
	"net/http"
	"os"

	"github.com/sirupsen/logrus"
)

func main() {
	log := logrus.New()
	log.SetOutput(os.Stdout)
	log.SetLevel(logrus.InfoLevel)
	log.SetFormatter(&logrus.JSONFormatter{})
	config := configs.LoadConfig()
	db := db.NewDb(config)
	router := http.NewServeMux()
	handler := middleware.Logging(log)(router)
	prodRepo := product.NewProductrRepo(db)
	product.NewProdHandler(router, product.ProductHandlerDeps{
		ProdRepo: prodRepo,
	})
	server := &http.Server{Addr: ":8081", Handler: handler}
	log.WithField("addr", server.Addr).Info("Server is listening")

	if err := server.ListenAndServe(); err != nil {
		log.WithError(err).Fatal("Server stopped with error")
	}
}
