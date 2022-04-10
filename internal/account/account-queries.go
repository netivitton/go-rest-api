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
	return "ta da! \n"
}
