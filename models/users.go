package models

import (
	"RESTApi/db"
	"RESTApi/utils"
	"errors"
)

type Users struct {
	ID       int    `json:"id"`
	Email    string `json:"email" binding: "required"`
	Password string `json:"password" binding: "required"`
}

func (u *Users) Save() error {
	insertQuery:= `INSERT INTO users (email, password) VALUES (?,?)`
	HashedPassword, err:= utils.HashPassword(u.Password)
	if err != nil{
		return err
	}

	_, err = db.DB.Exec(insertQuery, u.Email, HashedPassword)
	if err != nil {
		return errors.New("unable to create table")
	}

	return nil
}

func (u *Users) Validate() error {
	getUserQuery:= `select * from users where email=?`
	row, err := db.DB.Query(getUserQuery,u.Email)
	if err != nil {
		return err
	}

	var userDetails Users
	for row.Next() {
		err = row.Scan(&userDetails.ID, &userDetails.Email, &userDetails.Password)
	}
	if err != nil {
		return err
	}

	u.ID = userDetails.ID
	err= utils.ValidatePassword([]byte(userDetails.Password), []byte(u.Password))
	if err != nil{
		return err
	}
	
	return nil
}


