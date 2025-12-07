package models

import (
	"RESTApi/db"
	"errors"
	"fmt"
	//"fmt"
)

type Registration struct {
    ID      int `json:"id"`
    UserID  int `json:"userid" binding:"required"`
    EventID int `json:"eventid" binding:"required"`
}

func RegisterForEvent(userID int64, eventID int64) error {
    registered, err := IsUserRegistered(userID, eventID)
    

    if registered || err != nil {
        return errors.New("user already registered for the event")
    } 

    insertQuery := "INSERT INTO registration (userId, eventId) VALUES (?, ?)"
    _, err = db.DB.Exec(insertQuery, userID, eventID)
    if err != nil {
        return errors.New("unable to register user for event")
    }

    return nil
}

func IsUserRegistered(userID int64, eventID int64) (bool, error) {
    rows, err := db.DB.Query(
        "SELECT * FROM registration WHERE userId=? AND eventId=? ",
        userID, eventID,
    )

    defer rows.Close()
	
    if err != nil {
        return false, err
    }

	var registeredUser Registration
	for rows.Next() {
		_ = rows.Scan(&registeredUser.ID, &registeredUser.EventID, &registeredUser.UserID)
		fmt.Println(registeredUser)
		return true, errors.New("user already registered")
	}

    // Row exists → registered
    return false, nil
}

func CancelEvent(userID int64, eventID int64) error {
    userRegistered, _ := IsUserRegistered(userID, eventID)

    if !userRegistered {
        fmt.Println("user is not registered")
        return errors.New("user is not registered for event")
    }

    deleteQuery := "DELETE FROM registration WHERE userId=? AND eventId=?"
    _, err := db.DB.Exec(deleteQuery, userID, eventID)
    if err != nil {
        fmt.Println(err)
        return errors.New("unable to cancel user for event")
    }
    return nil
}