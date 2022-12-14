package service

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"is-theatre/internal/datastruct"
	"is-theatre/internal/repository"
	"strconv"
)

type AuthService interface {
	SignUp(user datastruct.User) (*int64, error)
	SignIn(email, password string) (*string, error)
	Logout(userId int64) error
}

type authService struct {
	dao          repository.DAO
	tokenManager TokenManager
}

func NewAuthService(dao repository.DAO, manager TokenManager) AuthService {
	return &authService{
		dao:          dao,
		tokenManager: manager,
	}
}

func (a *authService) SignUp(user datastruct.User) (*int64, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.MinCost)
	if err != nil {
		return nil, err
	}
	user.Password = string(hashedPassword)
	if len(user.Role) == 0 {
		user.Role = "user"
	}
	id, err := a.dao.NewUserQuery().CreateUser(user)
	if err != nil {
		return nil, err
	}
	return id, nil
}

func (a *authService) SignIn(email, reqPassword string) (*string, error) {
	password, err := a.dao.NewUserQuery().GetUserPasswordByEmail(email)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(*password), []byte(reqPassword))
	if err != nil {
		return nil, fmt.Errorf("passwords don't match %v", err)
	} else {
		userID, err := a.dao.NewUserQuery().GetUserIdByEmail(email)
		if err != nil {
			return nil, err
		}

		jwt, err := a.tokenManager.NewJWT(strconv.Itoa(int(*userID)))
		if err != nil {
			return nil, err
		}

		return &jwt, nil
	}
}

func (a *authService) Logout(userID int64) error {
	_, err := a.tokenManager.NewJWT(strconv.Itoa(int(userID)))
	if err != nil {
		return err
	}

	return nil
}
