package services

import (
	"github.com/uwaifo/introserver/mvc/domain"
	"github.com/uwaifo/introserver/mvc/utils"
)

func GetUser( userId int64) (*domain.User, *utils.ApplicationError)  {
return domain.GetUser(userId)
}