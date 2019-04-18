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

func GetTaskListHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
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
