package execute

import (
	"os"
	"mime"
	"fmt"
	"dlser/mysql"
	"net/http"
	"io/ioutil"
	"io"
	"bytes"
)

var (
	dlch = make(chan mysql.DlTask, 10)
)

func init() {
	go scheduler()
}

func scheduler() {	
	for val := range dlch {
		fmt.Println("执行下载")
		go executeTask(&val)
	}
}

func AddTask(task mysql.DlTask) {
	dlch <- task
}

func executeTask(task *mysql.DlTask) {
	url, fileName := task.Addr, task.Name
	task.Status = 0;
	fmt.Println("开始下载")
	task.CreateTask()

	// 请求http文件
	res, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		task.Status = -1
		task.UpdateTask()
		return
	}

	// 读取http文件字节流
	fileData, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		fmt.Println(err)
		task.Status = -1
		task.UpdateTask()
		return
	}

	// 获取文件类型
	contentType := http.DetectContentType(fileData)
	postfixSlice, err := mime.ExtensionsByType(contentType)
	postfix := postfixSlice[0]
	if err != nil {
		fmt.Println(err)
		task.Status = -1
		task.UpdateTask()
		return
	}

	// 创建本地空文件
	path := "/file/" + fileName + postfix
	f, err := os.Create("." + path)
	defer f.Close()
	if err != nil {
		fmt.Println(err)
		task.Status = -1
		task.UpdateTask()
		return
	}

	// 因为之前的res.Body reader已经被读取清空 重新生成一个新的reader
	rawData := bytes.NewReader(fileData)
	io.Copy(f, rawData)

	task.Status = 1
	task.Path = "/download" + path
	task.ContentType = contentType
	task.UpdateTask()
	fmt.Println("下载完成")
}
