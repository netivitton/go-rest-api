package utils

import (
	"os"
)

type Config struct {
	// The path to the config file
	DB_HOST_ALL string `json:"db_host_all"`
	DB_HOST     string `json:"db_host"`
	DB_USER     string `json:"db_user"`
	DB_PASS     string `json:"db_pass"`
	DB_NAME     string `json:"db_name"`
}

func LoadConfig() (config Config, err error) {
	config.DB_HOST_ALL = os.Getenv("DB_HOST_ALL")
	config.DB_HOST = os.Getenv("DB_HOST")
	config.DB_USER = os.Getenv("DB_USER")
	config.DB_PASS = os.Getenv("DB_PASS")
	config.DB_NAME = os.Getenv("DB_NAME")
	if err != nil {
		return config, err
	}
	return config, nil
}
