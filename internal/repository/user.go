package repository

import (
	"fmt"
	"github.com/Masterminds/squirrel"
	"is-theatre/internal/datastruct"
)

type UserQuery interface {
	GetUser(id int64) (*datastruct.User, error)
	GetUsers(limit, offset uint64) ([]datastruct.User, error)
	CreateUser(user datastruct.User) (*int64, error)
	UpdateUser(person *datastruct.User) (*datastruct.User, error)
	DeleteUser(userID int64) error
	GetUserPasswordByEmail(email string) (*string, error)
	GetEmailByUserID(id int64) (string, error)
	GetUserIdByEmail(email string) (*int64, error)
}

type userQuery struct{}

func (u *userQuery) CreateUser(user datastruct.User) (*int64, error) {
	if len(user.Role) == 0 {
		user.Role = "user"
	}

	qb := dbQueryBuilder().
		Insert(datastruct.UserTableName).
		Columns("nickName", "email", "password", "role", "verified", "emailCode").
		Values(user.NickName, user.Email, user.Password, user.Role, user.Verified, user.EmailCode).
		Suffix("RETURNING id")

	var id int64
	err := qb.QueryRow().Scan(&id)
	if err != nil {
		return nil, err
	}

	return &id, nil
}

func (u *userQuery) GetUser(id int64) (*datastruct.User, error) {
	db := dbQueryBuilder().
		Select("nickName", "email", "role", "verified", "emailCode", "id").
		From(datastruct.UserTableName).
		Where(squirrel.Eq{"id": id})

	us := datastruct.User{}
	err := db.QueryRow().Scan(&us.NickName, &us.Email, &us.Role, &us.Verified, &us.EmailCode, &us.ID)
	if err != nil {
		return nil, fmt.Errorf("connot scan user: %w", err)
	}

	return &us, nil
}

func (u *userQuery) GetUsers(limit, offset uint64) ([]datastruct.User, error) {
	db := dbQueryBuilder().
		Select("nickName",
			"email",
			"password",
			"role",
			"verified",
			"emailCode",
			"id").
		From(datastruct.UserTableName).
		Limit(limit).Offset(offset)

	var users []datastruct.User
	var user datastruct.User
	rows, err := db.Query()
	if err != nil {
		return nil, fmt.Errorf("get users error: %w", err)
	}

	for rows.Next() {
		err = rows.Scan(&user.NickName,
			&user.Email,
			&user.Password,
			&user.Role,
			&user.Verified,
			&user.EmailCode,
			&user.ID)
		if err != nil {
			return nil, fmt.Errorf("get users error: %w", err)
		}
		users = append(users, user)
	}

	return users, nil
}

func (u *userQuery) DeleteUser(userID int64) error {
	db := dbQueryBuilder().
		Delete(datastruct.UserTableName).
		From(datastruct.UserTableName).
		Where(squirrel.Eq{"id": userID})

	_, err := db.Exec()
	if err != nil {
		return err
	}
	return nil
}

func (u *userQuery) UpdateUser(user *datastruct.User) (*datastruct.User, error) {

	fromDB, err := u.GetUser(user.ID)
	if err != nil {
		return nil, fmt.Errorf("update user error: %w", err)
	}

	updated := updateUser(fromDB, user)

	db := dbQueryBuilder().
		Update(datastruct.UserTableName).
		SetMap(map[string]interface{}{
			"nickName":  updated.NickName,
			"email":     updated.Email,
			"password":  updated.Password,
			"role":      updated.Role,
			"verified":  updated.Verified,
			"emailCode": updated.EmailCode,
			"id":        updated.ID,
		}).Where(squirrel.Eq{"id": user.ID}).
		Suffix("RETURNING nickName, email, password, role, verified, emailCode, id")

	var updatedPerson datastruct.User
	err = db.QueryRow().Scan(
		&updatedPerson.NickName,
		&updatedPerson.Email,
		&updatedPerson.Password,
		&updatedPerson.Role,
		&updatedPerson.Verified,
		&updatedPerson.EmailCode,
		&updatedPerson.ID)
	if err != nil {
		return nil, err
	}

	return &updatedPerson, nil
}

func (u *userQuery) GetUserPasswordByEmail(email string) (*string, error) {
	db := dbQueryBuilder().
		Select("password").
		From(datastruct.UserTableName).
		Where(squirrel.Eq{"email": email})

	var password string
	err := db.QueryRow().Scan(&password)
	if err != nil {
		return nil, fmt.Errorf("email and password don't match %w", err)
	}
	return &password, nil
}

func (u *userQuery) GetEmailByUserID(id int64) (string, error) {
	db := dbQueryBuilder().
		Select("email").
		From(datastruct.UserTableName).
		Where(squirrel.Eq{"id": id})

	var email string
	err := db.QueryRow().Scan(&email)
	if err != nil {
		return "", fmt.Errorf("get email by id: %w", err)
	}

	return email, nil
}

func (u *userQuery) GetUserIdByEmail(email string) (*int64, error) {
	db := dbQueryBuilder().Select("id").
		From(datastruct.UserTableName).
		Where(squirrel.Eq{"email": email})

	var id int64
	err := db.QueryRow().Scan(&id)
	if err != nil {
		return nil, fmt.Errorf("get user id by email: %w", err)
	}

	return &id, nil
}

func updateUser(fromDB, new *datastruct.User) (updated datastruct.User) {
	updated = *fromDB

	if new.ID > 0 {
		updated.ID = new.ID
	}

	if len(new.NickName) > 0 {
		updated.NickName = new.NickName
	}

	if len(new.Email) > 0 {
		updated.Email = new.Email
	}

	if len(new.Password) > 0 {
		updated.Password = new.Password
	}

	if len(new.Role) > 0 {
		updated.Role = new.Role
	}

	if new.Verified {
		updated.Verified = new.Verified
	}

	if new.EmailCode > 0 {
		updated.EmailCode = new.EmailCode
	}

	return
}
