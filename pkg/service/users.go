package service

import (
	"eShop/errs"
	"eShop/models"
	"eShop/pkg/repository"
	"eShop/utils"
	"errors"
)

func GetAllUsers() (users []models.User, err error) {
	users, err = repository.GetAllUsers()
	if err != nil {
		return nil, err
	}

	return users, nil
}

func GetUserByID(id uint) (user models.User, err error) {
	user, err = repository.GetUserByID(id)
	if err != nil {
		return user, err
	}

	return user, nil
}

func CreateUser(user models.User) error {
	// 1. Check username uniqueness
	userFromDB, err := repository.GetUserByUsername(user.Username)
	if err != nil && !errors.Is(err, errs.ErrRecordNotFound) {
		return err
	}

	if userFromDB.ID > 0 {
		return errs.ErrUsernameUniquenessFailed
	}

	user.Role = "user"

	// 2. Generate password hash
	user.Password = utils.GenerateHash(user.Password)

	// 3. Repository call
	err = repository.CreateUser(user)
	if err != nil {
		return err
	}

	return nil
}
