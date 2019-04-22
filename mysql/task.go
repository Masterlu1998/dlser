package mysql

import (
	"time"
)

type jsonTime time.Time

type DlTask struct {
	ID          int       `gorm:"primary_key; AUTO_INCREMENT" json:"id"`
	Addr        string    `json:"addr"`
	Name        string    `json:"name"`
	Status      int       `json:"status"`
	Path        string    `json:"path"`
	ContentType string    `json:"contentType"`
	CreateTime  time.Time `gorm:"default:'CURRENT_TIMESTAMP'" json:"createTime"`
}

func (this *DlTask) CreateTask() {
	db := GetDbConnection()
	db.Create(this)
}

func (this *DlTask) UpdateTask() {
	db := GetDbConnection()
	db.Model(this).Updates(*this)
}

func (_ *DlTask) FindTaskInfoList(index int, pageSize int, keywords string, startTime time.Time, endTime time.Time) []DlTask {
	db := GetDbConnection()
	keywords = "%" + keywords + "%"
	var dlTasks []DlTask
	startTimeFlag, endTimeFlag := time.Date(0001, time.January, 1, 0, 0, 0, 0, time.UTC), time.Date(0001, time.January, 1, 0, 0, 0, 0, time.UTC)
	isUseTime := !(startTime.Equal(startTimeFlag) && endTime.Equal(endTimeFlag))
	if isUseTime {
		db.Offset((index-1)*pageSize).Limit(pageSize).Where(`name LIKE ? AND status != -2`, keywords).Where(`create_time BETWEEN ? AND ?`, startTime, endTime).Find(&dlTasks)
	} else {
		db.Offset((index-1)*pageSize).Limit(pageSize).Where(`name LIKE ? AND status != -2`, keywords).Find(&dlTasks)
	}
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
