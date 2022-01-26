package server

import (
	"blockchain/models"
	"blockchain/store/mongodb"
	"context"
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"net/http"
)

const (
	userURL = "/transactions"
)

//func (h *Handler) Register(router *httprouter.Router) {
//	router.GET(userURL, h.GetTransactions)
//}

func (h *Handler) GetTrx(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		response(w, http.StatusBadRequest, "invalid method")
		return
	}

	collection := new(mongo.Collection)

	w.Header().Set("Content-Type", "application/json")

	var trxs []models.Transaction

	cur, err := collection.Find(context.TODO(), bson.M{})

	if err != nil {
		mongodb.GetError(err, w)
		return
	}

	defer cur.Close(context.TODO())

	for cur.Next(context.TODO()) {

		var trx models.Transaction
		err := cur.Decode(&trx)
		if err != nil {
			log.Fatal(err)
		}

		trxs = append(trxs, trx)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(trxs) // encode similar to serialize process.
}

//func (h *Handler) GetTransactions(w http.ResponseWriter, r *http.Request) {
//
//	if r.Method != http.MethodGet {
//		response(w, http.StatusBadRequest, "invalid method")
//		return
//	}
//	transaction := new(Transaction)
//	err := json.NewDecoder(r.Body).Decode(transaction)
//	if err != nil {
//		response(w, http.StatusBadRequest, err)
//		return
//	}
//
//	transaction, err = h.GetTransactions()
//	if err != nil {
//		response(w, http.StatusInternalServerError, err.Error())
//		return
//	}
//
//	response(w, http.StatusOK, transaction)
//}
