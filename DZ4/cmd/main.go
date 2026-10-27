package main

import (
	"dz4/configs"
	"dz4/pkg/db"
	"net/http"
)

func main() {
	config := configs.LoadConfig()
	_ = db.NewDb(config)
	router := http.NewServeMux()
	server := http.Server{Addr: ":8081", Handler: router}
	server.ListenAndServe()
}
