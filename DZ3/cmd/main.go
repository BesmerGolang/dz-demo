package main

import (
	"dz3/configs"
	"dz3/internal/verify"
	"net/http"
)

func main() {
	cfg := configs.LoadConfig()
	h := &verify.Handler{
		Config: cfg,
	}

	router := http.NewServeMux()
	server := http.Server{
		Addr:    ":8081",
		Handler: router,
	}
	router.HandleFunc("POST /send", h.SendHandler)
	router.HandleFunc("GET /verify/{hash}", h.VerifyHandler)
	server.ListenAndServe()

}
