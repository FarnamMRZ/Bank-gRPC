package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
)

func NewDB() (*mongo.Client, error) {
	return mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost"))
}

func CloseDB(client *mongo.Client) {
	err := client.Disconnect(context.TODO())
	if err != nil {
		log.Fatalf("Faild to disconnect from database: %v", err)
	}
}

type DAO interface {
	NewAccountQuery() AccountQuery
	NewBankQuery() BankQuery
	NewCustomerQuery() CustomerQuery
}

type dao struct {
	db *mongo.Client
}

var txnOpts = options.Transaction()

func NewDAO(c *mongo.Client) DAO {
	return &dao{c}
}

func (d *dao) NewBankQuery() BankQuery {
	return &bankQuery{d.db}
}

func (d *dao) NewCustomerQuery() CustomerQuery {
	return &customerQuery{d.db}
}

func (d *dao) NewAccountQuery() AccountQuery {
	return &accountQuery{d.db}
}
