package datastruct

import "time"

const WorkerTableName = "Worker"

type Worker struct {
	Id         int64
	Name       string
	SecondName string
	BirthDate  time.Time
	Salary     float64
}
