package mysql

import(
	"fmt"
)

type DlTask struct {
	ID int `gorm:"primary_key; AUTO_INCREMENT"`
	Addr string
	Name   string
	Status int
}

func (*DlTask) CreateTask(task *DlTask) {
	db, err := GetDbConnection()
	if err != nil {
		fmt.Println(err)
	}
	db.Create(task)
}

func (*DlTask) UpdateTask(task *DlTask) {
	db, err := GetDbConnection()
	if err != nil {
		fmt.Println(err)
	}
	db.Save(task)
}