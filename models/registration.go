package models

import (
	"RESTApi/db"
	"errors"
)

type Registration struct {
	ID      int `json:"id"`
	UserID  int `json:"userid" binding: "required"`
	EventID int `json:"eventid" binding:"required"`
}

func RegisterForEvent(userID int64, eventID int64) error {
	insertQuery := `INSERT INTO registration (userId,eventId) VALUES (?,?)`
	_, err := db.DB.Exec(insertQuery, userID, eventID)
	if err != nil {
		return errors.New("unable to create table")
	}

	return nil

}