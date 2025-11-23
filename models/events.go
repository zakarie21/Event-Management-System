package models

import (
	//"fmt"
	"RESTApi/db"
	"errors"
	"fmt"
	"time"
)

type Events struct {
	ID          int       `json:"id"`
	UserID      int       `json:"userid"`
	Name        string    `json:"name" binding: "required"`
	Description string    `json:"description" binding: "required"`
	Date        time.Time `json:"date" binding: "required"`
	Location    string    `json:"location" binding: "required"`
}



func GetAllEvents() ([]Events, error) {
	var events []Events

	rows, err := db.DB.Query(`SELECT * FROM events`)
	if err != nil {
		return events, errors.New("command failed to exectue")
	}

	for rows.Next() {
		var event Events
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.Date, &event.UserID)
		if err != nil {
			return events, errors.New("unable to scan rows")
		}
		events = append(events, event)

	}

	return events, nil
}

func (e *Events) Save() error {
	insertQuery := `INSERT INTO events (name, description,location,dateTime, userId) VALUES (?,?,?,?,?)`
	_, err := db.DB.Exec(insertQuery, e.Name, e.Description, e.Location, e.Date, e.UserID)
	if err != nil {
		return errors.New("unable to create table")
	}

	return nil
}

func GetEvent(id int64) (Events, error) {
	eventRow := db.DB.QueryRow(`SELECT * FROM events WHERE ID=(?)`, id)

	var event Events
	err := eventRow.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.Date, &event.UserID)
	if err != nil {
		return event, err
	}

	return event, nil

}

func DeleteEvent(id int64) error {
	deleteQuery := `DELETE FROM events WHERE ID=(?)`

	_, err := db.DB.Exec(deleteQuery, id)

	if err != nil {
		fmt.Println(err)
		return errors.New("unable to delete table")
	}

	return nil
}

func (e Events) UpdateEvent(id int64) error {
	
	
	updateQuery:= `UPDATE events SET name=?,description=?,location=?,dateTime=? WHERE id=?`

	_, err := db.DB.Exec(updateQuery, e.Name, e.Description,e.Location, e.Date, id)
	if err != nil {
		return errors.New("command failed to execute")
	}
	
	return  nil
}
	

