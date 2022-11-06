package repository

type UserQuery interface {
	GetUser(id int64)
}
