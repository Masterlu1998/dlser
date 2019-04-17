package mysql

type DlTask struct {
	ID          int    `gorm:"primary_key; AUTO_INCREMENT" json:"id"`
	Addr        string `json:"addr"`
	Name        string `json:"name"`
	Status      int    `json:"status"`
	Path        string `json:"path"`
	ContentType string `json:"contentType`
}

func (this *DlTask) CreateTask() {
	db := GetDbConnection()
	db.Create(this)
}

func (this *DlTask) UpdateTask() {
	db := GetDbConnection()
	db.Save(this)
}

func (_ *DlTask) FindTaskList(index int, pageSize int, keywords string) []DlTask {
	db := GetDbConnection()
	keywords = "%" + keywords + "%"
	var dlTasks []DlTask
	db.Offset((index-1)*pageSize).Limit(pageSize).Where(`name LIKE ?`, keywords).Find(&dlTasks)
	return dlTasks
}

func (this *DlTask) FindOneTask() *DlTask {
	db := GetDbConnection()
	var resultTask = new(DlTask)
	db.Where(this).Find(resultTask)
	return resultTask
}
