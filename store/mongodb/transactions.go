package mongodb

import (
	"blockchain/models"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (d *db) GetTransactions(ctx context.Context, skip, limit int64) ([]models.Transaction, error) {
	filter := bson.M{"_id": 1} // todo

	findOpt := options.Find()

	findOpt.SetSkip(skip)
	findOpt.SetLimit(limit)
	findOpt.SetSort(bson.D{})
	cursor, err := d.conn.Collection(collectionTransactions).Find(ctx, filter, findOpt)
	if err != nil {
		return nil, err
	}

	transactions := make([]models.Transaction, limit)
	if err = cursor.All(ctx, &transactions); err != nil {
		return nil, err
	}

	return transactions, nil
}

func (d *db) InsertTransactions(ctx context.Context, trx []models.Transaction) error {
	transactions := make([]mongo.WriteModel, len(trx))

	for i, tx := range trx {
		transactions[i] = mongo.NewInsertOneModel().SetDocument(tx)
	}

	_, err := d.conn.Collection(collectionTransactions).BulkWrite(ctx, transactions)
	if err != nil {
		return err
	}

	return nil
}
