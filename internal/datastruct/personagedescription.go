package datastruct

const PersonageDescriptionTableName = "PersonageDescription"

type PersonageDescription struct {
	Id          int64
	PersonageId int64
	Age         string
	Height      string
	Weight      string
	Voice       string
}
