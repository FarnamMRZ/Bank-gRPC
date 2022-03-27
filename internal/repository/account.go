package repository

import (
	"bank/internal/datastruct"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type AccountQuery interface {
	CreateAccount(customerUserName, bankName, number string, amount int64) error
	UpdateAccount(customerUserName, accountNumber string, amount int64) error
	GetAccountAmount(customerUserName, accountNumber string) (*int64, error)
	AccountExist(customerUserName, accountNumber string) (bool, error)
}

type accountQuery struct {
	db *mgo.Session
}

func (ac *accountQuery) CreateAccount(customerUserName, bankName, number string, amount int64) error {
	account := datastruct.Account{
		Bank:     bankName,
		Customer: customerUserName,
		Number:   number,
		Amount:   amount,
	}

	return ac.db.DB("bank").C("accounts").Insert(account)
}

func (ac *accountQuery) UpdateAccount(customerUserName, accountNumber string, amount int64) error {
	return ac.db.DB("bank").C("accounts").Update(bson.M{"customer": customerUserName, "number": accountNumber}, bson.M{"$set": bson.M{"amount": amount}})
}

func (ac *accountQuery) GetAccountAmount(customerUserName, accountNumber string) (*int64, error) {
	var account datastruct.Account
	err := ac.db.DB("bank").C("accounts").Find(bson.M{"customer": customerUserName, "number": accountNumber}).One(&account)
	if err != nil {
		return nil, err
	}
	return &account.Amount, nil
}

func (ac *accountQuery) AccountExist(customerUserName, accountNumber string) (bool, error) {
	count, err := ac.db.DB("bank").C("accounts").Find(bson.M{"customer": customerUserName, "number": accountNumber}).Count()
	if err != nil {
		return false, err
	}
	if count < 1 {
		return false, nil
	}
	return true, err
}
