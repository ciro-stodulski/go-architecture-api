package userservice

import (
	"go-api/cmd/core/entities/user"
)

func (cuuc *userService) Register(user *user.User) (*user.User, error) {

	err := cuuc.SqlUser.Create(user)

	return user, err
}
