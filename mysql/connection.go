package mysql

import (
	"github.com/jinzhu/gorm"
)

func GetDbConnection() (*gorm.DB, error) {
	db, err := gorm.Open("mysql", "root:62795828lovE@tcp(116.62.156.102:3306)/dl_app?charset=utf8")
	if err != nil {
		return nil, err
	}

	// grom全局设置
	db.SingularTable(true)
	return db, nil
}	
