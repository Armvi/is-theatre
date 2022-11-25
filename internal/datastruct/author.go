package datastruct

const AuthorTableName = "Author"

type Author struct {
	Id         int64
	Name       string
	SecondName string
	Country    string
	Century    int
}
