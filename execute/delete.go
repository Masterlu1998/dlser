package execute

import (
	"fmt"
	"os"
	"strings"
	"dlser/mysql"
)

type DeleteTask struct {
	FilePathSli []string
	IdSli []int
}

var mvch = make(chan DeleteTask, 10)

func init() {
	go deletor()
}

func deletor() {
	for deleteTaskObj := range mvch {
		fmt.Println("开始删除")
		go executeDelete(deleteTaskObj)
	}
}

func AddDelete(deleteTaskObj DeleteTask) {
	mvch <- deleteTaskObj
}

func executeDelete(deleteTask DeleteTask) {
	filePathSli, idSli := deleteTask.FilePathSli, deleteTask.IdSli
	for _, filePath := range filePathSli {
		realFilePath := "./" + strings.Join(strings.Split(filePath, "/")[2 :], "/")
		err := os.Remove(realFilePath)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	var dlTask mysql.DlTask
	dlTask.DeleteFileById(idSli)
	fmt.Println("删除完成")
}
