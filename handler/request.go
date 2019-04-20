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
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	fmt.Println("收到下载请求！")

	// 解析请求
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		res := common.HandleRes(-1, "解析错误", nil, err)
		fmt.Fprintln(w, res)
		return
	}
	defer r.Body.Close()
 	reqBody := make(map[string]interface{})
	json.Unmarshal(data, &reqBody)

	// 赋值
	url, name := reqBody["addr"].(string), reqBody["name"].(string)

	// 往任务通道中添加下载任务
	execute.AddTask(mysql.DlTask{Addr: url, Name: name})

	res := common.HandleRes(0, "开始下载", nil, nil)
	fmt.Fprintln(w, res)
}

// 获取文件处理器
func GetFileHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// 先设置响应头
	w.Header().Add("Content-Type", "application/json")

	// 解析请求
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		resJSON := common.HandleRes(-1, "解析请求失败", nil, nil)
		fmt.Fprintln(w, resJSON)
		return
	}
	defer r.Body.Close()
	reqObj := make(map[string]interface{})
	json.Unmarshal(data, &reqObj)

	// 请求参数赋值
	id := int(reqObj["id"].(float64))
	findTask := mysql.DlTask{ ID: id }

	// 查询数据库
	resultTask := findTask.FindOneTask()

	// 返回相应
	resObj := map[string]interface{}{"path": resultTask.Path}
	resJSON := common.HandleRes(0, "查询成功", resObj, nil)
	fmt.Fprintln(w, resJSON)
}

// 删除文件处理器
func DeleteFileHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// 先设置响应头
	w.Header().Add("Content-Type", "application/json")

	// 解析请求
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		resJSON := common.HandleRes(-1, "解析失败", nil, nil)
		fmt.Fprintln(w, resJSON)
		return
	}
	defer r.Body.Close()
  reqObj := make(map[string]interface{})
	json.Unmarshal(data, &reqObj)

	// 入参赋值
	// TODO: 优化写法
	idInterfaceSli := reqObj["id"].([]interface{})
	var idSli []int
	for _, val := range idInterfaceSli {
		idSli = append(idSli, int(val.(float64)))
	}

	// 数据库操作
	findTask := new(mysql.DlTask)
	filePaths := findTask.FindFilePath(idSli)
	if len(filePaths) == 0 {
		fmt.Println(err)
		resJSON := common.HandleRes(-1, "文件已删除", nil, nil)
		fmt.Fprintln(w, resJSON)
		return
	}
	deleteTask := execute.DeleteTask{ FilePathSli: filePaths ,IdSli: idSli }
	execute.AddDelete(deleteTask)

	// 返回响应
	resJSON := common.HandleRes(0, "删除成功", nil, nil)
	fmt.Fprintln(w, resJSON)
}
