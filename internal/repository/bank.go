package repository

import (
	"bank/internal/datastruct"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type BankQuery interface {
	CreateBank(bankName string) error
	BankExist(bankName string) (bool, error)
}

type bankQuery struct {
	db *mgo.Session
}

func (bq *bankQuery) CreateBank(bankName string) error {
	bank := datastruct.Bank{Name: bankName}
	err := bq.db.DB("bank").C("banks").Insert(bank)
	return err
}

func (bq *bankQuery) BankExist(bankName string) (bool, error) {
	count, err := bq.db.DB("bank").C("banks").Find(bson.M{"name": bankName}).Count()
	if err != nil {
		return false, err
	}
	if count < 1 {
		return false, nil
	}
	return true, nil
}
