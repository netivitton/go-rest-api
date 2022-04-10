package internal

import (
	"time"

	"gorm.io/gorm"
)

// Defining a struct type
type Account struct {
	gorm.Model
	Name      string
	Last_name string
	Email     string
	AddressID int
	Address   Address `gorm:"foreignkey:AddressID"`
}

type Address struct {
	gorm.Model
	Home_no string
}

type Company struct {
	gorm.Model
	Company_Name string
	AddressID    int
	Address      Address `gorm:"foreignkey:AddressID"`
}

type Oauth2_client struct {
	ID     string `db:"id"`
	Secret string `db:"secret"`
	Domain string `db:"domain"`
	Data   []byte `db:"data"`
}

type Oauth2_token struct {
	ID        int64     `db:"id"`
	CreatedAt time.Time `db:"created_at"`
	ExpiresAt time.Time `db:"expires_at"`
	Code      string    `db:"code"`
	Access    string    `db:"access"`
	Refresh   string    `db:"refresh"`
	Data      []byte    `db:"data"`
}
