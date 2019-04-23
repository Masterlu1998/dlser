package handler

import (
	"dlser/common"
	"dlser/execute"
	"dlser/mysql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// 请求下载处理器
func RequestHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// 返回响应
	common.WriteHeader(w, "Content-Type", "application/json")
	fmt.Println("收到下载请求！")

	// 解析请求
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		common.HandleErr(w, -1, err, "解析错误")
		return
	}
	defer r.Body.Close()
	var reqObj struct {
		Addr string `json:"addr"`
		Name string `json:"name"`
	}
	err = json.Unmarshal(data, &reqObj)
	if err != nil {
		common.HandleErr(w, -1, nil, "解析失败")
		return
	}

	// 赋值
	url, name := reqObj.Addr, reqObj.Name

	// 往任务通道中添加下载任务
	execute.AddTask(mysql.DlTask{Addr: url, Name: name})

	common.HandleSuc(w, 0, nil, "开始下载")
}

// 获取文件处理器
func GetFileHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// 先设置响应头
	common.WriteHeader(w, "Content-Type", "application/json")

	// 解析请求
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		common.HandleErr(w, -1, err, "解析请求失败")
		return
	}
	defer r.Body.Close()
	var reqObj struct {
		ID int `json:"id"`
	}
	err = json.Unmarshal(data, &reqObj)
	if err != nil {
		common.HandleErr(w, -1, err, "json解析失败")
		return
	}

	// 请求参数赋值
	id := reqObj.ID
	findTask := mysql.DlTask{ID: id}

	// 查询数据库
	resultTask := findTask.FindOneTask()

	// 返回相应
	resObj := map[string]interface{}{"path": resultTask.Path}
	common.HandleSuc(w, 0, resObj, "查询成功")
}

// 删除文件处理器
func DeleteFileHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// 先设置响应头
	common.WriteHeader(w, "Content-Type", "application/json")

	// 解析请求
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		common.HandleErr(w, -1, err, "请求读取失败")
		return
	}
	defer r.Body.Close()
	var reqObj struct {
		ID []int `json:"id"`
	}
	err = json.Unmarshal(data, &reqObj)
	if err != nil {
		common.HandleErr(w, -1, err, "json解析失败")
		return
	}

	// 入参赋值
	idInterfaceSli := reqObj.ID

	// 数据库操作
	findTask := new(mysql.DlTask)
	filePaths := findTask.FindFilePath(idInterfaceSli)
	if len(filePaths) == 0 {
		common.HandleErr(w, -1, err, "文件已删除")
		return
	}
	deleteTask := execute.DeleteTask{FilePathSli: filePaths, IdSli: idInterfaceSli}
	execute.AddDelete(deleteTask)

	// 返回响应
	common.HandleSuc(w, 0, nil, "删除成功")
}
