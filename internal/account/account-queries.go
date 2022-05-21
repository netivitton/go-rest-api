package internal

import (
	"encoding/json"
	"fmt"
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
	tx.Create(address)
	tx.Create(&Account{UserName: "BBB", Password: "CCC", AddressID: address.ID})
	tx.Commit()
	utils.CloseDB(gormdb)
	return "ta da! \n"
}

func LoginAccount() string {
	gormdb, err := utils.Connection()
	if err != nil {
		log.Fatal("cannot connect DB", err)
	}
	var result Account
	gormdb.Raw("SELECT id, created_at, updated_at, deleted_at, user_name, password, name, last_name, email, fa_code, address_id FROM accounts where user_name = ? AND password = ?;", "AAA", "BBB").Scan(&result)
	marshal_struct, _ := json.Marshal(result)
	fmt.Println("Marshaled string: ", string(marshal_struct))
	utils.CloseDB(gormdb)
	return "ta da! \n"
}
