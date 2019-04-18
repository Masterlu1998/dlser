package mysql

type DlTask struct {
	ID          int    `gorm:"primary_key; AUTO_INCREMENT" json:"id"`
	Addr        string `json:"addr"`
	Name        string `json:"name"`
	Status      int    `json:"status"`
	Path        string `json:"path"`
	ContentType string `json:"contentType"`
}

func (this *DlTask) CreateTask() {
	db := GetDbConnection()
	db.Create(this)
}

func (this *DlTask) UpdateTask() {
	db := GetDbConnection()
	db.Model(this).Updates(*this)
}

func (_ *DlTask) FindTaskInfoList(index int, pageSize int, keywords string) []DlTask {
	db := GetDbConnection()
	keywords = "%" + keywords + "%"
	var dlTasks []DlTask
	db.Offset((index-1)*pageSize).Limit(pageSize).Where(`name LIKE ? AND status != -2`, keywords).Find(&dlTasks)
	return dlTasks
}

func (this *DlTask) FindOneTask() *DlTask {
	db := GetDbConnection()
	var resultTask = new(DlTask)
	db.Where(this).Find(resultTask)
	return resultTask
}

func (this *DlTask) FindFilePath(idSli []int) []string {
	db := GetDbConnection()
	var filePaths []string
	db.Model(this).Where("id in (?) AND status != -2", idSli).Pluck("path", &filePaths)
	return filePaths
}

func (this *DlTask) DeleteFileById(idSli []int) {
	db := GetDbConnection()
	db.Model(this).Where("id in (?)", idSli).Update("status", -2)
}
