package internal

import (
	"log"

	"github.com/netivitton/go-rest-api/utils"
)

// Defining a struct type
func InitAccountModel() string {
	gormdb, err := utils.Connection()
	if err != nil {
		log.Fatal("cannot connect DB", err)
	}
	gormdb.AutoMigrate(&Address{})
	gormdb.AutoMigrate(&Account{})
	gormdb.AutoMigrate(&Company{})
	gormdb.AutoMigrate(&Oauth2_client{})
	gormdb.AutoMigrate(&Oauth2_token{})
	utils.CloseDB(gormdb)
	return "ta da! \n"
}

func InsertAccount() string {
	gormdb, err := utils.Connection()
	if err != nil {
		log.Fatal("cannot connect DB", err)
	}
	tx := gormdb.Begin()
	address := &Address{Home_no: "123"}
	tx.Create(&address)
	tx.Create(&Account{UserName: "AAA", Password: "BBB", AddressID: address.ID})
	tx.Commit()
	utils.CloseDB(gormdb)
	return "ta da! \n"
}
