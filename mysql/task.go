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
	DB.Create(this)
}

func (this *DlTask) UpdateTask() {
	DB.Model(this).Updates(*this)
}

func (_ *DlTask) FindTaskInfoList(index int, pageSize int, keywords string, startTime time.Time, endTime time.Time) []DlTask {
	keywords = "%" + keywords + "%"
	var dlTasks []DlTask
	startTimeFlag, endTimeFlag := time.Date(0001, time.January, 1, 0, 0, 0, 0, time.UTC), time.Date(0001, time.January, 1, 0, 0, 0, 0, time.UTC)
	isUseTime := !(startTime.Equal(startTimeFlag) && endTime.Equal(endTimeFlag))
	if isUseTime {
		DB.Offset((index-1)*pageSize).Limit(pageSize).Where(`name LIKE ? AND status != -2`, keywords).Where(`create_time BETWEEN ? AND ?`, startTime, endTime).Find(&dlTasks)
	} else {
		DB.Offset((index-1)*pageSize).Limit(pageSize).Where(`name LIKE ? AND status != -2`, keywords).Find(&dlTasks)
	}
	return dlTasks
}

func (this *DlTask) FindOneTask() *DlTask {
	var resultTask = new(DlTask)
	DB.Where(this).Find(resultTask)
	return resultTask
}

func (this *DlTask) FindFilePath(idSli []int) []string {
	var filePaths []string
	DB.Model(this).Where("id in (?) AND status != -2", idSli).Pluck("path", &filePaths)
	return filePaths
}

func (this *DlTask) DeleteFileById(idSli []int) {
	DB.Model(this).Where("id in (?)", idSli).Update("status", -2)
}
