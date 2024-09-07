package service

import (
	"eShop/errs"
	"eShop/pkg/repository"
	"eShop/utils"
	"errors"
)

func SignIn(username, password string) (accessToken string, err error) {
	password = utils.GenerateHash(password)
	user, err := repository.GetUserByUsernameAndPassword(username, password)
	if err != nil {
		if errors.Is(err, errs.ErrRecordNotFound) {
			return "", errs.ErrIncorrectUsernameOrPassword
		}
		return "", err
	}

	accessToken, err = GenerateToken(user.ID, user.UserName, user.Role)
	if err != nil {
		return "", err
	}

	return accessToken, nil
}
