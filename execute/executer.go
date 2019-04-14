package execute

import (
	"os/exec"
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

	cmd := exec.Command("wget", url, "-P", "./file")
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
		task.Status = -1
		return
	}

	task.Status = 1
	fmt.Println("下载完成")
}
