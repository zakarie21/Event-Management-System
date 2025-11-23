 USERS - ID, NAME, AGE, EMAIL, PASSWORD,

 EVENTS - ID, USERID, NAME, DESCRIPTION, DATE, LOCATION,


Endpoints of the service ->

GET -> /events -> Get a list of available events
response : user = [{id: 20, "userID" : 300, "name" : "jack", "Description" : "Concert in London", "Date" : "11/11/25", "Location" : "London"},
{id: 30, "userID" : 400, "name" : "Steve", "Description" : "Concert in Liverpool", "Date" : "11/12/25", "Location" : "Liverpool"}]


POST ->/events -> register an event
request body:  {"userID" : 750, "Name" : "Lucy", "Description" : "Cinema in town centtre" , "date" : "26/12/250" , "location" : "Manchester"}
response: {id : 600, "userID" : 750, "Name" : "Lucy", "Description" : "Cinema in town centtre" , "date" : "26/12/250" , "location" : "Manchester"}
error scenario: HTTP.BadRequest (400), HTTP.InternalError = (500)

GET -> /events/id -> Get the event associated with the id
response: {id : 600, "userID" : 750, "Name" : "Lucy", "Description" : "Cinema in town centtre" , "date" : "26/12/250" , "location" : "Manchester"}

PUT -> /events/<id> ->  update an event with id
request body: {id : 20, "userID" : 750, "Name" : "Lucy", "Description" : "Cinema in Glasgow" , "date" : "31/12/250" , "location" : "Glasgow"} 
response: "Your event has been updated" {id : 20, "userID" : 750, "Name" : "Lucy", "Description" : "Cinema in Glasgow" , "date" : "31/12/250" , "location" : "Glasgow"} 
error scenario: HTTP.BadRequest (400), HTTP.InternalError = (500)

DELETE -> /events/<id> -> Delete an event with specific id
response: {message : "Event deleted successfully"}
error scenario: HTTP.BadRequest (400), HTTP.InternalError = (500)

POST ->/signup 
Request: {Name: "Sarah" , Age: 21, Email: "Sarah@gmail.com", Password: "Purple"}
Response: {id: 2000, Name: "Sarah" , Age: 21, Email: "Sarah@gmail.com", Password: "Purple"}

POST ->/login
Request: {Email: "Sarah@gmail.com", Password: "Purple"}
response: {
	"message": "Logged In Successfully"
}{
	"token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6Im5hbWFuQGlkZmNiYW5rLmNvbSIsImV4cCI6MTc1NjU1ODYxNiwiaWQiOjB9.5m0k2aMmq0i1ER_iEvigw9djycY5dhmGC2fxCtRgFVI"
}

POST ->/events/<id>/register
DELETE ->/events/<id>/register

