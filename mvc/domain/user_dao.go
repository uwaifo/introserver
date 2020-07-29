package domain

import (
	"fmt"
	"github.com/uwaifo/introserver/mvc/utils"
	"net/http"
)

var (
	users = map[int64]*User{
		001: {Id: 001, FirstName: "Frank", LastName: "White", Email: "myemail@gmail.com"},
		212: {Id: 212, FirstName: "James", LastName: "May", Email: "myemail@gmail.com"},
		123: {Id: 123, FirstName: "Fede", LastName: "Leon", Email: "myemail@gmail.com"},
	}

	//UserDao userDaoInterface
)
func GetUser(userId int64) (*User, *utils.ApplicationError)  {

	if  user := users[userId];user != nil {
		return user, nil
		
	}
	///return user, nil
	return  nil, &utils.ApplicationError{
		Message: fmt.Sprintf("a user with the id of %v was not found", userId),
		StatusCode: http.StatusNotFound,
		Code: "not_found",
	}


	//errors.New(fmt.Sprintf("a user with the id of %v was not found", userId))

}
