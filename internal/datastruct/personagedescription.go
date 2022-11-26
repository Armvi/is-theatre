package datastruct

const PersonageDescriptionTableName = "PersonageDescription"

type PersonageDescription struct {
	Id          int64
	Age         string
	Voice       string
	Height      string
	Weight      string
	Gender      bool
	Description string
}
