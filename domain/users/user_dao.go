package users

import (
	"fmt"
	"github.com/garyjdn/bookstore_users-api/utils/errors"
)

var (
	userDB = make(map[int64]*User)
)

func (user *User) Get() *errors.RestErr {
	res := userDB[user.Id]
	if res == nil {
		return errors.NotFoundError(fmt.Sprintf("user %d not found", user.Id))
	}
	user.Id = res.Id
	user.FirstName = res.FirstName
	user.LastName = res.LastName
	user.Email = res.Email
	user.DateCreated = res.DateCreated
	return nil
}

func (user *User) Save() *errors.RestErr {
	current := userDB[user.Id]
	if current != nil {
		if current.Email == user.Email {
			return errors.BadRequestError(fmt.Sprintf("email %s already exist", user.Email))
		}
		return errors.BadRequestError(fmt.Sprintf("user %d already exist", user.Id))
	}
	userDB[user.Id] = user
	return nil
}
