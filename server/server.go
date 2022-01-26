package server

import (
	"blockchain/config"
	"blockchain/store"
	"blockchain/store/mongodb"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Handler struct {
	store store.Store
}

// NewHandler returns *Handler
func NewHandler(store store.Store) (*Handler, error) {
	return &Handler{
		store: store,
	}, nil
}

// Run runs the server.
func Run() error {
	cfg, err := config.LoadCfg()
	if err != nil {
		return err
	}

	mongo, err := mongodb.NewConnection(context.Background(), &mongodb.Config{
		URL:      cfg.MongoURI,
		DbName:   cfg.MongoDB,
		UserName: cfg.MongoUser,
		Password: cfg.MongoPassword,
	})
	if err != nil {
		return err
	}

	h, err := NewHandler(mongo)
	if err != nil {
		return err
	}

	http.HandleFunc("/transactions/get", h.GetTrx)

	log.Println("REST server is running")
	return http.ListenAndServe(fmt.Sprintf("%s:%d", cfg.ServiceHost, cfg.ServicePort), nil)
}

func response(w http.ResponseWriter, statusCode int, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(v)
	if err != nil {
		return
	}
}
