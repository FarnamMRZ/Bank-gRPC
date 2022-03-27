package repository

import (
	"bank/internal/datastruct"
	"gopkg.in/mgo.v2"
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
	db *mgo.Session
}

func (cq *customerQuery) CreateCustomer(userName string, initMoney int64) error {
	customer := datastruct.Customer{Name: userName, Safe: initMoney}
	return cq.db.DB("bank").C("customers").Insert(customer)
}

func (cq *customerQuery) CustomerExist(userName string) (bool, error) {
	count, err := cq.db.DB("bank").C("customers").Find(bson.M{"name": userName}).Count()
	if err != nil {
		return false, err
	}
	if count < 1 {
		return false, nil
	}
	return true, nil
}

func (cq *customerQuery) GetSafeAmount(userName string) (*int64, error) {
	var customer datastruct.Customer
	err := cq.db.DB("bank").C("customers").Find(bson.M{"name": userName}).One(&customer)
	if err != nil {
		return nil, err
	}
	return &customer.Safe, nil
}

func (cq *customerQuery) UpdateSafeAmount(userName string, newAmount int64) error {
	return cq.db.DB("bank").C("customers").Update(bson.M{"name": userName}, bson.M{"$set": bson.M{"safe": newAmount}})
}

func (cq *customerQuery) GetNumberOfAccounts(userName string) (*int64, error) {
	var customer datastruct.Customer
	err := cq.db.DB("bank").C("customers").Find(bson.M{"name": userName}).One(&customer)
	if err != nil {
		return nil, err
	}
	return &customer.Accounts, nil
}

func (cq *customerQuery) AddAccount(userName string) error {
	return cq.db.DB("bank").C("customers").Update(bson.M{"name": userName}, bson.M{"$inc": bson.M{"accounts": 1}})
}
