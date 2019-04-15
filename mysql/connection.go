package mysql

import (
	"fmt"
	"github.com/jinzhu/gorm"
)


func GetDbConnection() (*gorm.DB, error) {
	fmt.Println("生成一个连接")
	db, err := gorm.Open("mysql", "root:62795828lovE@tcp(116.62.156.102:3306)/dl_app?charset=utf8")
	if err != nil {
		return nil, err
	}
	
	// grom全局设置
	db.SingularTable(true)
	return db, nil
}	
