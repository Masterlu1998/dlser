package mysql

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB
var err error

func init() {
	DB, err = gorm.Open("mysql", "root:62795828lovE@tcp(116.62.156.102:3306)/dl_app?charset=utf8&parseTime=true")
	if err != nil {
		panic(err)
	}
	DB.SingularTable(true)
	fmt.Println("新建mysql连接")
}

