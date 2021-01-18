package users

import (
	"strings"
	"github.com/uuthman/bookstore_users-api/services/utils/errors"
)

type User struct{
	ID int64 `json:"id"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Email string `json:"email"`
	DateCreated string `json:"dated_created"`
} 

func (user *User) Validate() *errors.RestErr{
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == ""{
		return errors.NewBadRequestError("invalid email address")
	}
	return nil
}