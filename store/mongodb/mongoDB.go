package mongodb

import (
	"context"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"net/http"
)

const collectionTransactions = "transactions"

type db struct {
	conn *mongo.Database
}

type Config struct {
	URL      string
	DbName   string
	UserName string
	Password string
}

func NewConnection(ctx context.Context, cfg *Config) (*db, error) {
	clientOptions := options.Client().ApplyURI(cfg.URL)
	clientOptions.SetAuth(options.Credential{
		AuthSource: cfg.DbName,
		Username:   cfg.UserName,
		Password:   cfg.Password,
	})

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, fmt.Errorf("connection failed %v", err)
	}

	if err = client.Ping(ctx, readpref.Primary()); err != nil {
		return nil, fmt.Errorf("failed to ping %v", err)
	}

	return &db{conn: client.Database(cfg.DbName)}, nil
}

// ErrorResponse : This is error model.
type ErrorResponse struct {
	StatusCode   int    `json:"status"`
	ErrorMessage string `json:"message"`
}

// GetError : This is helper function to prepare error model.
// If you want to export your function. You must to start upper case function name. Otherwise you won't see your function when you import that on other class.
func GetError(err error, w http.ResponseWriter) {

	var response = ErrorResponse{
		ErrorMessage: err.Error(),
		StatusCode:   http.StatusInternalServerError,
	}

	message, _ := json.Marshal(response)

	w.WriteHeader(response.StatusCode)
	w.Write(message)
	log.Fatal(err.Error())
}
