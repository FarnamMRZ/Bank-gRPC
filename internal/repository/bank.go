package repository

import (
	"context"
	"github.com/FarnamMRZ/Bank-gRPC/internal/datastruct"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

type BankQuery interface {
	CreateBank(bankName string) error
	BankExist(bankName string) (bool, error)
}

type bankQuery struct {
	db *mongo.Client
}

func (bq *bankQuery) CreateBank(bankName string) error {
	bank := datastruct.Bank{Name: bankName}
	_, err := bq.db.Database("bank").Collection("banks").InsertOne(context.TODO(), bank)
	return err
}

func (bq *bankQuery) BankExist(bankName string) (bool, error) {
	ss := bq.db.Database("bank").Collection("banks").FindOne(context.TODO(), bson.M{"name": bankName})
	if ss.Err() != nil {
		if ss.Err() == mongo.ErrNoDocuments {
			return false, nil
		}
		return false, ss.Err()
	}
	return true, nil
}
