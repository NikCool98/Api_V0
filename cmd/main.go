package main

import (
	"fmt"
	"github.com/NickCool98/Api_V0/internal/config"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func main() {
	cfg := config.MustLoad()
	fmt.Println(cfg) // не забыть удалить

	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hi, server is working!"))
	})

	if err := http.ListenAndServe(":8000", r); err != nil {
		fmt.Printf("Start server error: %s", err.Error())
		return
	}
}

// Здесь можно выполнять запросы к базе данных
