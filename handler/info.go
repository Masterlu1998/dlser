package handler

import (
	"dlser/common"
	"dlser/mysql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func GetTaskList(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// 读取请求参数
	data, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		fmt.Println(err)
		resJSON := common.HandleRes(-1, "读取失败", nil, nil)
		fmt.Fprintln(w, resJSON)
		return
	}
	var reqObj map[string]interface{}
	json.Unmarshal(data, &reqObj)

	// 赋值参数
	index, pageSize, keywords := int(reqObj["index"].(float64)), int(reqObj["pageSize"].(float64)), reqObj["keywords"].(string)
	
	// 调用接口获取查询结果
	dlTask := new(mysql.DlTask)
	dlTasks := dlTask.FindTaskList(index, pageSize, keywords)

	// 设置响应头，返回响应
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	resObj := map[string]interface{}{"taskList": dlTasks}
	resJSON := common.HandleRes(0, "查询成功", resObj, nil)
	fmt.Fprintln(w, resJSON)
}

func GetFile(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// 解析请求
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		resJSON := common.HandleRes(-1, "解析请求失败", nil, nil)
		fmt.Fprintln(w, resJSON)
		return
	}
	var reqObj map[string]interface{}
	json.Unmarshal(data, &reqObj)

	// 请求参数赋值
	id := int(reqObj["id"].(float64))
	findTask := mysql.DlTask{ ID: id }

	// 查询数据库
	resultTask := findTask.FindOneTask()

	// 返回相应
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	resObj := map[string]interface{}{ "path": resultTask.Path }
	resJSON := common.HandleRes(0, "查询成功", resObj, nil)
	fmt.Fprintln(w, resJSON)
}
