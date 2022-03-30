package repository

import (
	"context"
	"github.com/FarnamMRZ/Bank-gRPC/internal/datastruct"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

type CustomerQuery interface {
	CreateCustomer(userName string, initMoney int64) error
	CustomerExist(userName string) (bool, error)
	GetSafeAmount(userName string) (*int64, error)
	UpdateSafeAmount(userName string, newAmount int64) error
	GetNumberOfAccounts(userName string) (*int64, error)
	AddAccount(userName string) error
}

type customerQuery struct {
	db *mongo.Client
}

func (cq *customerQuery) CreateCustomer(userName string, initMoney int64) error {
	customer := datastruct.Customer{Name: userName, Safe: initMoney}
	_, err := cq.db.Database("bank").Collection("customers").InsertOne(context.TODO(), customer)
	return err
}

func (cq *customerQuery) CustomerExist(userName string) (bool, error) {
	ss := cq.db.Database("bank").Collection("customers").FindOne(context.TODO(), bson.M{"name": userName})
	if ss.Err() != nil {
		if ss.Err() == mongo.ErrNoDocuments {
			return false, nil
		}
		return false, ss.Err()
	}
	return true, nil
}

func (cq *customerQuery) GetSafeAmount(userName string) (*int64, error) {
	var customer datastruct.Customer
	err := cq.db.Database("bank").Collection("customers").FindOne(context.TODO(), bson.M{"name": userName}).Decode(&customer)
	if err != nil {
		return nil, err
	}
	return &customer.Safe, nil
}

func (cq *customerQuery) UpdateSafeAmount(userName string, newAmount int64) error {
	_, err := cq.db.Database("bank").Collection("customers").UpdateOne(context.TODO(), bson.M{"name": userName}, bson.M{"$set": bson.M{"safe": newAmount}})
	return err
}

func (cq *customerQuery) GetNumberOfAccounts(userName string) (*int64, error) {
	var customer datastruct.Customer
	err := cq.db.Database("bank").Collection("customers").FindOne(context.TODO(), bson.M{"name": userName}).Decode(&customer)
	if err != nil {
		return nil, err
	}
	return &customer.Accounts, nil
}

func (cq *customerQuery) AddAccount(userName string) error {
	_, err := cq.db.Database("bank").Collection("customers").UpdateOne(context.TODO(), bson.M{"name": userName}, bson.M{"$inc": bson.M{"accounts": 1}})
	return err
}
