package models

import (
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Transaction struct {
	ID               primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	TransactionID    int64              `json:"transactionID" bson:"transactionID"`
	SenderAddress    string             `json:"senderAddress" bson:"senderAddress"`
	RecipientAddress string             `json:"recipientAddress" bson:"recipientAddress"`
	BlockNum         int64              `json:"blockNum" bson:"blockNum"`
	ConfirmNum       int64              `json:"confirmNum" bson:"confirmNum"`
	TrxDate          string             `json:"trxDate" bson:"trxDate"`
	Amount           float64            `json:"amount" bson:"amount"`
	Commission       float64            `json:"commission" bson:"commission"`
}

func (t *Transaction) Validation() error {
	if t.TransactionID == 0 {
		return errors.New("transactionID cannot be zero")
	}
	if t.SenderAddress == "" {
		return errors.New("senderAddress cannot be empty")
	}
	if t.RecipientAddress == "" {
		return errors.New("recipientAddress cannot be empty")
	}
	if t.BlockNum == 0 {
		return errors.New("blockNum cannot be zero")
	}
	if t.ConfirmNum == 0 {
		return errors.New("confirmNum cannot be zero")
	}
	if t.TrxDate == "" {
		return errors.New("trxDate cannot be empty")
	}
	if t.Amount == 0 {
		return errors.New("amount cannot be zero")
	}
	if t.Commission == 0 {
		return errors.New("commission cannot be zero")
	}
	return nil
}
