package db

import (
	"database/sql"
	_"modernc.org/sqlite"
)

var DB *sql.DB
var err error

func InitDB() {
	DB, err = sql.Open("sqlite", "api.db")
	if err != nil {
		panic(err)
	}
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	CreateTable()
}

func CreateTable() {
	createTableQuery := `CREATE TABLE IF NOT EXISTS events  (
     id         INTEGER PRIMARY KEY AUTOINCREMENT,
    name        TEXT NOT NULL ,
    description TEXT NOT NULL ,
    location    TEXT NOT NULL ,
    dateTime    DATETIME NOT NULL ,
    userId INTEGER NOT NULL 
    
     )`

	 _, err:= DB.Exec(createTableQuery) 

	 if err != nil{
		panic(err)
	 } 
	
	createUserTableQuery:= `CREATE TABLE IF NOT EXISTS users (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	email TEXT NOT NULL,
	password TEXT NOT NULL
	)`

	_, err = DB.Exec(createUserTableQuery)

	if err != nil {
		panic(err)
	}


}

