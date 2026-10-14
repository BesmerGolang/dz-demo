package main

import (
	"fmt"
	"math/rand/v2"
	"net/http"
	"strconv"
)

func returnNum(w http.ResponseWriter, r *http.Request) {
	resp := rand.IntN(7)
	w.Write([]byte(strconv.Itoa(resp)))
}
func main() {
	router := http.NewServeMux()
	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	router.HandleFunc("/rand", returnNum)
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("Ошибка запуска свервера:", err)
	}
}

//Второе дз готово !!!
