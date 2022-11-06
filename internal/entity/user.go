package entities

const UserTableName = "User"

type User struct {
	ID          int64  `db:"id"`
	NickName    string `db:"first_name"`
	Email       string `db:"email"`
	Password    string `db:"password"`
	PhoneNumber string `db:"phone_number"`
	Role        Role   `db:"role"`
	Verified    bool   `db:"verified"`
	EmailCode   int32  `db:"email_code"`
}

type Role string

const (
	ADMIN Role = "admin"
	USER  Role = "user"
)
