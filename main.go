package main

import (
	"blockchain/config"
	"blockchain/store/mongodb"
	"context"
	"fmt"
	"log"
)

func main() {
	cfg, err := config.LoadCfg()
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(cfg)

	db, err := mongodb.NewConnection(context.Background(), &mongodb.Config{
		URL:      cfg.MongoURI,
		DbName:   cfg.MongoDB,
		UserName: cfg.MongoUser,
		Password: cfg.MongoPassword,
	})
	if err != nil {
		panic(err)
	}

	//trx := []models.Transaction{
	//	{
	//		TransactionID:    228,
	//		SenderAddress:    "lesnaya",
	//		RecipientAddress: "adsmskam",
	//		BlockNum:         2,
	//		ConfirmNum:       3,
	//		TrxDate:          "546ryut,",
	//		Amount:           231,
	//		Commission:       99999,
	//	},
	//	{
	//		TransactionID:    23112,
	//		SenderAddress:    "ledasdasdasaya",
	//		RecipientAddress: "adsmska213123123123m",
	//		BlockNum:         4,
	//		ConfirmNum:       678,
	//		TrxDate:          "542312dsa6ryut,",
	//		Amount:           231231231231,
	//		Commission:       9999933213123,
	//	},
	//	{
	//		TransactionID:    2282312,
	//		SenderAddress:    "lesnaydsdada",
	//		RecipientAddress: "adsmskam",
	//		BlockNum:         2,
	//		ConfirmNum:       3,
	//		TrxDate:          "546ryut,",
	//		Amount:           231,
	//		Commission:       99999,
	//	},
	//}
	//
	//err = db.InsertTransactions(context.Background(), trx)
	//if err != nil {
	//	panic(err)
	//}
	trx, err := db.GetTransactions(context.Background(), 1, 2)
	if err != nil {
		panic(err)
	}
	fmt.Println(trx)
}
