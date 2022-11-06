package entity

const UserTableName = "User"

type User struct {
	Id          int64  `json:"id,omitempty"`
	NickName    string `json:"nick_name,omitempty"`
	Email       string `json:"email,omitempty"`
	Password    string `json:"password,omitempty"`
	PhoneNumber string `json:"phone_number,omitempty"`
	Role        Role   `json:"role,omitempty"`
	Verified    bool   `json:"verified,omitempty"`
	EmailCode   int32  `json:"email_code,omitempty"`
}

type Role string

const (
	ADMIN Role = "admin"
	USER  Role = "user"
)
