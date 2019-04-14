package execute

import (
	"io"
	"net/http"
	"os"
	"fmt"
)

type DlTask struct {
	// id     int
	Addr   string
	Name   string
	Status int
}

var (
	dlch chan DlTask = make(chan DlTask)
)

func init() {
	go scheduler()
}

func scheduler() {	
	for val := range dlch {
		fmt.Println("执行下载", val)
		executeTask(&val)
	}
}

func AddTask(task DlTask) {
	dlch <- task
}

func executeTask(task *DlTask) {
	url := task.Addr
	fmt.Println("开始下载")
	res, err := http.Get(url)
	fmt.Println("1111")

	if err != nil {
		task.Status = -1
		return
	}
	downloadPath := "./file/" + task.Name
	f, err := os.Create(downloadPath)
	defer f.Close()
	if err != nil {
		fmt.Println(err)
		task.Status = -1
		return
	}
	io.Copy(f, res.Body)
	fmt.Println("下载完成")
}
