package service

import (
	"fmt"
	"is-theatre/internal/datastruct"
	"is-theatre/internal/entity"
	"is-theatre/internal/repository"
)

type UserService interface {
	GetUser(requestedUserId int64, userId int64) (*entity.User, error)
	DeleteUser(id int64, userId int64) error
	UpdateUser(user entity.User) (*entity.User, error)
}

type userService struct {
	dao repository.DAO
}

func NewUserService(dao repository.DAO) UserService {
	return &userService{dao: dao}
}

func (u *userService) GetUser(requestedUserId int64, userId int64) (*entity.User, error) {
	var userBySession *datastruct.User
	var err error

	userBySession, err = u.dao.NewUserQuery().GetUser(userId)
	if err != nil {
		return nil, fmt.Errorf("service error: %w", err)
	}

	userByRequest, err := u.dao.NewUserQuery().GetUser(requestedUserId)
	if err != nil {
		return nil, fmt.Errorf("service error: %w", err)
	}

	if userByRequest.ID == userBySession.ID || userBySession.Role == datastruct.ADMIN {
		return userRepToEntity(userByRequest), nil
	}

	return &entity.User{
		ID:       userByRequest.ID,
		NickName: userByRequest.NickName,
		Email:    userByRequest.Email,
	}, nil
}

func (u *userService) DeleteUser(id int64, userId int64) error {
	user, err := u.dao.NewUserQuery().GetUser(userId)
	if err != nil {
		return err
	}

	if user.Role == datastruct.ADMIN || id == user.ID {
		err = u.dao.NewUserQuery().DeleteUser(id)
		if err != nil {
			return err
		}
		return nil
	}

	return fmt.Errorf("you have no access")
}

func (u *userService) UpdateUser(person entity.User) (*entity.User, error) {
	// email checking
	// phone number checking
	user, err := u.dao.NewUserQuery().GetUser(person.ID)
	if err != nil {
		return nil, err
	}

	if user.Role == datastruct.ADMIN || user.ID == person.ID {

		tmp, err := u.dao.NewUserQuery().GetUser(person.ID)
		if err != nil {
			return nil, fmt.Errorf("update user error: %w", err)
		}

		updateUser(person, tmp)

		updatedUser, err := u.dao.NewUserQuery().UpdateUser(tmp)
		if err != nil {
			return nil, err
		}
		return userRepToEntity(updatedUser), nil
	}
	return nil, fmt.Errorf("you have no access")
}

func userRepToEntity(u *datastruct.User) *entity.User {
	return &entity.User{
		ID:       u.ID,
		NickName: u.NickName,
		Email:    u.Email,
		Password: u.Password,
	}
}

func updateUser(eu entity.User, user *datastruct.User) {
	if len(eu.Email) > 0 {
		user.Email = eu.Email
	}
	if len(eu.NickName) > 0 {
		user.NickName = eu.NickName
	}
	if len(eu.Password) > 0 {
		user.Password = eu.Password
	}
}
