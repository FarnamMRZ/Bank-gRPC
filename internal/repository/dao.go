package repository

import (
	"gopkg.in/mgo.v2"
)

func NewDB() (*mgo.Session, error) {
	return mgo.Dial("localhost")
}

type DAO interface {
	NewAccountQuery() AccountQuery
	NewBankQuery() BankQuery
	NewCustomerQuery() CustomerQuery
}

type dao struct {
	db *mgo.Session
}

func NewDAO(s *mgo.Session) DAO {
	return &dao{s}
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
