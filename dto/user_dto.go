package dto

import "zzidun.tech/vforum0/model"

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func UserDto_Make(user model.User) User {
	return User{
		Name:  user.Name,
		Email: user.Email,
	}
}
