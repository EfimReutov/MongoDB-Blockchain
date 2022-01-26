package store

import (
	"blockchain/models"
	"context"
)

type Store interface {
	GetTransactions(ctx context.Context, skip, limit int64) ([]models.Transaction, error)
	InsertTransactions(ctx context.Context, trx []models.Transaction) error
}
