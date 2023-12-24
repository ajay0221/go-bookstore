package config

import (
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func Connect() {
	cfg := &mysql.Config{
		User:   "root",
		Passwd: "pass", // type password here
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "testing",
		Params: map[string]string{
			"charset":   "utf8",
			"parseTime": "True",
			"loc":       "Local",
		},
	}
	d, err := gorm.Open("mysql", cfg.FormatDSN())
	if err != nil {
		fmt.Printf("Error connecting to database: %v\n", err)
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
