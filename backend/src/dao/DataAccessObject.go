package dao

import (
	"gopkg.in/mgo.v2"
	"log"
)

// BankMockupDAO struct defines mongoDB storage location
type BankMockupDAO struct {
	Server   string
	Database string
}

var db *mgo.Database

// Collections
const (
	// AccountsCollection cointains the name of accounts collection name in mongoDB
	AccountsCollection = "accounts"
	// TransactionsCollection cointains the name of transaction collection name in mongoDB
	TransactionsCollection = "transactions"
	// AuthorizationsCollection cointains authorized keys
	AuthorizationsCollection = "authorizations"
)

// ConnectToDatabase function establishes a connection to MongoDB databse
func (dao *BankMockupDAO) ConnectToDatabase() {
	session, err := mgo.Dial(dao.Server)
	if err != nil {
		log.Fatal(err)
	}

	db = session.DB(dao.Database)
}
