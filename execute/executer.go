package execute

import (
	"os/exec"
	"fmt"
	"dlser/mysql"
)


var (
	dlch = make(chan mysql.DlTask)
)

func init() {
	go scheduler()
}

func scheduler() {	
	for val := range dlch {
		fmt.Println("执行下载")
		executeTask(&val)
	}
}

func AddTask(task mysql.DlTask) {
	dlch <- task
}

func executeTask(task *mysql.DlTask) {
	url := task.Addr
	task.Status = 0;
	fmt.Println("开始下载")

	task.CreateTask(task)
	cmd := exec.Command("wget", url, "-P", "./file")
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
		task.Status = -1
		return
	}

	task.Status = 1
	task.UpdateTask(task)
	fmt.Println("下载完成")
}
