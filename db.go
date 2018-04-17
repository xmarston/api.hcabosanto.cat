package main

import (
	"os"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"errors"
	"github.com/jmoiron/sqlx"
)

type Database struct {
	Host       string
	Name       string
	User       string
	Password   string
	Connection *sqlx.DB
}

func (db *Database) Init() {
	db.Host = os.Getenv("DB_HOST")
	db.Name = os.Getenv("DB_NAME")
	db.User = os.Getenv("DB_USER")
	db.Password = os.Getenv("DB_PASSWORD")
}

func (db *Database) StartConnection() error {
	var err error
	if db.Name == "" {
		err = errors.New("database struct no initialized")
	} else {
		db.Connection, err = sqlx.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s", db.User, db.Password, db.Host, db.Name))
	}

	return err
}

func GetDb() (*sqlx.DB, error) {
	db := Database{}
	db.Init()
	err := db.StartConnection()

	return db.Connection, err
}
