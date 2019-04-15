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

func (this *DlTask) CreateTask() {
	db, err := GetDbConnection()
	if err != nil {
		fmt.Println(err)
	}
	db.Create(this)
}

func (this *DlTask) UpdateTask() {
	db, err := GetDbConnection()
	if err != nil {
		fmt.Println(err)
	}
	db.Save(this)
}

func (_ *DlTask) FindTaskList(index int, pageSize int, keywords string) []DlTask  {
	db, err := GetDbConnection()
	if err != nil {
		fmt.Println(err)
	}
	keywords = "%" + keywords + "%"
	var dlTasks []DlTask 
	db.Offset((index - 1) * pageSize).Limit(pageSize).Where(`name LIKE ?`, keywords).Find(&dlTasks)
	return dlTasks
}