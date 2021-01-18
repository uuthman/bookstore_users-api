package services

import (
	"github.com/uuthman/bookstore_users-api/services/utils/errors"
	"github.com/uuthman/bookstore_users-api/domain/users"
)

func GetUser(userId int64) (*users.User,*errors.RestErr){
	result := &users.User{ID:userId}

	if err := result.Get(); err != nil{
		return nil,err
	}

	return result,nil

}

func CreateUser(user users.User) (*users.User,*errors.RestErr){

	if err := user.Validate(); err != nil{
		return nil,err
	}

	if err := user.Save(); err != nil{
		return nil,err
	}
	return &user,nil
}