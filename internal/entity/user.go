package entity

const UserTableName = "User"

type User struct {
	ID       int64  `json:"id,omitempty"`
	NickName string `json:"nick_name,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

type Role string

const (
	ADMIN Role = "admin"
	USER  Role = "user"
)
