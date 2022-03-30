package repository

import (
	"context"
	"github.com/FarnamMRZ/Bank-gRPC/internal/datastruct"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type AccountQuery interface {
	CreateAccount(customerUserName, bankName, number string, amount int64) error
	UpdateAccount(customerUserName, accountNumber string, amount int64) error
	GetAccountAmount(customerUserName, accountNumber string) (*int64, error)
	AccountExist(customerUserName, accountNumber string) (bool, error)
	Withdraw(userName, customerUserName, accountNumber string, amount int64) error
}

type accountQuery struct {
	db *mongo.Client
}

func (ac *accountQuery) CreateAccount(customerUserName, bankName, number string, amount int64) error {
	account := datastruct.Account{
		Bank:     bankName,
		Customer: customerUserName,
		Number:   number,
		Amount:   amount,
	}

	_, err := ac.db.Database("bank").Collection("accounts").InsertOne(context.TODO(), account)
	return err
}

func (ac *accountQuery) UpdateAccount(customerUserName, accountNumber string, amount int64) error {
	_, err := ac.db.Database("bank").Collection("accounts").UpdateOne(context.TODO(), bson.M{"customer": customerUserName, "number": accountNumber}, bson.M{"$set": bson.M{"amount": amount}})
	return err
}

func (ac *accountQuery) GetAccountAmount(customerUserName, accountNumber string) (*int64, error) {
	var account datastruct.Account
	err := ac.db.Database("bank").Collection("accounts").FindOne(context.TODO(), bson.M{"customer": customerUserName, "number": accountNumber}).Decode(&account)
	if err != nil {
		return nil, err
	}
	return &account.Amount, nil
}

func (ac *accountQuery) AccountExist(customerUserName, accountNumber string) (bool, error) {
	ss := ac.db.Database("bank").Collection("accounts").FindOne(context.TODO(), bson.M{"customer": customerUserName, "number": accountNumber})
	if ss.Err() != nil {
		if ss.Err() == mongo.ErrNoDocuments {
			return false, nil
		}
		return false, ss.Err()
	}
	return true, nil
}

func (ac *accountQuery) Withdraw(userName, customerUserName, accountNumber string, amount int64) error {
	session, err := ac.db.StartSession()
	if err != nil {
		return err
	}
	defer session.EndSession(context.Background())

	callback := func(sc mongo.SessionContext) (interface{}, error) {
		result, err := ac.db.Database("bank").Collection("accounts").UpdateOne(
			context.TODO(),
			bson.M{"customer": customerUserName, "number": accountNumber},
			bson.M{"$inc": bson.M{"amount": -amount}},
		)
		if err != nil {
			return nil, err
		}

		result, err = ac.db.Database("bank").Collection("customers").UpdateOne(
			context.TODO(),
			bson.M{"name": userName},
			bson.M{"$inc": bson.M{"safe": amount}},
		)
		if err != nil {
			return nil, err
		}

		return result, nil
	}

	_, err = session.WithTransaction(context.Background(), callback, txnOpts)
	if err != nil {
		return err
	}
	return nil
}
