package mysql

import (
	"fmt"
	"sync"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB
var l sync.Mutex
var err error

func GetDbConnection() (*gorm.DB, error) {
	if db == nil {
		fmt.Println("创建一个新的连接")
		l.Lock()
		defer l.Unlock()
		if db == nil {
			db, err = gorm.Open("mysql", "root:62795828lovE@tcp(116.62.156.102:3306)/dl_app?charset=utf8")
			if err != nil {
				fmt.Println(err)
				return nil, err
			}
			// gorm全局设置
			db.SingularTable(true)
		}
	}
	return db, nil
}
