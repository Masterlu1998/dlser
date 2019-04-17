package mysql

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB
var err error

func init() {
	fmt.Println("新建mysql连接")
	db, err = gorm.Open("mysql", "root:62795828lovE@tcp(116.62.156.102:3306)/dl_app?charset=utf8")
	if err != nil {
		panic(err)
	}
	db.SingularTable(true)
}

func GetDbConnection() *gorm.DB {
	return db
}
