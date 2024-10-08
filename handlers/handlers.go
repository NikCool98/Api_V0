package handlers

import (
	"encoding/json"
	"fmt"
	config2 "github.com/NickCool98/Api_V0/internal/config"
	"github.com/NickCool98/Api_V0/internal/storage"
	"github.com/gorilla/mux"
	"net/http"
)

type Server struct {
	cfg   config2.ConfigHttp
	Cache *storage.OrderCache
}

func New(cfgPath string, cache *storage.OrderCache) *Server {
	cfg := config2.MustLoad(cfgPath)
	return &Server{
		cfg:   cfg.HTTPServer,
		Cache: cache,
	}
}

func (s *Server) Launch() error {
	r := mux.NewRouter()
	r.HandleFunc("/order/{order_uid}", s.GetOrderHandler)
	http.Handle("/", r)
	err := http.ListenAndServe(s.cfg.Port+":"+s.cfg.Port, nil)
	if err != nil {
		return fmt.Errorf("failed to launch server: %w", err)
	}
	return nil
}

func (s *Server) GetOrderHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orderUID := vars["order_uid"]
	if order, ok := s.Cache.GetOrd(orderUID); ok {
		orderJSON, _ := json.MarshalIndent(order, "", "    ")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(orderJSON)
	} else {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(fmt.Sprintf("\nOrderUID: <%s> not found!\n", orderUID)))
	}
}
